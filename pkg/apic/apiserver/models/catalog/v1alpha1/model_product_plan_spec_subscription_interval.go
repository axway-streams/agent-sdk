/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ProductPlanSpecSubscriptionInterval The subscription interval (catalog.v1alpha1.ProductPlan)
type ProductPlanSpecSubscriptionInterval struct {
	// The type of the interval
	Type string `json:"type,omitempty"`
	// The subscription inverval length
	Length float32 `json:"length,omitempty"`
}
