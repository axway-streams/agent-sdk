/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ExternalSecretSpecBasicAuthData Data for Basic Auth.
type ExternalSecretSpecBasicAuthData struct {
	Kind string `json:"kind"`
	// Alias for the Basic Auth username in the secret.
	UsernameAlias string `json:"usernameAlias"`
	// Alias for the Basic Auth password in the secret.
	PasswordAlias string `json:"passwordAlias"`
}
