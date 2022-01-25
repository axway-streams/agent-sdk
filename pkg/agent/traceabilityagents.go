package agent

import (
	"fmt"
	"time"

	"github.com/Axway/agent-sdk/pkg/apic"
	apiV1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	"github.com/Axway/agent-sdk/pkg/config"
	utilErrors "github.com/Axway/agent-sdk/pkg/util/errors"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

func traceabilityAgent(res *apiV1.ResourceInstance) *v1alpha1.TraceabilityAgent {
	agentRes := &v1alpha1.TraceabilityAgent{}
	agentRes.FromInstance(res)

	return agentRes
}

func createTraceabilityAgentStatusResource(status, prevStatus, message string) *v1alpha1.TraceabilityAgent {
	agentRes := v1alpha1.TraceabilityAgent{}
	agentRes.Name = agent.cfg.GetAgentName()
	agentRes.Status.Version = config.AgentVersion
	agentRes.Status.LatestAvailableVersion = config.AgentLatestVersion
	agentRes.Status.State = status
	agentRes.Status.PreviousState = prevStatus
	agentRes.Status.Message = message
	agentRes.Status.LastActivityTime = getTimestamp()
	agentRes.Status.SdkVersion = config.SDKVersion

	return &agentRes
}

func mergeTraceabilityAgentWithConfig(cfg *config.CentralConfiguration) {
	ta := traceabilityAgent(GetAgentResource())
	resCfgTeamName := ta.Spec.Config.OwningTeam
	resCfgLogLevel := ta.Spec.Logging.Level
	applyResConfigToCentralConfig(cfg, "", resCfgTeamName, resCfgLogLevel)
}

func (j *discoveryCache) updateAccessRequestCache() {
	log.Trace("updating AccessRequest cache")

	// Update cache with published resources
	existingApps := make(map[string]bool)
	query := map[string]string{}

	if !j.lastAccessRequestTime.IsZero() && !j.refreshAll {
		query[apic.QueryKey] = fmt.Sprintf("state.name==%s", string(apic.AccessRequestProvisioned))
	}
	accessRequestInstances, _ := GetCentralClient().GetAPIV1ResourceInstancesWithPageSize(query, agent.cfg.GetSubscriptionURL(), apiServerPageSize)
	accessRequests, err := v1alpha1.AccessRequestFromInstanceArray(accessRequestInstances)
	if err != nil {
		log.Error(utilErrors.Wrap(ErrUnableToGetAPIV1Resources, err.Error()).FormatError("AccessRequests"))
		return
	}

	for _, accessReq := range accessRequests {
		// Update the lastAccessRequestTime based on the newest access request found
		thisTime := time.Time(accessReq.Metadata.Audit.CreateTimestamp)
		if j.lastAccessRequestTime.Before(thisTime) {
			j.lastAccessRequestTime = thisTime
		}

		externalAppID := addAccessRequestToAPICache(*accessReq)
		existingApps[externalAppID] = true
	}

	if j.refreshAll {
		// Remove items that are not published as Resources
		cacheKeys := agent.accessRequestMap.GetKeys()
		for _, key := range cacheKeys {
			if _, ok := existingApps[key]; !ok {
				agent.accessRequestMap.Delete(key)
			}
		}
	}
}

func addAccessRequestToAPICache(accessReq v1alpha1.AccessRequest) string {
	var (
		externalAppID   string
		externalAppName string
	)
	appID, ok := accessReq.Spec.Data[apic.DataplaneAppID]
	externalAppID = appID.(string)
	if ok {
		appName := accessReq.Spec.Data[apic.DataplaneAppName]
		externalAppName = appName.(string)
		agent.accessRequestMap.SetWithSecondaryKey(externalAppID, externalAppName, accessReq)
		log.Tracef("added app name: %s, id %s to App cache", externalAppName, externalAppID)
	}
	return externalAppID
}

// GetAccessRequestByAppName - finds the api by the ID from cache or API Server query
func GetAccessRequestByAppName(externalAppName string) v1alpha1.AccessRequest {
	accessRequest := v1alpha1.AccessRequest{}
	if agent.accessRequestMap != nil {
		cachedAccessRequest, err := agent.accessRequestMap.GetBySecondaryKey(externalAppName) // try to get the AccessRequest by a secondary key, App Name
		if err == nil && cachedAccessRequest != nil {
			accessRequest = cachedAccessRequest.(v1alpha1.AccessRequest)
		}
	}
	return accessRequest
}

// GetAccessRequestByAppID - finds the api by the ID from cache or API Server query
func GetAccessRequestByAppID(externalAppID string) v1alpha1.AccessRequest {
	accessRequest := v1alpha1.AccessRequest{}
	if agent.accessRequestMap != nil {
		cachedAccessRequest, err := agent.accessRequestMap.Get(externalAppID) // try to get the AccessRequest by primary key, App ID
		if err == nil && cachedAccessRequest != nil {
			accessRequest = cachedAccessRequest.(v1alpha1.AccessRequest)
		}
	}
	return accessRequest
}
