/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// MeshWorkloadSpec  (management.v1alpha1.MeshWorkload)
type MeshWorkloadSpec struct {
	// References to k8sresources making up the workload.
	Resources []string `json:"resources,omitempty"`
	// Labels from the k8sresources making up the workload. (management.v1alpha1.MeshWorkload)
	Labels map[string]string `json:"labels,omitempty"`
	// Discovered ports on the workload.
	Ports []MeshWorkloadSpecPorts `json:"ports,omitempty"`
}
