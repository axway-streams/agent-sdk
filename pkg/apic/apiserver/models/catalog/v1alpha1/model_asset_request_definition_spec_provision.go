/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AssetRequestDefinitionSpecProvision  (catalog.v1alpha1.AssetRequestDefinition)
type AssetRequestDefinitionSpecProvision struct {
	// JSON Schema draft \\#7 for describing the data to be sent back after access has been provisioned. (catalog.v1alpha1.AssetRequestDefinition)
	Schema map[string]interface{} `json:"schema"`
}
