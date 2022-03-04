/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ApiKeyAuthRuleSpec  (management.v1alpha1.APIKeyAuthRule)
type ApiKeyAuthRuleSpec struct {
	// The description of the authentication rule.
	Description string `json:"description,omitempty"`
	// Where to look for the API key, defaults to header.
	In string `json:"in,omitempty"`
	// name of the API key field, defaults to x-api-key.
	Name string `json:"name,omitempty"`
}
