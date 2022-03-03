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

	MeshWorkloadScopes = []string{"Mesh"}
)

const MeshWorkloadResourceName = "meshworkloads"

func MeshWorkloadGVK() apiv1.GroupVersionKind {
	return _MeshWorkloadGVK
}

func init() {
	apiv1.RegisterGVK(_MeshWorkloadGVK, MeshWorkloadScopes[0], MeshWorkloadResourceName)
}

// MeshWorkload Resource
type MeshWorkload struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner     `json:"owner"`
	Spec  MeshWorkloadSpec `json:"spec"`
}

// MeshWorkloadFromInstanceArray converts a []*ResourceInstance to a []*MeshWorkload
func MeshWorkloadFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*MeshWorkload, error) {
	newArray := make([]*MeshWorkload, 0)
	for _, item := range fromArray {
		res := &MeshWorkload{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*MeshWorkload, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a MeshWorkload to a ResourceInstance
func (res *MeshWorkload) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = MeshWorkloadGVK()
	res.ResourceMeta = meta

	m, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	instance := apiv1.ResourceInstance{}
	err = json.Unmarshal(m, &instance)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}

// FromInstance converts a ResourceInstance to a MeshWorkload
func (res *MeshWorkload) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}
	var err error
	rawResource := ri.GetRawResource()
	if rawResource == nil {
		rawResource, err = json.Marshal(ri)
		if err != nil {
			return err
		}
	}
	err = json.Unmarshal(rawResource, res)
	return err
}

// MarshalJSON custom marshaller to handle sub resources
func (res *MeshWorkload) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["owner"] = res.Owner
	out["spec"] = res.Spec

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *MeshWorkload) UnmarshalJSON(data []byte) error {
	var err error

	aux := &apiv1.ResourceInstance{}
	err = json.Unmarshal(data, aux)
	if err != nil {
		return err
	}

	res.ResourceMeta = aux.ResourceMeta
	res.Owner = aux.Owner

	// ResourceInstance holds the spec as a map[string]interface{}.
	// Convert it to bytes, then convert to the spec type for the resource.
	sr, err := json.Marshal(aux.Spec)
	if err != nil {
		return err
	}

	err = json.Unmarshal(sr, &res.Spec)
	if err != nil {
		return err
	}

	return nil
}

// PluralName returns the plural name of the resource
func (res *MeshWorkload) PluralName() string {
	return MeshWorkloadResourceName
}
