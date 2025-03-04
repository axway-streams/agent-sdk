/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ProductPlanSpecSubscription Defines Plan's subscription information (catalog.v1alpha1.ProductPlan)
type ProductPlanSpecSubscription struct {
	Interval ProductPlanSpecSubscriptionInterval `json:"interval,omitempty"`
	Renewal  string                              `json:"renewal,omitempty"`
	Approval string                              `json:"approval,omitempty"`
}
