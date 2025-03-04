package servicenowsdkgo

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewCredentialAuthenticationProviderAdapter(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "successful",
			test: func(t *testing.T) {
				cred := mocking.NewMockCredential()

				authenticationProvider, err := newCredentialAuthenticationProviderAdapter(cred)
				assert.Nil(t, err)
				assert.Equal(t, &credentialAuthenticationProviderAdapter{
					cred: cred,
				}, authenticationProvider)
			},
		},
		{
			name: "nil credential",
			test: func(t *testing.T) {
				authenticationProvider, err := newCredentialAuthenticationProviderAdapter(nil)
				assert.Equal(t, errors.New("credential is nil"), err)
				assert.Nil(t, authenticationProvider)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCredentialAuthenticationProviderAdapter_AuthenticateRequest(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
