/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AssetMappingTemplateGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "AssetMappingTemplate",
		},
		APIVersion: "v1alpha1",
	}

	AssetMappingTemplateScopes = []string{"Environment"}
)

const AssetMappingTemplateResourceName = "assetmappingtemplates"

func AssetMappingTemplateGVK() apiv1.GroupVersionKind {
	return _AssetMappingTemplateGVK
}

func init() {
	apiv1.RegisterGVK(_AssetMappingTemplateGVK, AssetMappingTemplateScopes[0], AssetMappingTemplateResourceName)
}

// AssetMappingTemplate Resource
type AssetMappingTemplate struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner             `json:"owner"`
	Spec  AssetMappingTemplateSpec `json:"spec"`
}

// AssetMappingTemplateFromInstanceArray converts a []*ResourceInstance to a []*AssetMappingTemplate
func AssetMappingTemplateFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AssetMappingTemplate, error) {
	newArray := make([]*AssetMappingTemplate, 0)
	for _, item := range fromArray {
		res := &AssetMappingTemplate{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AssetMappingTemplate, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AssetMappingTemplate to a ResourceInstance
func (res *AssetMappingTemplate) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = AssetMappingTemplateGVK()
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

// FromInstance converts a ResourceInstance to a AssetMappingTemplate
func (res *AssetMappingTemplate) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *AssetMappingTemplate) MarshalJSON() ([]byte, error) {
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
func (res *AssetMappingTemplate) UnmarshalJSON(data []byte) error {
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
func (res *AssetMappingTemplate) PluralName() string {
	return AssetMappingTemplateResourceName
}
