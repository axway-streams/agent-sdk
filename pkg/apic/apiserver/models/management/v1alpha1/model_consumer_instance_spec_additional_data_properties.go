/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ConsumerInstanceSpecAdditionalDataProperties The Catalog Item properties for additional data. (management.v1alpha1.ConsumerInstance)
type ConsumerInstanceSpecAdditionalDataProperties struct {
	// Defines the unstrucured data type. Example 'APIBuilderConnector' or 'SDK'.
	Type string `json:"type"`
	// Defines the Content Type of the data.
	ContentType string `json:"contentType"`
	// Short name for the unstructured data.
	Label string `json:"label,omitempty"`
	// File name used to be sent as part of the content disposition header for revision unstructured data download.
	FileName string `json:"fileName,omitempty"`
	// Base64 encoded data for the Catalog Item revision content.
	Data string `json:"data"`
}
