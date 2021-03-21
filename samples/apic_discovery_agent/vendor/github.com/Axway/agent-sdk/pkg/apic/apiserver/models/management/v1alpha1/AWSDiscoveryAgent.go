/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AWSDiscoveryAgentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "AWSDiscoveryAgent",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	AWSDiscoveryAgentScope = "Environment"

	AWSDiscoveryAgentResource = "awsdiscoveryagents"
)

func AWSDiscoveryAgentGVK() apiv1.GroupVersionKind {
	return _AWSDiscoveryAgentGVK
}

func init() {
	apiv1.RegisterGVK(_AWSDiscoveryAgentGVK, AWSDiscoveryAgentScope, AWSDiscoveryAgentResource)
}

// AWSDiscoveryAgent Resource
type AWSDiscoveryAgent struct {
	apiv1.ResourceMeta

	Spec AwsDiscoveryAgentSpec `json:"spec"`

	Status AwsDiscoveryAgentStatus `json:"status"`
}

// FromInstance converts a ResourceInstance to a AWSDiscoveryAgent
func (res *AWSDiscoveryAgent) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &AwsDiscoveryAgentSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = AWSDiscoveryAgent{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a AWSDiscoveryAgent to a ResourceInstance
func (res *AWSDiscoveryAgent) AsInstance() (*apiv1.ResourceInstance, error) {
	m, err := json.Marshal(res.Spec)
	if err != nil {
		return nil, err
	}

	spec := map[string]interface{}{}
	err = json.Unmarshal(m, &spec)
	if err != nil {
		return nil, err
	}

	meta := res.ResourceMeta
	meta.GroupVersionKind = AWSDiscoveryAgentGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
