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
)

type ServiceNowClient struct {
	Credential core.Credential
	BaseUrl    string //nolint:stylecheck
	Session    http.Client
}

// Now returns a NowRequestBuilder associated with the Client.
// It prepares the NowRequestBuilder with the base URL for the ServiceNow instance.
func (c *ServiceNowClient) Now(ctx context.Context) *NowRequestBuilder {
	return NewNowRequestBuilder(ctx, c.BaseUrl+"/now", c)
}

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

	return &ServiceNowClient{
		Credential: credential,
		BaseUrl:    instance,
		Session:    http.Client{},
	}
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

	request, err := requestInfo.ToRequestWithContext(ctx)
	if err != nil {
		return nil, err
	}

	authHeader, err := c.Credential.GetAuthentication()
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", authHeader)
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

func (c *ServiceNowClient) Send(ctx context.Context, requestInfo core.IRequestInformation,
	errorMapping core.ErrorMapping) (*http.Response, error) {
	return c.SendWithContext(ctx, requestInfo, errorMapping)
}
