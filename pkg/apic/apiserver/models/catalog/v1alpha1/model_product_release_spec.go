/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ProductReleaseSpec struct for ProductReleaseSpec
type ProductReleaseSpec struct {
	// Description of the release.
	Description string `json:"description,omitempty"`
	// Version of the release.
	Version    string                  `json:"version"`
	Product    string                  `json:"product"`
	Asset      ProductReleaseSpecAsset `json:"asset,omitempty"`
	ReleaseTag string                  `json:"releaseTag"`
	State      string                  `json:"state,omitempty"`
	// list of categories for the released product.
	Categories []string `json:"categories,omitempty"`
}
