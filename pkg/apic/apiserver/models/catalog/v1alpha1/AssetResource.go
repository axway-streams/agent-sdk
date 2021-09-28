/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AssetResourceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "AssetResource",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	AssetResourceScope = "Asset"

	AssetResourceResourceName = "assetresources"
)

func AssetResourceGVK() apiv1.GroupVersionKind {
	return _AssetResourceGVK
}

func init() {
	apiv1.RegisterGVK(_AssetResourceGVK, AssetResourceScope, AssetResourceResourceName)
}

// AssetResource Resource
type AssetResource struct {
	apiv1.ResourceMeta

	Owner *apiv1.Owner `json:"owner"`

	References AssetResourceReferences `json:"references"`

	Spec AssetResourceSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a AssetResource
func (res *AssetResource) FromInstance(ri *apiv1.ResourceInstance) error {
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

// AssetResourceFromInstanceArray converts a []*ResourceInstance to a []*AssetResource
func AssetResourceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AssetResource, error) {
	newArray := make([]*AssetResource, 0)
	for _, item := range fromArray {
		res := &AssetResource{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AssetResource, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AssetResource to a ResourceInstance
func (res *AssetResource) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = AssetResourceGVK()
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
