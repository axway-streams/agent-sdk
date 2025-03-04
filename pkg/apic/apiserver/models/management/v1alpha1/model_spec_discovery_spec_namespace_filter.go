/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// SpecDiscoverySpecNamespaceFilter a list of namespace names to follow. If not set, follows all namespaces. (management.v1alpha1.SpecDiscovery)
type SpecDiscoverySpecNamespaceFilter struct {
	Names []string `json:"names,omitempty"`
}
