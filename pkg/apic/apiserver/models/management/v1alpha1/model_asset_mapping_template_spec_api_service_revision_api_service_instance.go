/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AssetMappingTemplateSpecApiServiceRevisionApiServiceInstance  (management.v1alpha1.AssetMappingTemplate)
type AssetMappingTemplateSpecApiServiceRevisionApiServiceInstance struct {
	// Attributes used to filter the APIServiceInstances for the API Service on which the template applies. (management.v1alpha1.AssetMappingTemplate)
	Attributes map[string]string `json:"attributes,omitempty"`
}
