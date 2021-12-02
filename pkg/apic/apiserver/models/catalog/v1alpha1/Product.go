/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ProductGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "Product",
		},
		APIVersion: "v1alpha1",
	}

	ProductScopes = []string{""}
)

const ProductResourceName = "products"

func ProductGVK() apiv1.GroupVersionKind {
	return _ProductGVK
}

func init() {
	apiv1.RegisterGVK(_ProductGVK, ProductScopes[0], ProductResourceName)
}

// Product Resource
type Product struct {
	apiv1.ResourceMeta

	Icon interface{} `json:"icon"`

	Owner *apiv1.Owner `json:"owner"`

	References ProductReferences `json:"references"`

	Spec ProductSpec `json:"spec"`

	State ProductState `json:"state"`
}

// FromInstance converts a ResourceInstance to a Product
func (res *Product) FromInstance(ri *apiv1.ResourceInstance) error {
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

// ProductFromInstanceArray converts a []*ResourceInstance to a []*Product
func ProductFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Product, error) {
	newArray := make([]*Product, 0)
	for _, item := range fromArray {
		res := &Product{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Product, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Product to a ResourceInstance
func (res *Product) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ProductGVK()
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
