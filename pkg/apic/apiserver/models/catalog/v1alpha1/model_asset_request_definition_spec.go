/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AssetRequestDefinitionSpec  (catalog.v1alpha1.AssetRequestDefinition)
type AssetRequestDefinitionSpec struct {
	// JSON Schema draft \\#7 for defining the AssetRequest properties needed to get access to an APIServiceInstance. (catalog.v1alpha1.AssetRequestDefinition)
	Schema    map[string]interface{}              `json:"schema"`
	Provision AssetRequestDefinitionSpecProvision `json:"provision,omitempty"`
}
