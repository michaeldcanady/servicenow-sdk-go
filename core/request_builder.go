package core

import (
	"encoding/json"
	"fmt"
)

// RequestBuilder represents a builder for constructing HTTP request information.
type RequestBuilder struct {
	// PathParameters is a map of path parameters used in the URL template.
	PathParameters map[string]string
	// Client is an instance of the HTTP client used to send requests.
	Client Client
	// UrlTemplate is the URL template for constructing the request URL.
	UrlTemplate string
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
	return rB.ToRequestInformation(HEAD, nil, nil)
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
func (rB *RequestBuilder) ToGetRequestInformation(params interface{}) (*RequestInformation, error) {
	return rB.ToRequestInformation(GET, nil, params)
}

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
	return rB.ToRequestInformation(GET, data, params)
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
func (rB *RequestBuilder) ToPostRequestInformation(data map[string]string, params interface{}) (*RequestInformation, error) {
	return rB.ToRequestInformation(POST, data, params)
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
func (rB *RequestBuilder) ToDeleteRequestInformation(params interface{}) (*RequestInformation, error) {
	return rB.ToRequestInformation(DELETE, nil, params)
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
