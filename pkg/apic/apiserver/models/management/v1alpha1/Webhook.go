/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_WebhookGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "Webhook",
		},
		APIVersion: "v1alpha1",
	}

	WebhookScopes = []string{"Environment", "Integration"}
)

const WebhookResourceName = "webhooks"

func WebhookGVK() apiv1.GroupVersionKind {
	return _WebhookGVK
}

func init() {
	apiv1.RegisterGVK(_WebhookGVK, WebhookScopes[0], WebhookResourceName)
}

// Webhook Resource
type Webhook struct {
	apiv1.ResourceMeta

	Owner *apiv1.Owner `json:"owner"`

	Spec WebhookSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a Webhook
func (res *Webhook) FromInstance(ri *apiv1.ResourceInstance) error {
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

// WebhookFromInstanceArray converts a []*ResourceInstance to a []*Webhook
func WebhookFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*Webhook, error) {
	newArray := make([]*Webhook, 0)
	for _, item := range fromArray {
		res := &Webhook{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*Webhook, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a Webhook to a ResourceInstance
func (res *Webhook) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = WebhookGVK()
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
