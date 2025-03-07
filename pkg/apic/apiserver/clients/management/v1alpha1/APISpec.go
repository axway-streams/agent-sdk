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

type APISpecMergeFunc func(*m.APISpec, *m.APISpec) (*m.APISpec, error)

// APISpecMerge builds a merge option for an update operation
func APISpecMerge(f APISpecMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.APISpec{}, &m.APISpec{}

		switch t := prev.(type) {
		case *m.APISpec:
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
		case *m.APISpec:
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

// APISpecClient - rest client for APISpec resources that have a defined resource scope
type APISpecClient struct {
	client v1.Scoped
}

// UnscopedAPISpecClient - rest client for APISpec resources that do not have a defined scope
type UnscopedAPISpecClient struct {
	client v1.Unscoped
}

// NewAPISpecClient - creates a client that is not scoped to any resource
func NewAPISpecClient(c v1.Base) (*UnscopedAPISpecClient, error) {

	client, err := c.ForKind(m.APISpecGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedAPISpecClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedAPISpecClient) WithScope(scope string) *APISpecClient {
	return &APISpecClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedAPISpecClient) Get(name string) (*m.APISpec, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.APISpec{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedAPISpecClient) Update(res *m.APISpec, opts ...v1.UpdateOption) (*m.APISpec, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.APISpec{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *APISpecClient) List(options ...v1.ListOptions) ([]*m.APISpec, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.APISpec, len(riList))

	for i := range riList {
		result[i] = &m.APISpec{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *APISpecClient) Get(name string) (*m.APISpec, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.APISpec{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *APISpecClient) Delete(res *m.APISpec) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *APISpecClient) Create(res *m.APISpec, opts ...v1.CreateOption) (*m.APISpec, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.APISpec{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *APISpecClient) Update(res *m.APISpec, opts ...v1.UpdateOption) (*m.APISpec, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.APISpec{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
