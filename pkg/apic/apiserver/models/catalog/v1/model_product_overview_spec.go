/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

// ProductOverviewSpec  (catalog.v1.ProductOverview)
type ProductOverviewSpec struct {
	// Defines all the documents and order for marketplace.
	Documents []ProductOverviewSpecDocuments `json:"documents,omitempty"`
}
