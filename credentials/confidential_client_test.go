package credentials

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfidentialClient(t *testing.T) {
	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		wantErr      error
	}{
		{
			name:         "empty client id returns error",
			clientID:     "",
			clientSecret: "secret",
			wantErr:      ErrEmptyClientID,
		},
		{
			name:         "empty client secret returns error",
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
			client, err := newConfidentialClient(test.clientID, test.clientSecret, Authority("https://dev12345.service-now.com"))

			if test.wantErr != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.wantErr)
				assert.Nil(t, client)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, client)
			assert.NotNil(t, client.oauthClient)
		})
	}

	t.Run("empty authority does not initialize the client", func(t *testing.T) {
		client, err := newConfidentialClient("id", "secret", "")
		require.NoError(t, err)
		require.NotNil(t, client)
		assert.Nil(t, client.oauthClient)
	})
}

func TestConfidentialClient_AcquireTokenByCode(t *testing.T) {
	tests := []struct {
		name       string
		initClient bool
		fake       *baseClientTestOAuthClient
		wantErr    string
	}{
		{
			name:    "uninitialized client returns error",
			wantErr: "OAuth2 client not initialized",
		},
		{
			name:       "exchange error propagates",
			initClient: true,
			fake:       &baseClientTestOAuthClient{err: errors.New("code exchange failed")},
			wantErr:    "code exchange failed",
		},
		{
			name:       "successful exchange returns converted token",
			initClient: true,
			fake:       &baseClientTestOAuthClient{token: &oauth2.Token{AccessToken: "code-access"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &confidentialClient{baseClient: &baseClient{clientID: "id", clientSecret: "secret"}}
			if test.initClient {
				c.oauthClient = test.fake
			}

			token, err := c.acquireTokenByCode(context.Background(), "code", "https://redirect", "state")

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "code-access", token.AccessToken)
			assert.Equal(t, "code", test.fake.lastCode)
		})
	}
}

func TestConfidentialClient_AcquireTokenByClientCredentials(t *testing.T) {
	tests := []struct {
		name       string
		initClient bool
		fake       *baseClientTestOAuthClient
		wantErr    string
	}{
		{
			name:    "uninitialized client returns error",
			wantErr: "OAuth2 client not initialized",
		},
		{
			name:       "exchange error propagates",
			initClient: true,
			fake:       &baseClientTestOAuthClient{err: errors.New("client credentials failed")},
			wantErr:    "client credentials failed",
		},
		{
			name:       "successful exchange returns converted token",
			initClient: true,
			fake:       &baseClientTestOAuthClient{token: &oauth2.Token{AccessToken: "cc-access"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &confidentialClient{baseClient: &baseClient{clientID: "id", clientSecret: "secret"}}
			if test.initClient {
				c.oauthClient = test.fake
			}

			token, err := c.acquireTokenByClientCredentials(context.Background(), []string{"scope1"})

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "cc-access", token.AccessToken)
			assert.Equal(t, []string{"scope1"}, test.fake.lastScopes)
		})
	}
}

func TestConfidentialClient_GetAuthorizationURL(t *testing.T) {
	tests := []struct {
		name       string
		initClient bool
		fake       *baseClientTestOAuthClient
		wantErr    string
	}{
		{
			name:    "uninitialized client returns error",
			wantErr: "OAuth2 client not initialized",
		},
		{
			name:       "auth code url error propagates",
			initClient: true,
			fake:       &baseClientTestOAuthClient{authCodeErr: errors.New("auth url failed")},
			wantErr:    "auth url failed",
		},
		{
			name:       "successful url generation",
			initClient: true,
			fake:       &baseClientTestOAuthClient{authCodeURL: "https://auth.example.com"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &confidentialClient{baseClient: &baseClient{clientID: "id", clientSecret: "secret"}}
			if test.initClient {
				c.oauthClient = test.fake
			}

			url, err := c.getAuthorizationURL("https://redirect", "state", nil)

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "https://auth.example.com", url)
		})
	}
}

func TestConfidentialClient_Initialize(t *testing.T) {
	c := &confidentialClient{baseClient: &baseClient{clientID: "id", clientSecret: "secret"}}
	c.Initialize("https://dev12345.service-now.com")

	_, err := c.getOAuthClient()
	require.NoError(t, err)
}
