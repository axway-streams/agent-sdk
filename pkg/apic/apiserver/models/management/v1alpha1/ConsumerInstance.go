/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ConsumerInstanceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "ConsumerInstance",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	ConsumerInstanceScope = "Environment"

	ConsumerInstanceResourceName = "consumerinstances"
)

func ConsumerInstanceGVK() apiv1.GroupVersionKind {
	return _ConsumerInstanceGVK
}

func init() {
	apiv1.RegisterGVK(_ConsumerInstanceGVK, ConsumerInstanceScope, ConsumerInstanceResourceName)
}

// ConsumerInstance Resource
type ConsumerInstance struct {
	apiv1.ResourceMeta

	Owner *apiv1.Owner `json:"owner"`

	References ConsumerInstanceReferences `json:"references"`

	Spec ConsumerInstanceSpec `json:"spec"`

	Status ConsumerInstanceStatus `json:"status"`
}

// FromInstance converts a ResourceInstance to a ConsumerInstance
func (res *ConsumerInstance) FromInstance(ri *apiv1.ResourceInstance) error {
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

// ConsumerInstanceFromInstanceArray converts a []*ResourceInstance to a []*ConsumerInstance
func ConsumerInstanceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ConsumerInstance, error) {
	newArray := make([]*ConsumerInstance, 0)
	for _, item := range fromArray {
		res := &ConsumerInstance{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ConsumerInstance, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ConsumerInstance to a ResourceInstance
func (res *ConsumerInstance) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ConsumerInstanceGVK()
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
