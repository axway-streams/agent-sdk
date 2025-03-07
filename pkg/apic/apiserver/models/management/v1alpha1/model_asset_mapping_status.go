/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AssetMappingStatus  (management.v1alpha1.AssetMapping)
type AssetMappingStatus struct {
	Level  string                   `json:"level,omitempty"`
	Source AssetMappingStatusSource `json:"source,omitempty"`
	// Generated catalog resources.
	Outputs []AssetMappingStatusOutputs `json:"outputs,omitempty"`
}
