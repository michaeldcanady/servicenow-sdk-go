package credentials

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var mockAnyContext = mock.Anything

// TODO: should be a Mock.mock
type authCodeTestClient struct {
	initializedWith string
	authURL         string
	authURLErr      error
	acquireToken    *AccessToken
	acquireErr      error
	refreshToken    *AccessToken
	refreshErr      error
	revokeErr       error
}

func (c *authCodeTestClient) Initialize(baseURL string) {
	c.initializedWith = baseURL
}

func (c *authCodeTestClient) getAuthorizationURL(_, _ string, _ []string) (string, error) {
	return c.authURL, c.authURLErr
}

func (c *authCodeTestClient) acquireTokenByCode(_ context.Context, _, _, _ string) (*AccessToken, error) {
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
			client: &authCodeTestClient{authURL: "https://auth.example.com", acquireToken: &AccessToken{AccessToken: "final"}},
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
			client: &authCodeTestClient{authURL: "https://auth.example.com", acquireToken: &AccessToken{AccessToken: "final"}},
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

			credential, err := NewAuthorizationCodeCredential(test.client, nil, 0, func() string { return "state" }, urlOpener, test.serverFactory)
			require.NoError(t, err)

			token, err := credential.GetToken(context.Background(), nil, nil)

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.wantToken, token)
		})
	}
}

func TestAuthorizationCodeCredential_Initialize(t *testing.T) {
	client := &authCodeTestClient{}
	credential, err := NewAuthorizationCodeCredential(client, nil, 0, nil, nil, nil)
	require.NoError(t, err)

	credential.Initialize("https://dev12345.service-now.com")

	assert.Equal(t, "https://dev12345.service-now.com", client.initializedWith)
}

func TestNewAuthorizationCodeCredential_Defaults(t *testing.T) {
	client := &authCodeTestClient{}
	credential, err := NewAuthorizationCodeCredential(client, nil, 0, nil, nil, nil)

	require.NoError(t, err)
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
			wantErr:      ErrEmptyClientID,
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
			provider, err := NewPrivateAuthorizationCodeProvider(test.clientID, test.clientSecret, WithInstance("dev12345"))

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

func TestNewPublicAuthorizationCodeProvider(t *testing.T) {
	tests := []struct {
		name     string
		clientID string
		wantErr  error
	}{
		{
			name:     "empty client id fails",
			clientID: "",
			wantErr:  ErrEmptyClientID,
		},
		{
			name:     "valid client id succeeds",
			clientID: "id",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider, err := NewPublicAuthorizationCodeProvider(test.clientID, WithInstance("dev12345"))

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
