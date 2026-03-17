package credentials

import (
	"testing"

	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationProviders_Initialization(t *testing.T) {
	instance := "dev12345"
	baseURL := "https://dev12345.service-now.com"

	tests := []struct {
		name     string
		provider func() (authentication.AuthenticationProvider, error)
		verify   func(t *testing.T, p authentication.AuthenticationProvider)
	}{
		{
			name: "ROPC Provider",
			provider: func() (authentication.AuthenticationProvider, error) {
				return NewROPCProvider("client-id", "client-secret", "username", "password")
			},
			verify: func(t *testing.T, p authentication.AuthenticationProvider) {
				bearer := p.(*BearerTokenAuthenticationProvider)
				tp := bearer.tokenProvider.(*ROPCCredential)
				assert.Equal(t, instance, tp.instance)
				assert.Equal(t, baseURL, tp.baseURL)
				assert.NotNil(t, tp.GetAllowedHostsValidator())
			},
		},
		{
			name: "Client Credentials Provider",
			provider: func() (authentication.AuthenticationProvider, error) {
				return NewClientCredentialsProvider("client-id", "client-secret")
			},
			verify: func(t *testing.T, p authentication.AuthenticationProvider) {
				bearer := p.(*BearerTokenAuthenticationProvider)
				tp := bearer.tokenProvider.(*ClientCredentialsCredential)
				assert.Equal(t, instance, tp.instance)
				assert.Equal(t, baseURL, tp.baseURL)
				assert.NotNil(t, tp.GetAllowedHostsValidator())
			},
		},
		{
			name: "Authorization Code Provider",
			provider: func() (authentication.AuthenticationProvider, error) {
				return NewAuthorizationCodeProvider("client-id", "client-secret")
			},
			verify: func(t *testing.T, p authentication.AuthenticationProvider) {
				bearer := p.(*BearerTokenAuthenticationProvider)
				tp := bearer.tokenProvider.(*AuthorizationCodeCredential)
				assert.Equal(t, instance, tp.instance)
				assert.Equal(t, baseURL, tp.baseURL)
				assert.NotNil(t, tp.GetAllowedHostsValidator())
			},
		},
		{
			name: "JWT Provider",
			provider: func() (authentication.AuthenticationProvider, error) {
				// Mock token provider for JWT
				mockTP := &BaseAccessTokenProvider{}
				return NewJWTProvider("client-id", "client-secret", mockTP)
			},
			verify: func(t *testing.T, p authentication.AuthenticationProvider) {
				bearer := p.(*BearerTokenAuthenticationProvider)
				tp := bearer.tokenProvider.(*JWTCredential)
				assert.Equal(t, instance, tp.instance)
				assert.Equal(t, baseURL, tp.baseURL)
				assert.NotNil(t, tp.GetAllowedHostsValidator())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := tt.provider()
			assert.NoError(t, err)

			preparable, ok := p.(Preparable)
			assert.True(t, ok, "Provider should be Preparable")

			preparable.Initialize(instance, baseURL)
			tt.verify(t, p)
		})
	}
}
