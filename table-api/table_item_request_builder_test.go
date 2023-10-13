package tableapi

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/abstraction"
	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

func (c *MockClient) Send(requestInfo *abstraction.RequestInformation, errorMapping abstraction.ErrorMapping) (*http.Response, error) {
	// Mock the client's behavior here.
	// You can create a mock response for testing purposes.
	response := &http.Response{
		StatusCode: 200,                                 // Mock the status code you expect.
		Body:       io.NopCloser(strings.NewReader("")), // Mock an empty response body.
	}
	return response, nil
}

func TestNewTableItemRequestBuilder(t *testing.T) {
	client := &MockClient{} // Use a pointer to the mock client.

	pathParameters := map[string]string{"baseurl": "instance.service-now.com", "table": "table1", "sysId": "sysid"}

	req := NewTableItemRequestBuilder(client, pathParameters)

	assert.NotNil(t, req)
}

func TestTableItemRequestBuilderGet(t *testing.T) {
	client := &MockClient{} // Use a pointer to the mock client.

	pathParameters := map[string]string{"baseurl": "instance.service-now.com", "table": "table1", "sysId": "sysid"}
	req := NewTableItemRequestBuilder(client, pathParameters)

	params := &TableItemRequestBuilderGetQueryParameters{
		DisplayValue:         "true",
		ExcludeReferenceLink: true,
		Fields:               []string{"field1", "field2"},
		QueryNoDomain:        true,
		View:                 "desktop",
	}

	response, err := req.Get(params)

	assert.Nil(t, err)         // Assert that there's no error.
	assert.NotNil(t, response) // Assert that the response is not nil.
}

func TestTableItemRequestBuilderDelete(t *testing.T) {
	client := &MockClient{} // Use a pointer to the mock client.

	pathParameters := map[string]string{"baseurl": "instance.service-now.com", "table": "table1", "sysId": "sysid"}
	req := NewTableItemRequestBuilder(client, pathParameters)

	params := &TableItemRequestBuilderDeleteQueryParameters{
		QueryNoDomain: true,
	}

	err := req.Delete(params)

	assert.Nil(t, err) // Assert that there's no error.
}
