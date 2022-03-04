/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ProductSpecFilters Filters what AssetRelease the Product points to. (catalog.v1alpha1.Product)
type ProductSpecFilters struct {
	// The AssetRelease version to use. Examples:   - 1.0.1 for a specific asset release version   - 1.* for all minor and patch versions of version 1   - 1.2.* for all the patch version for version 1.2
	Version string `json:"version,omitempty"`
}
