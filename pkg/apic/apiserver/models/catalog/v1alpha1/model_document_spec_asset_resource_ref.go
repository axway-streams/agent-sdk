/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// DocumentSpecAssetResourceRef struct for DocumentSpecAssetResourceRef
type DocumentSpecAssetResourceRef struct {
	Kind string `json:"kind"`
	Name string `json:"name,omitempty"`
	// Title for the article.
	Title string `json:"title"`
}
