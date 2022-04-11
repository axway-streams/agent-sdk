/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// VirtualServiceSpecApikeycredential The API Key credential.
type VirtualServiceSpecApikeycredential struct {
	Kind string `json:"kind"`
	// The name of the credential eg. external secret containing the api key.
	Name string `json:"name"`
	// The location of the api key eg. header, query or cookie.
	In string `json:"in"`
}
