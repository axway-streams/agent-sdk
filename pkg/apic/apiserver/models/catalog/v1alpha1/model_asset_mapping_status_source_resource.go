/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AssetMappingStatusSourceResource The resource that triggered the Asset Mapping. (catalog.v1alpha1.AssetMapping)
type AssetMappingStatusSourceResource struct {
	ApiService           AssetMappingStatusSourceResourceApiService           `json:"apiService,omitempty"`
	ApiServiceRevision   AssetMappingStatusSourceResourceApiServiceRevision   `json:"apiServiceRevision,omitempty"`
	ApiServiceInstance   AssetMappingStatusSourceResourceApiServiceInstance   `json:"apiServiceInstance,omitempty"`
	AssetMappingTemplate AssetMappingStatusSourceResourceAssetMappingTemplate `json:"assetMappingTemplate,omitempty"`
}
