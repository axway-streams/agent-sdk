/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ManagedApplicationGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "ManagedApplication",
		},
		APIVersion: "v1alpha1",
	}

	ManagedApplicationScopes = []string{"Environment"}
)

const ManagedApplicationResourceName = "managedapplications"

func ManagedApplicationGVK() apiv1.GroupVersionKind {
	return _ManagedApplicationGVK
}

func init() {
	apiv1.RegisterGVK(_ManagedApplicationGVK, ManagedApplicationScopes[0], ManagedApplicationResourceName)
}

// ManagedApplication Resource
type ManagedApplication struct {
	apiv1.ResourceMeta
	Marketplace ManagedApplicationMarketplace `json:"marketplace"`
	Owner       *apiv1.Owner                  `json:"owner"`
	References  ManagedApplicationReferences  `json:"references"`
	Spec        ManagedApplicationSpec        `json:"spec"`
	// 	Status      ManagedApplicationStatus      `json:"status"`
	Status *apiv1.ResourceStatus `json:"status"`
}

// NewManagedApplication creates an empty *ManagedApplication
func NewManagedApplication(name, scopeName string) *ManagedApplication {
	return &ManagedApplication{
		ResourceMeta: apiv1.ResourceMeta{
			Name:             name,
			GroupVersionKind: _ManagedApplicationGVK,
			Metadata: apiv1.Metadata{
				Scope: apiv1.MetadataScope{
					Name: scopeName,
					Kind: ManagedApplicationScopes[0],
				},
			},
		},
	}
}

// ManagedApplicationFromInstanceArray converts a []*ResourceInstance to a []*ManagedApplication
func ManagedApplicationFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*ManagedApplication, error) {
	newArray := make([]*ManagedApplication, 0)
	for _, item := range fromArray {
		res := &ManagedApplication{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*ManagedApplication, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a ManagedApplication to a ResourceInstance
func (res *ManagedApplication) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = ManagedApplicationGVK()
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

// FromInstance converts a ResourceInstance to a ManagedApplication
func (res *ManagedApplication) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *ManagedApplication) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(&res.ResourceMeta)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.Unmarshal(m, &out)
	if err != nil {
		return nil, err
	}

	out["marketplace"] = res.Marketplace
	out["owner"] = res.Owner
	out["references"] = res.References
	out["spec"] = res.Spec
	out["status"] = res.Status

	return json.Marshal(out)
}

// UnmarshalJSON custom unmarshaller to handle sub resources
func (res *ManagedApplication) UnmarshalJSON(data []byte) error {
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

	// marshalling subresource Marketplace
	if v, ok := aux.SubResources["marketplace"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "marketplace")
		err = json.Unmarshal(sr, &res.Marketplace)
		if err != nil {
			return err
		}
	}

	// marshalling subresource References
	if v, ok := aux.SubResources["references"]; ok {
		sr, err = json.Marshal(v)
		if err != nil {
			return err
		}

		delete(aux.SubResources, "references")
		err = json.Unmarshal(sr, &res.References)
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
func (res *ManagedApplication) PluralName() string {
	return ManagedApplicationResourceName
}
