package migrate

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/Axway/agent-sdk/pkg/apic"
	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	mv1a "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	"github.com/Axway/agent-sdk/pkg/apic/definitions"
	"github.com/Axway/agent-sdk/pkg/apic/provisioning"
	"github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/util"
	"github.com/Axway/agent-sdk/pkg/util/log"
)

const (
	serviceName  = "service-name"
	instanceName = "instance-name"
	revisionName = "revision-name"
)

var apiserviceName string

// Migrator interface for performing a migration on a ResourceInstance
type Migrator interface {
	Migrate(ri *v1.ResourceInstance) (*v1.ResourceInstance, error)
}

type ardCache interface {
	GetCredentialRequestDefinitionByName(name string) (*v1.ResourceInstance, error)
	AddAccessRequestDefinition(resource *v1.ResourceInstance)
	GetAccessRequestDefinitionByName(name string) (*v1.ResourceInstance, error)
}

// NewMarketplaceMigration - creates a new MarketplaceMigration
func NewMarketplaceMigration(client client, cfg config.CentralConfig, cache ardCache) *MarketplaceMigration {
	logger := log.NewFieldLogger().
		WithPackage("sdk.migrate").
		WithComponent("MarketplaceMigration")

	return &MarketplaceMigration{
		logger: logger,
		client: client,
		cfg:    cfg,
		cache:  cache,
	}
}

// MarketplaceMigration - used for migrating attributes to subresource
type MarketplaceMigration struct {
	logger log.FieldLogger
	client client
	cfg    config.CentralConfig
	cache  ardCache
}

// Migrate -
func (m *MarketplaceMigration) Migrate(ri *v1.ResourceInstance) (*v1.ResourceInstance, error) {
	if ri.Kind != mv1a.APIServiceGVK().Kind {
		return ri, nil
	}

	// check resource to see if this apiservice has already been run through migration
	apiSvc, err := ri.AsInstance()
	if err != nil {
		return nil, err
	}

	// get x-agent-details and determine if we need to process this apiservice for marketplace provisioning
	details := util.GetAgentDetails(apiSvc)
	if len(details) > 0 {
		completed := details[definitions.MarketplaceMigration]
		if completed == definitions.MigrationCompleted {
			// migration ran already
			m.logger.
				WithField(serviceName, apiSvc).
				Debugf("marketplace provision migration already completed")
			return ri, nil
		}
	}

	m.logger.
		WithField(serviceName, ri.Name).
		Tracef("perform marketplace provision")

	m.UpdateService(ri)

	if err != nil {
		return ri, fmt.Errorf("migration marketplace provisioning failed: %s", err)
	}

	return ri, nil
}

// UpdateService - gets a list of instances for the service and updates their request definitions.
func (m *MarketplaceMigration) UpdateService(ri *v1.ResourceInstance) error {
	revURL := m.cfg.GetRevisionsURL()
	q := map[string]string{
		"query": queryFunc(ri.Name),
	}

	revs, err := m.client.GetAPIV1ResourceInstancesWithPageSize(q, revURL, 100)
	if err != nil {
		return err
	}

	apiserviceName = ri.Name
	m.logger.
		WithField(serviceName, apiserviceName).
		Tracef("found %d revisions for api", len(revs))

	errCh := make(chan error, len(revs))
	wg := &sync.WaitGroup{}

	// query for api service instances by reference to a revision name
	for _, rev := range revs {
		wg.Add(1)

		go func(revision *v1.ResourceInstance) {
			defer wg.Done()

			q := map[string]string{
				"query": queryFunc(revision.Name),
			}
			url := m.cfg.GetInstancesURL()

			// Passing down apiservice name (ri.Name) for logging purposes
			// Possible future refactor to send context through to get proper resources downstream
			err := m.updateSvcInstance(url, q, revision)
			errCh <- err
		}(rev)
	}

	wg.Wait()
	close(errCh)

	for e := range errCh {
		if e != nil {
			return e
		}
	}

	return nil
}

func (m *MarketplaceMigration) updateSvcInstance(
	resourceURL string, query map[string]string, revision *v1.ResourceInstance) error {
	resources, err := m.client.GetAPIV1ResourceInstancesWithPageSize(query, resourceURL, 100)
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	errCh := make(chan error, len(resources))

	for _, resource := range resources {
		wg.Add(1)

		go func(svcInstance *v1.ResourceInstance) {
			defer wg.Done()

			err := m.handleSvcInstance(svcInstance, revision)
			if err != nil {
				errCh <- err
			}

		}(resource)
	}

	wg.Wait()
	close(errCh)

	for e := range errCh {
		if e != nil {
			return e
		}
	}

	return nil
}

func (m *MarketplaceMigration) processAccessRequestDefinition(oauthScopes map[string]string) (string, error) {
	ard, err := m.registerAccessRequestDefinition(oauthScopes)
	if err != nil {
		return "", err
	}

	newARD, err := m.client.CreateOrUpdateResource(ard)
	if err != nil {
		return "", err
	}
	var ardRIName string
	if newARD != nil {
		ard, err := newARD.AsInstance()
		if err == nil {
			m.cache.AddAccessRequestDefinition(ard)
		} else {
			return "", err
		}
		ardRIName = ard.Name
	}
	return ardRIName, nil
}

func getCredentialRequestPolicies(authPolicies []string) ([]string, error) {
	var credentialRequestPolicies []string

	for _, policy := range authPolicies {
		if policy == apic.Apikey {
			credentialRequestPolicies = append(credentialRequestPolicies, provisioning.APIKeyCRD)
		}
		if policy == apic.Oauth {
			credentialRequestPolicies = append(credentialRequestPolicies, []string{provisioning.OAuthPublicKeyCRD, provisioning.OAuthSecretCRD}...)
		}
	}

	return credentialRequestPolicies, nil

}

func (m *MarketplaceMigration) checkCredentialRequestDefinitions(credentialRequestPolicies []string) []string {
	// remove any crd not in the cache
	knownCRDs := make([]string, 0)
	for _, policy := range credentialRequestPolicies {
		if def, err := m.cache.GetCredentialRequestDefinitionByName(policy); err == nil && def != nil {
			knownCRDs = append(knownCRDs, policy)
		}
	}

	return knownCRDs
}

func (m *MarketplaceMigration) registerAccessRequestDefinition(scopes map[string]string) (*mv1a.AccessRequestDefinition, error) {

	callback := func(ard *mv1a.AccessRequestDefinition) (*mv1a.AccessRequestDefinition, error) {
		return ard, nil
	}

	var ard *mv1a.AccessRequestDefinition
	var err error
	if len(scopes) > 0 {
		ard, err = provisioning.NewAccessRequestBuilder(callback).Register()
		if err != nil {
			return nil, err
		}
	}
	return ard, nil
}

// updateRI updates the resource, and the sub resource
func (m *MarketplaceMigration) updateRI(ri *v1.ResourceInstance) error {
	_, err := m.client.UpdateResourceInstance(ri)
	if err != nil {
		return err
	}

	return nil
}

func (m *MarketplaceMigration) handleSvcInstance(
	svcInstance *v1.ResourceInstance, revision *v1.ResourceInstance) error {
	logger := m.logger.
		WithField(serviceName, apiserviceName).
		WithField(instanceName, svcInstance.Name).
		WithField(revisionName, revision.Name)

	apiSvcInst := mv1a.NewAPIServiceInstance(svcInstance.Name, svcInstance.Metadata.Scope.Name)
	apiSvcInst.FromInstance(svcInstance)

	specProcessor, err := getSpecParser(revision)
	if err != nil {
		return err
	}

	var i interface{} = specProcessor

	if processor, ok := i.(apic.OasSpecProcessor); ok {
		ardRIName := apiSvcInst.Spec.AccessRequestDefinition
		credentialRequestPolicies := apiSvcInst.Spec.CredentialRequestDefinitions

		processor.ParseAuthInfo()

		// get the auth policy from the spec
		authPolicies := processor.GetAuthPolicies()

		// get the apikey info
		apiKeyInfo := processor.GetAPIKeyInfo()
		if len(apiKeyInfo) > 0 {
			logger.Trace("instance has a spec definition type of apiKey")
			ardRIName = provisioning.APIKeyARD
		}

		// get oauth scopes
		oauthScopes := processor.GetOAuthScopes()

		var updateRequestDefinition = false

		// Check if ARD exists
		if apiSvcInst.Spec.AccessRequestDefinition == "" && len(oauthScopes) > 0 {
			// Only migrate resource with oauth scopes. Spec with type apiKey will be handled on startup
			logger.Trace("instance has a spec definition type of oauth")
			ardRIName, err = m.processAccessRequestDefinition(oauthScopes)
			if err != nil {
				return err
			}
		}

		// Check if CRD exists
		credentialRequestPolicies, err = getCredentialRequestPolicies(authPolicies)
		if err != nil {
			return err
		}

		// Find only the known CRDs
		credentialRequestDefinitions := m.checkCredentialRequestDefinitions(credentialRequestPolicies)
		if len(credentialRequestDefinitions) > 0 && !sortCompare(apiSvcInst.Spec.CredentialRequestDefinitions, credentialRequestDefinitions) {
			logger.Debugf("adding the following credential request definitions %s to apiserviceinstance %s", credentialRequestDefinitions, apiSvcInst.Name)
			updateRequestDefinition = true
		}

		existingARD, _ := m.cache.GetAccessRequestDefinitionByName(ardRIName)
		if existingARD == nil {
			ardRIName = ""
		} else {
			if apiSvcInst.Spec.AccessRequestDefinition == "" {
				logger.Debugf("adding the following access request definition %s to apiserviceinstance %s", ardRIName, apiSvcInst.Name)
				updateRequestDefinition = true
			}
		}

		// update apiserivceinstane spec with necessary request definitions
		if updateRequestDefinition {
			inInterface := m.newInstanceSpec(apiSvcInst.Spec.Endpoint, revision.Name, ardRIName, credentialRequestDefinitions)
			svcInstance.Spec = inInterface

			err = m.updateRI(svcInstance)
			if err != nil {
				return err
			}

			logger.Debugf("migrated instance %s with the necessary request definitions", apiSvcInst.Name)
		}
	}

	return nil
}

func (m *MarketplaceMigration) newInstanceSpec(
	endpoints []mv1a.ApiServiceInstanceSpecEndpoint,
	revisionName,
	ardRIName string,
	credentialRequestDefinitions []string,
) map[string]interface{} {
	newSpec := mv1a.ApiServiceInstanceSpec{
		Endpoint:                     endpoints,
		ApiServiceRevision:           revisionName,
		CredentialRequestDefinitions: credentialRequestDefinitions,
		AccessRequestDefinition:      ardRIName,
	}
	// convert to set ri.Spec
	var inInterface map[string]interface{}
	in, _ := json.Marshal(newSpec)
	json.Unmarshal(in, &inInterface)
	return inInterface
}

func sortCompare(apiSvcInstCRDs, knownCRDs []string) bool {
	if len(apiSvcInstCRDs) != len(knownCRDs) {
		return false
	}

	sort.Strings(apiSvcInstCRDs)
	sort.Strings(knownCRDs)

	for i, v := range apiSvcInstCRDs {
		if v != knownCRDs[i] {
			return false
		}
	}
	return true
}

func getSpecParser(revision *v1.ResourceInstance) (apic.SpecProcessor, error) {
	specDefinitionType, ok := revision.Spec["definition"].(map[string]interface{})["type"].(string)
	if !ok {
		return nil, fmt.Errorf("could not get the spec definition type from apiservicerevision %s", revision.Name)
	}

	specDefinitionValue, ok := revision.Spec["definition"].(map[string]interface{})["value"].(string)
	if !ok {
		return nil, fmt.Errorf("could not get the spec definition value from apiservicerevision %s", revision.Name)
	}

	specDefinition, _ := base64.StdEncoding.DecodeString(specDefinitionValue)

	specParser := apic.NewSpecResourceParser(specDefinition, specDefinitionType)
	err := specParser.Parse()
	if err != nil {
		return nil, err
	}

	err = specParser.Parse()
	if err != nil {
		return nil, err
	}

	specProcessor := specParser.GetSpecProcessor()
	return specProcessor, nil
}
