package agent

import (
	"sync"

	"github.com/Axway/agent-sdk/pkg/util"

	defs "github.com/Axway/agent-sdk/pkg/apic/definitions"

	apiV1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/jobs"
	utilErrors "github.com/Axway/agent-sdk/pkg/util/errors"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

type instanceValidator struct {
	jobs.Job
	cacheLock       *sync.Mutex
	isAgentPollMode bool
}

func newInstanceValidator(cacheLock *sync.Mutex, isAgentPollMode bool) *instanceValidator {
	return &instanceValidator{cacheLock: cacheLock, isAgentPollMode: isAgentPollMode}
}

// Ready -
func (j *instanceValidator) Ready() bool {
	return true
}

// Status -
func (j *instanceValidator) Status() error {
	return nil
}

// Execute -
func (j *instanceValidator) Execute() error {
	j.validateAPIOnDataplane()
	return nil
}

func (j *instanceValidator) validateAPIOnDataplane() {
	j.cacheLock.Lock()
	defer j.cacheLock.Unlock()

	log.Info("validating api service instance on dataplane")
	// Validate the API on dataplane.  If API is not valid, mark the consumer instance as "DELETED"
	for _, key := range agent.cacheManager.GetAPIServiceInstanceKeys() {
		instance, err := agent.cacheManager.GetAPIServiceInstanceByID(key)
		if err != nil {
			continue
		}

		externalAPIID, _ := util.GetAgentDetailsValue(instance, defs.AttrExternalAPIID)
		if externalAPIID == "" {
			continue // skip service instances without external api id
		}
		externalAPIStage, _ := util.GetAgentDetailsValue(instance, defs.AttrExternalAPIStage)
		externalPrimaryKey, _ := util.GetAgentDetailsValue(instance, defs.AttrExternalAPIPrimaryKey)
		// Check if the consumer instance was published by agent, i.e. following attributes are set
		// - externalAPIID should not be empty
		// - externalAPIStage could be empty for dataplanes that do not support it
		if externalAPIID != "" && !agent.apiValidator(externalAPIID, externalAPIStage) {
			j.deleteServiceInstanceOrService(instance, externalPrimaryKey, externalAPIID)
		}
	}
}

func (j *instanceValidator) shouldDeleteService(primaryKey, apiID string) bool {
	count := 0
	if primaryKey != "" {
		count = j.getServiceInstanceCount(defs.AttrExternalAPIPrimaryKey, primaryKey)
		log.Tracef("Query instances with externalPrimaryKey attribute : %s", primaryKey)
	} else {
		count = j.getServiceInstanceCount(defs.AttrExternalAPIID, apiID)
		log.Tracef("Query instances with externalAPIID attribute : %s", apiID)
	}

	log.Tracef("Instances count : %d", count)

	return count <= 1
}

func (j *instanceValidator) getServiceInstanceCount(attName, attValue string) int {
	count := 0
	for _, key := range agent.cacheManager.GetAPIServiceInstanceKeys() {
		instance, _ := agent.cacheManager.GetAPIServiceInstanceByID(key)
		if instance != nil {
			v, _ := util.GetAgentDetailsValue(instance, attName)
			if attValue == v {
				count++
			}
		}
	}
	return count
}

func (j *instanceValidator) deleteServiceInstanceOrService(ri *apiV1.ResourceInstance, primaryKey, apiID string) {
	msg := ""
	var err error
	var agentError *utilErrors.AgentError

	defer func() {
		// remove the api service instance from the cache for both scenarios
		if j.isAgentPollMode {
			// In GRPC mode delete is done on receiving delete event from service
			agent.cacheManager.DeleteAPIServiceInstance(ri.Metadata.ID)
		}
	}()

	// delete if it is an api service
	if j.shouldDeleteService(primaryKey, apiID) {
		log.Infof("API no longer exists on the dataplane; deleting the API Service and corresponding catalog item %s", ri.Title)
		agentError = ErrDeletingService
		msg = "Deleted API Service for catalog item %s from Amplify Central"

		svc := agent.cacheManager.GetAPIServiceWithAPIID(apiID)
		if svc == nil {
			log.Errorf("api service %s not found in cache. unable to delete it from central", apiID)
			return
		}

		// deleting the service will delete all associated resources, including the consumerInstance
		err = agent.apicClient.DeleteServiceByName(svc.Name)
		if j.isAgentPollMode {
			agent.cacheManager.DeleteAPIService(apiID)
		}
	} else {
		// delete if it is an api service instance
		log.Infof("API no longer exists on the dataplane, deleting the catalog item %s", ri.Title)
		agentError = ErrDeletingCatalogItem
		msg = "Deleted catalog item %s from Amplify Central"

		err = agent.apicClient.DeleteAPIServiceInstance(ri.Name)
	}

	if err != nil {
		log.Error(utilErrors.Wrap(agentError, err.Error()).FormatError(ri.Title))
		return
	}

	log.Debugf(msg, ri.Title)
}
