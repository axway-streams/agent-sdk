/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_EdgeDataplaneGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "EdgeDataplane",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	EdgeDataplaneScope = "Environment"

	EdgeDataplaneResource = "edgedataplanes"
)

func EdgeDataplaneGVK() apiv1.GroupVersionKind {
	return _EdgeDataplaneGVK
}

func init() {
	apiv1.RegisterGVK(_EdgeDataplaneGVK, EdgeDataplaneScope, EdgeDataplaneResource)
}

// EdgeDataplane Resource
type EdgeDataplane struct {
	apiv1.ResourceMeta

	Spec EdgeDataplaneSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a EdgeDataplane
func (res *EdgeDataplane) FromInstance(ri *apiv1.ResourceInstance) error {
	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &EdgeDataplaneSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = EdgeDataplane{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a EdgeDataplane to a ResourceInstance
func (res *EdgeDataplane) AsInstance() (*apiv1.ResourceInstance, error) {
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
