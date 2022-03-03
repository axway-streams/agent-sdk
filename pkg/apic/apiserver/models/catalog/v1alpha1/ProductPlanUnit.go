/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ProductPlanUnitGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "catalog",
			Kind:  "ProductPlanUnit",
		},
		APIVersion: "v1alpha1",
	}

	ProductPlanUnitScopes = []string{""}
)

const ProductPlanUnitResourceName = "productplanunits"

func ProductPlanUnitGVK() apiv1.GroupVersionKind {
	return _ProductPlanUnitGVK
}

func init() {
	apiv1.RegisterGVK(_ProductPlanUnitGVK, ProductPlanUnitScopes[0], ProductPlanUnitResourceName)
}

// ProductPlanUnit Resource
type ProductPlanUnit struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner        `json:"owner"`
	Spec  ProductPlanUnitSpec `json:"spec"`
}

// ProductPlanUnitFromInstanceArray converts a []*ResourceInstance to a []*ProductPlanUnit
func ProductPlanUnitFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ProductPlanUnit, error) {
	newArray := make([]*ProductPlanUnit, 0)
	for _, item := range fromArray {
		res := &ProductPlanUnit{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ProductPlanUnit, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ProductPlanUnit to a ResourceInstance
func (res *ProductPlanUnit) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ProductPlanUnitGVK()
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

// FromInstance converts a ResourceInstance to a ProductPlanUnit
func (res *ProductPlanUnit) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *ProductPlanUnit) MarshalJSON() ([]byte, error) {
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
func (res *ProductPlanUnit) UnmarshalJSON(data []byte) error {
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
func (res *ProductPlanUnit) PluralName() string {
	return ProductPlanUnitResourceName
}
