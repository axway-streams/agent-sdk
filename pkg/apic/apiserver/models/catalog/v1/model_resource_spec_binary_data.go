/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

// ResourceSpecBinaryData struct for ResourceSpecBinaryData
type ResourceSpecBinaryData struct {
	Type string `json:"type"`
	// Base64 encoded value of the file.
	Content string `json:"content"`
}
