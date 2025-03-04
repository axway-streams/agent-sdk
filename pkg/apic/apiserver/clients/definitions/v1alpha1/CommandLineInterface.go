/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	m "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/definitions/v1alpha1"
)

type CommandLineInterfaceMergeFunc func(*m.CommandLineInterface, *m.CommandLineInterface) (*m.CommandLineInterface, error)

// CommandLineInterfaceMerge builds a merge option for an update operation
func CommandLineInterfaceMerge(f CommandLineInterfaceMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &m.CommandLineInterface{}, &m.CommandLineInterface{}

		switch t := prev.(type) {
		case *m.CommandLineInterface:
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
		case *m.CommandLineInterface:
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

// CommandLineInterfaceClient - rest client for CommandLineInterface resources that have a defined resource scope
type CommandLineInterfaceClient struct {
	client v1.Scoped
}

// UnscopedCommandLineInterfaceClient - rest client for CommandLineInterface resources that do not have a defined scope
type UnscopedCommandLineInterfaceClient struct {
	client v1.Unscoped
}

// NewCommandLineInterfaceClient - creates a client that is not scoped to any resource
func NewCommandLineInterfaceClient(c v1.Base) (*UnscopedCommandLineInterfaceClient, error) {

	client, err := c.ForKind(m.CommandLineInterfaceGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedCommandLineInterfaceClient{client}, nil

}

// WithScope - sets the resource scope for the client
func (c *UnscopedCommandLineInterfaceClient) WithScope(scope string) *CommandLineInterfaceClient {
	return &CommandLineInterfaceClient{
		c.client.WithScope(scope),
	}
}

// Get - gets a resource by name
func (c *UnscopedCommandLineInterfaceClient) Get(name string) (*m.CommandLineInterface, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.CommandLineInterface{}
	service.FromInstance(ri)

	return service, nil
}

// Update - updates a resource
func (c *UnscopedCommandLineInterfaceClient) Update(res *m.CommandLineInterface, opts ...v1.UpdateOption) (*m.CommandLineInterface, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.CommandLineInterface{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List - gets a list of resources
func (c *CommandLineInterfaceClient) List(options ...v1.ListOptions) ([]*m.CommandLineInterface, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*m.CommandLineInterface, len(riList))

	for i := range riList {
		result[i] = &m.CommandLineInterface{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get - gets a resource by name
func (c *CommandLineInterfaceClient) Get(name string) (*m.CommandLineInterface, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &m.CommandLineInterface{}
	service.FromInstance(ri)

	return service, nil
}

// Delete - deletes a resource
func (c *CommandLineInterfaceClient) Delete(res *m.CommandLineInterface) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create - creates a resource
func (c *CommandLineInterfaceClient) Create(res *m.CommandLineInterface, opts ...v1.CreateOption) (*m.CommandLineInterface, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &m.CommandLineInterface{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update - updates a resource
func (c *CommandLineInterfaceClient) Update(res *m.CommandLineInterface, opts ...v1.UpdateOption) (*m.CommandLineInterface, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &m.CommandLineInterface{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
