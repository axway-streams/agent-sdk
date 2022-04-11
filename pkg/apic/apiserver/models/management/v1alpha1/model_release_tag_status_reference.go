/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ReleaseTagStatusReference The resource reference that was created/updated.
type ReleaseTagStatusReference struct {
	// message of the status of the reference
	Message string `json:"message,omitempty"`
	Kind    string `json:"kind,omitempty"`
	// The name of the resource that got created.
	Name      string `json:"name,omitempty"`
	ScopeKind string `json:"scopeKind,omitempty"`
	ScopeName string `json:"scopeName,omitempty"`
}
