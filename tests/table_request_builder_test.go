package tests

import (
	"testing"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestNewTableRequestBuilder(t *testing.T) {

	cred := servicenowsdkgo.NewUsernamePasswordCredential("username", "password")

	client := servicenowsdkgo.NewClient(cred, "instance")

	req := client.Now().Table("table1")

	assert.NotNil(t, req)
}

func TestTableUrl(t *testing.T) {

	cred := servicenowsdkgo.NewUsernamePasswordCredential("username", "password")

	client := servicenowsdkgo.NewClient(cred, "instance")

	req := client.Now().Table("table1")

	assert.Equal(t, req.Url, "instance.service-now.com/api/now/table/table1")
}
