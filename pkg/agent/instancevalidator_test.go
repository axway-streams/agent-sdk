package agent

import (
	"net/http"
	"testing"

	"github.com/Axway/agent-sdk/pkg/apic/definitions"

	agentcache "github.com/Axway/agent-sdk/pkg/agent/cache"
	"github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/apic"
	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/config"
	"github.com/stretchr/testify/assert"
)

func setupCache(externalAPIID, externalAPIName string) (*v1.ResourceInstance, *v1.ResourceInstance) {
	svc := &v1.ResourceInstance{
		ResourceMeta: v1.ResourceMeta{
			Metadata: v1.Metadata{
				ID: "svc-" + externalAPIID,
			},
			SubResources: map[string]interface{}{
				definitions.XAgentDetails: map[string]interface{}{
					definitions.AttrExternalAPIID:         externalAPIID,
					definitions.AttrExternalAPIPrimaryKey: "primary-" + externalAPIID,
					definitions.AttrExternalAPIName:       externalAPIName,
				},
			},
		},
	}
	instance := &v1.ResourceInstance{
		ResourceMeta: v1.ResourceMeta{
			Metadata: v1.Metadata{
				ID: "instance-" + externalAPIID,
			},
			SubResources: map[string]interface{}{
				definitions.XAgentDetails: map[string]interface{}{
					definitions.AttrExternalAPIID:         externalAPIID,
					definitions.AttrExternalAPIPrimaryKey: "primary-" + externalAPIID,
					definitions.AttrExternalAPIName:       externalAPIName,
				},
			},
		},
	}

	agent.cacheManager = agentcache.NewAgentCacheManager(&config.CentralConfiguration{}, false)
	agent.cacheManager.AddAPIServiceInstance(instance)
	agent.cacheManager.AddAPIService(svc)
	return svc, instance
}

func setupAPICClient(mockResponse []api.MockResponse) {
	client, httpClient := apic.GetTestServiceClient()
	httpClient.SetResponses(mockResponse)
	agent.apicClient = client
}

func setupAPIValidator(apiValidation bool) {
	agent.apiValidator = func(apiID, stageName string) bool {
		return apiValidation
	}
}

func TestValidatorAPIExistsOnDataplane(t *testing.T) {
	// Setup
	instanceValidator := newInstanceValidator()
	setupCache("12345", "test")
	setupAPIValidator(true)
	instanceValidator.Execute()
	i, err := agent.cacheManager.GetAPIServiceInstanceByID("instance-12345")
	assert.Nil(t, err)
	assert.NotNil(t, i)

	s := agent.cacheManager.GetAPIServiceWithPrimaryKey("primary-12345")
	assert.NotNil(t, s)
}

func TestValidatorAPIDoesExistsDeleteService(t *testing.T) {
	// Setup
	instanceValidator := newInstanceValidator()
	setupCache("12345", "test")
	setupAPICClient([]api.MockResponse{
		{
			RespCode: http.StatusOK, // get instance finalizers
		},
		{
			RespCode: http.StatusNoContent, // delete instance
		},
		{
			RespCode: http.StatusNoContent, // delete service
		},
	})
	setupAPIValidator(false)
	instanceValidator.Execute()
	i, err := agent.cacheManager.GetAPIServiceInstanceByID("instance-12345")
	assert.NotNil(t, err)
	assert.Nil(t, i)

	s := agent.cacheManager.GetAPIServiceWithPrimaryKey("primary-12345")
	assert.Nil(t, s)
}

func TestValidatorAPIDoesExistsDeleteInstance(t *testing.T) {
	instanceValidator := newInstanceValidator()

	setupCache("12345", "test")
	instance := &v1.ResourceInstance{
		ResourceMeta: v1.ResourceMeta{
			Metadata: v1.Metadata{
				ID: "instance-" + "123456",
			},
			SubResources: map[string]interface{}{
				definitions.XAgentDetails: map[string]interface{}{
					definitions.AttrExternalAPIID:         "123456",
					definitions.AttrExternalAPIPrimaryKey: "primary-12345",
					definitions.AttrExternalAPIName:       "test",
				},
			},
		},
	}
	agent.cacheManager.AddAPIServiceInstance(instance)
	setupAPICClient([]api.MockResponse{
		{
			RespCode: http.StatusOK, // get finalizers
		},
		{
			RespCode: http.StatusNoContent, // delete instance
		},
	})
	agent.apiValidator = func(apiID, stageName string) bool {
		return apiID != "12345"
	}
	instanceValidator.Execute()
	i, err := agent.cacheManager.GetAPIServiceInstanceByID("instance-12345")
	assert.NotNil(t, err)
	assert.Nil(t, i)

	s := agent.cacheManager.GetAPIServiceWithPrimaryKey("primary-12345")
	assert.NotNil(t, s)
}
