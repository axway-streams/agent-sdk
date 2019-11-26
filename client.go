package apic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	corecfg "git.ecd.axway.int/apigov/aws_apigw_discovery_agent/core/config"
	"git.ecd.axway.int/apigov/service-mesh-agent/pkg/apicauth"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

//CatalogCreator - interface
type CatalogCreator interface {
	CreateService(serviceBody ServiceBody) ([]byte, error)
	ExecuteService(service Service) (string, error)
	DeployAPI(service Service)
}

//ServiceBody -
type ServiceBody struct {
	NameToPush    string `json:",omitempty"`
	APIName       string `json:",omitempty"`
	URL           string `json:",omitempty"`
	TeamID        string `json:",omitempty"`
	Description   string `json:",omitempty"`
	Version       string `json:",omitempty"`
	AuthPolicy    string `json:",omitempty"`
	Swagger       []byte `json:",omitempty"`
	Documentation []byte `json:",omitempty"`
	Tags          map[string]interface{}
	AgentMode     corecfg.AgentMode `json:",omitempty"`
}

//Service - Used for both adding and updating of catalog item
type Service struct {
	Method    string            `json:",omitempty"`
	URL       string            `json:",omitempty"`
	Buffer    []byte            `json:",omitempty"`
	AgentMode corecfg.AgentMode `json:",omitempty"`
}

// Client -
type Client struct {
	tokenRequester *apicauth.PlatformTokenGetter
	cfg            corecfg.CentralConfig
}

// New -
func New(cfg corecfg.CentralConfig) *Client {
	tokenURL := cfg.GetAuthConfig().GetTokenURL()
	aud := cfg.GetAuthConfig().GetAudience()
	priKey := cfg.GetAuthConfig().GetPrivateKey()
	pubKey := cfg.GetAuthConfig().GetPublicKey()
	keyPwd := cfg.GetAuthConfig().GetKeyPassword()
	clientID := cfg.GetAuthConfig().GetClientID()
	authTimeout := cfg.GetAuthConfig().GetTimeout()

	return &Client{
		cfg:            cfg,
		tokenRequester: apicauth.NewPlatformTokenGetter(priKey, pubKey, keyPwd, tokenURL, aud, clientID, authTimeout),
	}
}

// MapToStringArray -
func (c *Client) MapToStringArray(m map[string]interface{}) []string {
	strArr := []string{}

	for key, val := range m {
		v := val.(*string)
		if *v == "" {
			strArr = append(strArr, key)
		} else {
			strArr = append(strArr, key+"_"+*v)
		}
	}
	return strArr
}

var httpClient = http.DefaultClient
var log logrus.FieldLogger = logrus.WithField("package", "apic")

// SetLog sets the logger for the package.
func SetLog(newLog logrus.FieldLogger) {
	log = newLog
	return
}

// DeployAPI -
func (c *Client) DeployAPI(service Service) (string, error) {
	request, err := setHeader(c, service.Method, service.URL, bytes.NewBuffer(service.Buffer))

	if err != nil {
		return "", err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return "", err
	}

	if !(response.StatusCode == http.StatusOK || response.StatusCode == http.StatusCreated) {
		detail := make(map[string]*json.RawMessage)
		json.NewDecoder(response.Body).Decode(&detail)
		for k, v := range detail {
			buffer, _ := v.MarshalJSON()
			log.Debugf("HTTP response key %v: %v", k, string(buffer))
		}
		return "", errors.New(response.Status)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	return handleResponse(service.Method, service.AgentMode, body)
}

func handleResponse(method string, agentMode corecfg.AgentMode, body []byte) (string, error) {

	itemID := ""

	// Connected Mode
	if agentMode == corecfg.Connected {
		metadata := gjson.Get(string(body), "metadata").String()
		if metadata != "" {
			itemID = gjson.Get(string(metadata), "id").String()
		}
		// Disconnected Mode
	} else {
		itemID = gjson.Get(string(body), "id").String()
	}

	log.Debugf("HTTP response returning itemID: [%v]", itemID)
	return itemID, nil
}

// SetHeader - set header
func setHeader(c *Client, method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	var token string

	if token, err = c.tokenRequester.GetToken(); err != nil {
		return nil, err
	}

	request.Header.Add("X-Axway-Tenant-Id", c.cfg.GetTenantID())
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Add("Content-Type", "application/json")
	return request, nil
}
