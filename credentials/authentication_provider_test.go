package credentials

import (
	"testing"

	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticationProviders_Initialization(t *testing.T) {
	baseURL := "https://dev12345.service-now.com"

	tests := []struct {
		name     string
		provider func() (authentication.AuthenticationProvider, error)
		verify   func(t *testing.T, p authentication.AuthenticationProvider)
	}{
		{
			name: "ROPC Provider with Options",
			provider: func() (authentication.AuthenticationProvider, error) {
				return NewROPCProvider("client-id", "client-secret", "username", "password",
					WithURL(baseURL),
				)
			},
			verify: func(t *testing.T, p authentication.AuthenticationProvider) {
				bearer := p.(*BearerTokenAuthenticationProvider)
				tp := bearer.tokenProvider.(*ROPCCredential)
				assert.Equal(t, baseURL, tp.baseURL)
				assert.NotNil(t, tp.GetAllowedHostsValidator())
			},
		},
		{
			name: "Client Credentials Provider with Options",
			provider: func() (authentication.AuthenticationProvider, error) {
				return NewClientCredentialsProvider("client-id", "client-secret",
					WithURL(baseURL),
				)
			},
			verify: func(t *testing.T, p authentication.AuthenticationProvider) {
				bearer := p.(*BearerTokenAuthenticationProvider)
				tp := bearer.tokenProvider.(*ClientCredentialsCredential)
				assert.Equal(t, baseURL, tp.baseURL)
				assert.NotNil(t, tp.GetAllowedHostsValidator())
			},
		},
		{
			name: "Authorization Code Provider with Options",
			provider: func() (authentication.AuthenticationProvider, error) {
				return NewAuthorizationCodeProvider("client-id", "client-secret",
					WithURL(baseURL),
				)
			},
			verify: func(t *testing.T, p authentication.AuthenticationProvider) {
				bearer := p.(*BearerTokenAuthenticationProvider)
				tp := bearer.tokenProvider.(*AuthorizationCodeCredential)
				assert.Equal(t, baseURL, tp.baseURL)
				assert.NotNil(t, tp.GetAllowedHostsValidator())
				assert.Equal(t, 5001, tp.port) // Default port
			},
		},
		{
			name: "Authorization Code Provider with Custom Options",
			provider: func() (authentication.AuthenticationProvider, error) {
				customStateGen := func() string { return "custom-state" }
				customURLOpener := func(url string) error { return nil }
				return NewAuthorizationCodeProvider("client-id", "client-secret",
					WithPort(8080),
					WithStateGenerator(customStateGen),
					WithURLOpener(customURLOpener),
				)
			},
			verify: func(t *testing.T, p authentication.AuthenticationProvider) {
				bearer := p.(*BearerTokenAuthenticationProvider)
				tp := bearer.tokenProvider.(*AuthorizationCodeCredential)
				assert.Equal(t, 8080, tp.port)
				assert.NotNil(t, tp.stateGenerator)
				assert.Equal(t, "custom-state", tp.stateGenerator())
				assert.NotNil(t, tp.urlOpener)
			},
		},
		{
			name: "Authorization Code Provider with Custom Server Factory",
			provider: func() (authentication.AuthenticationProvider, error) {
				customServerFactory := func(state string, port int) (AuthorizationCodeServer, error) {
					m := &mockAuthorizationCodeServer{}
					m.On("GetAddr").Return("http://localhost:1234")
					return m, nil
				}
				return NewAuthorizationCodeProvider("client-id", "client-secret",
					WithServerFactory(customServerFactory),
				)
			},
			verify: func(t *testing.T, p authentication.AuthenticationProvider) {
				bearer := p.(*BearerTokenAuthenticationProvider)
				tp := bearer.tokenProvider.(*AuthorizationCodeCredential)
				assert.NotNil(t, tp.serverFactory)
				s, _ := tp.serverFactory("state", 0)
				assert.Equal(t, "http://localhost:1234", s.GetAddr())
			},
		},
		{
			name: "JWT Provider with Options",
			provider: func() (authentication.AuthenticationProvider, error) {
				// Mock token provider for JWT
				mockTP := &BaseAccessTokenProvider{}
				return NewJWTProvider("client-id", "client-secret", mockTP,
					WithURL(baseURL),
				)
			},
			verify: func(t *testing.T, p authentication.AuthenticationProvider) {
				bearer := p.(*BearerTokenAuthenticationProvider)
				tp := bearer.tokenProvider.(*JWTCredential)
				assert.Equal(t, baseURL, tp.baseURL)
				assert.NotNil(t, tp.GetAllowedHostsValidator())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := tt.provider()
			assert.NoError(t, err)

			tt.verify(t, p)
		})
	}
}
