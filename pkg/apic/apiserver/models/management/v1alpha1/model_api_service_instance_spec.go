/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ApiServiceInstanceSpec  (management.v1alpha1.APIServiceInstance)
type ApiServiceInstanceSpec struct {
	ApiServiceRevision           string   `json:"apiServiceRevision"`
	AccessRequestDefinition      string   `json:"accessRequestDefinition,omitempty"`
	CredentialRequestDefinitions []string `json:"credentialRequestDefinitions,omitempty"`
	// A list of locations where the api is deployed.
	Endpoint []ApiServiceInstanceSpecEndpoint `json:"endpoint"`
}
