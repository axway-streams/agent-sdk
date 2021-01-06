/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_MeshDiscoveryGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "MeshDiscovery",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	MeshDiscoveryScope = "Mesh"

	MeshDiscoveryResource = "meshdiscoveries"
)

func MeshDiscoveryGVK() apiv1.GroupVersionKind {
	return _MeshDiscoveryGVK
}

func init() {
	apiv1.RegisterGVK(_MeshDiscoveryGVK, MeshDiscoveryScope, MeshDiscoveryResource)
}

// MeshDiscovery Resource
type MeshDiscovery struct {
	apiv1.ResourceMeta

	Spec MeshDiscoverySpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a MeshDiscovery
func (res *MeshDiscovery) FromInstance(ri *apiv1.ResourceInstance) error {
	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &MeshDiscoverySpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = MeshDiscovery{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a MeshDiscovery to a ResourceInstance
func (res *MeshDiscovery) AsInstance() (*apiv1.ResourceInstance, error) {
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
