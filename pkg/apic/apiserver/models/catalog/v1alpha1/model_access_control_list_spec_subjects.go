/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AccessControlListSpecSubjects  (catalog.v1alpha1.AccessControlList)
type AccessControlListSpecSubjects struct {
	// Type of the subject
	Type string `json:"type,omitempty"`
	// ID of the subject
	Id string `json:"id,omitempty"`
}
