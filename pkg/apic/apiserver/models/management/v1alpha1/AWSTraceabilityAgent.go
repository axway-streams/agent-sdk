/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AWSTraceabilityAgentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "AWSTraceabilityAgent",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	AWSTraceabilityAgentScope = "Environment"

	AWSTraceabilityAgentResource = "awstraceabilityagents"
)

func AWSTraceabilityAgentGVK() apiv1.GroupVersionKind {
	return _AWSTraceabilityAgentGVK
}

func init() {
	apiv1.RegisterGVK(_AWSTraceabilityAgentGVK, AWSTraceabilityAgentScope, AWSTraceabilityAgentResource)
}

// AWSTraceabilityAgent Resource
type AWSTraceabilityAgent struct {
	apiv1.ResourceMeta

	Spec AwsTraceabilityAgentSpec `json:"spec"`

	Status AwsTraceabilityAgentStatus `json:"status"`
}

// FromInstance converts a ResourceInstance to a AWSTraceabilityAgent
func (res *AWSTraceabilityAgent) FromInstance(ri *apiv1.ResourceInstance) error {
	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &AwsTraceabilityAgentSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = AWSTraceabilityAgent{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a AWSTraceabilityAgent to a ResourceInstance
func (res *AWSTraceabilityAgent) AsInstance() (*apiv1.ResourceInstance, error) {
	m, err := json.Marshal(res.Spec)
	if err != nil {
		return nil, err
	}

	spec := map[string]interface{}{}
	err = json.Unmarshal(m, &spec)
	if err != nil {
		return nil, err
	}

	return &apiv1.ResourceInstance{ResourceMeta: res.ResourceMeta, Spec: spec}, nil
}
