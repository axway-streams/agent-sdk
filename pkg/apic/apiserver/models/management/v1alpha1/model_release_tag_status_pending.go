/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

import (
	// GENERATE: The following code has been modified after code generation
	// 	"time"
	time "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

// ReleaseTagStatusPending struct for ReleaseTagStatusPending
type ReleaseTagStatusPending struct {
	Type string `json:"type"`
	// Time when the change occured.
	Timestamp time.Time `json:"timestamp"`
	// message of the pending status
	Detail string                            `json:"detail"`
	Meta   map[string]map[string]interface{} `json:"meta,omitempty"`
}
