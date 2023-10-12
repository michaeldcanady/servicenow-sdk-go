package tests

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewClient(cred, "instance")

	assert.NotNil(t, client)
}

func TestClientURL(t *testing.T) {

	cred := credentials.NewUsernamePasswordCredential("username", "password")

	client := NewClient(cred, "instance")

	assert.Equal(t, client.BaseUrl, "https://instance.service-now.com/api")

}
