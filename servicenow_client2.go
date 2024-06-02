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
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

type ServiceNowClient2 struct {
	authProvider *internal.BaseAuthorizationProvider
	baseURL      string
	session      http.Client
}

// Now returns a NowRequestBuilder associated with the Client.
// It prepares the NowRequestBuilder with the base URL for the ServiceNow instance.
func (c *ServiceNowClient2) Now() *NowRequestBuilder2 {
	return NewNowRequestBuilder2(c, map[string]string{"baseurl": c.GetBaseURL()})
}

// NewServiceNowClient2 creates a new instance of the ServiceNow client.
// It accepts a UsernamePasswordCredential and an instance URL.
// If the instance URL does not end with ".service-now.com", it appends the suffix.
// It returns a pointer to the Client.
func NewServiceNowClient3(credential core.Credential, instance string) (*ServiceNowClient2, error) {
	if !strings.HasSuffix(instance, ".service-now.com") {
		instance += ".service-now.com"
	}

	if !strings.HasPrefix(instance, "https://") {
		instance = "https://" + instance
	}

	authProvider, err := internal.NewBaseAuthorizationProvider(credential)
	if err != nil {
		return nil, err
	}

	return &ServiceNowClient2{
		authProvider: authProvider,
		baseURL:      instance,
		session:      http.Client{},
	}, nil
}

func (c *ServiceNowClient2) unmarshallError(response *http.Response) error {
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

func (c *ServiceNowClient2) throwIfFailedResponse(response *http.Response, errorMappings intCore.ErrorMapping) error {
	if response.StatusCode < 400 {
		return nil
	}

	statusAsString := strconv.Itoa(response.StatusCode)
	var errorCtor interface{} = nil

	if errorMappings.Len() != 0 {
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

func (c *ServiceNowClient2) toRequestWithContext(ctx context.Context, requestInfo intCore.RequestInformation) (*http.Request, error) {
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

func (c *ServiceNowClient2) toRequest(requestInfo intCore.RequestInformation) (*http.Request, error) {
	return c.toRequestWithContext(context.Background(), requestInfo)
}

func (c *ServiceNowClient2) SendWithContext(ctx context.Context, requestInfo intCore.RequestInformation, errorMapping intCore.ErrorMapping) (*http.Response, error) {
	request, err := c.toRequestWithContext(ctx, requestInfo)
	if err != nil {
		return nil, err
	}

	response, err := c.session.Do(request)
	if err != nil {
		return nil, fmt.Errorf("unable to complete request: %s", err)
	}

	err = c.throwIfFailedResponse(response, errorMapping)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ServiceNowClient2) Send(requestInfo intCore.RequestInformation, errorMapping intCore.ErrorMapping) (*http.Response, error) {
	return c.SendWithContext(context.Background(), requestInfo, errorMapping)
}

func (c *ServiceNowClient2) GetBaseURL() string {
	return c.baseURL
}
