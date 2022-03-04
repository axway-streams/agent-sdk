/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AssetSpec  (catalog.v1alpha1.Asset)
type AssetSpec struct {
	// description of the asset.
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	// list of categories for the asset.
	Categories []string `json:"categories,omitempty"`
}
