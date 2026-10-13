// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"

	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type clientCredentialsClient interface {
	Initialize(baseURL string)
	acquireTokenByClientCredentials(ctx context.Context, scopes []string) (*AccessToken, error)
	acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// clientCredentialsCredential implements the OAuth2 Client Credentials flow.
type clientCredentialsCredential struct {
	*baseAccessTokenProvider
	client clientCredentialsClient
	scopes []string
}

// newClientCredentialsCredential creates a new clientCredentialsCredential.
func newClientCredentialsCredential(client clientCredentialsClient, scopes []string) *clientCredentialsCredential {
	c := &clientCredentialsCredential{
		client: client,
		scopes: scopes,
	}

	base := newBaseAccessTokenProvider()
	base.retrieveInitialToken = c.GetToken
	base.refreshToken = client.acquireTokenByRefreshToken
	base.revokeToken = client.revokeToken

	c.baseAccessTokenProvider = base

	return c
}

// Initialize initializes the provider and its internal client with the base URL.
func (c *clientCredentialsCredential) Initialize(baseURL string) {
	c.baseAccessTokenProvider.Initialize(baseURL)
	c.client.Initialize(baseURL)
}

// GetToken acquires a token using the client credentials.
func (c *clientCredentialsCredential) GetToken(ctx context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
	return c.client.acquireTokenByClientCredentials(ctx, c.scopes)
}

// NewClientCredentialsProvider creates a new AuthenticationProvider for the Client Credentials flow.
func NewClientCredentialsProvider(clientID, clientSecret string, opts ...AuthOption) (authentication.AuthenticationProvider, error) {
	config := buildAuthConfig(opts)

	if clientID == "" {
		return nil, errors.ErrEmptyClientID
	}
	if clientSecret == "" {
		return nil, errors.ErrEmptyClientSecret
	}

	client, err := newConfidentialClient(clientID, clientSecret, authority(""), func(co *clientOptions) {
		co.httpClient = config.httpClient
	})
	if err != nil {
		return nil, err
	}

	tokenProvider := newClientCredentialsCredential(client, config.scopes)

	tokenProvider.clientID = clientID
	tokenProvider.SetAllowedHosts(config.allowedHosts)

	if config.tokenStore != nil {
		tokenProvider.setKeyedTokenStore(config.tokenStore)
	}

	tokenProvider.Initialize(config.baseURL)

	return NewBearerTokenAuthenticationProvider(tokenProvider, BearerOptions{})
}
