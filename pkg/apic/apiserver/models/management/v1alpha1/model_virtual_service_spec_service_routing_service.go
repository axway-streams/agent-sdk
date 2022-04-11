/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// VirtualServiceSpecServiceRoutingService struct for VirtualServiceSpecServiceRoutingService
type VirtualServiceSpecServiceRoutingService struct {
	Prefix    string                                             `json:"prefix"`
	Protocol  string                                             `json:"protocol,omitempty"`
	Endpoints []VirtualServiceSpecServiceRoutingServiceEndpoints `json:"endpoints"`
	Codec     string                                             `json:"codec"`
	// Timeout in seconds.
	ConnectTimeout int32 `json:"connectTimeout,omitempty"`
	// The backend credentials.
	// GENERATE: The following code has been modified after code generation
	Credentials []interface{} `json:"credentials,omitempty"`
}
