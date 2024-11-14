package tableapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func builder(uri string) (*TableRequestBuilder, error) {
	client := &MockClient{}

	parsedURL, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	pathParameters := map[string]string{"baseurl": "http://" + parsedURL.Host, "table": parsedURL.Path}

	return NewTableRequestBuilder(context.Background(), client, pathParameters), nil
}

func TestNewTableRequestBuilder(t *testing.T) {
	client := MockClient{}

	pathParameters := map[string]string{"baseurl": "https://instance.service-now.com/api/now", "table": "table1"}

	req := NewTableItemRequestBuilder(context.Background(), &client, pathParameters)

	assert.NotNil(t, req)
}

func TestTableUrl(t *testing.T) {
	client := MockClient{}

	pathParameters := map[string]string{"baseurl": "https://instance.service-now.com/api/now", "table": "table1"}

	req := NewTableItemRequestBuilder(context.Background(), &client, pathParameters)

	assert.Equal(t, req.PathParameters, pathParameters)

	if !reflect.DeepEqual(req.PathParameters, pathParameters) {
		t.Errorf("excepted: %s, got: %s", pathParameters, req.PathParameters)
	}
}

func TestTableRequestBuilder_Get(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(getFakeCollectionJSON())
	}))

	client := &MockClient{}

	parsedURL, err := url.Parse(mockServer.URL)
	assert.Nil(t, err)

	pathParameters := map[string]string{"baseurl": "http://" + parsedURL.Host, "table": parsedURL.Path}

	builder := NewTableRequestBuilder(context.Background(), client, pathParameters)

	// Call the Get method
	resp, err := builder.Get(context.Background(), nil)
	assert.Nil(t, err)

	assert.NotNil(t, resp)
	assert.IsType(t, &TableCollectionResponse{}, resp)
	assert.Len(t, resp.Result, 1)
	assert.Equal(t, &fakeEntry, resp.Result[0])
}

//nolint:dupl
func TestTableRequestBuilder_Post(t *testing.T) {
	// Create a mock mockServer
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"some": "response"}`))
	}))
	defer mockServer.Close()

	builder, err := builder(mockServer.URL)
	assert.Nil(t, err)

	queryParameters := &TableRequestBuilderPostQueryParamters{
		DisplayValue:         "true",
		ExcludeReferenceLink: true,
		Fields:               []string{"field1", "field2"},
		InputDisplayValue:    true,
		View:                 "desktop",
	}

	tests := []test[*TableItemResponse]{
		{
			title:     "ValidRequest",
			value:     map[string]string{"key": "value"},
			expected:  nil,
			expectErr: false,
			err:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			response, err := builder.Post(tt.value.(map[string]string), queryParameters)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)

			assert.IsType(t, &TableItemResponse{}, response)
		})
	}
}

//nolint:dupl
func TestTableRequestBuilder_Post2(t *testing.T) {
	// Create a mock mockServer
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"some": "response"}`))
	}))
	defer mockServer.Close()

	builder, err := builder(mockServer.URL)
	assert.Nil(t, err)

	queryParameters := &TableRequestBuilderPostQueryParameters{
		DisplayValue:         "true",
		ExcludeReferenceLink: true,
		Fields:               []string{"field1", "field2"},
		InputDisplayValue:    true,
		View:                 "desktop",
	}

	tests := []test[*TableItemResponse]{
		{
			title:     "ValidRequest",
			value:     map[string]string{"key": "value"},
			expected:  nil,
			expectErr: false,
			err:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			response, err := builder.Post2(context.Background(), tt.value.(map[string]string), queryParameters)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)

			assert.IsType(t, &TableItemResponse{}, response)
		})
	}

	t.Run("ValidRequest", func(t *testing.T) {
		// Call the Post method with valid parameters
		response, err := builder.Post2(context.Background(), map[string]string{"key": "value"}, queryParameters)
		assert.Nil(t, err)

		assert.IsType(t, &TableItemResponse{}, response)
	})
}

//nolint:dupl
func TestTableRequestBuilder_Post3(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"some": "response"}`))
	}))
	defer mockServer.Close()

	builder, err := builder(mockServer.URL)
	assert.Nil(t, err)

	queryParameters := &TableRequestBuilderPostQueryParameters{
		DisplayValue:         "true",
		ExcludeReferenceLink: true,
		Fields:               []string{"field1", "field2"},
		InputDisplayValue:    true,
		View:                 "desktop",
	}

	tests := []test[*TableItemResponse2[TableEntry]]{
		{
			title:     "ValidRequest-map[string]string",
			value:     map[string]string{"key": "value"},
			expected:  nil,
			expectErr: false,
			err:       nil,
		},
		{
			title:     "ValidRequest-tableEntry",
			value:     TableEntry{"key": "value"},
			expected:  nil,
			expectErr: false,
			err:       nil,
		},
		{
			title:     "InvalidRequest",
			value:     "bad",
			expected:  nil,
			expectErr: true,
			err:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			response, err := builder.Post3(tt.value, queryParameters)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)

			assert.IsType(t, &TableItemResponse{}, response)
		})
	}
}

func TestTableRequestBuilder_Count(t *testing.T) {
	client := &mockClient{}

	pathParameters := map[string]string{"baseurl": fakeItemCountLinkKey, "table": "table1"}

	builder := NewTableRequestBuilder(context.Background(), client, pathParameters)

	// Call the Get method
	count, err := builder.Count()
	assert.Nil(t, err)
	assert.NotEqual(t, -1, count)
	assert.IsType(t, 1, count)
}
