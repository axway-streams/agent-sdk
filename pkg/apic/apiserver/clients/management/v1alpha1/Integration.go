/*
 * This file is automatically generated
 */

package v1alpha1

import (
	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

// IntegrationClient -
type IntegrationClient struct {
	client v1.Scoped
}

// NewIntegrationClient -
func NewIntegrationClient(c v1.Base) (*IntegrationClient, error) {

	client, err := c.ForKind(v1alpha1.IntegrationGVK())
	if err != nil {
		return nil, err
	}

	return &IntegrationClient{client}, nil

}

// List -
func (c *IntegrationClient) List(options ...v1.ListOptions) ([]*v1alpha1.Integration, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.Integration, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.Integration{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *IntegrationClient) Get(name string) (*v1alpha1.Integration, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.Integration{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *IntegrationClient) Delete(res *v1alpha1.Integration) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *IntegrationClient) Create(res *v1alpha1.Integration) (*v1alpha1.Integration, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.Integration{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *IntegrationClient) Update(res *v1alpha1.Integration) (*v1alpha1.Integration, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.Integration{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
