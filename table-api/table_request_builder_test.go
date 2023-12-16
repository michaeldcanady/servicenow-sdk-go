package tableapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTableRequestBuilder(t *testing.T) {
	client := MockClient{}

	pathParameters := map[string]string{"baseurl": "https://instance.service-now.com/api/now", "table": "table1"}

	req := NewTableItemRequestBuilder(&client, pathParameters)

	assert.NotNil(t, req)
}

func TestTableUrl(t *testing.T) {
	client := MockClient{}

	pathParameters := map[string]string{"baseurl": "https://instance.service-now.com/api/now", "table": "table1"}

	req := NewTableItemRequestBuilder(&client, pathParameters)

	assert.Equal(t, req.PathParameters, pathParameters)

	if !reflect.DeepEqual(req.PathParameters, pathParameters) {
		t.Errorf("excepted: %s, got: %s", pathParameters, req.PathParameters)
	}
}

func TestTableRequestBuilderGet(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		// Simulate successful response with the provided JSON
		responseJSON, err := json.Marshal(fakeCollectionResult)
		assert.Nil(t, err)

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseJSON) //nolint:errcheck
	}))

	client := &MockClient{}

	parsedURL, err := url.Parse(mockServer.URL)
	assert.Nil(t, err)

	pathParameters := map[string]string{"baseurl": "http://" + parsedURL.Host, "table": parsedURL.Path}

	builder := NewTableRequestBuilder(client, pathParameters)

	// Call the Get method
	resp, err := builder.Get(nil)
	assert.Nil(t, err)

	assert.NotNil(t, resp)
	assert.IsType(t, &TableCollectionResponse{}, resp)
	assert.Len(t, resp.Result, 1)
	//assert.Equal(t, expectedTableEntry, resp.Result[0])
}

//nolint:dupl
func TestTableRequestBuilderPost(t *testing.T) {
	t.Run("ValidRequest", func(t *testing.T) {
		// Create a mock mockServer
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "POST", r.Method)

			// Handle the request and send a mock response
			// You can customize this based on your actual implementation
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"some": "response"}`))
		}))
		defer mockServer.Close()

		client := &MockClient{}

		parsedURL, err := url.Parse(mockServer.URL)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
			return
		}

		pathParameters := map[string]string{"baseurl": "http://" + parsedURL.Host, "table": parsedURL.Path}

		builder := NewTableRequestBuilder(client, pathParameters)

		// Call the Post method with valid parameters
		response, err := builder.Post(map[string]string{"key": "value"}, &TableRequestBuilderPostQueryParamters{
			DisplayValue:         "true",
			ExcludeReferenceLink: true,
			Fields:               []string{"field1", "field2"},
			InputDisplayValue:    true,
			View:                 "desktop",
		})

		// Check if there are no errors and the response is as expected
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		assert.IsType(t, &TableItemResponse{}, response)
	})

	t.Run("ValidRequestWithNilParams", func(t *testing.T) {
		// ... Similar to the previous test but with nil parameters ...
	})

	t.Run("ServerError", func(t *testing.T) {
		// ... Test when the server returns an error (e.g., 500 Internal Server Error) ...
	})
}

//nolint:dupl
func TestTableRequestBuilderPost2(t *testing.T) {
	t.Run("ValidRequest", func(t *testing.T) {
		// Create a mock mockServer
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "POST", r.Method)

			// Handle the request and send a mock response
			// You can customize this based on your actual implementation
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"some": "response"}`))
		}))
		defer mockServer.Close()

		client := &MockClient{}

		parsedURL, err := url.Parse(mockServer.URL)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
			return
		}

		pathParameters := map[string]string{"baseurl": "http://" + parsedURL.Host, "table": parsedURL.Path}

		builder := NewTableRequestBuilder(client, pathParameters)

		// Call the Post method with valid parameters
		response, err := builder.Post2(map[string]string{"key": "value"}, &TableRequestBuilderPostQueryParameters{
			DisplayValue:         "true",
			ExcludeReferenceLink: true,
			Fields:               []string{"field1", "field2"},
			InputDisplayValue:    true,
			View:                 "desktop",
		})

		// Check if there are no errors and the response is as expected
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		assert.IsType(t, &TableItemResponse{}, response)
	})

	t.Run("ValidRequestWithNilParams", func(t *testing.T) {
		// ... Similar to the previous test but with nil parameters ...
	})

	t.Run("ServerError", func(t *testing.T) {
		// ... Test when the server returns an error (e.g., 500 Internal Server Error) ...
	})
}

func TestTableRequestBuilderCount(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate successful response with the provided JSON
		responseJSON := ``

		w.WriteHeader(http.StatusOK)
		w.Header().Add("X-Total-Count", "1")
		_, _ = w.Write([]byte(responseJSON)) //nolint:errcheck
	}))
	defer mockServer.Close()

	client := &MockClient{}

	// Create an instance of TableRequestBuilder using the mock server URL

	parsedURL, err := url.Parse(mockServer.URL)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	pathParameters := map[string]string{"baseurl": "http://" + parsedURL.Host, "table": parsedURL.Path}

	builder := NewTableRequestBuilder(client, pathParameters)

	// Call the Get method
	count, err := builder.Count()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
		return
	}

	// You can further validate the response if needed
	if count == -1 {
		t.Error("Expected a non-nil response, but got nil")
	}

	expectedType := reflect.Int

	if reflect.TypeOf(count).Kind() != expectedType {
		t.Errorf("Expected response of type %v, but got type %v", expectedType, reflect.TypeOf(count))
	}
}
