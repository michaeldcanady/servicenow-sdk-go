package tests

import (
	"testing"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/abstraction"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	cred := abstraction.NewUsernamePasswordCredential("username", "password")

	client := servicenowsdkgo.NewClient(cred, "instance")

	assert.NotNil(t, client)
}

func TestClientURL(t *testing.T) {

	cred := abstraction.NewUsernamePasswordCredential("username", "password")

	client := servicenowsdkgo.NewClient(cred, "instance")

	assert.Equal(t, client.BaseUrl, "https://instance.service-now.com/api")

}
