/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AmplifyJWTAuthRuleGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "AmplifyJWTAuthRule",
		},
		APIVersion: "v1alpha1",
	}

	AmplifyJWTAuthRuleScopes = []string{"VirtualAPI", "VirtualAPIRelease"}
)

const AmplifyJWTAuthRuleResourceName = "amplifyjwtauthrules"

func AmplifyJWTAuthRuleGVK() apiv1.GroupVersionKind {
	return _AmplifyJWTAuthRuleGVK
}

func init() {
	apiv1.RegisterGVK(_AmplifyJWTAuthRuleGVK, AmplifyJWTAuthRuleScopes[0], AmplifyJWTAuthRuleResourceName)
}

// AmplifyJWTAuthRule Resource
type AmplifyJWTAuthRule struct {
	apiv1.ResourceMeta
	Owner *apiv1.Owner `json:"owner"`
	Spec  interface{}  `json:"spec"`
}

// AmplifyJWTAuthRuleFromInstanceArray converts a []*ResourceInstance to a []*AmplifyJWTAuthRule
func AmplifyJWTAuthRuleFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AmplifyJWTAuthRule, error) {
	newArray := make([]*AmplifyJWTAuthRule, 0)
	for _, item := range fromArray {
		res := &AmplifyJWTAuthRule{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AmplifyJWTAuthRule, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AmplifyJWTAuthRule to a ResourceInstance
func (res *AmplifyJWTAuthRule) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = AmplifyJWTAuthRuleGVK()
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

// FromInstance converts a ResourceInstance to a AmplifyJWTAuthRule
func (res *AmplifyJWTAuthRule) FromInstance(ri *apiv1.ResourceInstance) error {
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
func (res *AmplifyJWTAuthRule) MarshalJSON() ([]byte, error) {
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
func (res *AmplifyJWTAuthRule) UnmarshalJSON(data []byte) error {
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
func (res *AmplifyJWTAuthRule) PluralName() string {
	return AmplifyJWTAuthRuleResourceName
}
