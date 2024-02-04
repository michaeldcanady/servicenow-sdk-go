package tableapi

import (
	"encoding/json"
	"io"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

var (
	testResult = map[string]interface{}{
		"result": fakeResultItem,
	}
)

func (c *MockClient) Send(requestInfo core.IRequestInformation, errorMapping core.ErrorMapping) (*http.Response, error) {
	req, err := requestInfo.ToRequest()
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func TestNewTableItemRequestBuilder(t *testing.T) {
	client := &MockClient{}

	pathParameters := map[string]string{"baseurl": "instance.service-now.com", "table": "table1", "sysId": "sysid"}

	req := NewTableItemRequestBuilder(client, pathParameters)

	assert.NotNil(t, req)
}

func TestTableItemRequestBuilderGet(t *testing.T) {
	client := &MockClient{}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with the provided JSON

		responseJSON, err := json.Marshal(testResult)
		assert.Nil(t, err)

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseJSON) //nolint:errcheck
	}))

	parsedURL, err := url.Parse(mockServer.URL)
	assert.Nil(t, err)

	pathParameters := map[string]string{"baseurl": "http://" + parsedURL.Host, "table": parsedURL.Path, "sysId": "sysid"}

	req := NewTableItemRequestBuilder(client, pathParameters)

	params := &TableItemRequestBuilderGetQueryParameters{
		DisplayValue:         "true",
		ExcludeReferenceLink: true,
		Fields:               []string{"field1", "field2"},
		QueryNoDomain:        true,
		View:                 "desktop",
	}

	response, err := req.Get(params)

	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestTableItemRequestBuilderDelete(t *testing.T) {
	client := &MockClient{}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with the provided JSON
		responseJSON := ``

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(responseJSON)) //nolint:errcheck
	}))

	parsedURL, err := url.Parse(mockServer.URL)
	assert.Nil(t, err)

	pathParameters := map[string]string{"baseurl": "http://" + parsedURL.Host, "table": parsedURL.Path, "sysId": "sysid"}
	req := NewTableItemRequestBuilder(client, pathParameters)

	params := &TableItemRequestBuilderDeleteQueryParameters{
		QueryNoDomain: true,
	}

	err = req.Delete(params)

	assert.Nil(t, err)
}

//nolint:dupl
func TestTableItemRequestBuilderPut(t *testing.T) {
	// Create a mock client for testing
	client := &MockClient{}

	// Create a mock server for testing
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := testResult

		resp = maps.Clone[map[string]interface{}](resp)

		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			break
		case http.MethodPut:
			body := make(map[string]interface{})

			// Read the request body into a []byte
			requestBody, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = json.Unmarshal(requestBody, &body) // Use &body to correctly update the 'body' map
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			result := resp["result"].(map[string]interface{})

			for key, value := range body {
				result[key] = value
			}

			resp["result"] = result // Update the 'result' in 'resp'

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		jsonData, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonData)
	}))

	// Parse the URL of the mock server
	parsedURL, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	// Prepare path parameters
	pathParameters := map[string]string{
		"baseurl": "http://" + parsedURL.Host,
		"table":   parsedURL.Path,
		"sysId":   "sysid",
	}

	// Create a request builder
	req := NewTableItemRequestBuilder(client, pathParameters)

	// Prepare values to update the record
	values := map[string]string{
		// Provide values to change in the record here
		"location": "home",
	}

	// Send the PUT request to update the record
	updatedRecord, err := req.Put(values, nil)

	// Perform assertions and test the response
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// You should assert that the updatedRecord matches your expected response here.
	// You may need to unmarshal the JSON response and compare specific fields.
	// For example: assert.Equal(t, expectedValue, updatedRecord.Field)

	assert.Equal(t, fakeTableItemResponse, updatedRecord)
}

//nolint:dupl
func TestTableItemRequestBuilderPut2(t *testing.T) {
	// Create a mock client for testing
	client := &MockClient{}

	// Create a mock server for testing
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := testResult

		resp = maps.Clone[map[string]interface{}](resp)

		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			break
		case http.MethodPut:
			body := make(map[string]interface{})

			// Read the request body into a []byte
			requestBody, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = json.Unmarshal(requestBody, &body) // Use &body to correctly update the 'body' map
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			result := resp["result"].(map[string]interface{})

			for key, value := range body {
				result[key] = value
			}

			resp["result"] = result // Update the 'result' in 'resp'

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		jsonData, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jsonData)
	}))

	// Parse the URL of the mock server
	parsedURL, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	// Prepare path parameters
	pathParameters := map[string]string{
		"baseurl": "http://" + parsedURL.Host,
		"table":   parsedURL.Path,
		"sysId":   "sysid",
	}

	// Create a request builder
	req := NewTableItemRequestBuilder(client, pathParameters)

	// Prepare values to update the record
	values := map[string]string{
		// Provide values to change in the record here
		"location": "home",
	}

	// Send the PUT request to update the record
	updatedRecord, err := req.Put2(values, nil)

	// Perform assertions and test the response
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// You should assert that the updatedRecord matches your expected response here.
	// You may need to unmarshal the JSON response and compare specific fields.
	// For example: assert.Equal(t, expectedValue, updatedRecord.Field)

	assert.Equal(t, fakeTableItemResponse, updatedRecord)
}
