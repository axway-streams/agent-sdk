/*
 * This file is automatically generated
 */

package v1alpha1

import (
	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

// AWSTraceabilityAgentClient -
type AWSTraceabilityAgentClient struct {
	client v1.Scoped
}

// UnscopedAWSTraceabilityAgentClient -
type UnscopedAWSTraceabilityAgentClient struct {
	client v1.Unscoped
}

// NewAWSTraceabilityAgentClient -
func NewAWSTraceabilityAgentClient(c v1.Base) (*UnscopedAWSTraceabilityAgentClient, error) {

	client, err := c.ForKind(v1alpha1.AWSTraceabilityAgentGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedAWSTraceabilityAgentClient{client}, nil

}

// WithScope -
func (c *UnscopedAWSTraceabilityAgentClient) WithScope(scope string) *AWSTraceabilityAgentClient {
	return &AWSTraceabilityAgentClient{
		c.client.WithScope(scope),
	}
}

// List -
func (c *AWSTraceabilityAgentClient) List(options ...v1.ListOptions) ([]*v1alpha1.AWSTraceabilityAgent, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.AWSTraceabilityAgent, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.AWSTraceabilityAgent{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *AWSTraceabilityAgentClient) Get(name string) (*v1alpha1.AWSTraceabilityAgent, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.AWSTraceabilityAgent{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *AWSTraceabilityAgentClient) Delete(res *v1alpha1.AWSTraceabilityAgent) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *AWSTraceabilityAgentClient) Create(res *v1alpha1.AWSTraceabilityAgent) (*v1alpha1.AWSTraceabilityAgent, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.AWSTraceabilityAgent{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *AWSTraceabilityAgentClient) Update(res *v1alpha1.AWSTraceabilityAgent) (*v1alpha1.AWSTraceabilityAgent, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.AWSTraceabilityAgent{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
