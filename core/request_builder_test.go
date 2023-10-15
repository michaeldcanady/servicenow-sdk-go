package core

import (
	"net/http"
	"testing"
)

// Create a mock client for testing purposes
type mockClient struct{}

func (c *mockClient) Send(requestInfo *RequestInformation, errorMapping ErrorMapping) (*http.Response, error) {
	// Implement a mock Send function for testing
	return nil, nil
}

func TestToHeadRequestInformation(t *testing.T) {
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

func TestToGetRequestInformation(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Create mock query parameters
	params := map[string]interface{}{
		"param1": "value1",
		"param2": "value2",
	}

	// Call the ToGetRequestInformation method
	requestInfo, err := builder.ToGetRequestInformation(params)

	// Perform assertions to check the result
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if requestInfo.Method != GET {
		t.Errorf("Expected method to be GET, but got %s", requestInfo.Method)
	}
	// Add more assertions as needed
}

func TestToPostRequestInformation(t *testing.T) {
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
	requestInfo, err := builder.ToPostRequestInformation(data, params)

	// Perform assertions to check the result
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if requestInfo.Method != POST {
		t.Errorf("Expected method to be POST, but got %s", requestInfo.Method)
	}
	// Add more assertions as needed
}

func TestToDeleteRequestInformation(t *testing.T) {
	// Create a mock RequestBuilder with a mock client
	builder := NewRequestBuilder(&mockClient{}, "https://example.com", nil)

	// Create mock query parameters
	params := map[string]interface{}{
		"param1": "value1",
		"param2": "value2",
	}

	// Call the ToDeleteRequestInformation method
	requestInfo, err := builder.ToDeleteRequestInformation(params)

	// Perform assertions to check the result
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if requestInfo.Method != DELETE {
		t.Errorf("Expected method to be DELETE, but got %s", requestInfo.Method)
	}
	// Add more assertions as needed
}

func TestToRequestInformation(t *testing.T) {
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
	// Add more assertions as needed

	// Call the ToRequestInformation method for a POST request
	requestInfo, err = builder.ToRequestInformation(POST, data, params)

	// Perform assertions to check the result
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if requestInfo.Method != POST {
		t.Errorf("Expected method to be POST, but got %s", requestInfo.Method)
	}
	// Add more assertions as needed
}
