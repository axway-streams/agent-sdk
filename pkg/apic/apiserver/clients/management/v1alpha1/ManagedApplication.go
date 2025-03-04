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

type ManagedApplicationMergeFunc func(*m.ManagedApplication, *m.ManagedApplication) (*m.ManagedApplication, error)

// ManagedApplicationMerge builds a merge option for an update operation
func ManagedApplicationMerge(f ManagedApplicationMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.ManagedApplication{}, &m.ManagedApplication{}

		switch t := prev.(type) {
		case *m.ManagedApplication:
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
		case *m.ManagedApplication:
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

// ManagedApplicationClient - rest client for ManagedApplication resources that have a defined resource scope
type ManagedApplicationClient struct {
	client v1.Scoped
}

// UnscopedManagedApplicationClient - rest client for ManagedApplication resources that do not have a defined scope
type UnscopedManagedApplicationClient struct {
	client v1.Unscoped
}

// NewManagedApplicationClient - creates a client that is not scoped to any resource
func NewManagedApplicationClient(c v1.Base) (*UnscopedManagedApplicationClient, error) {

	client, err := c.ForKind(m.ManagedApplicationGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedManagedApplicationClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedManagedApplicationClient) WithScope(scope string) *ManagedApplicationClient {
	return &ManagedApplicationClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedManagedApplicationClient) Get(name string) (*m.ManagedApplication, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.ManagedApplication{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedManagedApplicationClient) Update(res *m.ManagedApplication, opts ...v1.UpdateOption) (*m.ManagedApplication, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.ManagedApplication{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *ManagedApplicationClient) List(options ...v1.ListOptions) ([]*m.ManagedApplication, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.ManagedApplication, len(riList))

	for i := range riList {
		result[i] = &m.ManagedApplication{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *ManagedApplicationClient) Get(name string) (*m.ManagedApplication, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.ManagedApplication{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *ManagedApplicationClient) Delete(res *m.ManagedApplication) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *ManagedApplicationClient) Create(res *m.ManagedApplication, opts ...v1.CreateOption) (*m.ManagedApplication, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.ManagedApplication{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *ManagedApplicationClient) Update(res *m.ManagedApplication, opts ...v1.UpdateOption) (*m.ManagedApplication, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.ManagedApplication{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
