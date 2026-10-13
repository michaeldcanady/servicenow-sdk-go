// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2/pkce"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var mockAnyContext = mock.Anything

// TODO: should be a Mock.mock
type authCodeTestClient struct {
	initializedWith string
	authURL         string
	authVerifier    string
	authURLErr      error
	acquireToken    *AccessToken
	acquireErr      error
	refreshToken    *AccessToken
	refreshErr      error
	revokeErr       error
	lastVerifier    string
}

func (c *authCodeTestClient) Initialize(baseURL string) {
	c.initializedWith = baseURL
}

func (c *authCodeTestClient) getAuthorizationURL(_, _ string, _ []string) (string, string, error) {
	return c.authURL, c.authVerifier, c.authURLErr
}

func (c *authCodeTestClient) acquireTokenByCode(_ context.Context, _, _, _, verifier string) (*AccessToken, error) {
	c.lastVerifier = verifier
	return c.acquireToken, c.acquireErr
}

func (c *authCodeTestClient) acquireTokenByRefreshToken(_ context.Context, _ string) (*AccessToken, error) {
	return c.refreshToken, c.refreshErr
}

func (c *authCodeTestClient) revokeToken(_ context.Context, _, _ string) error {
	return c.revokeErr
}

func authCodeTestServerFactory(server *mockAuthorizationCodeServer) ServerFactory {
	return func(_ string, _ int) (AuthorizationCodeServer, error) {
		return server, nil
	}
}

func TestAuthorizationCodeCredential_GetToken(t *testing.T) {
	tests := []struct {
		name          string
		client        *authCodeTestClient
		serverFactory ServerFactory
		urlOpener     func(string) error
		wantErr       string
		wantToken     *AccessToken
	}{
		{
			name:   "server factory error propagates",
			client: &authCodeTestClient{},
			serverFactory: func(_ string, _ int) (AuthorizationCodeServer, error) {
				return nil, errors.New("server start failed")
			},
			wantErr: "server start failed",
		},
		{
			name:   "authorization url error propagates",
			client: &authCodeTestClient{authURLErr: errors.New("auth url failed")},
			serverFactory: authCodeTestServerFactory(func() *mockAuthorizationCodeServer {
				m := &mockAuthorizationCodeServer{}
				m.On("GetAddr").Return("http://localhost:5001")
				m.On("Shutdown", mockAnyContext).Return(nil)
				return m
			}()),
			wantErr: "auth url failed",
		},
		{
			name:   "url opener error propagates",
			client: &authCodeTestClient{authURL: "https://auth.example.com"},
			serverFactory: authCodeTestServerFactory(func() *mockAuthorizationCodeServer {
				m := &mockAuthorizationCodeServer{}
				m.On("GetAddr").Return("http://localhost:5001")
				m.On("Shutdown", mockAnyContext).Return(nil)
				return m
			}()),
			urlOpener: func(_ string) error { return errors.New("open failed") },
			wantErr:   "open failed",
		},
		{
			name:   "server result error propagates",
			client: &authCodeTestClient{authURL: "https://auth.example.com"},
			serverFactory: authCodeTestServerFactory(func() *mockAuthorizationCodeServer {
				m := &mockAuthorizationCodeServer{}
				m.On("GetAddr").Return("http://localhost:5001")
				m.On("Result", mockAnyContext).Return("", "", errors.New("callback failed"))
				m.On("Shutdown", mockAnyContext).Return(nil)
				return m
			}()),
			wantErr: "callback failed",
		},
		{
			name:   "code exchange error propagates",
			client: &authCodeTestClient{authURL: "https://auth.example.com", acquireErr: errors.New("exchange failed")},
			serverFactory: authCodeTestServerFactory(func() *mockAuthorizationCodeServer {
				m := &mockAuthorizationCodeServer{}
				m.On("GetAddr").Return("http://localhost:5001")
				m.On("Result", mockAnyContext).Return("code", "state", nil)
				m.On("Shutdown", mockAnyContext).Return(nil)
				return m
			}()),
			wantErr: "exchange failed",
		},
		{
			name:   "successful flow returns token",
			client: &authCodeTestClient{authURL: "https://auth.example.com", authVerifier: "test-verifier", acquireToken: &AccessToken{AccessToken: "final"}},
			serverFactory: authCodeTestServerFactory(func() *mockAuthorizationCodeServer {
				m := &mockAuthorizationCodeServer{}
				m.On("GetAddr").Return("http://localhost:5001")
				m.On("Result", mockAnyContext).Return("code", "state", nil)
				m.On("Shutdown", mockAnyContext).Return(nil)
				return m
			}()),
			wantToken: &AccessToken{AccessToken: "final"},
		},
		{
			name:   "shutdown error propagates when no other error occurred",
			client: &authCodeTestClient{authURL: "https://auth.example.com", authVerifier: "test-verifier", acquireToken: &AccessToken{AccessToken: "final"}},
			serverFactory: authCodeTestServerFactory(func() *mockAuthorizationCodeServer {
				m := &mockAuthorizationCodeServer{}
				m.On("GetAddr").Return("http://localhost:5001")
				m.On("Result", mockAnyContext).Return("code", "state", nil)
				m.On("Shutdown", mockAnyContext).Return(errors.New("shutdown failed"))
				return m
			}()),
			wantErr: "shutdown failed",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			urlOpener := test.urlOpener
			if urlOpener == nil {
				urlOpener = func(_ string) error { return nil }
			}

			credential := newAuthorizationCodeCredential(test.client, 0, func() string { return "state" }, urlOpener, test.serverFactory)

			token, err := credential.GetToken(context.Background(), nil, nil)

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.wantToken, token)
			assert.Equal(t, test.client.authVerifier, test.client.lastVerifier, "GetToken must thread the verifier from the authorization URL to the code exchange")
		})
	}
}

func TestAuthorizationCodeCredential_Initialize(t *testing.T) {
	client := &authCodeTestClient{}
	credential := newAuthorizationCodeCredential(client, 0, nil, nil, nil)

	credential.Initialize("https://dev12345.service-now.com")

	assert.Equal(t, "https://dev12345.service-now.com", client.initializedWith)
}

func TestNewAuthorizationCodeCredential_Defaults(t *testing.T) {
	client := &authCodeTestClient{}
	credential := newAuthorizationCodeCredential(client, 0, nil, nil, nil)
	require.NotNil(t, credential)
	assert.Equal(t, 5001, credential.port)
	assert.NotNil(t, credential.stateGenerator)
	assert.NotNil(t, credential.urlOpener)
	assert.NotNil(t, credential.serverFactory)
}

func TestNewPrivateAuthorizationCodeProvider(t *testing.T) {
	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		wantErr      error
	}{
		{
			name:         "empty client id fails",
			clientID:     "",
			clientSecret: "secret",
			wantErr:      snerrors.ErrEmptyClientID,
		},
		{
			// NewPrivateAuthorizationCodeProvider still falls back to a public (PKCE) client
			// when the secret is blank - only NewConfidentialClient rejects an empty secret.
			name:         "empty client secret falls back to a public client",
			clientID:     "id",
			clientSecret: "",
		},
		{
			name:         "valid credentials succeed",
			clientID:     "id",
			clientSecret: "secret",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider, err := NewPrivateAuthorizationCodeProvider(test.clientID, test.clientSecret, AuthorizationCodeOptions{
				CommonOptions: CommonOptions{
					AllowedHosts: []string{"dev12345.service-now.com"},
				},
			})

			if test.wantErr != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.wantErr)
				assert.Nil(t, provider)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, provider)
		})
	}
}

func TestPublicAuthorizationCodeProvider_EndToEnd(t *testing.T) {
	rt := &capturingTokenEndpoint{body: []byte(`{"access_token":"e2e-token","expires_in":3600}`)}

	var authURL string
	urlOpener := func(u string) error {
		authURL = u
		return nil
	}

	server := &mockAuthorizationCodeServer{}
	server.On("GetAddr").Return("http://localhost:5001")
	server.On("Result", mockAnyContext).Return("e2e-code", "state", nil)
	server.On("Shutdown", mockAnyContext).Return(nil)

	provider, err := NewPublicAuthorizationCodeProvider("public-client",
		WithURL("https://dev12345.service-now.com"),
		WithHTTPClient(&http.Client{Transport: rt}),
		WithURLOpener(urlOpener),
		WithServerFactory(func(_ string, _ int) (AuthorizationCodeServer, error) { return server, nil }),
	)
	require.NoError(t, err)

	request := abstractions.NewRequestInformation()
	request.SetUri(*mustParseURL(t, "https://dev12345.service-now.com/now/table/incident"))
	require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))

	header := request.Headers.Get("Authorization")
	require.Len(t, header, 1)
	assert.Equal(t, "Bearer e2e-token", header[0])

	// The token request must present the verifier matching the authorization
	// URL's challenge and must not carry a client secret.
	require.NotNil(t, rt.lastForm)
	assert.Equal(t, "public-client", rt.lastForm.Get("client_id"))
	assert.Empty(t, rt.lastForm.Get("client_secret"))

	u, err := url.Parse(authURL)
	require.NoError(t, err)
	challenge := u.Query().Get("code_challenge")
	require.NotEmpty(t, challenge)
	verifier := rt.lastForm.Get("code_verifier")
	require.NotEmpty(t, verifier)
	want, err := pkce.NewPKCEChallenge(pkce.MethodS256, verifier)
	require.NoError(t, err)
	assert.Equal(t, want, challenge)
}

func TestNewPublicAuthorizationCodeProvider(t *testing.T) {
	tests := []struct {
		name     string
		clientID string
		wantErr  error
	}{
		{
			name:     "empty client id fails",
			clientID: "",
			wantErr:  snerrors.ErrEmptyClientID,
		},
		{
			name:     "valid client id succeeds",
			clientID: "id",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider, err := NewPublicAuthorizationCodeProvider(test.clientID, AuthorizationCodeOptions{
				CommonOptions: CommonOptions{
					AllowedHosts: []string{"dev12345.service-now.com"},
				},
			})

			if test.wantErr != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.wantErr)
				assert.Nil(t, provider)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, provider)
		})
	}
}
