/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_TraceabilityAgentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "TraceabilityAgent",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	TraceabilityAgentScope = "Environment"

	TraceabilityAgentResourceName = "traceabilityagents"
)

func TraceabilityAgentGVK() apiv1.GroupVersionKind {
	return _TraceabilityAgentGVK
}

func init() {
	apiv1.RegisterGVK(_TraceabilityAgentGVK, TraceabilityAgentScope, TraceabilityAgentResourceName)
}

// TraceabilityAgent Resource
type TraceabilityAgent struct {
	apiv1.ResourceMeta

	Owner *apiv1.Owner `json:"owner"`

	Spec TraceabilityAgentSpec `json:"spec"`

	Status TraceabilityAgentStatus `json:"status"`
}

// FromInstance converts a ResourceInstance to a TraceabilityAgent
func (res *TraceabilityAgent) FromInstance(ri *apiv1.ResourceInstance) error {
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

// TraceabilityAgentFromInstanceArray converts a []*ResourceInstance to a []*TraceabilityAgent
func TraceabilityAgentFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*TraceabilityAgent, error) {
	newArray := make([]*TraceabilityAgent, 0)
	for _, item := range fromArray {
		res := &TraceabilityAgent{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*TraceabilityAgent, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a TraceabilityAgent to a ResourceInstance
func (res *TraceabilityAgent) AsInstance() (*apiv1.ResourceInstance, error) {
	meta := res.ResourceMeta
	meta.GroupVersionKind = TraceabilityAgentGVK()
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
