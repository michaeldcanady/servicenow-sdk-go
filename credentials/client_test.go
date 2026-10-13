// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// baseClientTestOAuthClient is a fake satisfying the package-local oauth2Client interface.
type baseClientTestOAuthClient struct {
	token         *oauth2.Token
	err           error
	authCodeURL   string
	authCodeErr   error
	revokeErr     error
	lastUsername  string
	lastPassword  string
	lastRefresh   string
	lastAssertion string
	lastScopes    []string
	lastCode      string
	lastRevoked   string
	lastTypeHint  string
}

func (c *baseClientTestOAuthClient) ExchangePassword(_ context.Context, user, pass string, _ []string) (*oauth2.Token, error) {
	c.lastUsername = user
	c.lastPassword = pass
	return c.token, c.err
}

func (c *baseClientTestOAuthClient) ExchangeRefreshToken(_ context.Context, refresh string) (*oauth2.Token, error) {
	c.lastRefresh = refresh
	return c.token, c.err
}

func (c *baseClientTestOAuthClient) ExchangeJWT(_ context.Context, assertion string) (*oauth2.Token, error) {
	c.lastAssertion = assertion
	return c.token, c.err
}

func (c *baseClientTestOAuthClient) Revoke(_ context.Context, token, tokenTypeHint string) error {
	c.lastRevoked = token
	c.lastTypeHint = tokenTypeHint
	return c.revokeErr
}

func (c *baseClientTestOAuthClient) ExchangeCode(_ context.Context, code, _, _, _ string) (*oauth2.Token, error) {
	c.lastCode = code
	return c.token, c.err
}

func (c *baseClientTestOAuthClient) ExchangeClientCredentials(_ context.Context, scopes []string) (*oauth2.Token, error) {
	c.lastScopes = scopes
	return c.token, c.err
}

func (c *baseClientTestOAuthClient) AuthCodeURL(_, _, _, _ string, _ []string) (string, error) {
	return c.authCodeURL, c.authCodeErr
}

func TestBaseClient_GetOAuthClient(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(c *baseClient)
		wantErr bool
	}{
		{
			name:    "uninitialized client returns error",
			setup:   func(_ *baseClient) {},
			wantErr: true,
		},
		{
			name: "initialized client is returned",
			setup: func(c *baseClient) {
				c.Initialize("https://dev12345.service-now.com")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &baseClient{clientID: "id", clientSecret: "secret"}
			test.setup(c)

			client, err := c.getOAuthClient()

			if test.wantErr {
				require.Error(t, err)
				assert.Nil(t, client)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, client)
		})
	}
}

func TestBaseClient_Initialize(t *testing.T) {
	t.Run("empty base url is a no-op", func(t *testing.T) {
		c := &baseClient{clientID: "id"}
		c.Initialize("")

		_, err := c.getOAuthClient()
		require.Error(t, err)
	})

	t.Run("already-initialized client is not overwritten", func(t *testing.T) {
		fake := &baseClientTestOAuthClient{}
		c := &baseClient{clientID: "id", oauthClient: fake}
		c.Initialize("https://dev12345.service-now.com")

		got, err := c.getOAuthClient()
		require.NoError(t, err)
		assert.Same(t, fake, got)
	})

	t.Run("first initialization builds a client", func(t *testing.T) {
		c := &baseClient{clientID: "id"}
		c.Initialize("https://dev12345.service-now.com")

		got, err := c.getOAuthClient()
		require.NoError(t, err)
		assert.NotNil(t, got)
	})
}

func TestBaseClient_AcquireTokenByUsernamePassword(t *testing.T) {
	exchangeErr := errors.New("exchange failed")

	tests := []struct {
		name    string
		client  *baseClientTestOAuthClient
		initErr bool
		wantErr string
	}{
		{
			name:    "uninitialized client returns error",
			initErr: true,
			wantErr: "OAuth2 client not initialized",
		},
		{
			name:    "exchange error propagates",
			client:  &baseClientTestOAuthClient{err: exchangeErr},
			wantErr: "exchange failed",
		},
		{
			name:   "successful exchange returns converted token",
			client: &baseClientTestOAuthClient{token: &oauth2.Token{AccessToken: "access"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &baseClient{clientID: "id"}
			if !test.initErr {
				c.oauthClient = test.client
			}

			token, err := c.acquireTokenByUsernamePassword(context.Background(), "user", "pass")

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "access", token.AccessToken)
			assert.Equal(t, "user", test.client.lastUsername)
			assert.Equal(t, "pass", test.client.lastPassword)
		})
	}
}

func TestBaseClient_AcquireTokenByRefreshToken(t *testing.T) {
	tests := []struct {
		name    string
		client  *baseClientTestOAuthClient
		initErr bool
		wantErr string
	}{
		{
			name:    "uninitialized client returns error",
			initErr: true,
			wantErr: "OAuth2 client not initialized",
		},
		{
			name:    "exchange error propagates",
			client:  &baseClientTestOAuthClient{err: errors.New("refresh failed")},
			wantErr: "refresh failed",
		},
		{
			name:   "successful exchange returns converted token",
			client: &baseClientTestOAuthClient{token: &oauth2.Token{AccessToken: "new-access"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &baseClient{clientID: "id"}
			if !test.initErr {
				c.oauthClient = test.client
			}

			token, err := c.acquireTokenByRefreshToken(context.Background(), "refresh")

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "new-access", token.AccessToken)
			assert.Equal(t, "refresh", test.client.lastRefresh)
		})
	}
}

func TestBaseClient_AcquireTokenByJWT(t *testing.T) {
	tests := []struct {
		name    string
		client  *baseClientTestOAuthClient
		initErr bool
		wantErr string
	}{
		{
			name:    "uninitialized client returns error",
			initErr: true,
			wantErr: "OAuth2 client not initialized",
		},
		{
			name:    "exchange error propagates",
			client:  &baseClientTestOAuthClient{err: errors.New("jwt exchange failed")},
			wantErr: "jwt exchange failed",
		},
		{
			name:   "successful exchange returns converted token",
			client: &baseClientTestOAuthClient{token: &oauth2.Token{AccessToken: "jwt-access"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &baseClient{clientID: "id"}
			if !test.initErr {
				c.oauthClient = test.client
			}

			token, err := c.acquireTokenByJWT(context.Background(), "assertion")

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "jwt-access", token.AccessToken)
			assert.Equal(t, "assertion", test.client.lastAssertion)
		})
	}
}

func TestBaseClient_RevokeToken(t *testing.T) {
	tests := []struct {
		name    string
		client  *baseClientTestOAuthClient
		initErr bool
		wantErr string
	}{
		{
			name:    "uninitialized client returns error",
			initErr: true,
			wantErr: "OAuth2 client not initialized",
		},
		{
			name:    "revoke error propagates",
			client:  &baseClientTestOAuthClient{revokeErr: errors.New("revoke failed")},
			wantErr: "revoke failed",
		},
		{
			name:   "successful revoke",
			client: &baseClientTestOAuthClient{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &baseClient{clientID: "id"}
			if !test.initErr {
				c.oauthClient = test.client
			}

			err := c.revokeToken(context.Background(), "token", "access_token")

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "token", test.client.lastRevoked)
			assert.Equal(t, "access_token", test.client.lastTypeHint)
		})
	}
}
