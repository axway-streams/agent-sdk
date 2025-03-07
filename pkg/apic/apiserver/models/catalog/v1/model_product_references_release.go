/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

// ProductReferencesRelease  (catalog.v1.Product)
type ProductReferencesRelease struct {
	// The latest AssetRelease computed based on the provided Asset filters.
	Name string `json:"name,omitempty"`
	// The AssetRelease version that the Product currently points to.
	Version string `json:"version,omitempty"`
	// The AssetRelease state.
	State string `json:"state,omitempty"`
}
