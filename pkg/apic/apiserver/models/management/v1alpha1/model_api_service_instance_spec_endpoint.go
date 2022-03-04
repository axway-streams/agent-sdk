/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ApiServiceInstanceSpecEndpoint  (management.v1alpha1.APIServiceInstance)
type ApiServiceInstanceSpecEndpoint struct {
	Host     string                        `json:"host"`
	Port     int32                         `json:"port,omitempty"`
	Protocol string                        `json:"protocol"`
	Routing  ApiServiceInstanceSpecRouting `json:"routing,omitempty"`
}
