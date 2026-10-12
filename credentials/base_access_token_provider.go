// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// BaseAccessTokenProvider provides common functionality for OAuth2 token providers.
type BaseAccessTokenProvider struct {
	// token the current valid access token
	token *AccessToken
	// mutex ensures only one token operation is happening at a time
	mutex                 sync.RWMutex
	allowedHostsValidator *authentication.AllowedHostsValidator
	// retrieveInitialToken is a function to get the ServiceNow token.
	retrieveInitialToken func(ctx context.Context, url *url.URL, additionalAuthenticationContext map[string]interface{}) (*AccessToken, error)
	// refreshToken is a function to refresh the token.
	refreshToken func(ctx context.Context, refreshToken string) (*AccessToken, error)
	// revokeToken is a function to revoke the token.
	revokeToken func(ctx context.Context, token, tokenTypeHint string) error
	// tokenStore is an optional store for persisting and retrieving access tokens.
	tokenStore TokenStore
	// clientID is the key under which tokens are persisted in the token store.
	clientID string
	// baseURL is the base URL for the ServiceNow instance.
	baseURL string
}

func newBaseAccessTokenProvider(allowedHosts []string) *BaseAccessTokenProvider {
	var validator *authentication.AllowedHostsValidator
	if len(allowedHosts) > 0 {
		validator, _ = authentication.NewAllowedHostsValidatorErrorCheck(allowedHosts)
	}

	return &BaseAccessTokenProvider{
		allowedHostsValidator: validator,
	}
}

// Initialize initializes the provider with the base URL.
func (p *BaseAccessTokenProvider) Initialize(baseURL string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.baseURL = baseURL

	if p.allowedHostsValidator == nil && baseURL != "" {
		u, err := url.Parse(baseURL)
		if err == nil && u.Host != "" {
			validator, _ := authentication.NewAllowedHostsValidatorErrorCheck([]string{u.Host})
			p.allowedHostsValidator = validator
		}
	}
}

// SetTokenStore sets the token store for the provider.
func (p *BaseAccessTokenProvider) SetTokenStore(store TokenStore) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.tokenStore = store
}

// setClientID sets the identifier used as the token store key. Tokens are only
// persisted when both a store and a client ID are set.
func (p *BaseAccessTokenProvider) setClientID(clientID string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.clientID = clientID
}

// Revoker is implemented by token providers that can revoke their current token.
type Revoker interface {
	// Revoke revokes the current access token and refresh token.
	Revoke(ctx context.Context) error
}

// Revoke revokes the current access token and refresh token and then clears the
// cached and persisted token. It is best-effort: when the identity provider
// cannot be asked to revoke (for example it advertises no revocation endpoint),
// the token is still cleared locally and revocation succeeds.
func (p *BaseAccessTokenProvider) Revoke(ctx context.Context) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.token == nil {
		p.deleteToken(ctx)
		return nil
	}

	if p.revokeToken == nil {
		return fmt.Errorf("token revocation failed: no revocation function available")
	}

	// Revoke refresh token if present
	if p.token.RefreshToken != "" {
		if err := p.revokeToken(ctx, p.token.RefreshToken, "refresh_token"); err != nil {
			return err
		}
	}

	// Revoke access token
	if err := p.revokeToken(ctx, p.token.AccessToken, "access_token"); err != nil {
		return err
	}

	p.token = nil
	p.deleteToken(ctx)
	return nil
}

// GetAllowedHostsValidator returns the allowed hosts validator.
func (p *BaseAccessTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return p.allowedHostsValidator
}

// GetAuthorizationToken gets the authorization token.
func (p *BaseAccessTokenProvider) GetAuthorizationToken(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if uri != nil && p.allowedHostsValidator != nil && !p.allowedHostsValidator.IsUrlHostValid(uri) {
		return "", nil
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.token == nil {
		p.token = p.loadToken(ctx)
	}

	if p.token != nil && !p.token.IsExpired() {
		return p.token.AccessToken, nil
	}

	if p.token != nil && p.token.RefreshToken != "" && p.refreshToken != nil {
		if newToken, err := p.refreshToken(ctx, p.token.RefreshToken); err == nil {
			p.token = newToken
			p.saveToken(ctx)
			return p.token.AccessToken, nil
		}
	}

	if p.retrieveInitialToken == nil {
		return "", fmt.Errorf("token acquisition failed: no initial retrieval function available")
	}

	newToken, err := p.retrieveInitialToken(ctx, uri, additionalAuthenticationContext)
	if err != nil {
		return "", fmt.Errorf("initial token acquisition failed: %w", err)
	}

	p.token = newToken
	p.saveToken(ctx)
	return p.token.AccessToken, nil
}

// loadToken retrieves a persisted token. Persistence is best-effort, so a
// missing store or client ID, or a store error, yields no token.
func (p *BaseAccessTokenProvider) loadToken(ctx context.Context) *AccessToken {
	if p.tokenStore == nil || p.clientID == "" {
		return nil
	}

	token, err := p.tokenStore.Load(ctx, p.clientID)
	if err != nil {
		return nil
	}

	return token
}

// saveToken persists the current token. Persistence is best-effort: store
// failures are ignored so they never break authentication.
func (p *BaseAccessTokenProvider) saveToken(ctx context.Context) {
	if p.tokenStore == nil || p.clientID == "" || p.token == nil {
		return
	}

	_ = p.tokenStore.Save(ctx, p.clientID, p.token)
}

// deleteToken removes the persisted token. Persistence is best-effort: store
// failures are ignored.
func (p *BaseAccessTokenProvider) deleteToken(ctx context.Context) {
	if p.tokenStore == nil || p.clientID == "" {
		return
	}

	_ = p.tokenStore.Delete(ctx, p.clientID)
}
