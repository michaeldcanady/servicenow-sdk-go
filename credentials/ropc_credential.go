// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"

	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type ropcClient interface {
	Initialize(baseURL string)
	acquireTokenByUsernamePassword(ctx context.Context, username, password string) (*AccessToken, error)
	acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// ROPCCredential implements the ROPC (Resource Owner Password Credentials) flow.
type ROPCCredential struct {
	*BaseAccessTokenProvider
	client   ropcClient
	username string
	password string
}

// NewROPCCredential creates a new ROPCCredential.
func NewROPCCredential(client ropcClient, username, password string) (*ROPCCredential, error) {
	c := &ROPCCredential{
		client:   client,
		username: username,
		password: password,
	}

	base := newBaseAccessTokenProvider()
	base.retrieveInitialToken = c.GetToken
	base.refreshToken = client.acquireTokenByRefreshToken
	base.revokeToken = client.revokeToken

	c.BaseAccessTokenProvider = base

	return c, nil
}

// Initialize initializes the provider and its internal client with the base URL.
func (c *ROPCCredential) Initialize(baseURL string) {
	c.BaseAccessTokenProvider.Initialize(baseURL)
	c.client.Initialize(baseURL)
}

// GetToken acquires a token using the username and password.
func (c *ROPCCredential) GetToken(ctx context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
	return c.client.acquireTokenByUsernamePassword(ctx, c.username, c.password)
}

// NewROPCProvider creates a new AuthenticationProvider for the ROPC flow.
// The ROPC flow is available but discouraged; see the package documentation
// for why you should prefer other flows (ADR-016).
func NewROPCProvider(clientID, clientSecret, username, password string, opts ROPCOptions) (authentication.AuthenticationProvider, error) {
	if clientID == "" {
		return nil, errors.ErrEmptyClientID
	}
	if clientSecret == "" {
		return nil, errors.ErrEmptyClientSecret
	}
	if username == "" {
		return nil, errors.ErrEmptyUsername
	}
	if password == "" {
		return nil, errors.ErrEmptyPassword
	}

	client, err := newConfidentialClient(clientID, clientSecret, "", func(co *clientOptions) {
		co.httpClient = opts.HTTPClient
	})
	if err != nil {
		return nil, err
	}

	tokenProvider, err := NewROPCCredential(client, username, password)
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
