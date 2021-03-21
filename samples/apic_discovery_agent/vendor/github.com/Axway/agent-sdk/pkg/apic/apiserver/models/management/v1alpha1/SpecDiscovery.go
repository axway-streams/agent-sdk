/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_SpecDiscoveryGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "SpecDiscovery",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	SpecDiscoveryScope = "K8SCluster"

	SpecDiscoveryResource = "specdiscoveries"
)

func SpecDiscoveryGVK() apiv1.GroupVersionKind {
	return _SpecDiscoveryGVK
}

func init() {
	apiv1.RegisterGVK(_SpecDiscoveryGVK, SpecDiscoveryScope, SpecDiscoveryResource)
}

// SpecDiscovery Resource
type SpecDiscovery struct {
	apiv1.ResourceMeta

	Spec SpecDiscoverySpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a SpecDiscovery
func (res *SpecDiscovery) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &SpecDiscoverySpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = SpecDiscovery{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a SpecDiscovery to a ResourceInstance
func (res *SpecDiscovery) AsInstance() (*apiv1.ResourceInstance, error) {
	m, err := json.Marshal(res.Spec)
	if err != nil {
		return nil, err
	}

	spec := map[string]interface{}{}
	err = json.Unmarshal(m, &spec)
	if err != nil {
		return nil, err
	}

	meta := res.ResourceMeta
	meta.GroupVersionKind = SpecDiscoveryGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
