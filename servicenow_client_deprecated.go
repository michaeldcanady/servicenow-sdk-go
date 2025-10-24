package servicenowsdkgo

import (
	"net/http"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type webClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ServiceNowClient struct {
	// Deprecated: deprecated since v1.6.0.
	Credential     core.Credential
	authProvider   internal.AuthorizationProvider
	BaseUrl        string //nolint:stylecheck
	Session        webClient
	requestAdapter abstractions.RequestAdapter
}

// Deprecated: deprecated since v1.8.0. Please use Now2.
//
// Now returns a NowRequestBuilder associated with the Client.
// It prepares the NowRequestBuilder with the base URL for the ServiceNow instance.
func (c *ServiceNowClient) Now() *NowRequestBuilder {
	return NewNowRequestBuilder(c.BaseUrl+"/now", c)
}

// Deprecated: deprecated since v1.6.0. Please use `NewServiceNowClient2` instead.
//
// NewServiceNowClient creates a new instance of the ServiceNow client.
// It accepts a UsernamePasswordCredential and an instance URL.
// If the instance URL does not end with ".service-now.com/api", it appends the suffix.
// It returns a pointer to the Client.
func NewServiceNowClient(credential core.Credential, instance string) *ServiceNowClient {
	if !strings.HasSuffix(instance, ".service-now.com/api") {
		instance += ".service-now.com/api"
	}

	if !strings.HasPrefix(instance, "https://") {
		instance = "https://" + instance
	}

	authProvider, _ := internal.NewBaseAuthorizationProvider(credential)

	return &ServiceNowClient{
		Credential:   credential,
		authProvider: authProvider,
		BaseUrl:      instance,
		Session:      &http.Client{},
	}
}
