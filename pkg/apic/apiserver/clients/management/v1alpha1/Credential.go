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

type CredentialMergeFunc func(*m.Credential, *m.Credential) (*m.Credential, error)

// CredentialMerge builds a merge option for an update operation
func CredentialMerge(f CredentialMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.Credential{}, &m.Credential{}

		switch t := prev.(type) {
		case *m.Credential:
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
		case *m.Credential:
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

// CredentialClient - rest client for Credential resources that have a defined resource scope
type CredentialClient struct {
	client v1.Scoped
}

// UnscopedCredentialClient - rest client for Credential resources that do not have a defined scope
type UnscopedCredentialClient struct {
	client v1.Unscoped
}

// NewCredentialClient - creates a client that is not scoped to any resource
func NewCredentialClient(c v1.Base) (*UnscopedCredentialClient, error) {

	client, err := c.ForKind(m.CredentialGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedCredentialClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedCredentialClient) WithScope(scope string) *CredentialClient {
	return &CredentialClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedCredentialClient) Get(name string) (*m.Credential, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Credential{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedCredentialClient) Update(res *m.Credential, opts ...v1.UpdateOption) (*m.Credential, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Credential{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *CredentialClient) List(options ...v1.ListOptions) ([]*m.Credential, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.Credential, len(riList))

	for i := range riList {
		result[i] = &m.Credential{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *CredentialClient) Get(name string) (*m.Credential, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Credential{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *CredentialClient) Delete(res *m.Credential) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *CredentialClient) Create(res *m.Credential, opts ...v1.CreateOption) (*m.Credential, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.Credential{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *CredentialClient) Update(res *m.Credential, opts ...v1.UpdateOption) (*m.Credential, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Credential{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
