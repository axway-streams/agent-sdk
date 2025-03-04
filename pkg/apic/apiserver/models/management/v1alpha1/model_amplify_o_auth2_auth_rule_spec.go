/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AmplifyOAuth2AuthRuleSpec  (management.v1alpha1.AmplifyOAuth2AuthRule)
type AmplifyOAuth2AuthRuleSpec struct {
	// The description of the authentication rule.
	Description        string `json:"description,omitempty"`
	Amplifyjwtauthrule string `json:"amplifyjwtauthrule"`
	// GENERATE: The following code has been modified after code generation
	Flows AmplifyOAuth2AuthRuleSpecOAuthFlows `json:"flows"`
}
