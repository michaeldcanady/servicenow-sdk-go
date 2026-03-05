package credentials

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUsernamePasswordCredential(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
	}{
		{
			name:     "Standard credentials",
			username: "testuser",
			password: "testpassword",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			credential := NewUsernamePasswordCredential(test.username, test.password)
			assert.Equal(t, test.username, credential.Username)
			assert.Equal(t, test.password, credential.Password)
		})
	}
}

func TestUsernamePasswordCredential_GetAuthentication(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
	}{
		{
			name:     "Standard credentials",
			username: "testuser",
			password: "testpassword",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			credential := NewUsernamePasswordCredential(test.username, test.password)
			authHeader, err := credential.GetAuthentication()
			assert.NoError(t, err)

			expectedAuthHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(test.username+":"+test.password))
			assert.Equal(t, expectedAuthHeader, authHeader)
		})
	}
}

func TestUsernamePasswordCredential_BasicAuth(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
	}{
		{
			name:     "Standard credentials",
			username: "testuser",
			password: "testpassword",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			credential := NewUsernamePasswordCredential(test.username, test.password)
			authHeader := credential.BasicAuth(test.username, test.password)

			expectedEncodedAuth := base64.StdEncoding.EncodeToString([]byte(test.username + ":" + test.password))
			assert.Equal(t, expectedEncodedAuth, authHeader)
		})
	}
}
