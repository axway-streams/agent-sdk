/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AccessRequestDefinitionGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "AccessRequestDefinition",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	AccessRequestDefinitionScope = "Environment"

	AccessRequestDefinitionResourceName = "accessrequestdefinitions"
)

func AccessRequestDefinitionGVK() apiv1.GroupVersionKind {
	return _AccessRequestDefinitionGVK
}

func init() {
	apiv1.RegisterGVK(_AccessRequestDefinitionGVK, AccessRequestDefinitionScope, AccessRequestDefinitionResourceName)
}

// AccessRequestDefinition Resource
type AccessRequestDefinition struct {
	apiv1.ResourceMeta

	Owner *apiv1.Owner `json:"owner"`

	Spec AccessRequestDefinitionSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a AccessRequestDefinition
func (res *AccessRequestDefinition) FromInstance(ri *apiv1.ResourceInstance) error {
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

// AccessRequestDefinitionFromInstanceArray converts a []*ResourceInstance to a []*AccessRequestDefinition
func AccessRequestDefinitionFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AccessRequestDefinition, error) {
	newArray := make([]*AccessRequestDefinition, 0)
	for _, item := range fromArray {
		res := &AccessRequestDefinition{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AccessRequestDefinition, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AccessRequestDefinition to a ResourceInstance
func (res *AccessRequestDefinition) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = AccessRequestDefinitionGVK()
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
