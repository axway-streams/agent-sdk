/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AmplifyOAuth2AuthRuleGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "AmplifyOAuth2AuthRule",
		},
		APIVersion: "v1alpha1",
	}

	AmplifyOAuth2AuthRuleScopes = []string{"VirtualAPI", "VirtualAPIRelease"}
)

const AmplifyOAuth2AuthRuleResourceName = "amplifyoauth2authrules"

func AmplifyOAuth2AuthRuleGVK() apiv1.GroupVersionKind {
	return _AmplifyOAuth2AuthRuleGVK
}

func init() {
	apiv1.RegisterGVK(_AmplifyOAuth2AuthRuleGVK, AmplifyOAuth2AuthRuleScopes[0], AmplifyOAuth2AuthRuleResourceName)
}

// AmplifyOAuth2AuthRule Resource
type AmplifyOAuth2AuthRule struct {
	apiv1.ResourceMeta

	Owner *apiv1.Owner `json:"owner"`

	Spec AmplifyOAuth2AuthRuleSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a AmplifyOAuth2AuthRule
func (res *AmplifyOAuth2AuthRule) FromInstance(ri *apiv1.ResourceInstance) error {
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

// AmplifyOAuth2AuthRuleFromInstanceArray converts a []*ResourceInstance to a []*AmplifyOAuth2AuthRule
func AmplifyOAuth2AuthRuleFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AmplifyOAuth2AuthRule, error) {
	newArray := make([]*AmplifyOAuth2AuthRule, 0)
	for _, item := range fromArray {
		res := &AmplifyOAuth2AuthRule{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AmplifyOAuth2AuthRule, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AmplifyOAuth2AuthRule to a ResourceInstance
func (res *AmplifyOAuth2AuthRule) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = AmplifyOAuth2AuthRuleGVK()
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
