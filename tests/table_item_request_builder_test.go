package tests

import (
	"testing"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestNewTableItemRequestBuilder(t *testing.T) {

	cred := servicenowsdkgo.NewUsernamePasswordCredential("username", "password")

	client := servicenowsdkgo.NewClient(cred, "instance")

	req := client.Now().Table("table1").ById("sysid")

	assert.NotNil(t, req)
}

func TestTableItemUrl(t *testing.T) {

	cred := servicenowsdkgo.NewUsernamePasswordCredential("username", "password")

	client := servicenowsdkgo.NewClient(cred, "instance")

	req := client.Now().Table("table1").ById("sysid")

	assert.Equal(t, req.PathParameters, map[string]string{"baseurl": "https://instance.service-now.com/api/now", "table": "table1", "sysId": "sysid"})
}
