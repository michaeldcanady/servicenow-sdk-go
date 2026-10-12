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

// ClientCredentialsCredential implements the OAuth2 Client Credentials flow.
type ClientCredentialsCredential struct {
	*BaseAccessTokenProvider
	client clientCredentialsClient
	scopes []string
}

// NewClientCredentialsCredential creates a new ClientCredentialsCredential.
func NewClientCredentialsCredential(client clientCredentialsClient, scopes []string) (*ClientCredentialsCredential, error) {
	c := &ClientCredentialsCredential{
		client: client,
		scopes: scopes,
	}

	base := newBaseAccessTokenProvider()
	base.retrieveInitialToken = c.GetToken
	base.refreshToken = client.acquireTokenByRefreshToken
	base.revokeToken = client.revokeToken

	c.BaseAccessTokenProvider = base

	return c, nil
}

// Initialize initializes the provider and its internal client with the base URL.
func (c *ClientCredentialsCredential) Initialize(baseURL string) {
	c.BaseAccessTokenProvider.Initialize(baseURL)
	c.client.Initialize(baseURL)
}

// GetToken acquires a token using the client credentials.
func (c *ClientCredentialsCredential) GetToken(ctx context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
	return c.client.acquireTokenByClientCredentials(ctx, c.scopes)
}

// NewClientCredentialsProvider creates a new AuthenticationProvider for the Client Credentials flow.
func NewClientCredentialsProvider(clientID, clientSecret string, opts ClientCredentialsOptions) (authentication.AuthenticationProvider, error) {
	if clientID == "" {
		return nil, errors.ErrEmptyClientID
	}
	if clientSecret == "" {
		return nil, errors.ErrEmptyClientSecret
	}

	client, err := newConfidentialClient(clientID, clientSecret, "", func(co *clientOptions) {
		co.httpClient = opts.HTTPClient
	})
	if err != nil {
		return nil, err
	}

	tokenProvider, err := NewClientCredentialsCredential(client, opts.Scopes)
	if err != nil {
		return nil, err
	}

	tokenProvider.clientID = clientID
	tokenProvider.SetAllowedHosts(opts.AllowedHosts)

	if opts.TokenStore != nil {
		tokenProvider.SetTokenStore(opts.TokenStore)
	}

	return NewBearerTokenAuthenticationProvider(tokenProvider), nil
}
