/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ReleaseTagGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "ReleaseTag",
		},
		APIVersion: "v1alpha1",
	}

	ReleaseTagScopes = []string{"VirtualAPI"}
)

const ReleaseTagResourceName = "releasetags"

func ReleaseTagGVK() apiv1.GroupVersionKind {
	return _ReleaseTagGVK
}

func init() {
	apiv1.RegisterGVK(_ReleaseTagGVK, ReleaseTagScopes[0], ReleaseTagResourceName)
}

// ReleaseTag Resource
type ReleaseTag struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner   `json:"owner"`
	Spec  ReleaseTagSpec `json:"spec"`
	State interface{}    `json:"state"`
	// 	Status ReleaseTagStatus `json:"status"`
	Status *apiv1.ResourceStatus `json:"status"`
}

// ReleaseTagFromInstanceArray converts a []*ResourceInstance to a []*ReleaseTag
func ReleaseTagFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ReleaseTag, error) {
	newArray := make([]*ReleaseTag, 0)
	for _, item := range fromArray {
		res := &ReleaseTag{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ReleaseTag, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ReleaseTag to a ResourceInstance
func (res *ReleaseTag) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ReleaseTagGVK()
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

// FromInstance converts a ResourceInstance to a ReleaseTag
func (res *ReleaseTag) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *ReleaseTag) MarshalJSON() ([]byte, error) {
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
	out["state"] = res.State
	out["status"] = res.Status

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *ReleaseTag) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource State
	if v, ok := aux.SubResources["state"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "state")
		err = json.Unmarshal(sr, &res.State)
		if err != nil {
			return err
		}
	}

	// marshalling subresource Status
	if v, ok := aux.SubResources["status"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "status")
		// 		err = json.Unmarshal(sr, &res.Status)
		res.Status = &apiv1.ResourceStatus{}
		err = json.Unmarshal(sr, res.Status)
		if err != nil {
			return err
		}
	}

	return nil
}

// PluralName returns the plural name of the resource
func (res *ReleaseTag) PluralName() string {
	return ReleaseTagResourceName
}
