// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"strings"
	"time"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// ClientCredentialsConfig configures the OAuth 2.0 client credentials flow.
type ClientCredentialsConfig struct {
	ProviderConfig
	// ClientID is the OAuth client identifier.
	ClientID string
	// ClientSecret is the OAuth client secret. Client credentials require a
	// confidential client.
	ClientSecret string
	// Scopes requested from the token endpoint.
	Scopes []string
	// ClientAuthMethod controls how the client authenticates at the token
	// endpoint. The default is client_secret_post.
	ClientAuthMethod ClientAuthMethod
	// TokenStore persists tokens. A nil store uses an in-memory store.
	TokenStore TokenStore
}

// Validate checks the client credentials settings.
func (c ClientCredentialsConfig) Validate() error {
	if err := c.ProviderConfig.Validate(); err != nil {
		return err
	}
	if strings.TrimSpace(c.ClientID) == "" {
		return snerrors.ErrEmptyClientID
	}
	if strings.TrimSpace(c.ClientSecret) == "" {
		return snerrors.ErrEmptyClientSecret
	}
	if c.ClientAuthMethod == ClientAuthMethodNone {
		return snerrors.ErrInvalidAuthMethod
	}
	if _, err := c.ClientAuthMethod.oauth2Method(); err != nil {
		return err
	}
	return nil
}

// ClientCredentialsProvider authenticates with the OAuth 2.0 client
// credentials flow.
type ClientCredentialsProvider struct {
	config    ClientCredentialsConfig
	lifecycle *tokenLifecycle
	client    *oauth2.Client
}

// NewClientCredentialsProvider creates a client credentials provider. It does
// not contact ServiceNow.
func NewClientCredentialsProvider(config ClientCredentialsConfig) (*ClientCredentialsProvider, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	validator, err := newAllowedHostsValidator(config.ProviderConfig)
	if err != nil {
		return nil, err
	}

	provider := &ClientCredentialsProvider{
		config: config,
	}
	provider.lifecycle = newTokenLifecycle(
		FlowClientCredentials,
		config.ClientID,
		config.Scopes,
		"",
		config.TokenStore,
		validator,
	)
	provider.lifecycle.reacquireOnInvalidGrant = true
	provider.lifecycle.acquire = provider.acquireToken
	provider.lifecycle.refresh = provider.refreshToken
	provider.lifecycle.revokeHook = provider.revokeToken
	return provider, nil
}

// Initialize stores the ServiceNow base URL and prepares the token endpoint.
func (p *ClientCredentialsProvider) Initialize(baseURL string) error {
	if p == nil {
		return snerrors.ErrNilProvider
	}
	if err := p.lifecycle.initialize(baseURL); err != nil {
		return err
	}

	normalized, err := normalizeBaseURL(baseURL)
	if err != nil {
		return err
	}

	method, err := p.config.ClientAuthMethod.oauth2Method()
	if err != nil {
		return err
	}

	options := []oauth2.Option{
		oauth2.WithClientSecret(p.config.ClientSecret),
		oauth2.WithAuthMethod(method),
	}
	if p.config.HTTPClient != nil {
		options = append(options, oauth2.WithHTTPClient(p.config.HTTPClient))
	}

	p.client = oauth2.NewClient(p.config.ClientID, &oauth2.Endpoints{
		AuthURL:       normalized + "/oauth_auth.do",
		TokenURL:      normalized + "/oauth_token.do",
		RevocationURL: normalized + "/oauth_revoke.do",
	}, options...)
	return nil
}

// AuthenticateRequest adds a Bearer token to the request.
func (p *ClientCredentialsProvider) AuthenticateRequest(ctx context.Context, request *abstractions.RequestInformation, _ map[string]interface{}) error {
	if p == nil {
		return snerrors.ErrNilProvider
	}
	if request == nil {
		return snerrors.ErrNilRequestInfo
	}
	if request.Headers != nil && request.Headers.ContainsKey("Authorization") {
		return nil
	}

	uri, err := request.GetUri()
	if err != nil {
		return err
	}

	token, err := p.lifecycle.authorizationToken(ctx, uri)
	if err != nil {
		return err
	}
	if token == "" {
		return nil
	}

	if request.Headers == nil {
		request.Headers = abstractions.NewRequestHeaders()
	}
	setHeaderIfAbsent(request, "Authorization", "Bearer "+token)
	return nil
}

// Revoke revokes the current token and clears the local cache.
func (p *ClientCredentialsProvider) Revoke(ctx context.Context) error {
	if p == nil {
		return snerrors.ErrNilProvider
	}
	return p.lifecycle.revoke(ctx)
}

func (p *ClientCredentialsProvider) acquireToken(ctx context.Context, _ string) (*Token, error) {
	if p.client == nil {
		return nil, snerrors.ErrNotInitialized
	}

	token, err := p.client.ExchangeClientCredentials(ctx, p.lifecycle.scopes)
	if err != nil {
		return nil, err
	}
	return tokenFromOAuth2(token, p.lifecycle.now())
}

func (p *ClientCredentialsProvider) refreshToken(ctx context.Context, _ string, refreshToken string) (*Token, error) {
	if p.client == nil {
		return nil, snerrors.ErrNotInitialized
	}

	token, err := p.client.ExchangeRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return tokenFromOAuth2(token, p.lifecycle.now())
}

func (p *ClientCredentialsProvider) revokeToken(ctx context.Context, _ string, token *Token) error {
	if p.client == nil {
		return snerrors.ErrNotInitialized
	}

	if token.RefreshToken != "" {
		if err := p.client.Revoke(ctx, token.RefreshToken, oauth2.RefreshTokenKey); err != nil {
			return err
		}
	}
	if token.AccessToken != "" {
		if err := p.client.Revoke(ctx, token.AccessToken, oauth2.AccessTokenKey); err != nil {
			return err
		}
	}
	return nil
}

func tokenFromOAuth2(token *oauth2.Token, now time.Time) (*Token, error) {
	if token == nil {
		return nil, snerrors.ErrMissingAccessToken
	}

	result := &Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		Scope:        token.Scope,
	}
	if token.ExpiresIn > 0 {
		result.ExpiresAt = now.Add(time.Duration(token.ExpiresIn) * time.Second)
	}
	if idToken, ok := token.Raw["id_token"].(string); ok {
		result.IDToken = idToken
	}
	return result, nil
}
