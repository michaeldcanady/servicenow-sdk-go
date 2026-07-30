package credentials

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ropcTestClient struct {
	initializedWith string
	acquireToken    *AccessToken
	acquireErr      error
	lastUsername    string
	lastPassword    string
	refreshToken    *AccessToken
	refreshErr      error
	revokeErr       error
}

func (c *ropcTestClient) Initialize(baseURL string) {
	c.initializedWith = baseURL
}

func (c *ropcTestClient) acquireTokenByUsernamePassword(_ context.Context, username, password string) (*AccessToken, error) {
	c.lastUsername = username
	c.lastPassword = password
	return c.acquireToken, c.acquireErr
}

func (c *ropcTestClient) acquireTokenByRefreshToken(_ context.Context, _ string) (*AccessToken, error) {
	return c.refreshToken, c.refreshErr
}

func (c *ropcTestClient) revokeToken(_ context.Context, _, _ string) error {
	return c.revokeErr
}

func TestROPCCredential_GetToken(t *testing.T) {
	tests := []struct {
		name      string
		client    *ropcTestClient
		username  string
		password  string
		wantToken *AccessToken
		wantErr   string
	}{
		{
			name:     "acquisition error propagates",
			client:   &ropcTestClient{acquireErr: errors.New("acquire failed")},
			username: "user",
			password: "pass",
			wantErr:  "acquire failed",
		},
		{
			name:      "successful acquisition returns token and forwards credentials",
			client:    &ropcTestClient{acquireToken: &AccessToken{AccessToken: "ropc-token"}},
			username:  "user",
			password:  "pass",
			wantToken: &AccessToken{AccessToken: "ropc-token"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			credential, err := NewROPCCredential(test.client, test.username, test.password, nil)
			require.NoError(t, err)

			token, err := credential.GetToken(context.Background(), nil, nil)

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.wantToken, token)
			assert.Equal(t, test.username, test.client.lastUsername)
			assert.Equal(t, test.password, test.client.lastPassword)
		})
	}
}

func TestROPCCredential_Initialize(t *testing.T) {
	client := &ropcTestClient{}
	credential, err := NewROPCCredential(client, "user", "pass", nil)
	require.NoError(t, err)

	credential.Initialize("https://dev12345.service-now.com")

	assert.Equal(t, "https://dev12345.service-now.com", client.initializedWith)
}

func TestNewROPCProvider(t *testing.T) {
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
			name:         "empty client secret fails",
			clientID:     "id",
			clientSecret: "",
			wantErr:      ErrEmptyClientSecret,
		},
		{
			name:         "valid credentials succeed",
			clientID:     "id",
			clientSecret: "secret",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider, err := NewROPCProvider(test.clientID, test.clientSecret, "user", "pass", WithInstance("dev12345"))

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
