/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

// ProductReleaseStatus  (catalog.v1.ProductRelease)
type ProductReleaseStatus struct {
	// The current status level, indicating progress towards consistency.
	Level string `json:"level,omitempty"`
	// Reasons for the generated status.
	// GENERATE: The following code has been modified after code generation
	Reasons []interface{} `json:"reasons,omitempty"`
}
