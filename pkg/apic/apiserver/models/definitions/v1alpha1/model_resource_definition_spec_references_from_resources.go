/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ResourceDefinitionSpecReferencesFromResources struct for ResourceDefinitionSpecReferencesFromResources
type ResourceDefinitionSpecReferencesFromResources struct {
	// Defines the kind of the resource.
	Kind string `json:"kind,omitempty"`
	// The type of the reference.
	Types []string `json:"types,omitempty"`
}
