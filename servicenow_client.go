package servicenowsdkgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

var (
	ErrNilRequestInfo = errors.New("requestInfo cannot be nil")
)

type ServiceNowClient struct {
	Credential core.Credential
	BaseUrl    string
	Session    http.Client
}

// Now returns a NowRequestBuilder associated with the Client.
// It prepares the NowRequestBuilder with the base URL for the ServiceNow instance.
func (C *ServiceNowClient) Now() *NowRequestBuilder {
	return NewNowRequestBuilder(C.BaseUrl+"/now", C)
}

// NewClient creates a new instance of the ServiceNow client.
// It accepts a UsernamePasswordCredential and an instance URL.
// If the instance URL does not end with ".service-now.com/api", it appends the suffix.
// It returns a pointer to the Client.
func NewClient(credential core.Credential, instance string) *ServiceNowClient {
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

func (C *ServiceNowClient) unmarshallError(response *http.Response) error {
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

func (C *ServiceNowClient) throwIfFailedResponse(response *http.Response, errorMappings core.ErrorMapping) error {

	if response.StatusCode < 400 {
		return nil
	}

	statusAsString := strconv.Itoa(response.StatusCode)
	var errorCtor interface{} = nil

	if len(errorMappings) != 0 {
		if errorMappings[statusAsString] != "" {
			errorCtor = errorMappings[statusAsString]
		} else if response.StatusCode >= 400 && response.StatusCode < 500 && errorMappings["4XX"] != "" {
			errorCtor = errorMappings["4XX"]
		} else if response.StatusCode >= 500 && response.StatusCode < 600 && errorMappings["5XX"] != "" {
			errorCtor = errorMappings["5XX"]
		}
	}

	if errorCtor == nil {
		err := &core.ApiError{
			Message:            "The server returned an unexpected status code and no error factory is registered for this code: " + statusAsString,
			ResponseStatusCode: response.StatusCode,
		}
		return err
	}

	stringError := C.unmarshallError(response)

	return stringError
}

func (C *ServiceNowClient) toRequest(requestInfo *core.RequestInformation) (*http.Request, error) {
	if requestInfo == nil {
		return nil, ErrNilRequestInfo
	}

	request, err := requestInfo.ToRequest()
	if err != nil {
		return nil, err
	}

	authHeader, err := C.Credential.GetAuthentication()
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", authHeader)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	return request, nil
}

func (C *ServiceNowClient) Send(requestInfo *core.RequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {

	request, err := C.toRequest(requestInfo)
	if err != nil {
		return nil, err
	}

	response, err := C.Session.Do(request)
	if err != nil {
		return nil, fmt.Errorf("unable to complete request: %s", err)
	}

	err = C.throwIfFailedResponse(response, errorMapping)
	if err != nil {
		return nil, err
	}

	return response, nil
}
