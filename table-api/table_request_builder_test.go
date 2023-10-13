package tableapi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTableRequestBuilder(t *testing.T) {
	client := &MockClient{} // Use a pointer to the mock client.

	pathParameters := map[string]string{"baseurl": "instance.service-now.com", "table": "table1"}

	req := NewTableRequestBuilder(client, pathParameters)

	assert.NotNil(t, req)
}

func TestTableRequestBuilderById(t *testing.T) {
	client := &MockClient{} // Use a pointer to the mock client.

	pathParameters := map[string]string{"baseurl": "instance.service-now.com", "table": "table1"}
	req := NewTableRequestBuilder(client, pathParameters)

	sysID := "sysid"
	tableItemRequestBuilder := req.ById(sysID)

	assert.NotNil(t, tableItemRequestBuilder)
	assert.Equal(t, tableItemRequestBuilder.RequestBuilder.PathParameters["sysId"], sysID)
}

func TestTableRequestBuilderGet(t *testing.T) {
	client := &MockClient{} // Use a pointer to the mock client.

	pathParameters := map[string]string{"baseurl": "instance.service-now.com", "table": "table1"}
	req := NewTableRequestBuilder(client, pathParameters)

	params := &TableRequestBuilderGetQueryParameters{
		DisplayValue:             TRUE,
		ExcludeReferenceLink:     true,
		Fields:                   []string{"field1", "field2"},
		QueryNoDomain:            true,
		View:                     DESKTOP,
		Limit:                    10,
		NoCount:                  false,
		Offset:                   0,
		Query:                    "field=value",
		QueryCategory:            "category",
		SuppressPaginationHeader: true,
	}

	response, err := req.Get(params)

	assert.Nil(t, err)         // Assert that there's no error.
	assert.NotNil(t, response) // Assert that the response is not nil.
}

func TestTableRequestBuilderPOST(t *testing.T) {
	client := &MockClient{} // Use a pointer to the mock client.

	pathParameters := map[string]string{"baseurl": "instance.service-now.com", "table": "table1"}
	req := NewTableRequestBuilder(client, pathParameters)

	data := map[string]interface{}{"field1": "value1", "field2": "value2"}
	params := &TableRequestBuilderPostQueryParamters{
		DisplayValue:         FALSE,
		ExcludeReferenceLink: false,
		Fields:               []string{"field1", "field2"},
		InputDisplayValue:    true,
		View:                 MOBILE,
	}

	response, err := req.POST(data, params)

	assert.Nil(t, err)         // Assert that there's no error.
	assert.NotNil(t, response) // Assert that the response is not nil.
}

func TestTableRequestBuilderCount(t *testing.T) {
	client := &MockClient{} // Use a pointer to the mock client.

	pathParameters := map[string]string{"baseurl": "instance.service-now.com", "table": "table1"}
	req := NewTableRequestBuilder(client, pathParameters)

	// Create a mock response with a "X-Total-Count" header.
	response := &http.Response{
		StatusCode: 200,
		Header:     http.Header{},
	}
	response.Header.Add("X-Total-Count", "42")

	// Mock the Send method to return the response.
	//client.Send = func(requestInfo *abstraction.RequestInformation, errorMapping abstraction.ErrorMapping) (*http.Response, error) {
	//	return response, nil
	//}

	count, err := req.Count()

	assert.Nil(t, err)         // Assert that there's no error.
	assert.Equal(t, count, 42) // Assert that the count is as expected.
}
