/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// WatchTopicSpecScope Resource scope filter details. (management.v1alpha1.WatchTopic)
type WatchTopicSpecScope struct {
	// Value for the Kind of the scope of the resource. Use \"*\" for any.
	Kind string `json:"kind"`
	// Name of the scope of the resource. Use \"*\" for any.
	Name string `json:"name"`
}
