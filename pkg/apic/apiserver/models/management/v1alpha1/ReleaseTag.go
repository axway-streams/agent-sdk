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

	Owner *apiv1.Owner `json:"owner"`

	Spec ReleaseTagSpec `json:"spec"`

	State interface{} `json:"state"`

	Status ReleaseTagStatus `json:"status"`
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
