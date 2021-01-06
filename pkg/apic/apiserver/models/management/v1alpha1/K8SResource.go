/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_K8SResourceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "K8SResource",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	K8SResourceScope = "K8SCluster"

	K8SResourceResource = "k8sresources"
)

func K8SResourceGVK() apiv1.GroupVersionKind {
	return _K8SResourceGVK
}

func init() {
	apiv1.RegisterGVK(_K8SResourceGVK, K8SResourceScope, K8SResourceResource)
}

// K8SResource Resource
type K8SResource struct {
	apiv1.ResourceMeta

	Spec K8SResourceSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a K8SResource
func (res *K8SResource) FromInstance(ri *apiv1.ResourceInstance) error {
	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &K8SResourceSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = K8SResource{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a K8SResource to a ResourceInstance
func (res *K8SResource) AsInstance() (*apiv1.ResourceInstance, error) {
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
