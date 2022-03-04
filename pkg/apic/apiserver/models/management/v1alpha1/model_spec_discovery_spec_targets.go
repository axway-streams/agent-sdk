/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// SpecDiscoverySpecTargets  (management.v1alpha1.SpecDiscovery)
type SpecDiscoverySpecTargets struct {
	ExactPaths     []SpecDiscoverySpecTargetsExactPaths     `json:"exactPaths,omitempty"`
	FromAnnotation []SpecDiscoverySpecTargetsFromAnnotation `json:"fromAnnotation,omitempty"`
}
