/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/catalog/v1alpha1"
)

type AccessControlListMergeFunc func(*v1alpha1.AccessControlList, *v1alpha1.AccessControlList) (*v1alpha1.AccessControlList, error)

// Merge builds a merge option for an update operation
func AccessControlListMerge(f AccessControlListMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &v1alpha1.AccessControlList{}, &v1alpha1.AccessControlList{}

		switch t := prev.(type) {
		case *v1alpha1.AccessControlList:
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
		case *v1alpha1.AccessControlList:
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

// AccessControlListClient -
type AccessControlListClient struct {
	client v1.Scoped
}

// UnscopedAccessControlListClient -
type UnscopedAccessControlListClient struct {
	client v1.Unscoped
}

// NewAccessControlListClient -
func NewAccessControlListClient(c v1.Base) (*UnscopedAccessControlListClient, error) {

	client, err := c.ForKind(v1alpha1.AccessControlListGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedAccessControlListClient{client}, nil

}

// WithScope -
func (c *UnscopedAccessControlListClient) WithScope(scope string) *AccessControlListClient {
	return &AccessControlListClient{
		c.client.WithScope(scope),
	}
}

// Get -
func (c *UnscopedAccessControlListClient) Get(name string) (*v1alpha1.AccessControlList, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.AccessControlList{}
	service.FromInstance(ri)

	return service, nil
}

// Update -
func (c *UnscopedAccessControlListClient) Update(res *v1alpha1.AccessControlList, opts ...v1.UpdateOption) (*v1alpha1.AccessControlList, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.AccessControlList{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List -
func (c *AccessControlListClient) List(options ...v1.ListOptions) ([]*v1alpha1.AccessControlList, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.AccessControlList, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.AccessControlList{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *AccessControlListClient) Get(name string) (*v1alpha1.AccessControlList, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.AccessControlList{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *AccessControlListClient) Delete(res *v1alpha1.AccessControlList) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *AccessControlListClient) Create(res *v1alpha1.AccessControlList, opts ...v1.CreateOption) (*v1alpha1.AccessControlList, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.AccessControlList{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *AccessControlListClient) Update(res *v1alpha1.AccessControlList, opts ...v1.UpdateOption) (*v1alpha1.AccessControlList, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.AccessControlList{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
