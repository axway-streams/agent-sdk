/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_MeshWorkloadGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "MeshWorkload",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	MeshWorkloadScope = "Mesh"

	MeshWorkloadResource = "meshworkloads"
)

func MeshWorkloadGVK() apiv1.GroupVersionKind {
	return _MeshWorkloadGVK
}

func init() {
	apiv1.RegisterGVK(_MeshWorkloadGVK, MeshWorkloadScope, MeshWorkloadResource)
}

// MeshWorkload Resource
type MeshWorkload struct {
	apiv1.ResourceMeta

	Spec MeshWorkloadSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a MeshWorkload
func (res *MeshWorkload) FromInstance(ri *apiv1.ResourceInstance) error {
	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &MeshWorkloadSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = MeshWorkload{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a MeshWorkload to a ResourceInstance
func (res *MeshWorkload) AsInstance() (*apiv1.ResourceInstance, error) {
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
