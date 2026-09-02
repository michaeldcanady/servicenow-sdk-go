// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2/pkce"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPublicClient(t *testing.T) {
	tests := []struct {
		name     string
		clientID string
		wantErr  error
	}{
		{
			name:     "empty client id returns error",
			clientID: "",
			wantErr:  ErrEmptyClientID,
		},
		{
			name:     "valid client id succeeds",
			clientID: "client-id",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := newPublicClient(test.clientID, Authority("https://dev12345.service-now.com"))

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
		client, err := newPublicClient("client-id", "")
		require.NoError(t, err)
		require.NotNil(t, client)
		assert.Nil(t, client.oauthClient)
	})
}

func TestPublicClient_AcquireTokenByCode(t *testing.T) {
	tests := []struct {
		name       string
		initClient bool
		fake       *baseClientTestOAuthClient
		wantErr    string
	}{
		{
			name:       "uninitialized client returns error",
			initClient: false,
			wantErr:    "OAuth2 client not initialized",
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
			c := &publicClient{baseClient: &baseClient{clientID: "id"}}
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

func TestPublicClient_GenerateChallenge(t *testing.T) {
	tests := []struct {
		name      string
		method    pkce.Method
		wantEmpty bool
		wantErr   bool
	}{
		{
			name:      "unset method returns empty challenge",
			method:    pkce.MethodUnset,
			wantEmpty: true,
		},
		{
			name:   "S256 method generates a challenge and verifier",
			method: pkce.MethodS256,
		},
		{
			name:   "plain method generates a challenge and verifier",
			method: pkce.MethodPlain,
		},
		{
			name:    "unknown method returns error",
			method:  pkce.MethodUnknown,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &publicClient{baseClient: &baseClient{clientID: "id"}, method: test.method}

			challenge, err := c.generateChallenge()

			if test.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			if test.wantEmpty {
				assert.Empty(t, challenge)
				assert.Empty(t, c.verifier)
				return
			}
			assert.NotEmpty(t, challenge)
			assert.NotEmpty(t, c.verifier)
		})
	}
}

func TestPublicClient_GetAuthorizationURL(t *testing.T) {
	tests := []struct {
		name       string
		initClient bool
		fake       *baseClientTestOAuthClient
		method     pkce.Method
		wantErr    string
	}{
		{
			name:    "uninitialized client returns error",
			wantErr: "OAuth2 client not initialized",
		},
		{
			name:       "challenge generation error propagates",
			initClient: true,
			fake:       &baseClientTestOAuthClient{},
			method:     pkce.MethodUnknown,
			wantErr:    "unsupported PKCE method",
		},
		{
			name:       "auth code url error propagates",
			initClient: true,
			fake:       &baseClientTestOAuthClient{authCodeErr: errors.New("auth url failed")},
			method:     pkce.MethodS256,
			wantErr:    "auth url failed",
		},
		{
			name:       "successful url generation",
			initClient: true,
			fake:       &baseClientTestOAuthClient{authCodeURL: "https://auth.example.com"},
			method:     pkce.MethodS256,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &publicClient{baseClient: &baseClient{clientID: "id"}, method: test.method}
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

func TestPublicClient_Initialize(t *testing.T) {
	c := &publicClient{baseClient: &baseClient{clientID: "id"}}
	c.Initialize("https://dev12345.service-now.com")

	_, err := c.getOAuthClient()
	require.NoError(t, err)
}
