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

type SecretMergeFunc func(*m.Secret, *m.Secret) (*m.Secret, error)

// SecretMerge builds a merge option for an update operation
func SecretMerge(f SecretMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.Secret{}, &m.Secret{}

		switch t := prev.(type) {
		case *m.Secret:
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
		case *m.Secret:
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

// SecretClient - rest client for Secret resources that have a defined resource scope
type SecretClient struct {
	client v1.Scoped
}

// UnscopedSecretClient - rest client for Secret resources that do not have a defined scope
type UnscopedSecretClient struct {
	client v1.Unscoped
}

// NewSecretClient - creates a client that is not scoped to any resource
func NewSecretClient(c v1.Base) (*UnscopedSecretClient, error) {

	client, err := c.ForKind(m.SecretGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedSecretClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedSecretClient) WithScope(scope string) *SecretClient {
	return &SecretClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedSecretClient) Get(name string) (*m.Secret, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Secret{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedSecretClient) Update(res *m.Secret, opts ...v1.UpdateOption) (*m.Secret, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Secret{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *SecretClient) List(options ...v1.ListOptions) ([]*m.Secret, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.Secret, len(riList))

	for i := range riList {
		result[i] = &m.Secret{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *SecretClient) Get(name string) (*m.Secret, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Secret{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *SecretClient) Delete(res *m.Secret) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *SecretClient) Create(res *m.Secret, opts ...v1.CreateOption) (*m.Secret, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.Secret{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *SecretClient) Update(res *m.Secret, opts ...v1.UpdateOption) (*m.Secret, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Secret{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
