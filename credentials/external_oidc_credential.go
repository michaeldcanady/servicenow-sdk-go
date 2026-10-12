// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// defaultOIDCScope is the scope requested when the caller does not configure any
// scopes. It yields an ID token whose aud equals the client ID.
const defaultOIDCScope = "openid"

var _ authentication.AccessTokenProvider = (*ExternalOIDCCredential)(nil)

// ExternalOIDCCredential implements external OIDC authentication: it runs the
// authorization code + PKCE grant against an external identity provider and
// presents the resulting token verbatim to ServiceNow. ServiceNow's OIDP record
// is the validator (ADR-014).
//
// The interactive redirect machinery (redirect server, state, PKCE, code
// exchange) lives in the embedded authorizationCodeFlow, shared with the other
// browser-redirect credentials.
type ExternalOIDCCredential struct {
	*BaseAccessTokenProvider
	*authorizationCodeFlow

	client authorizationCodeClient
}

// newExternalOIDCCredential creates a new ExternalOIDCCredential.
func newExternalOIDCCredential(client authorizationCodeClient, allowedHosts []string, port int, stateGenerator func() string, urlOpener func(string) error, serverFactory ServerFactory) *ExternalOIDCCredential {
	c := &ExternalOIDCCredential{
		client:                client,
		authorizationCodeFlow: newAuthorizationCodeFlow(port, stateGenerator, urlOpener, serverFactory),
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = c.GetToken
	base.refreshToken = client.acquireTokenByRefreshToken
	base.revokeToken = client.revokeToken

	c.BaseAccessTokenProvider = base

	return c
}

// Initialize initializes the base token provider with the ServiceNow base URL.
// The external client ignores it: its endpoints come from OIDC discovery, not
// from the ServiceNow instance.
func (c *ExternalOIDCCredential) Initialize(baseURL string) {
	c.BaseAccessTokenProvider.Initialize(baseURL)
	c.client.Initialize(baseURL)
}

// GetToken acquires a token from the external identity provider using the
// authorization code + PKCE flow.
func (c *ExternalOIDCCredential) GetToken(ctx context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
	return c.acquire(ctx, c.client)
}

// NewExternalOIDCProvider creates an AuthenticationProvider that authenticates
// to ServiceNow with a token issued by an external OIDC identity provider.
//
// Configure the issuer with WithOIDCIssuer, as either a plain issuer URL or a
// full OIDC metadata URL. The provider is a public client (authorization code +
// PKCE, no secret) and presents the ID token by default; configuring scopes
// presents the audience-scoped access token instead. It performs no client-side
// signature or claims validation (ADR-014).
func NewExternalOIDCProvider(clientID string, opts ...AuthOption) (authentication.AuthenticationProvider, error) {
	config := &AuthConfig{
		httpClient: http.DefaultClient,
		port:       defaultAuthorizationCodePort,
	}
	for _, opt := range opts {
		opt(config)
	}

	if strings.TrimSpace(clientID) == "" {
		return nil, ErrEmptyClientID
	}

	if strings.TrimSpace(config.oidcIssuer) == "" {
		return nil, ErrEmptyOIDCIssuer
	}

	document, err := fetchOIDCDiscovery(context.Background(), config.oidcIssuer, config.httpClient)
	if err != nil {
		return nil, err
	}

	scopes := config.scopes
	useAccessToken := len(scopes) > 0
	if len(scopes) == 0 {
		scopes = []string{defaultOIDCScope}
	}

	client, err := newExternalOIDCClient(clientID, document.endpoints(), scopes, useAccessToken, config.httpClient)
	if err != nil {
		return nil, err
	}

	credential := newExternalOIDCCredential(client, config.allowedHosts, config.port, config.stateGenerator, config.urlOpener, config.serverFactory)
	credential.Initialize(config.baseURL)

	credential.setClientID(clientID)
	if config.tokenStore != nil {
		credential.SetTokenStore(config.tokenStore)
	}

	return NewBearerTokenAuthenticationProvider(credential), nil
}
