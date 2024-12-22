package servicenowsdkgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	now "github.com/michaeldcanady/servicenow-sdk-go/now"
)

// Deprecated: deprecated since v{unreleased}.
//
// ServiceNowClient ...
type ServiceNowClient struct {
	// Deprecated: deprecated since v1.6.0.
	Credential   core.Credential
	authProvider *internal.BaseAuthorizationProvider
	BaseUrl      string //nolint:stylecheck
	Session      http.Client
}

// Deprecated: deprecated since v{unreleased}.
//
// Now returns a NowRequestBuilder associated with the Client.
// It prepares the NowRequestBuilder with the base URL for the ServiceNow instance.
func (c *ServiceNowClient) Now() *NowRequestBuilder {
	return NewNowRequestBuilder(c.BaseUrl+"/now", c)
}

// Now provides entrypoint into Service-Now's APIs
func (c *ServiceNowClient) Now2() *now.NowRequestBuilder2 {
	pathParameters := map[string]string{
		"baseurl": c.BaseUrl,
	}
	return now.NewAPIV1CompatibleNowRequestBuilder2Internal(pathParameters, c)
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
		Session:      http.Client{},
	}
}

// Deprecated: deprecated since v{unreleased}.
//
// NewServiceNowClient2 creates a new instance of the ServiceNow client.
// It accepts a UsernamePasswordCredential and an instance URL.
// If the instance URL does not end with ".service-now.com/api", it appends the suffix.
// It returns a pointer to the Client.
func NewServiceNowClient2(credential core.Credential, instance string) (*ServiceNowClient, error) {
	if !strings.HasSuffix(instance, ".service-now.com/api") {
		instance += ".service-now.com/api"
	}

	if !strings.HasPrefix(instance, "https://") {
		instance = "https://" + instance
	}

	authProvider, err := internal.NewBaseAuthorizationProvider(credential)
	if err != nil {
		return nil, err
	}

	return &ServiceNowClient{
		Credential:   credential,
		authProvider: authProvider,
		BaseUrl:      instance,
		Session:      http.Client{},
	}, nil
}

func (c *ServiceNowClient) unmarshallError(response *http.Response) error {
	var stringError core.ServiceNowError

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &stringError); err != nil {
		return err
	}
	return &stringError
}

func (c *ServiceNowClient) throwIfFailedResponse(response *http.Response, errorMappings core.ErrorMapping) error {
	if response.StatusCode < 400 {
		return nil
	}

	statusAsString := strconv.Itoa(response.StatusCode)
	var errorCtor interface{} = nil

	if len(errorMappings) != 0 {
		var isOk bool
		errorCtor, isOk = errorMappings.Get(response.StatusCode)
		if !isOk {
			errorCtor = nil
		}
	}

	if errorCtor == nil {
		err := &core.ApiError{
			Message:            "The server returned an unexpected status code and no error factory is registered for this code: " + statusAsString,
			ResponseStatusCode: response.StatusCode,
		}
		return err
	}

	stringError := c.unmarshallError(response)

	return stringError
}

func (c *ServiceNowClient) toRequestWithContext(ctx context.Context, requestInfo core.IRequestInformation) (*http.Request, error) {
	if requestInfo == nil {
		return nil, ErrNilRequestInfo
	}

	if ctx == nil {
		return nil, ErrNilContext
	}

	err := c.authProvider.AuthorizeRequest(requestInfo)
	if err != nil {
		return nil, err
	}

	request, err := requestInfo.ToRequestWithContext(ctx)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	return request, nil
}

func (c *ServiceNowClient) toRequest(requestInfo core.IRequestInformation) (*http.Request, error) {
	return c.toRequestWithContext(context.Background(), requestInfo)
}

func (c *ServiceNowClient) SendWithContext(ctx context.Context, requestInfo core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {
	request, err := c.toRequestWithContext(ctx, requestInfo)
	if err != nil {
		return nil, err
	}

	response, err := c.Session.Do(request)
	if err != nil {
		return nil, fmt.Errorf("unable to complete request: %s", err)
	}

	err = c.throwIfFailedResponse(response, errorMapping)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ServiceNowClient) Send(requestInfo core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {
	return c.SendWithContext(context.Background(), requestInfo, errorMapping)
}

func (c *ServiceNowClient) GetBaseURL() string {
	return c.BaseUrl
}
