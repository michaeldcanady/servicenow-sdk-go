package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

func (c *MockClient) Send(requestInfo *servicenowsdkgo.RequestInformation, errorMapping servicenowsdkgo.ErrorMapping) (*http.Response, error) {
	// Mock the client's behavior here.
	// You can create a mock response for testing purposes.
	response := &http.Response{
		StatusCode: 200, // Mock the status code you expect.
		Body:       ioutil.NopCloser(strings.NewReader("")), // Mock an empty response body.
	}
	return response, nil
}

func TestNewTableItemRequestBuilder(t *testing.T) {
	client := MockClient{}

 pathParameters := map[string]string{"baseurl":"instance.service-now.com", "table":"table1", "sysId":"sysid"}

 req := NewTableItemRequestBuilder(client, pathParameters)

	assert.NotNil(t, req)

 assert.Equal(t, req.PathParameters, pathParameters)
}
