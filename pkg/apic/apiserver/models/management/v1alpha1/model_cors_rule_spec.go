/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// CorsRuleSpec  (management.v1alpha1.CorsRule)
type CorsRuleSpec struct {
	// CORS configuration rule.
	Description                   string   `json:"description,omitempty"`
	AccessControlAllowOrigin      []string `json:"accessControlAllowOrigin"`
	AccessControlMaxAge           int32    `json:"accessControlMaxAge,omitempty"`
	AccessControlAllowHeaders     []string `json:"accessControlAllowHeaders,omitempty"`
	AccessControlAllowMethods     []string `json:"accessControlAllowMethods,omitempty"`
	AccessControlExposeHeaders    []string `json:"accessControlExposeHeaders,omitempty"`
	AccessControlAllowCredentials bool     `json:"accessControlAllowCredentials,omitempty"`
}
