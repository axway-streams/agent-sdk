/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// VirtualServiceSpecApikeycredentialIn The object containing the api key name and location in the request.
type VirtualServiceSpecApikeycredentialIn struct {
	// The name of the api key header.
	Name     string `json:"name"`
	Location string `json:"location"`
}
