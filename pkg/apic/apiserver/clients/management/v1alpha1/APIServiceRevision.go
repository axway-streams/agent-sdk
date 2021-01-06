/*
 * This file is automatically generated
 */

package v1alpha1

import (
	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

// APIServiceRevisionClient -
type APIServiceRevisionClient struct {
	client v1.Scoped
}

// UnscopedAPIServiceRevisionClient -
type UnscopedAPIServiceRevisionClient struct {
	client v1.Unscoped
}

// NewAPIServiceRevisionClient -
func NewAPIServiceRevisionClient(c v1.Base) (*UnscopedAPIServiceRevisionClient, error) {

	client, err := c.ForKind(v1alpha1.APIServiceRevisionGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedAPIServiceRevisionClient{client}, nil

}

// WithScope -
func (c *UnscopedAPIServiceRevisionClient) WithScope(scope string) *APIServiceRevisionClient {
	return &APIServiceRevisionClient{
		c.client.WithScope(scope),
	}
}

// List -
func (c *APIServiceRevisionClient) List(options ...v1.ListOptions) ([]*v1alpha1.APIServiceRevision, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.APIServiceRevision, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.APIServiceRevision{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *APIServiceRevisionClient) Get(name string) (*v1alpha1.APIServiceRevision, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.APIServiceRevision{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *APIServiceRevisionClient) Delete(res *v1alpha1.APIServiceRevision) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *APIServiceRevisionClient) Create(res *v1alpha1.APIServiceRevision) (*v1alpha1.APIServiceRevision, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.APIServiceRevision{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *APIServiceRevisionClient) Update(res *v1alpha1.APIServiceRevision) (*v1alpha1.APIServiceRevision, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.APIServiceRevision{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
