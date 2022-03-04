/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

// ProductReleaseSpecRelease  (catalog.v1.ProductRelease)
type ProductReleaseSpecRelease struct {
	Name string `json:"name,omitempty"`
	// The AssetRelease version.
	Version string `json:"version,omitempty"`
	// The AssetRelease state.
	State string `json:"state,omitempty"`
}
