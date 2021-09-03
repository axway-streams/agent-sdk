package apic

import (
	"encoding/json"
	"fmt"
	"net/http"

	coreapi "github.com/Axway/agent-sdk/pkg/api"
	management "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
	agenterrors "github.com/Axway/agent-sdk/pkg/util/errors"
)

// SubscriptionState - Type definition for subscription state
type SubscriptionState string

// SubscriptionState
const (
	SubscriptionApproved              = AccessRequestProvisioning
	SubscriptionRequested             = AccessRequestFailedProvisioning
	SubscriptionRejected              = AccessRequestFailedProvisioning
	SubscriptionActive                = AccessRequestProvisioned
	SubscriptionUnsubscribed          = AccessRequestDeprovisioned
	SubscriptionUnsubscribeInitiated  = AccessRequestDeprovisioning
	SubscriptionFailedToSubscribe     = AccessRequestFailedProvisioning
	SubscriptionFailedToUnsubscribe   = AccessRequestFailedDeprovisioning
	AccessRequestProvisioning         = SubscriptionState("provisioning")
	AccessRequestProvisioned          = SubscriptionState("provisioned")
	AccessRequestFailedProvisioning   = SubscriptionState("failedProvisioning")
	AccessRequestDeprovisioning       = SubscriptionState("deprovisioning")
	AccessRequestDeprovisioned        = SubscriptionState("deprovisioned")
	AccessRequestFailedDeprovisioning = SubscriptionState("failedDeprovisioning")
)

const (
	appNameKey              = "appName"
	subscriptionAppNameType = "string"
	profileKey              = "profile"
)

// Subscription -
type Subscription interface {
	GetID() string
	GetName() string
	GetApicID() string
	GetRemoteAPIAttributes() map[string]string
	GetRemoteAPIID() string
	GetRemoteAPIStage() string
	GetCatalogItemID() string
	GetCreatedUserID() string
	GetState() SubscriptionState
	GetPropertyValue(propertyKey string) string
	UpdateState(newState SubscriptionState, description string) error
	UpdateStateWithProperties(newState SubscriptionState, description string, properties map[string]interface{}) error
	UpdateEnumProperty(key, value, dataType string) error
	UpdateProperties(appName string) error
	UpdatePropertyValues(values map[string]interface{}) error
}

// CentralSubscription -
type CentralSubscription struct {
	AccessRequest       *management.AccessRequest `json:"accessRequest"`
	ApicID              string                    `json:"-"`
	RemoteAPIID         string                    `json:"-"`
	RemoteAPIStage      string                    `json:"-"`
	apicClient          *ServiceClient
	RemoteAPIAttributes map[string]string
}

// GetRemoteAPIAttributes - Returns the attributes from the API that the subscription is tied to.
func (s *CentralSubscription) GetRemoteAPIAttributes() map[string]string {
	return s.RemoteAPIAttributes
}

// GetCreatedUserID - Returns ID of the user that created the subscription
func (s *CentralSubscription) GetCreatedUserID() string {
	return s.AccessRequest.Metadata.Audit.CreateUserID
}

// GetID - Returns ID of the subscription
func (s *CentralSubscription) GetID() string {
	return s.AccessRequest.Name
}

// GetName - Returns Name of the subscription
func (s *CentralSubscription) GetName() string {
	return s.AccessRequest.Name
}

// GetApicID - Returns ID of the Catalog Item or API Service instance
func (s *CentralSubscription) GetApicID() string {
	return s.ApicID
}

// GetRemoteAPIID - Returns ID of the API on remote gateway
func (s *CentralSubscription) GetRemoteAPIID() string {
	return s.RemoteAPIID
}

// GetRemoteAPIStage - Returns the stage name of the API on remote gateway
func (s *CentralSubscription) GetRemoteAPIStage() string {
	return s.RemoteAPIStage
}

// GetCatalogItemID - Returns ID of the Catalog Item
func (s *CentralSubscription) GetCatalogItemID() string {
	return s.AccessRequest.Spec.ApiServiceInstance
}

// GetState - Returns subscription state
func (s *CentralSubscription) GetState() SubscriptionState {
	return SubscriptionState(s.AccessRequest.State.Name)
}

// GetPropertyValue - Returns subscription Property value based on the key
func (s *CentralSubscription) GetPropertyValue(propertyKey string) string {
	if value, found := s.AccessRequest.Spec.Data[propertyKey]; found {
		return value.(string)
	}
	return ""
}

func (s *CentralSubscription) updateProperties(properties map[string]interface{}) error {
	if len(properties) == 0 {
		return nil
	}

	attributes := s.AccessRequest.Spec.Data

	// override with new values
	for k, v := range properties {
		attributes[k] = v
	}

	return s.updatePropertyValue(profileKey, attributes)
}

// UpdateStateWithProperties - Updates the state of subscription
func (s *CentralSubscription) UpdateStateWithProperties(newState SubscriptionState, description string, properties map[string]interface{}) error {
	headers, err := s.getServiceClient().createHeader()
	if err != nil {
		return err
	}

	subStateURL := s.getServiceClient().cfg.GetCatalogItemSubscriptionStatesURL(s.AccessRequest.Name)
	s.AccessRequest.State = management.AccessRequestState{
		Message: description,
		Name:    string(newState),
	}

	stateBody, err := json.Marshal(s.AccessRequest)
	if err != nil {
		return err
	}

	fmt.Println(string(stateBody))
	request := coreapi.Request{
		Method:      coreapi.PUT,
		URL:         subStateURL,
		QueryParams: nil,
		Headers:     headers,
		Body:        stateBody,
	}

	if err = s.updateProperties(properties); err != nil {
		return err
	}

	response, err := s.getServiceClient().apiClient.Send(request)
	if err != nil {
		return agenterrors.Wrap(ErrSubscriptionQuery, err.Error())
	}
	if !(response.Code == http.StatusOK || response.Code == http.StatusCreated) {
		readResponseErrors(response.Code, response.Body)
		return ErrSubscriptionResp.FormatError(response.Code)
	}
	return nil
}

// UpdateState - Updates the state of subscription
func (s *CentralSubscription) UpdateState(newState SubscriptionState, description string) error {
	return s.UpdateStateWithProperties(newState, description, map[string]interface{}{})
}

// getServiceClient - returns the apic client
func (s *CentralSubscription) getServiceClient() *ServiceClient {
	return s.apicClient
}

// getSubscriptions -
func (c *ServiceClient) getSubscriptions(states []string) ([]CentralSubscription, error) {
	queryParams := make(map[string]string)

	// searchQuery := ""
	// for _, state := range states {
	// 	if searchQuery != "" {
	// 		searchQuery += ","
	// 	}
	// 	searchQuery += "state.name==" + state
	// }

	// queryParams["query"] = searchQuery
	subs, err := c.sendSubscriptionsRequest(c.cfg.GetSubscriptionURL(), queryParams)
	if err != nil {
		return subs, err
	}

	matchingSubs := make([]CentralSubscription, 0)
	for _, sub := range subs {
		for _, state := range states {
			if state == sub.AccessRequest.State.Name {
				matchingSubs = append(matchingSubs, sub)
			}
		}
	}
	return matchingSubs, nil
}

func (c *ServiceClient) sendSubscriptionsRequest(url string, queryParams map[string]string) ([]CentralSubscription, error) {
	headers, err := c.createHeader()
	if err != nil {
		return nil, err
	}

	request := coreapi.Request{
		Method:      coreapi.GET,
		URL:         url,
		QueryParams: queryParams,
		Headers:     headers,
		Body:        nil,
	}

	response, err := c.apiClient.Send(request)
	if err != nil {
		return nil, agenterrors.Wrap(ErrSubscriptionQuery, err.Error())
	}
	if response.Code != http.StatusOK && response.Code != http.StatusNotFound {
		readResponseErrors(response.Code, response.Body)
		return nil, ErrSubscriptionResp.FormatError(response.Code)
	}

	subscriptions := make([]management.AccessRequest, 0)
	json.Unmarshal(response.Body, &subscriptions)

	// build the CentralSubscriptions from the UC ones
	centralSubscriptions := make([]CentralSubscription, 0)
	for i := range subscriptions {
		_ = i
		sub := CentralSubscription{
			AccessRequest: &subscriptions[i],
			apicClient:    c,
		}
		centralSubscriptions = append(centralSubscriptions, sub)
	}
	return centralSubscriptions, nil
}

// UpdateEnumProperty -
func (s *CentralSubscription) UpdateEnumProperty(key, newValue, dataType string) error {
	catalogItemID := s.GetCatalogItemID()

	// First need to get the subscriptionDefProperties for the catalog item
	ss, err := s.getServiceClient().GetSubscriptionDefinitionPropertiesForCatalogItem(catalogItemID, profileKey)
	if ss == nil || err != nil {
		return agenterrors.Wrap(ErrGetSubscriptionDefProperties, err.Error())
	}

	// update the appName in the enum
	prop := ss.GetProperty(key)

	// first check that the property is unique
	for _, ele := range prop.Enum {
		if ele == newValue {
			return nil
		}
	}
	newOptions := append(prop.Enum, newValue)

	ss.AddProperty(key, dataType, prop.Description, "", true, newOptions)
	// note: there will be a small time window where the enum items might be out-of-order. The agent will eventually
	// pick up the changes and update the schema, which will reorder them.

	// update the the subscriptionDefProperties for the catalog item. This MUST be done before updating the subscription
	err = s.getServiceClient().UpdateSubscriptionDefinitionPropertiesForCatalogItem(catalogItemID, profileKey, ss)
	if err != nil {
		return agenterrors.Wrap(ErrUpdateSubscriptionDefProperties, err.Error())
	}

	return nil
}

// UpdateProperties -
func (s *CentralSubscription) UpdateProperties(appName string) error {
	err := s.UpdateEnumProperty(appNameKey, appName, subscriptionAppNameType)
	if err != nil {
		return err
	}

	// Now we can update the appname in the subscription itself
	err = s.updatePropertyValue(profileKey, map[string]interface{}{appNameKey: appName})
	if err != nil {
		return agenterrors.Wrap(ErrUpdateSubscriptionDefProperties, err.Error())
	}

	return nil
}

// UpdatePropertyValue - Updates the property value of the subscription
func (s *CentralSubscription) updatePropertyValue(propertyKey string, value map[string]interface{}) error {
	headers, err := s.getServiceClient().createHeader()
	if err != nil {
		return err
	}

	url := s.getServiceClient().cfg.GetCatalogItemSubscriptionPropertiesURL(s.AccessRequest.Name)
	s.AccessRequest.Spec.Data = value
	body, err := json.Marshal(s.AccessRequest)
	if err != nil {
		return err
	}

	request := coreapi.Request{
		Method:  coreapi.PUT,
		URL:     url,
		Headers: headers,
		Body:    body,
	}

	response, err := s.getServiceClient().apiClient.Send(request)
	if err != nil {
		return err
	}

	if !(response.Code == http.StatusOK) {
		readResponseErrors(response.Code, response.Body)
		return ErrSubscriptionResp.FormatError(response.Code)
	}
	return nil
}

// UpdatePropertyValues - Updates the property values of the subscription
func (s *CentralSubscription) UpdatePropertyValues(values map[string]interface{}) error {
	// headers, err := s.getServiceClient().createHeader()
	// if err != nil {
	// 	return err
	// }

	// url := fmt.Sprintf("%s/%s", s.getServiceClient().cfg.GetCatalogItemSubscriptionPropertiesURL(s.GetCatalogItemID(), s.GetID()), profileKey)
	// body, err := json.Marshal(values)
	// if err != nil {
	// 	return err
	// }

	// request := coreapi.Request{
	// 	Method:  coreapi.PUT,
	// 	URL:     url,
	// 	Headers: headers,
	// 	Body:    body,
	// }

	// response, err := s.getServiceClient().apiClient.Send(request)
	// if err != nil {
	// 	return err
	// }

	// if !(response.Code == http.StatusOK) {
	// 	readResponseErrors(response.Code, response.Body)
	// 	return ErrSubscriptionResp.FormatError(response.Code)
	// }
	return nil
}
