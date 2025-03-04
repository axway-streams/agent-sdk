package apic

// TODO - this file should be able to be removed once Unified Catalog support has been removed
import (
	"encoding/json"
	"net/http"

	coreapi "github.com/Axway/agent-sdk/pkg/api"
	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/util"
	agenterrors "github.com/Axway/agent-sdk/pkg/util/errors"
)

// SubscriptionSchema -
type SubscriptionSchema interface {
	AddProperty(name, dataType, description, apicRefField string, isRequired bool, enums []string)
	GetProperty(name string) *SubscriptionSchemaPropertyDefinition
	AddUniqueKey(keyName string)
	GetSubscriptionName() string
	SetSubscriptionName(name string)
	mapStringInterface() (map[string]interface{}, error)
	rawJSON() (json.RawMessage, error)
	SetJSONDraft07SchemaVersion()
}

// AnyOfSubscriptionSchemaPropertyDefinitions - used for items of SubscriptionSchemaPropertyDefinition
type AnyOfSubscriptionSchemaPropertyDefinitions struct {
	AnyOf []SubscriptionSchemaPropertyDefinition `json:"anyOf,omitempty"`
}

// SubscriptionSchemaPropertyDefinition -
type SubscriptionSchemaPropertyDefinition struct {
	Type               string                                          `json:"type"`
	Title              string                                          `json:"title,omitempty"`
	Description        string                                          `json:"description,omitempty"`
	DefaultValue       interface{}                                     `json:"default,omitempty"`
	Enum               []string                                        `json:"enum,omitempty"`
	ReadOnly           bool                                            `json:"readOnly,omitempty"`
	Format             string                                          `json:"format,omitempty"`
	Properties         map[string]SubscriptionSchemaPropertyDefinition `json:"properties,omitempty"`
	RequiredProperties []string                                        `json:"required,omitempty"`
	Items              *AnyOfSubscriptionSchemaPropertyDefinitions     `json:"items,omitempty"`    // We use a pointer to avoid generating an empty struct if not set
	MinItems           *uint                                           `json:"minItems,omitempty"` // We use a pointer to differentiate the "blank value" from a choosen 0 min value
	MaxItems           *uint                                           `json:"maxItems,omitempty"` // We use a pointer to differentiate the "blank value" from a choosen 0 min value
	Minimum            *float64                                        `json:"minimum,omitempty"`  // We use a pointer to differentiate the "blank value" from a choosen 0 min value
	Maximum            *float64                                        `json:"maximum,omitempty"`  // We use a pointer to differentiate the "blank value" from a choosen 0 max value
	APICRef            string                                          `json:"x-axway-ref-apic,omitempty"`
	Name               string                                          `json:"-"`
	Required           bool                                            `json:"-"`
	//Pattern            string                                          `json:"pattern,omitempty"`
}

type subscriptionSchema struct {
	SubscriptionName  string                                          `json:"-"`
	SchemaType        string                                          `json:"type"`
	SchemaVersion     string                                          `json:"$schema"`
	SchemaDescription string                                          `json:"description"`
	Properties        map[string]SubscriptionSchemaPropertyDefinition `json:"properties"`
	Required          []string                                        `json:"required,omitempty"`
	UniqueKeys        []string                                        `json:"x-axway-unique-keys,omitempty"`
}

// NewSubscriptionSchema -
func NewSubscriptionSchema(name string) SubscriptionSchema {
	return &subscriptionSchema{
		SubscriptionName:  name,
		SchemaType:        "object",
		SchemaVersion:     "http://json-schema.org/draft-04/schema#",
		SchemaDescription: "Subscription specification for authentication",
		Properties:        make(map[string]SubscriptionSchemaPropertyDefinition),
		Required:          make([]string, 0),
	}
}

//SetJSONDraft07SchemaVersion - set the JSON schema for the subscription definition to Draft-07
func (ss *subscriptionSchema) SetJSONDraft07SchemaVersion() {
	ss.SchemaVersion = "http://json-schema.org/draft-07/schema#"
}

// AddProperty -
func (ss *subscriptionSchema) AddProperty(name, dataType, description, apicRefField string, isRequired bool, enums []string) {
	newProp := SubscriptionSchemaPropertyDefinition{
		Type:        dataType,
		Title:       name,
		Description: description,
		APICRef:     apicRefField,
	}

	if len(enums) > 0 {
		newProp.Enum = enums
	}
	ss.Properties[name] = newProp

	// required slice can't contain duplicates!
	if isRequired && !util.StringSliceContains(ss.Required, name) {
		ss.Required = append(ss.Required, name)
	}
}

// GetProperty -
func (ss *subscriptionSchema) GetProperty(name string) *SubscriptionSchemaPropertyDefinition {
	if val, ok := ss.Properties[name]; ok {
		return &val
	}
	return nil
}

// GetSubscriptionName -
func (ss *subscriptionSchema) GetSubscriptionName() string {
	return ss.SubscriptionName
}

// SetSubscriptionName -
func (ss *subscriptionSchema) SetSubscriptionName(name string) {
	ss.SubscriptionName = name
}

// AddUniqueKey -
func (ss *subscriptionSchema) AddUniqueKey(keyName string) {
	ss.UniqueKeys = append(ss.UniqueKeys, keyName)
}

// rawJSON -
func (ss *subscriptionSchema) rawJSON() (json.RawMessage, error) {
	schemaBuffer, err := json.Marshal(ss)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(schemaBuffer), nil
}

// mapStringInterface -
func (ss *subscriptionSchema) mapStringInterface() (map[string]interface{}, error) {
	schemaBuffer, err := json.Marshal(ss)
	if err != nil {
		return nil, err
	}

	var stringMap map[string]interface{}
	json.Unmarshal(schemaBuffer, &stringMap)
	if err != nil {
		return nil, err
	}

	return stringMap, nil
}

// RegisterSubscriptionSchema - Adds a new subscription schema for the specified auth type. In publishToEnvironment mode
// creates a API Server resource for subscription definition
func (c *ServiceClient) RegisterSubscriptionSchema(subscriptionSchema SubscriptionSchema, update bool) error {
	c.subscriptionRegistrationLock.Lock()
	defer c.subscriptionRegistrationLock.Unlock()

	return c.registerSubscriptionSchema(subscriptionSchema, update)
}

func (c *ServiceClient) registerSubscriptionSchema(subscriptionSchema SubscriptionSchema, update bool) error {
	var registeredSpecHash uint64
	registeredSchema := c.getCachedSubscriptionSchema(subscriptionSchema.GetSubscriptionName())

	if registeredSchema != nil {
		registeredSpecHash, _ = util.ComputeHash(registeredSchema.Spec)
	} else {
		update = true
	}

	spec, err := c.prepareSubscriptionDefinitionSpec(registeredSchema, subscriptionSchema)
	if err != nil {
		return err
	}

	// Create New definition
	if registeredSchema == nil {
		return c.createSubscriptionSchema(subscriptionSchema.GetSubscriptionName(), spec)
	}

	if update {
		// Check if the schema definitions changed before update
		if currentHash, _ := util.ComputeHash(spec); currentHash != registeredSpecHash {
			return c.updateSubscriptionSchema(subscriptionSchema.GetSubscriptionName(), spec)
		}
	}

	return nil
}

func (c *ServiceClient) getCachedSubscriptionSchema(defName string) *v1alpha1.ConsumerSubscriptionDefinition {
	cachedSchema, err := c.subscriptionSchemaCache.Get(defName)
	if err != nil {
		registeredSchema, _ := c.getSubscriptionSchema(defName)
		if registeredSchema != nil {
			c.subscriptionSchemaCache.Set(defName, registeredSchema)
		}
		return registeredSchema
	}
	return cachedSchema.(*v1alpha1.ConsumerSubscriptionDefinition)
}

func (c *ServiceClient) getSubscriptionSchema(schemaName string) (*v1alpha1.ConsumerSubscriptionDefinition, error) {
	headers, err := c.createHeader()
	if err != nil {
		return nil, err
	}

	request := coreapi.Request{
		Method:  coreapi.GET,
		URL:     c.cfg.GetAPIServerSubscriptionDefinitionURL() + "/" + schemaName,
		Headers: headers,
	}

	response, err := c.apiClient.Send(request)
	if err != nil {
		return nil, err
	}

	if response.Code != http.StatusOK {
		return nil, nil
	}
	registeredSchema := &v1alpha1.ConsumerSubscriptionDefinition{}
	json.Unmarshal(response.Body, registeredSchema)
	return registeredSchema, nil
}

func (c *ServiceClient) createSubscriptionSchema(defName string, spec *v1alpha1.ConsumerSubscriptionDefinitionSpec) error {
	//Add API Server resource - SubscriptionDefinition
	buffer, err := c.marshalSubscriptionDefinition(defName, spec)
	if err != nil {
		return err
	}

	headers, err := c.createHeader()
	if err != nil {
		return err
	}

	request := coreapi.Request{
		Method:  coreapi.POST,
		URL:     c.cfg.GetAPIServerSubscriptionDefinitionURL(),
		Headers: headers,
		Body:    buffer,
	}

	response, err := c.apiClient.Send(request)
	if err != nil {
		return agenterrors.Wrap(ErrSubscriptionSchemaCreate, err.Error())
	}
	if response.Code != http.StatusCreated {
		readResponseErrors(response.Code, response.Body)
		return agenterrors.Wrap(ErrSubscriptionSchemaResp, coreapi.POST).FormatError(response.Code)
	}
	registeredSchema := &v1alpha1.ConsumerSubscriptionDefinition{}
	json.Unmarshal(response.Body, registeredSchema)
	c.subscriptionSchemaCache.Set(defName, registeredSchema)
	return nil
}

func (c *ServiceClient) updateSubscriptionSchema(defName string, spec *v1alpha1.ConsumerSubscriptionDefinitionSpec) error {
	// Add API Server resource - SubscriptionDefinition
	buffer, err := c.marshalSubscriptionDefinition(defName, spec)
	if err != nil {
		return err
	}

	headers, err := c.createHeader()
	if err != nil {
		return err
	}
	request := coreapi.Request{
		Method:  coreapi.PUT,
		URL:     c.cfg.GetAPIServerSubscriptionDefinitionURL() + "/" + defName,
		Headers: headers,
		Body:    buffer,
	}

	response, err := c.apiClient.Send(request)
	if err != nil {
		return agenterrors.Wrap(ErrSubscriptionSchemaCreate, err.Error())
	}
	if !(response.Code == http.StatusOK) {
		readResponseErrors(response.Code, response.Body)
		return agenterrors.Wrap(ErrSubscriptionSchemaResp, coreapi.PUT).FormatError(response.Code)
	}
	registeredSchema := &v1alpha1.ConsumerSubscriptionDefinition{}
	json.Unmarshal(response.Body, registeredSchema)
	c.subscriptionSchemaCache.Set(defName, registeredSchema)
	return nil
}

func (c *ServiceClient) updateAccessRequestSubscriptionSchema(defName string, spec *v1alpha1.AccessRequestDefinitionSpec) error {
	// Add API Server resource - SubscriptionDefinition
	buffer, err := c.marshalAccessRequestSubscriptionDefinition(defName, spec)
	if err != nil {
		return err
	}

	headers, err := c.createHeader()
	if err != nil {
		return err
	}
	request := coreapi.Request{
		Method:  coreapi.PUT,
		URL:     c.cfg.GetAPIServerAccessRequestDefinitionURL() + "/" + defName,
		Headers: headers,
		Body:    buffer,
	}

	response, err := c.apiClient.Send(request)
	if err != nil {
		return agenterrors.Wrap(ErrSubscriptionSchemaCreate, err.Error())
	}
	if !(response.Code == http.StatusOK) {
		readResponseErrors(response.Code, response.Body)
		return agenterrors.Wrap(ErrAccessRequestSubscriptionSchemaResp, coreapi.PUT).FormatError(response.Code)
	}

	registeredSchema := &v1alpha1.AccessRequestDefinition{}
	json.Unmarshal(response.Body, registeredSchema)
	c.subscriptionSchemaCache.Set(defName, registeredSchema)
	return nil
}

// UpdateSubscriptionSchema - Updates a subscription schema in Publish to environment mode
// creates a API Server resource for subscription definition
func (c *ServiceClient) UpdateSubscriptionSchema(subscriptionSchema SubscriptionSchema) error {
	spec, err := c.prepareSubscriptionDefinitionSpec(nil, subscriptionSchema)
	if err != nil {
		return err
	}
	return c.updateSubscriptionSchema(subscriptionSchema.GetSubscriptionName(), spec)
}

// UpdateAccessRequestSubscriptionSchema - Updates an access request subscription schema in Publish to environment mode
// creates a API Server resource for subscription definition
func (c *ServiceClient) UpdateAccessRequestSubscriptionSchema(subscriptionSchema SubscriptionSchema) error {
	spec, err := c.prepareAccessRequestSubscriptionDefinitionSpec(subscriptionSchema)
	if err != nil {
		return err
	}
	return c.updateAccessRequestSubscriptionSchema(subscriptionSchema.GetSubscriptionName(), spec)
}

func (c *ServiceClient) prepareSubscriptionDefinitionSpec(registeredSchema *v1alpha1.ConsumerSubscriptionDefinition, subscriptionSchema SubscriptionSchema) (*v1alpha1.ConsumerSubscriptionDefinitionSpec, error) {
	catalogSubscriptionSchema, err := subscriptionSchema.mapStringInterface()
	if err != nil {
		return nil, err
	}

	webhooks := make([]string, 0)
	// use existing webhooks if present
	if registeredSchema != nil {
		webhooks = registeredSchema.Spec.Webhooks
	}

	if c.cfg.GetSubscriptionConfig().GetSubscriptionApprovalMode() == corecfg.WebhookApproval {
		found := false
		for _, webhook := range webhooks {
			if webhook == DefaultSubscriptionWebhookName {
				found = true
				break
			}
		}
		// Only add the default subscription webhook if it is not there
		if !found {
			webhooks = append(webhooks, DefaultSubscriptionWebhookName)
		}
	}

	return &v1alpha1.ConsumerSubscriptionDefinitionSpec{
		Webhooks: webhooks,
		Schema: v1alpha1.ConsumerSubscriptionDefinitionSpecSchema{
			Properties: []v1alpha1.ConsumerSubscriptionDefinitionSpecSchemaProperties{
				{
					Key:   profileKey,
					Value: catalogSubscriptionSchema,
				},
			},
		},
	}, nil
}

func (c *ServiceClient) prepareAccessRequestSubscriptionDefinitionSpec(subscriptionSchema SubscriptionSchema) (*v1alpha1.AccessRequestDefinitionSpec, error) {
	subscriptionSchema.SetJSONDraft07SchemaVersion()
	schema, err := subscriptionSchema.mapStringInterface()
	if err != nil {
		return nil, err
	}

	return &v1alpha1.AccessRequestDefinitionSpec{
		Schema: schema,
	}, nil
}

func (c *ServiceClient) marshalSubscriptionDefinition(defName string, spec *v1alpha1.ConsumerSubscriptionDefinitionSpec) ([]byte, error) {
	apiServerService := v1alpha1.ConsumerSubscriptionDefinition{
		ResourceMeta: v1.ResourceMeta{
			GroupVersionKind: v1alpha1.ConsumerSubscriptionDefinitionGVK(),
			Name:             defName,
			Title:            "Subscription definition created by agent",
			Attributes:       nil,
			Tags:             nil,
		},
		Spec: *spec,
	}

	return json.Marshal(apiServerService)
}

func (c *ServiceClient) marshalAccessRequestSubscriptionDefinition(defName string, spec *v1alpha1.AccessRequestDefinitionSpec) ([]byte, error) {
	apiServerService := v1alpha1.AccessRequestDefinition{

		ResourceMeta: v1.ResourceMeta{
			GroupVersionKind: v1alpha1.AccessRequestDefinitionGVK(),
			Name:             defName,
			Title:            "Subscription definition created by agent",
			Attributes:       nil,
			Tags:             nil,
		},
		Spec: *spec,
	}

	return json.Marshal(apiServerService)
}

func (c *ServiceClient) getProfilePropValue(subscriptionDef *v1alpha1.ConsumerSubscriptionDefinition) map[string]interface{} {
	for _, prop := range subscriptionDef.Spec.Schema.Properties {
		if prop.Key == profileKey {
			return prop.Value
		}
	}
	return nil
}
