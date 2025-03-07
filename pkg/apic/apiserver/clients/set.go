/*
 * This file is automatically generated
 */

package clients

import (
	"fmt"

	cAPIV1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	catalog_v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1"
	catalog_v1alpha1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1"
	definitions_v1alpha1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/definitions/v1alpha1"
	management_v1alpha1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1"
)

type Set struct {
	WatchTopicManagementV1alpha1                     *management_v1alpha1.WatchTopicClient
	DiscoveryAgentManagementV1alpha1                 *management_v1alpha1.UnscopedDiscoveryAgentClient
	TraceabilityAgentManagementV1alpha1              *management_v1alpha1.UnscopedTraceabilityAgentClient
	GovernanceAgentManagementV1alpha1                *management_v1alpha1.UnscopedGovernanceAgentClient
	EnvironmentManagementV1alpha1                    *management_v1alpha1.EnvironmentClient
	APIServiceManagementV1alpha1                     *management_v1alpha1.UnscopedAPIServiceClient
	APIServiceRevisionManagementV1alpha1             *management_v1alpha1.UnscopedAPIServiceRevisionClient
	APIServiceInstanceManagementV1alpha1             *management_v1alpha1.UnscopedAPIServiceInstanceClient
	ConsumerInstanceManagementV1alpha1               *management_v1alpha1.UnscopedConsumerInstanceClient
	ConsumerSubscriptionDefinitionManagementV1alpha1 *management_v1alpha1.UnscopedConsumerSubscriptionDefinitionClient
	IntegrationManagementV1alpha1                    *management_v1alpha1.IntegrationClient
	ResourceHookManagementV1alpha1                   *management_v1alpha1.UnscopedResourceHookClient
	K8SClusterManagementV1alpha1                     *management_v1alpha1.K8SClusterClient
	K8SResourceManagementV1alpha1                    *management_v1alpha1.UnscopedK8SResourceClient
	ResourceDiscoveryManagementV1alpha1              *management_v1alpha1.UnscopedResourceDiscoveryClient
	MeshManagementV1alpha1                           *management_v1alpha1.MeshClient
	SpecDiscoveryManagementV1alpha1                  *management_v1alpha1.UnscopedSpecDiscoveryClient
	APISpecManagementV1alpha1                        *management_v1alpha1.UnscopedAPISpecClient
	MeshWorkloadManagementV1alpha1                   *management_v1alpha1.UnscopedMeshWorkloadClient
	MeshServiceManagementV1alpha1                    *management_v1alpha1.UnscopedMeshServiceClient
	MeshDiscoveryManagementV1alpha1                  *management_v1alpha1.UnscopedMeshDiscoveryClient
	AssetMappingTemplateManagementV1alpha1           *management_v1alpha1.UnscopedAssetMappingTemplateClient
	AccessRequestDefinitionManagementV1alpha1        *management_v1alpha1.UnscopedAccessRequestDefinitionClient
	AccessRequestManagementV1alpha1                  *management_v1alpha1.UnscopedAccessRequestClient
	DeploymentManagementV1alpha1                     *management_v1alpha1.UnscopedDeploymentClient
	AmplifyConfigManagementV1alpha1                  *management_v1alpha1.UnscopedAmplifyConfigClient
	AmplifyRuntimeConfigManagementV1alpha1           *management_v1alpha1.UnscopedAmplifyRuntimeConfigClient
	VirtualHostManagementV1alpha1                    *management_v1alpha1.UnscopedVirtualHostClient
	ManagedApplicationManagementV1alpha1             *management_v1alpha1.UnscopedManagedApplicationClient
	CredentialManagementV1alpha1                     *management_v1alpha1.UnscopedCredentialClient
	VirtualAPIManagementV1alpha1                     *management_v1alpha1.VirtualAPIClient
	VirtualAPIReleaseManagementV1alpha1              *management_v1alpha1.VirtualAPIReleaseClient
	CorsRuleManagementV1alpha1                       *management_v1alpha1.UnscopedCorsRuleClient
	AmplifyJWTAuthRuleManagementV1alpha1             *management_v1alpha1.UnscopedAmplifyJWTAuthRuleClient
	AmplifyOAuth2AuthRuleManagementV1alpha1          *management_v1alpha1.UnscopedAmplifyOAuth2AuthRuleClient
	APIKeyAuthRuleManagementV1alpha1                 *management_v1alpha1.UnscopedAPIKeyAuthRuleClient
	AssetMappingManagementV1alpha1                   *management_v1alpha1.UnscopedAssetMappingClient
	ExternalSecretManagementV1alpha1                 *management_v1alpha1.UnscopedExternalSecretClient
	VirtualServiceManagementV1alpha1                 *management_v1alpha1.UnscopedVirtualServiceClient
	OAS3DocumentManagementV1alpha1                   *management_v1alpha1.UnscopedOAS3DocumentClient
	WebhookManagementV1alpha1                        *management_v1alpha1.UnscopedWebhookClient
	ReleaseTagManagementV1alpha1                     *management_v1alpha1.UnscopedReleaseTagClient
	CredentialRequestDefinitionManagementV1alpha1    *management_v1alpha1.UnscopedCredentialRequestDefinitionClient
	SecretManagementV1alpha1                         *management_v1alpha1.UnscopedSecretClient
	AccessControlListManagementV1alpha1              *management_v1alpha1.UnscopedAccessControlListClient
	StageCatalogV1alpha1                             *catalog_v1alpha1.StageClient
	AssetCatalogV1alpha1                             *catalog_v1alpha1.AssetClient
	AssetReleaseCatalogV1alpha1                      *catalog_v1alpha1.AssetReleaseClient
	CategoryCatalogV1alpha1                          *catalog_v1alpha1.CategoryClient
	AuthorizationProfileCatalogV1alpha1              *catalog_v1alpha1.AuthorizationProfileClient
	ApplicationCatalogV1alpha1                       *catalog_v1alpha1.ApplicationClient
	CredentialCatalogV1alpha1                        *catalog_v1alpha1.UnscopedCredentialClient
	SubscriptionCatalogV1alpha1                      *catalog_v1alpha1.SubscriptionClient
	MarketplaceCatalogV1alpha1                       *catalog_v1alpha1.MarketplaceClient
	PublishedProductCatalogV1alpha1                  *catalog_v1alpha1.UnscopedPublishedProductClient
	ProductCatalogV1alpha1                           *catalog_v1alpha1.ProductClient
	ProductReleaseCatalogV1alpha1                    *catalog_v1alpha1.ProductReleaseClient
	ProductPlanUnitCatalogV1alpha1                   *catalog_v1alpha1.ProductPlanUnitClient
	ProductPlanCatalogV1alpha1                       *catalog_v1alpha1.ProductPlanClient
	QuotaCatalogV1alpha1                             *catalog_v1alpha1.UnscopedQuotaClient
	AssetMappingCatalogV1alpha1                      *catalog_v1alpha1.UnscopedAssetMappingClient
	AssetResourceCatalogV1alpha1                     *catalog_v1alpha1.UnscopedAssetResourceClient
	AssetRequestDefinitionCatalogV1alpha1            *catalog_v1alpha1.UnscopedAssetRequestDefinitionClient
	AssetRequestCatalogV1alpha1                      *catalog_v1alpha1.UnscopedAssetRequestClient
	DocumentCatalogV1alpha1                          *catalog_v1alpha1.UnscopedDocumentClient
	ResourceCatalogV1alpha1                          *catalog_v1alpha1.UnscopedResourceClient
	ProductOverviewCatalogV1alpha1                   *catalog_v1alpha1.UnscopedProductOverviewClient
	WebhookCatalogV1alpha1                           *catalog_v1alpha1.UnscopedWebhookClient
	ReleaseTagCatalogV1alpha1                        *catalog_v1alpha1.UnscopedReleaseTagClient
	CredentialRequestDefinitionCatalogV1alpha1       *catalog_v1alpha1.UnscopedCredentialRequestDefinitionClient
	SecretCatalogV1alpha1                            *catalog_v1alpha1.UnscopedSecretClient
	AccessControlListCatalogV1alpha1                 *catalog_v1alpha1.UnscopedAccessControlListClient
	CategoryCatalogV1                                *catalog_v1.CategoryClient
	MarketplaceCatalogV1                             *catalog_v1.MarketplaceClient
	ProductCatalogV1                                 *catalog_v1.ProductClient
	ProductReleaseCatalogV1                          *catalog_v1.ProductReleaseClient
	QuotaCatalogV1                                   *catalog_v1.UnscopedQuotaClient
	DocumentCatalogV1                                *catalog_v1.UnscopedDocumentClient
	ResourceCatalogV1                                *catalog_v1.UnscopedResourceClient
	ProductOverviewCatalogV1                         *catalog_v1.UnscopedProductOverviewClient
	ResourceGroupDefinitionsV1alpha1                 *definitions_v1alpha1.ResourceGroupClient
	ResourceDefinitionDefinitionsV1alpha1            *definitions_v1alpha1.UnscopedResourceDefinitionClient
	ResourceDefinitionVersionDefinitionsV1alpha1     *definitions_v1alpha1.UnscopedResourceDefinitionVersionClient
	CommandLineInterfaceDefinitionsV1alpha1          *definitions_v1alpha1.UnscopedCommandLineInterfaceClient
	AccessControlListDefinitionsV1alpha1             *definitions_v1alpha1.UnscopedAccessControlListClient
}

func New(b cAPIV1.Base) *Set {
	s := &Set{}

	var err error

	s.WatchTopicManagementV1alpha1, err = management_v1alpha1.NewWatchTopicClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.WatchTopic: %s", err))
	}
	s.DiscoveryAgentManagementV1alpha1, err = management_v1alpha1.NewDiscoveryAgentClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.DiscoveryAgent: %s", err))
	}
	s.TraceabilityAgentManagementV1alpha1, err = management_v1alpha1.NewTraceabilityAgentClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.TraceabilityAgent: %s", err))
	}
	s.GovernanceAgentManagementV1alpha1, err = management_v1alpha1.NewGovernanceAgentClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.GovernanceAgent: %s", err))
	}
	s.EnvironmentManagementV1alpha1, err = management_v1alpha1.NewEnvironmentClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.Environment: %s", err))
	}
	s.APIServiceManagementV1alpha1, err = management_v1alpha1.NewAPIServiceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.APIService: %s", err))
	}
	s.APIServiceRevisionManagementV1alpha1, err = management_v1alpha1.NewAPIServiceRevisionClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.APIServiceRevision: %s", err))
	}
	s.APIServiceInstanceManagementV1alpha1, err = management_v1alpha1.NewAPIServiceInstanceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.APIServiceInstance: %s", err))
	}
	s.ConsumerInstanceManagementV1alpha1, err = management_v1alpha1.NewConsumerInstanceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.ConsumerInstance: %s", err))
	}
	s.ConsumerSubscriptionDefinitionManagementV1alpha1, err = management_v1alpha1.NewConsumerSubscriptionDefinitionClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.ConsumerSubscriptionDefinition: %s", err))
	}
	s.IntegrationManagementV1alpha1, err = management_v1alpha1.NewIntegrationClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.Integration: %s", err))
	}
	s.ResourceHookManagementV1alpha1, err = management_v1alpha1.NewResourceHookClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.ResourceHook: %s", err))
	}
	s.K8SClusterManagementV1alpha1, err = management_v1alpha1.NewK8SClusterClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.K8SCluster: %s", err))
	}
	s.K8SResourceManagementV1alpha1, err = management_v1alpha1.NewK8SResourceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.K8SResource: %s", err))
	}
	s.ResourceDiscoveryManagementV1alpha1, err = management_v1alpha1.NewResourceDiscoveryClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.ResourceDiscovery: %s", err))
	}
	s.MeshManagementV1alpha1, err = management_v1alpha1.NewMeshClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.Mesh: %s", err))
	}
	s.SpecDiscoveryManagementV1alpha1, err = management_v1alpha1.NewSpecDiscoveryClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.SpecDiscovery: %s", err))
	}
	s.APISpecManagementV1alpha1, err = management_v1alpha1.NewAPISpecClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.APISpec: %s", err))
	}
	s.MeshWorkloadManagementV1alpha1, err = management_v1alpha1.NewMeshWorkloadClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.MeshWorkload: %s", err))
	}
	s.MeshServiceManagementV1alpha1, err = management_v1alpha1.NewMeshServiceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.MeshService: %s", err))
	}
	s.MeshDiscoveryManagementV1alpha1, err = management_v1alpha1.NewMeshDiscoveryClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.MeshDiscovery: %s", err))
	}
	s.AssetMappingTemplateManagementV1alpha1, err = management_v1alpha1.NewAssetMappingTemplateClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.AssetMappingTemplate: %s", err))
	}
	s.AccessRequestDefinitionManagementV1alpha1, err = management_v1alpha1.NewAccessRequestDefinitionClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.AccessRequestDefinition: %s", err))
	}
	s.AccessRequestManagementV1alpha1, err = management_v1alpha1.NewAccessRequestClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.AccessRequest: %s", err))
	}
	s.DeploymentManagementV1alpha1, err = management_v1alpha1.NewDeploymentClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.Deployment: %s", err))
	}
	s.AmplifyConfigManagementV1alpha1, err = management_v1alpha1.NewAmplifyConfigClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.AmplifyConfig: %s", err))
	}
	s.AmplifyRuntimeConfigManagementV1alpha1, err = management_v1alpha1.NewAmplifyRuntimeConfigClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.AmplifyRuntimeConfig: %s", err))
	}
	s.VirtualHostManagementV1alpha1, err = management_v1alpha1.NewVirtualHostClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.VirtualHost: %s", err))
	}
	s.ManagedApplicationManagementV1alpha1, err = management_v1alpha1.NewManagedApplicationClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.ManagedApplication: %s", err))
	}
	s.CredentialManagementV1alpha1, err = management_v1alpha1.NewCredentialClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.Credential: %s", err))
	}
	s.VirtualAPIManagementV1alpha1, err = management_v1alpha1.NewVirtualAPIClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.VirtualAPI: %s", err))
	}
	s.VirtualAPIReleaseManagementV1alpha1, err = management_v1alpha1.NewVirtualAPIReleaseClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.VirtualAPIRelease: %s", err))
	}
	s.CorsRuleManagementV1alpha1, err = management_v1alpha1.NewCorsRuleClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.CorsRule: %s", err))
	}
	s.AmplifyJWTAuthRuleManagementV1alpha1, err = management_v1alpha1.NewAmplifyJWTAuthRuleClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.AmplifyJWTAuthRule: %s", err))
	}
	s.AmplifyOAuth2AuthRuleManagementV1alpha1, err = management_v1alpha1.NewAmplifyOAuth2AuthRuleClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.AmplifyOAuth2AuthRule: %s", err))
	}
	s.APIKeyAuthRuleManagementV1alpha1, err = management_v1alpha1.NewAPIKeyAuthRuleClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.APIKeyAuthRule: %s", err))
	}
	s.AssetMappingManagementV1alpha1, err = management_v1alpha1.NewAssetMappingClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.AssetMapping: %s", err))
	}
	s.ExternalSecretManagementV1alpha1, err = management_v1alpha1.NewExternalSecretClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.ExternalSecret: %s", err))
	}
	s.VirtualServiceManagementV1alpha1, err = management_v1alpha1.NewVirtualServiceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.VirtualService: %s", err))
	}
	s.OAS3DocumentManagementV1alpha1, err = management_v1alpha1.NewOAS3DocumentClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.OAS3Document: %s", err))
	}
	s.WebhookManagementV1alpha1, err = management_v1alpha1.NewWebhookClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.Webhook: %s", err))
	}
	s.ReleaseTagManagementV1alpha1, err = management_v1alpha1.NewReleaseTagClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.ReleaseTag: %s", err))
	}
	s.CredentialRequestDefinitionManagementV1alpha1, err = management_v1alpha1.NewCredentialRequestDefinitionClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.CredentialRequestDefinition: %s", err))
	}
	s.SecretManagementV1alpha1, err = management_v1alpha1.NewSecretClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.Secret: %s", err))
	}
	s.AccessControlListManagementV1alpha1, err = management_v1alpha1.NewAccessControlListClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/management/v1alpha1.AccessControlList: %s", err))
	}
	s.StageCatalogV1alpha1, err = catalog_v1alpha1.NewStageClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Stage: %s", err))
	}
	s.AssetCatalogV1alpha1, err = catalog_v1alpha1.NewAssetClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Asset: %s", err))
	}
	s.AssetReleaseCatalogV1alpha1, err = catalog_v1alpha1.NewAssetReleaseClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.AssetRelease: %s", err))
	}
	s.CategoryCatalogV1alpha1, err = catalog_v1alpha1.NewCategoryClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Category: %s", err))
	}
	s.AuthorizationProfileCatalogV1alpha1, err = catalog_v1alpha1.NewAuthorizationProfileClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.AuthorizationProfile: %s", err))
	}
	s.ApplicationCatalogV1alpha1, err = catalog_v1alpha1.NewApplicationClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Application: %s", err))
	}
	s.CredentialCatalogV1alpha1, err = catalog_v1alpha1.NewCredentialClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Credential: %s", err))
	}
	s.SubscriptionCatalogV1alpha1, err = catalog_v1alpha1.NewSubscriptionClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Subscription: %s", err))
	}
	s.MarketplaceCatalogV1alpha1, err = catalog_v1alpha1.NewMarketplaceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Marketplace: %s", err))
	}
	s.PublishedProductCatalogV1alpha1, err = catalog_v1alpha1.NewPublishedProductClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.PublishedProduct: %s", err))
	}
	s.ProductCatalogV1alpha1, err = catalog_v1alpha1.NewProductClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Product: %s", err))
	}
	s.ProductReleaseCatalogV1alpha1, err = catalog_v1alpha1.NewProductReleaseClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.ProductRelease: %s", err))
	}
	s.ProductPlanUnitCatalogV1alpha1, err = catalog_v1alpha1.NewProductPlanUnitClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.ProductPlanUnit: %s", err))
	}
	s.ProductPlanCatalogV1alpha1, err = catalog_v1alpha1.NewProductPlanClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.ProductPlan: %s", err))
	}
	s.QuotaCatalogV1alpha1, err = catalog_v1alpha1.NewQuotaClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Quota: %s", err))
	}
	s.AssetMappingCatalogV1alpha1, err = catalog_v1alpha1.NewAssetMappingClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.AssetMapping: %s", err))
	}
	s.AssetResourceCatalogV1alpha1, err = catalog_v1alpha1.NewAssetResourceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.AssetResource: %s", err))
	}
	s.AssetRequestDefinitionCatalogV1alpha1, err = catalog_v1alpha1.NewAssetRequestDefinitionClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.AssetRequestDefinition: %s", err))
	}
	s.AssetRequestCatalogV1alpha1, err = catalog_v1alpha1.NewAssetRequestClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.AssetRequest: %s", err))
	}
	s.DocumentCatalogV1alpha1, err = catalog_v1alpha1.NewDocumentClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Document: %s", err))
	}
	s.ResourceCatalogV1alpha1, err = catalog_v1alpha1.NewResourceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Resource: %s", err))
	}
	s.ProductOverviewCatalogV1alpha1, err = catalog_v1alpha1.NewProductOverviewClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.ProductOverview: %s", err))
	}
	s.WebhookCatalogV1alpha1, err = catalog_v1alpha1.NewWebhookClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Webhook: %s", err))
	}
	s.ReleaseTagCatalogV1alpha1, err = catalog_v1alpha1.NewReleaseTagClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.ReleaseTag: %s", err))
	}
	s.CredentialRequestDefinitionCatalogV1alpha1, err = catalog_v1alpha1.NewCredentialRequestDefinitionClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.CredentialRequestDefinition: %s", err))
	}
	s.SecretCatalogV1alpha1, err = catalog_v1alpha1.NewSecretClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.Secret: %s", err))
	}
	s.AccessControlListCatalogV1alpha1, err = catalog_v1alpha1.NewAccessControlListClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1alpha1.AccessControlList: %s", err))
	}
	s.CategoryCatalogV1, err = catalog_v1.NewCategoryClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1.Category: %s", err))
	}
	s.MarketplaceCatalogV1, err = catalog_v1.NewMarketplaceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1.Marketplace: %s", err))
	}
	s.ProductCatalogV1, err = catalog_v1.NewProductClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1.Product: %s", err))
	}
	s.ProductReleaseCatalogV1, err = catalog_v1.NewProductReleaseClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1.ProductRelease: %s", err))
	}
	s.QuotaCatalogV1, err = catalog_v1.NewQuotaClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1.Quota: %s", err))
	}
	s.DocumentCatalogV1, err = catalog_v1.NewDocumentClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1.Document: %s", err))
	}
	s.ResourceCatalogV1, err = catalog_v1.NewResourceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1.Resource: %s", err))
	}
	s.ProductOverviewCatalogV1, err = catalog_v1.NewProductOverviewClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/catalog/v1.ProductOverview: %s", err))
	}
	s.ResourceGroupDefinitionsV1alpha1, err = definitions_v1alpha1.NewResourceGroupClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/definitions/v1alpha1.ResourceGroup: %s", err))
	}
	s.ResourceDefinitionDefinitionsV1alpha1, err = definitions_v1alpha1.NewResourceDefinitionClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/definitions/v1alpha1.ResourceDefinition: %s", err))
	}
	s.ResourceDefinitionVersionDefinitionsV1alpha1, err = definitions_v1alpha1.NewResourceDefinitionVersionClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/definitions/v1alpha1.ResourceDefinitionVersion: %s", err))
	}
	s.CommandLineInterfaceDefinitionsV1alpha1, err = definitions_v1alpha1.NewCommandLineInterfaceClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/definitions/v1alpha1.CommandLineInterface: %s", err))
	}
	s.AccessControlListDefinitionsV1alpha1, err = definitions_v1alpha1.NewAccessControlListClient(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client for github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/definitions/v1alpha1.AccessControlList: %s", err))
	}
	return s
}
