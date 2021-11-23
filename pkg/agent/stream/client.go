package stream

import (
	"github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/apic/auth"
	hc "github.com/Axway/agent-sdk/pkg/util/healthcheck"
	"github.com/Axway/agent-sdk/pkg/util/log"
	wm "github.com/Axway/agent-sdk/pkg/watchmanager"
	"github.com/Axway/agent-sdk/pkg/watchmanager/proto"
)

// Client a client for opening up a grpc stream, and handling the received events on the stream.
type Client struct {
	apiClient       api.Client
	apisHost        string
	handlers        []Handler
	manager         wm.Manager
	newEventManager eventManagerFunc
	tenantID        string
	tokenGetter     auth.TokenGetter
	topic           string
}

// NewClient creates a Client
func NewClient(
	host string,
	tenantID string,
	topic string,
	tokenGetter auth.TokenGetter,
	apiClient api.Client,
	manager wm.Manager,
	handlers ...Handler,
) *Client {
	return &Client{
		apiClient:       apiClient,
		handlers:        handlers,
		apisHost:        host,
		newEventManager: NewEventListener,
		tenantID:        tenantID,
		tokenGetter:     tokenGetter,
		topic:           topic,
		manager:         manager,
	}
}

func (sc *Client) newStreamService() error {
	ric := newResourceClient(sc.apisHost, sc.tenantID, sc.apiClient, sc.tokenGetter)

	events, errors := make(chan *proto.Event), make(chan error)

	em := sc.newEventManager(
		events,
		ric,
		sc.handlers...,
	)

	id, err := sc.manager.RegisterWatch(sc.topic, events, errors)
	if err != nil {
		return err
	}

	log.Debugf("watch-controller subscription-id: %s", id)

	return em.Listen()
}

// Start starts the streaming client
func (sc *Client) Start() error {
	return sc.newStreamService()
}

// HealthCheck wraps a Watch Manager to provide a health check endpoint on the connection to central.
func HealthCheck(manager wm.Manager) hc.CheckStatus {
	return func(_ string) *hc.Status {
		ok := manager.Status()
		status := &hc.Status{
			Result: hc.OK,
		}

		if !ok {
			status.Result = hc.FAIL
			status.Details = "the stream to central is not open"
		}
		log.Infof("Stream status: %s", status.Result)
		return status
	}
}
