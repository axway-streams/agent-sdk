/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// SpecDiscoverySpecTargetsExactPaths  (management.v1alpha1.SpecDiscovery)
type SpecDiscoverySpecTargetsExactPaths struct {
	// path to api definition
	Path string `json:"path,omitempty"`
	// headers to add to the query (management.v1alpha1.SpecDiscovery)
	Headers map[string]string `json:"headers,omitempty"`
}
