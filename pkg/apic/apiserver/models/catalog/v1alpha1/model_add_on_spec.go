/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// AddOnSpec  (catalog.v1alpha1.AddOn)
type AddOnSpec struct {
	// description of the Plan.
	Description string `json:"description,omitempty"`
	// The unit used to measure the access to the resource.
	Unit string `json:"unit"`
	// The resources that the access is being granted to.
	// GENERATE: The following code has been modified after code generation
	Resources []interface{} `json:"resources,omitempty"`
}
