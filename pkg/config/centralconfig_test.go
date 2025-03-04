package config

import (
	"github.com/Axway/agent-sdk/pkg/util/exception"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDiscoveryAgentConfig(t *testing.T) {
	cfg := NewCentralConfig(DiscoveryAgent)
	centralConfig := cfg.(*CentralConfiguration)

	// Setup Auth config to ignore auth validation errors for this test
	authCfg := centralConfig.Auth.(*AuthConfiguration)
	authCfg.URL = "test"
	authCfg.Realm = "Broker"
	authCfg.ClientID = "aaaa"
	tmpFile, _ := ioutil.TempFile(".", "test*")
	authCfg.PrivateKey = "./" + tmpFile.Name()
	authCfg.PublicKey = "./" + tmpFile.Name()

	cfgValidator, ok := cfg.(IConfigValidator)
	assert.True(t, ok)
	assert.NotNil(t, cfgValidator)

	err := cfgValidator.ValidateCfg()

	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.organizationID, please set and/or check its value", err.Error())

	centralConfig.TenantID = "1111"
	err = cfgValidator.ValidateCfg()

	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.url, please set and/or check its value", err.Error())

	centralConfig.URL = "aaa"
	err = cfgValidator.ValidateCfg()

	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.url, please set and/or check its value", err.Error())

	centralConfig.URL = "http://localhost:8080"
	err = cfgValidator.ValidateCfg()

	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.environment, please set and/or check its value", err.Error())

	centralConfig.Environment = "eee"
	err = cfgValidator.ValidateCfg()
	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.auth.url, please set and/or check its value", err.Error())
	authCfg.URL = "http://localhost.com:8080"

	centralConfig.SingleURL = "malformed.singleURL.com"
	err = cfgValidator.ValidateCfg()
	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.singleURL, please set and/or check its value", err.Error())
	centralConfig.SingleURL = "https://ingestion.platform.axway.com"

	centralConfig.APIServerVersion = ""
	err = cfgValidator.ValidateCfg()
	assert.NotNil(t, err)
	// assert.Equal(t, "[Error Code 1401] - error with config central.apiServerVersion, please set and/or check its value", err.Error())

	centralConfig.APIServerVersion = "v1alpha1"

	assert.Equal(t, centralConfig.URL+"/api/unifiedCatalog/v1/catalogItems", cfg.GetCatalogItemsURL())
	assert.Equal(t, centralConfig.URL+"/apis/management/v1alpha1/environments/eee/apiservices", cfg.GetServicesURL())

	centralConfig.ReportActivityFrequency = 0
	err = cfgValidator.ValidateCfg()
	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.reportActivityFrequency, please set and/or check its value", err.Error())

	centralConfig.PollInterval = 0
	err = cfgValidator.ValidateCfg()
	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.pollInterval, please set and/or check its value", err.Error())
	centralConfig.PollInterval = 30

	err = cfgValidator.ValidateCfg()
	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.reportActivityFrequency, please set and/or check its value", err.Error())
	centralConfig.ReportActivityFrequency = time.Minute

	// validate mp and DOSA
	authCfg.ClientID = "DOSA_aaaa"
	centralConfig.isMarketplaceSubs = true
	err = cfgValidator.ValidateCfg()
	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1406] - using a DOSA_* service account with Marketplace Provisioning is unsupported", err.Error())

	cleanupFiles(tmpFile.Name())
}

func TestTraceabilityAgentConfig(t *testing.T) {
	cfg := NewCentralConfig(TraceabilityAgent)
	centralConfig := cfg.(*CentralConfiguration)

	// Setup Auth config to ignore auth validation errors for this test
	authCfg := centralConfig.Auth.(*AuthConfiguration)
	authCfg.URL = "test"
	authCfg.Realm = "Broker"
	authCfg.ClientID = "aaaa"
	tmpFile, _ := ioutil.TempFile(".", "test*")
	authCfg.PrivateKey = "./" + tmpFile.Name()
	authCfg.PublicKey = "./" + tmpFile.Name()

	cfgValidator, ok := cfg.(IConfigValidator)
	assert.True(t, ok)
	assert.NotNil(t, cfgValidator)
	err := cfgValidator.ValidateCfg()

	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.organizationID, please set and/or check its value", err.Error())

	centralConfig.TenantID = "1111"
	err = cfgValidator.ValidateCfg()

	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.url, please set and/or check its value", err.Error())

	centralConfig.URL = "aaa"
	err = cfgValidator.ValidateCfg()

	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.url, please set and/or check its value", err.Error())

	centralConfig.URL = "http://localhost.com"
	err = cfgValidator.ValidateCfg()

	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.deployment, please set and/or check its value", err.Error())

	centralConfig.APICDeployment = "aaa"
	err = cfgValidator.ValidateCfg()

	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.environment, please set and/or check its value", err.Error())

	centralConfig.Environment = "111111"
	err = cfgValidator.ValidateCfg()
	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.auth.url, please set and/or check its value", err.Error())

	authCfg.URL = "http://localhost.com:8080"

	centralConfig.SingleURL = "malformed.singleURL.com"
	err = cfgValidator.ValidateCfg()
	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.singleURL, please set and/or check its value", err.Error())
	centralConfig.SingleURL = "https://ingestion.platform.axway.com"

	centralConfig.ReportActivityFrequency = 0
	err = cfgValidator.ValidateCfg()
	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.reportActivityFrequency, please set and/or check its value", err.Error())

	cleanupFiles(tmpFile.Name())
}

func TestTraceabilityAgentOfflineConfig(t *testing.T) {
	cfg := NewCentralConfig(TraceabilityAgent)
	centralConfig := cfg.(*CentralConfiguration)

	// Set to offline mode
	centralConfig.UsageReporting.(*UsageReportingConfiguration).Offline = true

	cfgValidator, ok := cfg.(IConfigValidator)
	assert.True(t, ok)
	assert.NotNil(t, cfgValidator)
	err := cfgValidator.ValidateCfg()

	// Environment ID is the only config needed in offline mode
	assert.NotNil(t, err)
	assert.Equal(t, "[Error Code 1401] - error with config central.environmentID, please set and/or check its value", err.Error())
	centralConfig.EnvironmentID = "abc123"
	err = cfgValidator.ValidateCfg()
	assert.Nil(t, err)
}

func TestTeamConfig(t *testing.T) {
	cfg := NewCentralConfig(TraceabilityAgent)
	centralConfig := cfg.(*CentralConfiguration)

	// Setup Auth config to ignore auth validation errors for this test
	authCfg := centralConfig.Auth.(*AuthConfiguration)
	authCfg.URL = "test"
	authCfg.Realm = "Broker"
	authCfg.ClientID = "aaaa"
	tmpFile, _ := ioutil.TempFile(".", "test*")
	authCfg.PrivateKey = "./" + tmpFile.Name()
	authCfg.PublicKey = "./" + tmpFile.Name()
	centralConfig.TenantID = "1111"
	centralConfig.URL = "aaa"
	centralConfig.Environment = "111111"
	centralConfig.APICDeployment = "aaa"

	// Should be nil, not yet set
	assert.Equal(t, "", centralConfig.GetTeamID(), "Team ID was expected to be blank as it has not yet been set")

	//Set it and validate
	teamID := "abc12:34567:def89:12345:67890"
	centralConfig.SetTeamID(teamID)
	assert.Equal(t, teamID, centralConfig.GetTeamID(), "The Team ID was not set appropriately")

	cleanupFiles(tmpFile.Name())
}

func TestGRPCConfig(t *testing.T) {
	cfg := NewCentralConfig(GovernanceAgent)
	centralConfig := cfg.(*CentralConfiguration)
	centralConfig.GRPCCfg.FetchOnStartup.Enabled = true

	// by default it should work
	exception.Block{
		Try: func() {
			centralConfig.validateGRPC()
		}, Catch: func(err error) {
			assert.NoError(t, err)
		},
	}.Do()

	centralConfig.GRPCCfg.FetchOnStartup.PageSize = 0
	assertError(t, centralConfig)

	centralConfig.GRPCCfg.FetchOnStartup.PageSize = 1
	centralConfig.GRPCCfg.FetchOnStartup.Retention = 0
	assertError(t, centralConfig)
}

func assertError(t *testing.T, centralConfig *CentralConfiguration) {
	exception.Block{
		Try: func() {
			centralConfig.validateGRPC()
		}, Catch: func(err error) {
			assert.Error(t, err)
		},
	}.Do()
}

func cleanupFiles(fileName string) {
	// cleanup files
	os.Remove("./" + fileName)
}
