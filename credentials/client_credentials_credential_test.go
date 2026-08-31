// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type clientCredsTestClient struct {
	initializedWith string
	acquireToken    *AccessToken
	acquireErr      error
	lastScopes      []string
	refreshToken    *AccessToken
	refreshErr      error
	revokeErr       error
}

func (c *clientCredsTestClient) Initialize(baseURL string) {
	c.initializedWith = baseURL
}

func (c *clientCredsTestClient) acquireTokenByClientCredentials(_ context.Context, scopes []string) (*AccessToken, error) {
	c.lastScopes = scopes
	return c.acquireToken, c.acquireErr
}

func (c *clientCredsTestClient) acquireTokenByRefreshToken(_ context.Context, _ string) (*AccessToken, error) {
	return c.refreshToken, c.refreshErr
}

func (c *clientCredsTestClient) revokeToken(_ context.Context, _, _ string) error {
	return c.revokeErr
}

func TestClientCredentialsCredential_GetToken(t *testing.T) {
	tests := []struct {
		name      string
		client    *clientCredsTestClient
		scopes    []string
		wantToken *AccessToken
		wantErr   string
	}{
		{
			name:    "acquisition error propagates",
			client:  &clientCredsTestClient{acquireErr: errors.New("acquire failed")},
			wantErr: "acquire failed",
		},
		{
			name:      "successful acquisition returns token and forwards scopes",
			client:    &clientCredsTestClient{acquireToken: &AccessToken{AccessToken: "cc-token"}},
			scopes:    []string{"read", "write"},
			wantToken: &AccessToken{AccessToken: "cc-token"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			credential, err := NewClientCredentialsCredential(test.client, test.scopes, nil)
			require.NoError(t, err)

			token, err := credential.GetToken(context.Background(), nil, nil)

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.wantToken, token)
			assert.Equal(t, test.scopes, test.client.lastScopes)
		})
	}
}

func TestClientCredentialsCredential_Initialize(t *testing.T) {
	client := &clientCredsTestClient{}
	credential, err := NewClientCredentialsCredential(client, nil, nil)
	require.NoError(t, err)

	credential.Initialize("https://dev12345.service-now.com")

	assert.Equal(t, "https://dev12345.service-now.com", client.initializedWith)
}

func TestNewClientCredentialsProvider(t *testing.T) {
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
			provider, err := NewClientCredentialsProvider(test.clientID, test.clientSecret, WithInstance("dev12345"))

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
