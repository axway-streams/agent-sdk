/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// SubscriptionMarketplace Details about the marketplace Application. (catalog.v1alpha1.Subscription)
type SubscriptionMarketplace struct {
	// The name of the Marketplace.
	Name     string                          `json:"name"`
	Resource SubscriptionMarketplaceResource `json:"resource"`
}
