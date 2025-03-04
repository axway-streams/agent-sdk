package mock

import (
	"github.com/Axway/agent-sdk/pkg/apic/provisioning"
	"github.com/Axway/agent-sdk/pkg/authz/oauth"
)

type MockApplicationRequest struct {
	provisioning.ApplicationRequest
	ID       string
	AppName  string
	Details  map[string]string
	TeamName string
}

func (m MockApplicationRequest) GetID() string {
	return m.ID
}

func (m MockApplicationRequest) GetManagedApplicationName() string {
	return m.AppName
}

func (m MockApplicationRequest) GetApplicationDetailsValue(key string) string {
	return m.Details[key]
}
func (m MockApplicationRequest) GetTeamName() string {
	return m.TeamName
}

type MockCredentialRequest struct {
	provisioning.CredentialRequest
	ID          string
	AppDetails  map[string]string
	AppName     string
	CredDefName string
	Details     map[string]string
	CredData    map[string]interface{}
}

func (m MockCredentialRequest) GetApplicationName() string {
	return m.AppName
}

func (m MockCredentialRequest) GetID() string {
	return m.ID
}

func (m MockCredentialRequest) GetCredentialDetailsValue(key string) string {
	return m.Details[key]
}

func (m MockCredentialRequest) GetApplicationDetailsValue(key string) string {
	return m.AppDetails[key]
}

func (m MockCredentialRequest) GetCredentialType() string {
	return m.CredDefName
}

func (m MockCredentialRequest) GetCredentialData() map[string]interface{} {
	return m.CredData
}

func (m MockCredentialRequest) IsIDPCredential() bool {
	return false
}

func GetIDPProvider() oauth.Provider {
	return nil
}

func GetIDPCredentialData() provisioning.IDPCredentialData {
	return nil
}

type MockAccessRequest struct {
	provisioning.AccessRequest
	ID                            string
	AppDetails                    map[string]string
	AppName                       string
	Details                       map[string]string
	InstanceDetails               map[string]interface{}
	AccessRequestData             map[string]interface{}
	AccessRequestProvisioningData interface{}
}

func (m MockAccessRequest) GetID() string {
	return m.ID
}

func (m MockAccessRequest) GetAccessRequestData() map[string]interface{} {
	return m.AccessRequestData
}

func (m MockAccessRequest) GetAccessRequestProvisioningData() interface{} {
	return m.AccessRequestProvisioningData
}

func (m MockAccessRequest) GetApplicationName() string {
	return m.AppName
}

func (m MockAccessRequest) GetAccessRequestDetailsValue(key string) string {
	return m.Details[key]
}

func (m MockAccessRequest) GetApplicationDetailsValue(key string) string {
	return m.AppDetails[key]
}

func (m MockAccessRequest) GetInstanceDetails() map[string]interface{} {
	return m.InstanceDetails
}

type MockRequestStatus struct {
	Msg        string
	Properties map[string]string
	Status     provisioning.Status
}

func (m MockRequestStatus) GetStatus() provisioning.Status {
	return m.Status
}

func (m MockRequestStatus) GetMessage() string {
	return m.Msg
}

func (m MockRequestStatus) GetProperties() map[string]string {
	return m.Properties
}
