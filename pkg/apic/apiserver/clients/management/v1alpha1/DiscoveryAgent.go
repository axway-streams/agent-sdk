/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	m "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

type DiscoveryAgentMergeFunc func(*m.DiscoveryAgent, *m.DiscoveryAgent) (*m.DiscoveryAgent, error)

// DiscoveryAgentMerge builds a merge option for an update operation
func DiscoveryAgentMerge(f DiscoveryAgentMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.DiscoveryAgent{}, &m.DiscoveryAgent{}

		switch t := prev.(type) {
		case *m.DiscoveryAgent:
			p = t
		case *apiv1.ResourceInstance:
			err := p.FromInstance(t)
			if err != nil {
				return nil, fmt.Errorf("merge: failed to unserialise prev resource: %w", err)
			}
		default:
			return nil, fmt.Errorf("merge: failed to unserialise prev resource, unxexpected resource type: %T", t)
		}

		switch t := new.(type) {
		case *m.DiscoveryAgent:
			n = t
		case *apiv1.ResourceInstance:
			err := n.FromInstance(t)
			if err != nil {
				return nil, fmt.Errorf("merge: failed to unserialize new resource: %w", err)
			}
		default:
			return nil, fmt.Errorf("merge: failed to unserialise new resource, unxexpected resource type: %T", t)
		}

		return f(p, n)
	})
}

// DiscoveryAgentClient - rest client for DiscoveryAgent resources that have a defined resource scope
type DiscoveryAgentClient struct {
	client v1.Scoped
}

// UnscopedDiscoveryAgentClient - rest client for DiscoveryAgent resources that do not have a defined scope
type UnscopedDiscoveryAgentClient struct {
	client v1.Unscoped
}

// NewDiscoveryAgentClient - creates a client that is not scoped to any resource
func NewDiscoveryAgentClient(c v1.Base) (*UnscopedDiscoveryAgentClient, error) {

	client, err := c.ForKind(m.DiscoveryAgentGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedDiscoveryAgentClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedDiscoveryAgentClient) WithScope(scope string) *DiscoveryAgentClient {
	return &DiscoveryAgentClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedDiscoveryAgentClient) Get(name string) (*m.DiscoveryAgent, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.DiscoveryAgent{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedDiscoveryAgentClient) Update(res *m.DiscoveryAgent, opts ...v1.UpdateOption) (*m.DiscoveryAgent, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.DiscoveryAgent{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *DiscoveryAgentClient) List(options ...v1.ListOptions) ([]*m.DiscoveryAgent, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.DiscoveryAgent, len(riList))

	for i := range riList {
		result[i] = &m.DiscoveryAgent{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *DiscoveryAgentClient) Get(name string) (*m.DiscoveryAgent, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.DiscoveryAgent{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *DiscoveryAgentClient) Delete(res *m.DiscoveryAgent) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *DiscoveryAgentClient) Create(res *m.DiscoveryAgent, opts ...v1.CreateOption) (*m.DiscoveryAgent, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.DiscoveryAgent{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *DiscoveryAgentClient) Update(res *m.DiscoveryAgent, opts ...v1.UpdateOption) (*m.DiscoveryAgent, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.DiscoveryAgent{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
