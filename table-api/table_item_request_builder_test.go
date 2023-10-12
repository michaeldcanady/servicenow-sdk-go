package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/assert"
)

type MockCredential struct{}

func (c *MockCredential) GetAuthentication() (string, error) {
	// Mock authentication logic here, return a string (e.g., a token) and nil error.
	return "mocked-auth-token", nil
}

func TestNewTableItemRequestBuilder(t *testing.T) {
	cred := &MockCredential{}

	client := servicenowsdkgo.NewClient(cred, "instance")

  pathParameters := map[string]string{"baseurl":"instance.service-now.com", "table":"table1", "sysId":"sysid"}

  req := NewTableItemRequestBuilder(client, pathParameters)

	assert.NotNil(t, req)
}

func TestTableItemUrl(t *testing.T) {
	cred := &MockCredential{}

	client := servicenowsdkgo.NewClient(cred, "instance")

	req := client.Now().Table("table1").ById("sysid")

	assert.Equal(t, req.PathParameters, map[string]string{"baseurl": "https://instance.service-now.com/api/now", "table": "table1", "sysId": "sysid"})
}
