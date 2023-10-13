package abstraction

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type RequestBuilder struct {
	PathParameters map[string]string
	Client         Client
	UrlTemplate    string
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
func (T *RequestBuilder) ToHeadRequestInformation() (*RequestInformation, error) {
	return T.ToRequestInformation(HEAD, nil, nil)
}

// ToGetRequestInformation creates a new HTTP GET request's RequestInformation object.
// It sets the HTTP method to GET and includes the specified query parameters.
//
// Parameters:
//   - params: An interface representing query parameters for the GET request.
//
// Returns:
//   - *RequestInformation: A RequestInformation object representing the GET request.
//   - error: An error if there was an issue creating the request information.
func (T *RequestBuilder) ToGetRequestInformation(params interface{}) (*RequestInformation, error) {
	return T.ToRequestInformation(GET, nil, params)
}

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
func (T *RequestBuilder) ToPostRequestInformation(data map[string]interface{}, params interface{}) (*RequestInformation, error) {
	return T.ToRequestInformation(POST, data, params)
}

// ToDeleteRequestInformation creates a new HTTP DELETE request's RequestInformation object.
// It sets the HTTP method to DELETE and includes the specified query parameters.
//
// Parameters:
//   - params: An interface representing query parameters for the DELETE request.
//
// Returns:
//   - *RequestInformation: A RequestInformation object representing the DELETE request.
//   - error: An error if there was an issue creating the request information.
func (T *RequestBuilder) ToDeleteRequestInformation(params interface{}) (*RequestInformation, error) {
	return T.ToRequestInformation(DELETE, nil, params)
}

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
func (T *RequestBuilder) ToRequestInformation(method HttpMethod, data map[string]interface{}, params interface{}) (*RequestInformation, error) {
	requestInfo := NewRequestInformation()
	requestInfo.Method = method
	requestInfo.uri.PathParameters = T.PathParameters
	requestInfo.uri.UrlTemplate = T.UrlTemplate

	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal JSON: %s", err)
		}
		requestInfo.Content = jsonData
	}

	if params != nil {
		paramsValue := reflect.ValueOf(params)
		if paramsValue.Kind() == reflect.Ptr {
			paramsValue = paramsValue.Elem()
		}

		err := requestInfo.AddQueryParameters(paramsValue)
		if err != nil {
			return nil, err
		}
	}
	return requestInfo, nil
}
