/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AmplifyOAuth2AuthRuleSpecPasswordOAuthFlow struct for AmplifyOAuth2AuthRuleSpecPasswordOAuthFlow
type AmplifyOAuth2AuthRuleSpecPasswordOAuthFlow struct {
	TokenUrl   string `json:"tokenUrl"`
	RefreshUrl string `json:"refreshUrl,omitempty"`
}
