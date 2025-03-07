/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	m "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/catalog/v1alpha1"
)

type ApplicationMergeFunc func(*m.Application, *m.Application) (*m.Application, error)

// ApplicationMerge builds a merge option for an update operation
func ApplicationMerge(f ApplicationMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.Application{}, &m.Application{}

		switch t := prev.(type) {
		case *m.Application:
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
		case *m.Application:
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

// ApplicationClient - rest client for Application resources that have a defined resource scope
type ApplicationClient struct {
	client v1.Scoped
}

// NewApplicationClient - creates a client scoped to a particular resource
func NewApplicationClient(c v1.Base) (*ApplicationClient, error) {

	client, err := c.ForKind(m.ApplicationGVK())
	if err != nil {
		return nil, err
	}

	return &ApplicationClient{client}, nil

}

// List - gets a list of resources
func (c *ApplicationClient) List(options ...v1.ListOptions) ([]*m.Application, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.Application, len(riList))

	for i := range riList {
		result[i] = &m.Application{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *ApplicationClient) Get(name string) (*m.Application, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.Application{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *ApplicationClient) Delete(res *m.Application) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *ApplicationClient) Create(res *m.Application, opts ...v1.CreateOption) (*m.Application, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.Application{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *ApplicationClient) Update(res *m.Application, opts ...v1.UpdateOption) (*m.Application, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.Application{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
