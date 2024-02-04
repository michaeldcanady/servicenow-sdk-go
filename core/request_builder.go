package core

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gabriel-vasile/mimetype"
)

// RequestBuilder represents a builder for constructing HTTP request information.
type RequestBuilder struct {
	// PathParameters is a map of path parameters used in the URL template.
	PathParameters map[string]string
	// Client is an instance of the HTTP client used to send requests.
	Client Client
	// UrlTemplate is the URL template for constructing the request URL.
	UrlTemplate string //nolint:stylecheck
}

// NewRequestBuilder creates a new instance of the RequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created RequestBuilder.
func NewRequestBuilder(client Client, urlTemplate string, pathParameters map[string]string) *RequestBuilder {
	return &RequestBuilder{
		Client:         client,
		UrlTemplate:    urlTemplate,
		PathParameters: pathParameters,
	}
}

// ToHeadRequestInformation creates a new HTTP HEAD request's RequestInformation object.
// It sets the HTTP method to HEAD and includes no request data or query parameters.
//
// Returns:
//   - *RequestInformation: A RequestInformation object representing the HEAD request.
//   - error: An error if there was an issue creating the request information.
func (rB *RequestBuilder) ToHeadRequestInformation() (*RequestInformation, error) {
	return rB.ToRequestInformation2(HEAD, nil, nil)
}

// Deprecated: deprecated as of 1.4.0 please utilize `ToGetRequestInformation2`
// ToGetRequestInformation creates a new HTTP GET request's RequestInformation object.
// It sets the HTTP method to GET and includes the specified query parameters.
//
// Parameters:
//   - params: An interface representing query parameters for the GET request.
//
// Returns:
//   - *RequestInformation: A RequestInformation object representing the GET request.
//   - error: An error if there was an issue creating the request information.
func (rB *RequestBuilder) ToGetRequestInformation(params interface{}) (*RequestInformation, error) {
	return rB.ToRequestInformation2(GET, nil, params)
}

func (rB *RequestBuilder) ToGetRequestInformation2(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation3(GET, config)
}

// Deprecated: deprecated as of 1.4.0 please utilize `ToPutRequestInformation2`
// Put updates a table item using an HTTP PUT request.
// It takes a map of table entry data and optional query parameters to send in the request.
// The method returns a TableItemResponse representing the updated item or an error if the request fails.
//
// Parameters:
//   - tableEntry: A map containing the data to update the table item.
//   - params: An optional pointer to TableItemRequestBuilderPutQueryParameters, which can be used to specify query parameters for the request.
//
// Returns:
//   - *TableItemResponse: A TableItemResponse containing the updated item data.
//   - error: An error, if the request fails at any point, such as request information creation or JSON deserialization.
func (rB *RequestBuilder) ToPutRequestInformation(data map[string]string, params interface{}) (*RequestInformation, error) {
	return rB.ToRequestInformation2(PUT, data, params)
}

func (rB *RequestBuilder) ToPutRequestInformation2(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation3(PUT, config)
}

// Deprecated: deprecated as of 1.4.0 please utilize `ToPostRequestInformation3`
// ToPostRequestInformation2 creates a new HTTP POST request's RequestInformation object.
// It sets the HTTP method to POST and includes the specified data in the request body
// and query parameters.
//
// Parameters:
//   - data: A map[string]interface{} representing data to be included in the request body.
//   - params: An interface representing query parameters for the POST request.
//
// Returns:
//   - *RequestInformation: A RequestInformation object representing the POST request.
//   - error: An error if there was an issue creating the request information.
func (rB *RequestBuilder) ToPostRequestInformation2(data interface{}, params interface{}) (*RequestInformation, error) {
	return rB.ToRequestInformation2(POST, data, params)
}

func (rB *RequestBuilder) ToPostRequestInformation3(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation3(POST, config)
}

// Deprecated: deprecated as of 1.4.0 please utilize `ToPostRequestInformation2`
// ToPostRequestInformation creates a new HTTP POST request's RequestInformation object.
// It sets the HTTP method to POST and includes the specified data in the request body
// and query parameters.
//
// Parameters:
//   - data: A map[string]interface{} representing data to be included in the request body.
//   - params: An interface representing query parameters for the POST request.
//
// Returns:
//   - *RequestInformation: A RequestInformation object representing the POST request.
//   - error: An error if there was an issue creating the request information.
func (rB *RequestBuilder) ToPostRequestInformation(data map[string]string, params interface{}) (*RequestInformation, error) {
	return rB.ToRequestInformation2(POST, data, params)
}

// Deprecated: deprecated as of 1.4.0 please utilize `ToDeleteRequestInformation2`
// ToDeleteRequestInformation creates a new HTTP DELETE request's RequestInformation object.
// It sets the HTTP method to DELETE and includes the specified query parameters.
//
// Parameters:
//   - params: An interface representing query parameters for the DELETE request.
//
// Returns:
//   - *RequestInformation: A RequestInformation object representing the DELETE request.
//   - error: An error if there was an issue creating the request information.
func (rB *RequestBuilder) ToDeleteRequestInformation(params interface{}) (*RequestInformation, error) {
	return rB.ToRequestInformation2(DELETE, nil, params)
}

func (rB *RequestBuilder) ToDeleteRequestInformation2(config *RequestConfiguration) (*RequestInformation, error) {
	return rB.ToRequestInformation3(DELETE, config)
}

func (rB *RequestBuilder) prepareData(rawData interface{}) ([]byte, error) {
	var data []byte
	var err error

	if rawData == nil {
		return data, nil
	}

	if reflect.TypeOf(rawData) == reflect.TypeOf([]byte{}) {
		return rawData.([]byte), nil
	}

	if reflect.TypeOf(rawData) == reflect.TypeOf(map[string]string{}) {
		data, err = json.Marshal(rawData)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal JSON: %s", err)
		}
	}

	if len(data) == 0 {
		return data, fmt.Errorf("unsupported type: %T", rawData)
	}

	return data, nil
}

// Deprecated: deprecated as of 1.4.0 please utilize `ToRequestInformation3`
// ToRequestInformation2 creates a new HTTP request's RequestInformation object with the
// specified HTTP method, data in the request body, and query parameters.
//
// Parameters:
//   - method: The HTTP method for the request (e.g., "GET", "POST", "HEAD", "DELETE").
//   - data: A map[string]interface{} representing data to be included in the request body.
//   - params: An interface representing query parameters for the request.
//
// Returns:
//   - *RequestInformation: A RequestInformation object representing the HTTP request.
//   - error: An error if there was an issue creating the request information.
func (rB *RequestBuilder) ToRequestInformation2(method HttpMethod, rawData interface{}, params interface{}) (*RequestInformation, error) {
	requestInfo := NewRequestInformation()
	requestInfo.Method = method
	requestInfo.uri.PathParameters = rB.PathParameters
	requestInfo.uri.UrlTemplate = rB.UrlTemplate

	data, err := rB.prepareData(rawData)
	if err != nil {
		return nil, err
	}

	if len(data) != 0 {
		mime := mimetype.Detect(data)

		requestInfo.Content = data
		requestInfo.Headers.Add("Content-Type", mime.String())
	}

	if params != nil {
		err := requestInfo.AddQueryParameters(params)
		if err != nil {
			return nil, err
		}
	}
	return requestInfo, nil
}

func (rB *RequestBuilder) ToRequestInformation3(method HttpMethod, config *RequestConfiguration) (*RequestInformation, error) {
	requestInfo := NewRequestInformation()

	if config != nil {
		if config.QueryParameters != nil {
			err := requestInfo.AddQueryParameters(config.QueryParameters)
			if err != nil {
				return nil, err
			}
		}
		if config.Data != nil {
			data, err := rB.prepareData(config.Data)
			if err != nil {
				return nil, err
			}
			if len(data) != 0 {
				mime := mimetype.Detect(data)

				requestInfo.Content = data
				requestInfo.Headers.Add("Content-Type", mime.String())
			}
		}
	}

	requestInfo.Method = method
	requestInfo.uri.PathParameters = rB.PathParameters
	requestInfo.uri.UrlTemplate = rB.UrlTemplate

	return requestInfo, nil
}

// Deprecated: deprecated as of 1.4.0 please utilize `ToRequestInformation2`
// ToRequestInformation creates a new HTTP request's RequestInformation object with the
// specified HTTP method, data in the request body, and query parameters.
//
// Parameters:
//   - method: The HTTP method for the request (e.g., "GET", "POST", "HEAD", "DELETE").
//   - data: A map[string]interface{} representing data to be included in the request body.
//   - params: An interface representing query parameters for the request.
//
// Returns:
//   - *RequestInformation: A RequestInformation object representing the HTTP request.
//   - error: An error if there was an issue creating the request information.
func (rB *RequestBuilder) ToRequestInformation(method HttpMethod, data map[string]string, params interface{}) (*RequestInformation, error) {
	requestInfo := NewRequestInformation()
	requestInfo.Method = method
	requestInfo.uri.PathParameters = rB.PathParameters
	requestInfo.uri.UrlTemplate = rB.UrlTemplate

	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal JSON: %s", err)
		}
		requestInfo.Content = jsonData
	}

	if params != nil {
		err := requestInfo.AddQueryParameters(params)
		if err != nil {
			return nil, err
		}
	}
	return requestInfo, nil
}

// Deprecated: deprecated since v1.4.0. Please use SendGet2
func (rB *RequestBuilder) SendGet(params interface{}, errorMapping ErrorMapping, value Response) error {
	err := sendGet(rB, params, errorMapping, &value)
	if err != nil {
		return err
	}
	return nil
}

func (rB *RequestBuilder) SendGet2(ctx context.Context, config *RequestConfiguration) error {
	err := SendGet2(ctx, rB, config)
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: deprecated since v1.4.0. Please use SendPost3
func (rB *RequestBuilder) SendPost(data map[string]string, params interface{}, errorMapping ErrorMapping, value Response) error {
	return sendPost(rB, data, params, errorMapping, &value)
}

func (rB *RequestBuilder) SendPost3(ctx context.Context, config *RequestConfiguration) error {
	return SendPost2(ctx, rB, config)
}

// Deprecated: deprecated since v1.4.0. Please use SendPost3
func (rB *RequestBuilder) SendPost2(data interface{}, params interface{}, errorMapping ErrorMapping, value Response) error {
	return sendPost(rB, data, params, errorMapping, &value)
}

// Deprecated: deprecated since v1.4.0. Please use SendDelete2
func (rB *RequestBuilder) SendDelete(params interface{}, errorMapping ErrorMapping) error {
	return sendDelete(rB, params, errorMapping)
}

func (rB *RequestBuilder) SendDelete2(ctx context.Context, config *RequestConfiguration) error {
	return sendDelete2(ctx, rB, config)
}

// Deprecated: deprecated since v1.4.0. Please use SendPut2
func (rB *RequestBuilder) SendPut(data map[string]string, params interface{}, errorMapping ErrorMapping, value Response) error {
	return sendPut(rB, data, params, errorMapping, &value)
}

func (rB *RequestBuilder) SendPut2(ctx context.Context, config *RequestConfiguration) error {
	return sendPut2(ctx, rB, config)
}
