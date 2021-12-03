package watchmanager

import (
	"errors"
	"fmt"
	"sync"

	"google.golang.org/grpc/connectivity"

	"github.com/Axway/agent-sdk/pkg/util/log"
	"github.com/Axway/agent-sdk/pkg/watchmanager/proto"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// Manager - Interface to manage watch connections
type Manager interface {
	RegisterWatch(topic string, eventChan chan *proto.Event, errChan chan error) (string, error)
	CloseWatch(id string) error
	CloseAll()
	CloseConn()
	Status() bool
}

// TokenGetter - function to acquire token
type TokenGetter func() (string, error)

type watchManager struct {
	cfg                *Config
	clientMap          map[string]*watchClient
	connection         *grpc.ClientConn
	options            *watchOptions
	logger             logrus.FieldLogger
	newWatchClientFunc newWatchClientFunc
	mutex              sync.Mutex
}

// New - Creates a new watch manager
func New(cfg *Config, opts ...Option) (Manager, error) {
	err := cfg.validateCfg()
	if err != nil {
		return nil, err
	}

	entry := logrus.NewEntry(log.Get())

	manager := &watchManager{
		cfg:                cfg,
		logger:             entry.WithField("package", "watchmanager"),
		clientMap:          make(map[string]*watchClient),
		options:            newWatchOptions(),
		newWatchClientFunc: proto.NewWatchClient,
	}

	for _, opt := range opts {
		opt.apply(manager.options)
	}

	manager.connection, err = manager.createConnection()
	if err != nil {
		log.Errorf("failed to establish connection with watch service: %s", err.Error())
	}
	return manager, err
}

func (m *watchManager) createConnection() (*grpc.ClientConn, error) {
	grpcDialOptions := []grpc.DialOption{
		withKeepaliveParams(m.options.keepAlive.time, m.options.keepAlive.timeout),
		withRPCCredentials(m.cfg.TenantID, m.cfg.TokenGetter),
		withTLSConfig(m.options.tlsCfg),
		chainStreamClientInterceptor(
			logrusStreamClientInterceptor(m.options.loggerEntry),
		),
	}

	address := fmt.Sprintf("%s:%d", m.cfg.Host, m.cfg.Port)
	log.Infof("connecting to watch service. host: %s. port: %d", m.cfg.Host, m.cfg.Port)

	return grpc.Dial(address, grpcDialOptions...)
}

// RegisterWatch - Registers a subscription with watch service using topic
func (m *watchManager) RegisterWatch(link string, events chan *proto.Event, errors chan error) (string, error) {
	client, err := newWatchClient(
		m.connection,
		clientConfig{
			topicSelfLink: link,
			tokenGetter:   m.cfg.TokenGetter,
			events:        events,
			errors:        errors,
		},
		m.newWatchClientFunc,
	)
	if err != nil {
		return "", err
	}

	subscriptionID, _ := uuid.NewUUID()
	subID := subscriptionID.String()

	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.clientMap[subID] = client

	go client.processRequest()
	go client.processEvents()

	log.Infof("registered watch client. id: %s. watchtopic: %s", subID, link)

	return subID, nil
}

// CloseWatch closes the specified watch stream by id
func (m *watchManager) CloseWatch(id string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	client, ok := m.clientMap[id]
	if !ok {
		return errors.New("invalid watch subscription ID")
	}
	log.Infof("closing watch for subscription: %s", id)
	client.cancelStream()
	delete(m.clientMap, id)
	return nil
}

// CloseConn closes watch service connection, and all open streams
func (m *watchManager) CloseConn() {
	log.Info("closing watch service connection")

	m.connection.Close()
	for id := range m.clientMap {
		delete(m.clientMap, id)
	}
}

// CloseAll closes all streams, but leaves the connection open.
func (m *watchManager) CloseAll() {
	for id := range m.clientMap {
		m.CloseWatch(id)
	}
}

// Status returns a boolean to indicate if the clients connected to central are active.
func (m *watchManager) Status() bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	ok := true

	if len(m.clientMap) == 0 {
		ok = false
	}

	for k, c := range m.clientMap {
		if c.isRunning == false {
			log.Debugf("watchmanager: watch client is not running.")
			ok = false
			delete(m.clientMap, k)
		}
	}

	return ok && m.connection.GetState() == connectivity.Ready
}
