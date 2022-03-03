/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_SecretGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "Secret",
		},
		APIVersion: "v1alpha1",
	}

	SecretScopes = []string{"Environment", "Integration"}
)

const SecretResourceName = "secrets"

func SecretGVK() apiv1.GroupVersionKind {
	return _SecretGVK
}

func init() {
	apiv1.RegisterGVK(_SecretGVK, SecretScopes[0], SecretResourceName)
}

// Secret Resource
type Secret struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner `json:"owner"`
	Spec  SecretSpec   `json:"spec"`
}

// SecretFromInstanceArray converts a []*ResourceInstance to a []*Secret
func SecretFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Secret, error) {
	newArray := make([]*Secret, 0)
	for _, item := range fromArray {
		res := &Secret{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Secret, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Secret to a ResourceInstance
func (res *Secret) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = SecretGVK()
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

// FromInstance converts a ResourceInstance to a Secret
func (res *Secret) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *Secret) MarshalJSON() ([]byte, error) {
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
func (res *Secret) UnmarshalJSON(data []byte) error {
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
func (res *Secret) PluralName() string {
	return SecretResourceName
}
