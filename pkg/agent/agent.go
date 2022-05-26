package agent

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Axway/agent-sdk/pkg/agent/stream"

	agentcache "github.com/Axway/agent-sdk/pkg/agent/cache"
	"github.com/Axway/agent-sdk/pkg/agent/handler"
	"github.com/Axway/agent-sdk/pkg/agent/poller"
	"github.com/Axway/agent-sdk/pkg/agent/resource"
	"github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/apic"
	apiV1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/auth"
	"github.com/Axway/agent-sdk/pkg/apic/provisioning"
	"github.com/Axway/agent-sdk/pkg/cache"
	"github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/migrate"
	"github.com/Axway/agent-sdk/pkg/util"
	"github.com/Axway/agent-sdk/pkg/util/errors"
	hc "github.com/Axway/agent-sdk/pkg/util/healthcheck"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

// AgentStatus - status for Agent resource
const (
	AgentRunning   = "running"
	AgentStopped   = "stopped"
	AgentFailed    = "failed"
	AgentUnhealthy = "unhealthy"
)

// AgentResourceType - Holds the type for agent resource in Central
var AgentResourceType string

// APIValidator - Callback for validating the API
type APIValidator func(apiID, stageName string) bool

// ConfigChangeHandler - Callback for Config change event
type ConfigChangeHandler func()

type agentData struct {
	agentResourceManager resource.Manager

	apicClient       apic.Client
	cfg              config.CentralConfig
	agentFeaturesCfg config.AgentFeaturesConfig
	tokenRequester   auth.PlatformTokenGetter

	teamMap                    cache.Cache
	cacheManager               agentcache.Manager
	apiValidator               APIValidator
	configChangeHandler        ConfigChangeHandler
	agentResourceChangeHandler ConfigChangeHandler
	proxyResourceHandler       *handler.StreamWatchProxyHandler
	isInitialized              bool

	instanceValidatorJobID string
	provisioner            provisioning.Provisioning
	marketplaceMigration   migrate.Migrator
	streamer               *stream.StreamerClient
}

var agent agentData

var logger log.FieldLogger

func init() {
	logger = log.NewFieldLogger().
		WithPackage("sdk.agent").
		WithComponent("agent")
	agent.proxyResourceHandler = handler.NewStreamWatchProxyHandler()
}

// Initialize - Initializes the agent
func Initialize(centralCfg config.CentralConfig) error {
	return InitializeWithAgentFeatures(centralCfg, config.NewAgentFeaturesConfiguration())
}

// InitializeWithAgentFeatures - Initializes the agent with agent features
func InitializeWithAgentFeatures(centralCfg config.CentralConfig, agentFeaturesCfg config.AgentFeaturesConfig) error {

	if agent.teamMap == nil {
		agent.teamMap = cache.New()
	}

	err := checkRunningAgent()
	if err != nil {
		return err
	}

	err = config.ValidateConfig(agentFeaturesCfg)
	if err != nil {
		return err
	}
	agent.agentFeaturesCfg = agentFeaturesCfg
	centralCfg.SetIsMarketplaceSubsEnabled(agentFeaturesCfg.MarketplaceProvisioningEnabled())

	// validate the central config
	if agentFeaturesCfg.ConnectionToCentralEnabled() {
		err = config.ValidateConfig(centralCfg)
		if err != nil {
			return err
		}
	}

	// Only create the api map cache if it does not already exist
	if agent.cacheManager == nil {
		agent.cacheManager = agentcache.NewAgentCacheManager(centralCfg, agentFeaturesCfg.PersistCacheEnabled())
	}

	if centralCfg.GetUsageReportingConfig().IsOfflineMode() {
		// Offline mode does not need more initialization
		agent.cfg = centralCfg
		return nil
	}

	agent.cfg = centralCfg
	singleEntryFilter := []string{
		// Traceability host URL will be added by the traceability factory
		centralCfg.GetURL(),
		centralCfg.GetPlatformURL(),
		centralCfg.GetAuthConfig().GetTokenURL(),
		centralCfg.GetUsageReportingConfig().GetURL(),
	}
	api.SetConfigAgent(
		centralCfg.GetEnvironmentName(),
		centralCfg.IsUsingGRPC(),
		isRunningInDockerContainer(),
		centralCfg.GetAgentName(),
		centralCfg.GetSingleURL(),
		singleEntryFilter,
	)

	if agentFeaturesCfg.ConnectionToCentralEnabled() {
		err = initializeTokenRequester(centralCfg)
		if err != nil {
			return err
		}

		// Init apic client when the agent starts, and on config change.
		agent.apicClient = apic.New(centralCfg, agent.tokenRequester, agent.cacheManager)

		if util.IsNotTest() {
			err = initEnvResources(centralCfg, agent.apicClient)
			if err != nil {
				return err
			}
		}

		if centralCfg.GetAgentName() != "" {
			if agent.agentResourceManager == nil {
				agent.agentResourceManager, err = resource.NewAgentResourceManager(
					agent.cfg, agent.apicClient, agent.agentResourceChangeHandler,
				)
				if err != nil {
					return err
				}
			} else {
				agent.agentResourceManager.OnConfigChange(agent.cfg, agent.apicClient)
			}
		}
	}

	if !agent.isInitialized {
		setupSignalProcessor()
		// only do the periodic health check stuff if NOT in unit tests and running binary agents
		if util.IsNotTest() && !isRunningInDockerContainer() {
			hc.StartPeriodicHealthCheck()
		}

		if util.IsNotTest() && agent.agentFeaturesCfg.ConnectionToCentralEnabled() {
			StartAgentStatusUpdate()

			startTeamACLCache()

			err = registerSubscriptionWebhook(agent.cfg.GetAgentType(), agent.apicClient)
			if err != nil {
				return errors.Wrap(errors.ErrRegisterSubscriptionWebhook, err.Error())
			}

			// Set agent running
			if agent.agentResourceManager != nil {
				UpdateStatusWithPrevious(AgentRunning, "", "")
			}
		}
	}

	agent.isInitialized = true
	return nil
}

func initEnvResources(cfg config.CentralConfig, client apic.Client) error {
	env, err := client.GetEnvironment()
	if err != nil {
		return err
	}

	cfg.SetAxwayManaged(env.Spec.AxwayManaged)
	if cfg.GetEnvironmentID() == "" {
		// need to save this ID for the traceability agent for later
		cfg.SetEnvironmentID(env.Metadata.ID)
	}

	if cfg.GetTeamID() == "" {
		team, err := client.GetCentralTeamByName(cfg.GetTeamName())
		if err != nil {
			return err
		}

		cfg.SetTeamID(team.ID)
	}

	return nil
}

func checkRunningAgent() error {
	// Check only on startup of binary agents
	if !agent.isInitialized && util.IsNotTest() && !isRunningInDockerContainer() {
		return hc.CheckIsRunning()
	}
	return nil
}

// InitializeForTest - Initialize for test
func InitializeForTest(apicClient apic.Client) {
	if agent.cfg != nil {
		agent.cacheManager = agentcache.NewAgentCacheManager(agent.cfg, false)
	}
	agent.apicClient = apicClient
}

// GetConfigChangeHandler - returns registered config change handler
func GetConfigChangeHandler() ConfigChangeHandler {
	return agent.configChangeHandler
}

// OnConfigChange - Registers handler for config change event
func OnConfigChange(configChangeHandler ConfigChangeHandler) {
	agent.configChangeHandler = configChangeHandler
}

// OnAgentResourceChange - Registers handler for resource change event
func OnAgentResourceChange(agentResourceChangeHandler ConfigChangeHandler) {
	agent.agentResourceChangeHandler = agentResourceChangeHandler
}

// RegisterResourceEventHandler - Registers handler for resource events
func RegisterResourceEventHandler(name string, resourceEventHandler handler.Handler) {
	agent.proxyResourceHandler.RegisterTargetHandler(name, resourceEventHandler)
}

// UnregisterResourceEventHandler - removes the specified resource event handler
func UnregisterResourceEventHandler(name string) {
	agent.proxyResourceHandler.UnregisterTargetHandler(name)
}

// HandleFetchOnStartupResources to be called for fetch watched resource on startup, so that they are processed by handlers
// this operation is performed in a go routine
func HandleFetchOnStartupResources() {
	if agent.cfg != nil {
		if agent.cfg.IsFetchOnStartupEnabled() {
			if agent.streamer != nil {
				go agent.streamer.HandleFetchOnStartupResources()
			} else {
				log.Errorf("Handling fetch-on-startup resources will no occur as streamer is not initialized, check the logs for other errors")
			}
		} else {
			log.Warnf("Handling fetch-on-startup resources will no occur as central.grpc.fetchOnStartup.enabled=false")
		}
	} else {
		log.Warnf("Handling fetch-on-startup resources will no occur as config is not set. Test mode: %v", !util.IsNotTest())
	}
}

// SyncCache -
func SyncCache() error {
	migrations := []migrate.Migrator{
		migrate.NewAttributeMigration(agent.apicClient, agent.cfg),
	}

	if agent.agentFeaturesCfg.MarketplaceProvisioningEnabled() {
		marketplaceMigration := migrate.NewMarketplaceMigration(agent.apicClient, agent.cfg, agent.cacheManager)
		agent.marketplaceMigration = marketplaceMigration
		migrations = append(migrations, marketplaceMigration)
	}

	mig := migrate.NewMigrateAll(migrations...)
	isMpEnabled := agent.agentFeaturesCfg != nil && agent.agentFeaturesCfg.MarketplaceProvisioningEnabled()

	opts := []discoveryOpt{
		withMigration(mig),
		withMpEnabled(isMpEnabled),
	}

	if agent.agentResourceManager != nil {
		opts = append(opts, withAdditionalDiscoverFuncs(agent.agentResourceManager.FetchAgentResource))
	}

	discoveryCache := newDiscoveryCache(
		agent.cfg,
		GetCentralClient(),
		newHandlers(),
		opts...,
	)

	if !agent.cacheManager.HasLoadedPersistedCache() {
		err := discoveryCache.execute()
		if err != nil {
			return err
		}
		agent.cacheManager.SaveCache()
	}

	cacheSync := func() {
		agent.cacheManager.Flush()
		err := discoveryCache.execute()
		if err != nil {
			logger.WithError(err).Error("failed to re-sync cache after a cache flush")
		}
		agent.cacheManager.SaveCache()
	}

	return startCentralEventProcessor(cacheSync)
}

func registerSubscriptionWebhook(at config.AgentType, client apic.Client) error {
	if at == config.DiscoveryAgent {
		return client.RegisterSubscriptionWebhook()
	}
	return nil
}

func startTeamACLCache() {
	// Only discovery agents need to start the ACL handler
	if agent.cfg.GetAgentType() == config.DiscoveryAgent {
		registerAccessControlListHandler()
	}

	registerTeamMapCacheJob()
}

func isRunningInDockerContainer() bool {
	// Within the cgroup file, if you are not in a docker container all entries are like these devices:/
	// If in a docker container, entries are like this: devices:/docker/xxxxxxxxx.
	// So, all we need to do is see if ":/docker" exists somewhere in the file.
	bytes, err := ioutil.ReadFile("/proc/1/cgroup")
	if err != nil {
		return false
	}

	// Convert []byte to string and print to screen
	text := string(bytes)

	return strings.Contains(text, ":/docker")
}

// initializeTokenRequester - Create a new auth token requester
func initializeTokenRequester(centralCfg config.CentralConfig) error {
	var err error
	agent.tokenRequester = auth.NewPlatformTokenGetterWithCentralConfig(centralCfg)
	if util.IsNotTest() {
		_, err = agent.tokenRequester.GetToken()
	}
	return err
}

// GetCentralAuthToken - Returns the Auth token from AxwayID to make API call to Central
func GetCentralAuthToken() (string, error) {
	if agent.tokenRequester == nil {
		return "", apic.ErrAuthenticationCall
	}
	return agent.tokenRequester.GetToken()
}

// GetCentralClient - Returns the APIC Client
func GetCentralClient() apic.Client {
	return agent.apicClient
}

// GetCentralConfig - Returns the APIC Client
func GetCentralConfig() config.CentralConfig {
	return agent.cfg
}

// GetAPICache - Returns the cache
func GetAPICache() cache.Cache {
	if agent.cacheManager == nil {
		agent.cacheManager = agentcache.NewAgentCacheManager(agent.cfg, agent.agentFeaturesCfg.PersistCacheEnabled())
	}
	return agent.cacheManager.GetAPIServiceCache()
}

// GetCacheManager - Returns the cache
func GetCacheManager() agentcache.Manager {
	if agent.cacheManager == nil {
		agent.cacheManager = agentcache.NewAgentCacheManager(agent.cfg, agent.agentFeaturesCfg.PersistCacheEnabled())
	}
	return agent.cacheManager
}

// GetAgentResource - Returns Agent resource
func GetAgentResource() *apiV1.ResourceInstance {
	if agent.agentResourceManager == nil {
		return nil
	}
	return agent.agentResourceManager.GetAgentResource()
}

// UpdateStatus - Updates the agent state
func UpdateStatus(status, description string) {
	UpdateStatusWithPrevious(status, status, description)
}

// UpdateStatusWithPrevious - Updates the agent state providing a previous state
func UpdateStatusWithPrevious(status, prevStatus, description string) {
	if agent.agentResourceManager != nil {
		err := agent.agentResourceManager.UpdateAgentStatus(status, prevStatus, description)
		if err != nil {
			logger.Warnf("could not update the agent status reference, %s", err.Error())
		}
	}
}

func setupSignalProcessor() {
	if !agent.agentFeaturesCfg.ProcessSystemSignalsEnabled() {
		return
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-sigs
		cleanUp()
		logger.Info("Stopping agent")
		os.Exit(0)
	}()
}

// cleanUp - AgentCleanup
func cleanUp() {
	UpdateStatusWithPrevious(AgentStopped, AgentRunning, "")
}

func startCentralEventProcessor(cacheSyncFunc func()) error {
	if agent.cfg.IsUsingGRPC() {
		return startStreamMode(cacheSyncFunc)
	}
	return startPollMode(cacheSyncFunc)
}

func newHandlers() []handler.Handler {
	handlers := []handler.Handler{
		handler.NewAPISvcHandler(agent.cacheManager),
		handler.NewInstanceHandler(agent.cacheManager),
		handler.NewAgentResourceHandler(agent.agentResourceManager),
		agent.proxyResourceHandler,
	}

	if agent.cfg.GetAgentType() == config.DiscoveryAgent {
		handlers = append(
			handlers,
			handler.NewCategoryHandler(agent.cacheManager),
			handler.NewCRDHandler(agent.cacheManager),
			handler.NewARDHandler(agent.cacheManager),
			// handler.NewACLHandler(agent.cacheManager),
		)
	}

	// Register managed application and access handler for traceability agent
	// For discovery agent, the handlers gets registered while setting up provisioner
	if agent.cfg.GetAgentType() == config.TraceabilityAgent {
		handlers = append(
			handlers,
			handler.NewTraceAccessRequestHandler(agent.cacheManager, agent.apicClient),
			handler.NewTraceManagedApplicationHandler(agent.cacheManager),
		)
	}

	return handlers
}

func startPollMode(cacheSyncFunc func()) error {
	handlers := newHandlers()

	pc, err := poller.NewPollClient(
		agent.apicClient,
		agent.cfg,
		agent.tokenRequester,
		agent.cacheManager,
		func(p *poller.PollClient) {
			hc.RegisterHealthcheck(util.AmplifyCentral, "central", p.Healthcheck)
		},
		cacheSyncFunc,
		handlers...,
	)

	if err != nil {
		return fmt.Errorf("could not start the harvester poll client: %s", err)
	}

	newEventProcessorJob(pc, "Poll Client")

	return err
}

func startStreamMode(cacheSyncFunc func()) error {
	handlers := newHandlers()

	sc, err := stream.NewStreamerClient(
		agent.apicClient,
		agent.cfg,
		agent.tokenRequester,
		agent.cacheManager,
		func(s *stream.StreamerClient) {
			hc.RegisterHealthcheck(util.AmplifyCentral, "central", s.Healthcheck)
		},
		cacheSyncFunc,
		handlers...,
	)

	if err != nil {
		return fmt.Errorf("could not start the watch manager: %s", err)
	}

	agent.streamer = sc
	newEventProcessorJob(sc, "Stream Client")

	return err
}
