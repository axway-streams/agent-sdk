/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "git.ecd.axway.org/apigov/apic_agents_sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ResourceDefinitionVersionGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "definitions",
			Kind:  "ResourceDefinitionVersion",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	ResourceDefinitionVersionScope = "ResourceGroup"

	ResourceDefinitionVersionResource = "resourceversions"
)

func ResourceDefinitionVersionGVK() apiv1.GroupVersionKind {
	return _ResourceDefinitionVersionGVK
}

func init() {
	apiv1.RegisterGVK(_ResourceDefinitionVersionGVK, ResourceDefinitionVersionScope, ResourceDefinitionVersionResource)
}

// ResourceDefinitionVersion Resource
type ResourceDefinitionVersion struct {
	apiv1.ResourceMeta

	Spec ResourceDefinitionVersionSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a ResourceDefinitionVersion
func (res *ResourceDefinitionVersion) FromInstance(ri *apiv1.ResourceInstance) error {
	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &ResourceDefinitionVersionSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = ResourceDefinitionVersion{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a ResourceDefinitionVersion to a ResourceInstance
func (res *ResourceDefinitionVersion) AsInstance() (*apiv1.ResourceInstance, error) {
	m, err := json.Marshal(res.Spec)
	if err != nil {
		return nil, err
	}

	spec := map[string]interface{}{}
	err = json.Unmarshal(m, &spec)
	if err != nil {
		return nil, err
	}

	return &apiv1.ResourceInstance{ResourceMeta: res.ResourceMeta, Spec: spec}, nil
}
