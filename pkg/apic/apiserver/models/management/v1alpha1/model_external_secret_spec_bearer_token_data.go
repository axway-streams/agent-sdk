/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ExternalSecretSpecBearerTokenData Data for Bearer Token Auth.
type ExternalSecretSpecBearerTokenData struct {
	Kind string `json:"kind"`
	// Alias for the Bearer Token in the secret.
	Alias string `json:"alias"`
}
