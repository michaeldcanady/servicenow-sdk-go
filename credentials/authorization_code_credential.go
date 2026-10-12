// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2/pkce"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

var _ authentication.AccessTokenProvider = (*AuthorizationCodeCredential)(nil)

type authorizationCodeClient interface {
	authorizationCodeSource

	Initialize(baseURL string)
	acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// AuthorizationCodeCredential implements the OAuth2 Authorization Code flow.
//
// The interactive redirect machinery (redirect server, state, PKCE, code
// exchange) lives in the embedded authorizationCodeFlow, shared with the other
// browser-redirect credentials.
type AuthorizationCodeCredential struct {
	*BaseAccessTokenProvider
	*authorizationCodeFlow

	client authorizationCodeClient
}

// NewAuthorizationCodeCredential creates a new AuthorizationCodeCredential.
// This credential uses the provided authorization code to acquire an access token.
// The code is one-time use. Subsequent token acquisitions will use the refresh token.
func NewAuthorizationCodeCredential(client authorizationCodeClient, allowedHosts []string, port int, stateGenerator func() string, urlOpener func(string) error, serverFactory ServerFactory) (*AuthorizationCodeCredential, error) {
	c := &AuthorizationCodeCredential{
		client:                client,
		authorizationCodeFlow: newAuthorizationCodeFlow(port, stateGenerator, urlOpener, serverFactory),
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = c.GetToken
	base.refreshToken = client.acquireTokenByRefreshToken
	base.revokeToken = client.revokeToken

	c.BaseAccessTokenProvider = base

	return c, nil
}

// Initialize initializes the provider and its internal client with the base URL.
func (c *AuthorizationCodeCredential) Initialize(baseURL string) {
	c.BaseAccessTokenProvider.Initialize(baseURL)
	c.client.Initialize(baseURL)
}

// GetToken acquires a token using the authorization code flow.
func (c *AuthorizationCodeCredential) GetToken(ctx context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
	return c.acquire(ctx, c.client)
}

// NewAuthorizationCodeProvider creates a new AuthenticationProvider for the Authorization Code flow using functional options.
func NewAuthorizationCodeProvider(clientID, clientSecret string, opts ...AuthOption) (authentication.AuthenticationProvider, error) {
	config := &AuthConfig{
		httpClient: http.DefaultClient,
		port:       defaultAuthorizationCodePort,
	}
	for _, opt := range opts {
		opt(config)
	}

	authority := Authority(config.baseURL)

	var (
		client authorizationCodeClient
		err    error
	)

	if strings.TrimSpace(clientSecret) != "" {
		client, err = newConfidentialClient(clientID, clientSecret, authority, func(co *clientOptions) {
			co.httpClient = config.httpClient
		})
	} else {
		client, err = newPublicClient(clientID, authority, withPKCEChallenge(pkce.MethodS256), func(co *clientOptions) {
			co.httpClient = config.httpClient
		})
	}
	if err != nil {
		return nil, err
	}

	tokenProvider, err := NewAuthorizationCodeCredential(client, config.allowedHosts, config.port, config.stateGenerator, config.urlOpener, config.serverFactory)
	if err != nil {
		return nil, err
	}

	tokenProvider.Initialize(string(authority))

	tokenProvider.setClientID(clientID)
	if config.tokenStore != nil {
		tokenProvider.SetTokenStore(config.tokenStore)
	}

	return NewBearerTokenAuthenticationProvider(tokenProvider), nil
}

// NewPrivateAuthorizationCodeProvider creates a new AuthenticationProvider for the Authorization
// Code flow using a confidential (client-secret-bearing) client.
func NewPrivateAuthorizationCodeProvider(clientID, clientSecret string, opts ...AuthOption) (authentication.AuthenticationProvider, error) {
	return NewAuthorizationCodeProvider(clientID, clientSecret, opts...)
}

// NewPublicAuthorizationCodeProvider creates a new AuthenticationProvider for the Authorization
// Code flow using a public (no client-secret) client.
func NewPublicAuthorizationCodeProvider(clientID string, opts ...AuthOption) (authentication.AuthenticationProvider, error) {
	return NewAuthorizationCodeProvider(clientID, "", opts...)
}
