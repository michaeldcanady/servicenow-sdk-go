package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	collectionLinks = []string{fakeLastLink, fakeNextLink}
	itemLinks       = []string{}
	baseResp        = &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body:       http.NoBody,
		Request:    nil,
	}
)

// Create a mock client for testing purposes
type mockClient struct{}

func (c *mockClient) sendCollection(requestInformation IRequestInformation) (*http.Response, error) {
	url, err := requestInformation.Url()
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL: %s", err)
	}

	switch url {
	case fakeLastLink:
		baseResp.Body = io.NopCloser(strings.NewReader(string(sharedCurrentPageToJSON())))
	case fakeNextLink:
		baseResp.Body = io.NopCloser(strings.NewReader(string(sharedCurrentPageToJSON())))
	}

	return baseResp, nil
}

func (c *mockClient) sendItem(requestInformation IRequestInformation) (*http.Response, error) {
	uri, err := requestInformation.Url()
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL: %s", err)
	}

	parsedURI, _ := url.Parse(uri)

	switch parsedURI.Scheme + "://" + parsedURI.Host {
	}
	return nil, nil
}

func (c *mockClient) Send(ctx context.Context, requestInformation IRequestInformation, errorMapping ErrorMapping) (*http.Response, error) {
	uri, err := requestInformation.Url()
	if err != nil {
		return nil, fmt.Errorf("unable to parse URL: %s", err)
	}

	parsedURI, _ := url.Parse(uri)

	if uri == fakeLinkNilResp {
		return nil, nil
	}

	if slices.Contains[[]string, string](collectionLinks, uri) {
		return c.sendCollection(requestInformation)
	}

	if slices.Contains[[]string, string](itemLinks, parsedURI.Scheme+"://"+parsedURI.Host) {
		return c.sendItem(requestInformation)
	}

	return nil, nil
}

func TestRequestBuilderToHeadRequestInformation(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Call the ToHeadRequestInformation method
	requestInfo, err := builder.ToHeadRequestInformation()

	// Perform assertions to check the result
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if requestInfo.Method != HEAD {
		t.Errorf("Expected method to be HEAD, but got %s", requestInfo.Method)
	}
	// Add more assertions as needed
}

func TestRequestBuilderToGetRequestInformation(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Create mock query parameters
	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	// Call the ToGetRequestInformation method
	requestInfo, err := builder.ToGetRequestInformation(params)

	// Perform assertions to check the result
	assert.NoError(t, err)
	assert.Equal(t, GET, requestInfo.Method)
	assert.Equal(t, params, requestInfo.uri.QueryParameters)
}

func TestRequestBuilderToGetRequestInformation2(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Create mock query parameters
	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	config := RequestConfiguration{
		QueryParameters: params,
	}

	// Call the ToGetRequestInformation method
	requestInfo, err := builder.ToGetRequestInformation2(&config)

	// Perform assertions to check the result
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, GET, requestInfo.Method)
	assert.Equal(t, params, requestInfo.uri.QueryParameters)
}

func TestRequestBuilderToPostRequestInformation(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Create mock data and query parameters
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	expectedJSON, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	// Call the ToPostRequestInformation method
	requestInfo, err := builder.ToPostRequestInformation(data, params)

	// Perform assertions to check the result
	assert.NoError(t, err)
	assert.Equal(t, POST, requestInfo.Method)
	assert.Equal(t, expectedJSON, requestInfo.Content)
}

func TestRequestBuilder_ToPostRequestInformation2(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Create mock data and query parameters
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	expectedJSON, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	// Call the ToPostRequestInformation method
	requestInfo, err := builder.ToPostRequestInformation2(data, params)

	// Perform assertions to check the result
	assert.NoError(t, err)
	assert.Equal(t, POST, requestInfo.Method)
	assert.Equal(t, expectedJSON, requestInfo.Content)
}

func TestRequestBuilderToPostRequestInformation3(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Create mock data and query parameters
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	expectedJSON, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}

	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	config := RequestConfiguration{
		QueryParameters: params,
		Data:            data,
	}

	// Call the ToPostRequestInformation method
	requestInfo, err := builder.ToPostRequestInformation3(&config)

	// Perform assertions to check the result
	assert.NoError(t, err)
	assert.Equal(t, POST, requestInfo.Method)
	assert.Equal(t, expectedJSON, requestInfo.Content)
}

func TestRequestBuilderToDeleteRequestInformation(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Create mock query parameters
	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	// Call the ToDeleteRequestInformation method
	requestInfo, err := builder.ToDeleteRequestInformation(params)

	// Perform assertions to check the result
	assert.NoError(t, err)
	assert.Equal(t, params, requestInfo.uri.QueryParameters)
}

func TestRequestBuilderToDeleteRequestInformation2(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Create mock query parameters
	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	config := RequestConfiguration{
		QueryParameters: params,
	}

	// Call the ToDeleteRequestInformation method
	requestInfo, err := builder.ToDeleteRequestInformation2(&config)

	// Perform assertions to check the result
	assert.NoError(t, err)
	assert.Equal(t, params, requestInfo.uri.QueryParameters)
}

func TestRequestBuilderToPutRequestInformation2(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Create mock data and query parameters
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	// Call the ToPostRequestInformation method
	requestInfo, err := builder.ToPutRequestInformation(data, params)

	// Perform assertions to check the result
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if requestInfo.Method != PUT {
		t.Errorf("Expected method to be PUT, but got %s", requestInfo.Method)
	}
}

func TestRequestBuilderPrepareData(t *testing.T) {
	t.Run("NilData", func(t *testing.T) {
		rB := &RequestBuilder{}
		data, err := rB.prepareData(nil)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if data != nil {
			t.Fatalf("Expected nil data for nil input, got: %v", data)
		}
	})

	t.Run("ByteData", func(t *testing.T) {
		rB := &RequestBuilder{}
		rawData := []byte("test data")
		data, err := rB.prepareData(rawData)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(data, rawData) {
			t.Fatalf("Expected %v, got: %v", rawData, data)
		}
	})

	t.Run("MapData", func(t *testing.T) {
		rB := &RequestBuilder{}
		rawData := map[string]string{
			"key": "value",
		}
		data, err := rB.prepareData(rawData)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		expectedData, err := json.Marshal(rawData)
		if err != nil {
			t.Fatalf("Error marshalling map data: %v", err)
		}

		if !reflect.DeepEqual(data, expectedData) {
			t.Fatalf("Expected %v, got: %v", expectedData, data)
		}
	})
}

func TestRequestBuilderToRequestInformation(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Create mock data and query parameters
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	params := map[string]interface{}{
		"param1": "value1",
		"param2": "value2",
	}

	// Call the ToRequestInformation method for a GET request
	requestInfo, err := builder.ToRequestInformation(GET, nil, params)

	// Perform assertions to check the result
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if requestInfo.Method != GET {
		t.Errorf("Expected method to be GET, but got %s", requestInfo.Method)
	}

	// Call the ToRequestInformation method for a POST request
	requestInfo, err = builder.ToRequestInformation(POST, data, params)

	// Perform assertions to check the result
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if requestInfo.Method != POST {
		t.Errorf("Expected method to be POST, but got %s", requestInfo.Method)
	}
}

type requestBuilderTest struct {
	title  string
	method HttpMethod
	config RequestConfiguration
}

func TestRequestBuilderToRequestInformation3(t *testing.T) {
	tests := []requestBuilderTest{
		{
			title:  "Test GET",
			method: GET,
			config: RequestConfiguration{
				QueryParameters: map[string]interface{}{
					"param1": "value1",
					"param2": "value2",
				},
				Data: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			title:  "Test POST",
			method: POST,
			config: RequestConfiguration{
				QueryParameters: map[string]interface{}{
					"param1": "value1",
					"param2": "value2",
				},
				Data: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
	}

	for _, test := range tests {
		expectedJSON, err := json.Marshal(test.config.Data)
		assert.NoError(t, err)

		// Create a mock RequestBuilder with a mock client
		builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

		// Call the ToRequestInformation method for a GET request
		requestInfo, err := builder.ToRequestInformation3(test.method, &test.config) //nolint:gosec

		assert.NoError(t, err)
		assert.Equal(t, test.method, requestInfo.Method)
		assert.Equal(t, expectedJSON, requestInfo.Content)
	}
}
