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
	// inflight is the single token request shared by concurrent callers.
	inflight *tokenFlight
	// mutex guards the token, inflight, and configuration fields.
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
	// baseURL is the base URL for the ServiceNow instance.
	baseURL string
	// clientID is the OAuth2 client ID the tokens are requested for.
	clientID string
}

// tokenFlight tracks the single in-flight token request so concurrent callers
// coalesce onto it instead of stampeding the token endpoint.
type tokenFlight struct {
	// done is closed once the request completes and the result is published.
	done chan struct{}
	// token is the acquired token, valid after done is closed.
	token *AccessToken
	// err is the acquisition error, valid after done is closed.
	err error
}

func newBaseAccessTokenProvider() *BaseAccessTokenProvider {
	return &BaseAccessTokenProvider{}
}

// SetAllowedHosts configures the allowed hosts validator after construction.
// This allows the instance URL to be set separately from the allowed hosts.
func (p *BaseAccessTokenProvider) SetAllowedHosts(hosts []string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if len(hosts) > 0 {
		validator, _ := authentication.NewAllowedHostsValidatorErrorCheck(hosts)
		p.allowedHostsValidator = validator
	}
}

// storeKey returns the key used to persist tokens for this provider.
func (p *BaseAccessTokenProvider) storeKey() TokenKey {
	return TokenKey{
		Authority: p.baseURL,
		ClientID:  p.clientID,
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

// Revoke revokes the current access token and refresh token and deletes the
// persisted token.
func (p *BaseAccessTokenProvider) Revoke(ctx context.Context) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// With no in-memory token there is nothing to revoke at the endpoint, but
	// a token may still be persisted from an earlier run: delete it.
	if p.token == nil {
		return p.deleteStoredLocked(ctx)
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
	return p.deleteStoredLocked(ctx)
}

// deleteStoredLocked removes the persisted token. It must be called with the
// mutex held. A provider without a store deletes nothing and succeeds.
func (p *BaseAccessTokenProvider) deleteStoredLocked(ctx context.Context) error {
	if p.tokenStore == nil {
		return nil
	}
	return p.tokenStore.Delete(ctx, p.storeKey())
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

	// Load a persisted token when we hold nothing in memory yet.
	if p.token == nil && p.tokenStore != nil {
		stored, err := p.tokenStore.Load(ctx, p.storeKey())
		if err != nil {
			p.mutex.Unlock()
			return "", fmt.Errorf("token store load failed: %w", err)
		}
		p.token = stored
	}

	if p.token != nil && !p.token.IsExpired() {
		token := p.token.AccessToken
		p.mutex.Unlock()
		return token, nil
	}

	// Another caller is already fetching a token: wait for its outcome
	// instead of starting a second token request.
	if p.inflight != nil {
		flight := p.inflight
		p.mutex.Unlock()
		return p.await(ctx, flight)
	}

	// This caller becomes the leader and performs the single token request
	// without holding the lock.
	flight := &tokenFlight{done: make(chan struct{})}
	p.inflight = flight

	refreshToken := ""
	if p.token != nil {
		refreshToken = p.token.RefreshToken
	}
	p.mutex.Unlock()

	newToken, err := p.fetch(ctx, uri, additionalAuthenticationContext, refreshToken)

	p.mutex.Lock()
	if err == nil && newToken != nil {
		p.token = newToken
		p.persist(ctx, newToken)
	}
	flight.token = newToken
	flight.err = err
	p.inflight = nil
	p.mutex.Unlock()
	close(flight.done)

	if err != nil {
		return "", err
	}
	return newToken.AccessToken, nil
}

// fetch performs the single token request. It prefers a refresh when the
// caller holds a refresh token, and falls back to a fresh acquisition.
func (p *BaseAccessTokenProvider) fetch(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}, refreshToken string) (*AccessToken, error) {
	if refreshToken != "" && p.refreshToken != nil {
		if newToken, err := p.refreshToken(ctx, refreshToken); err == nil {
			return newToken, nil
		}
	}

	if p.retrieveInitialToken == nil {
		return nil, fmt.Errorf("token acquisition failed: no initial retrieval function available")
	}

	newToken, err := p.retrieveInitialToken(ctx, uri, additionalAuthenticationContext)
	if err != nil {
		return nil, fmt.Errorf("initial token acquisition failed: %w", err)
	}
	return newToken, nil
}

// await blocks until the in-flight token request completes or ctx is done.
func (p *BaseAccessTokenProvider) await(ctx context.Context, flight *tokenFlight) (string, error) {
	select {
	case <-flight.done:
		if flight.err != nil {
			return "", flight.err
		}
		return flight.token.AccessToken, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// persist stores the token for restart recovery. It must be called with the
// mutex held. Persistence is best-effort: the token was acquired successfully,
// so a store failure must not fail the request that triggered the acquisition.
func (p *BaseAccessTokenProvider) persist(ctx context.Context, token *AccessToken) {
	if p.tokenStore == nil {
		return
	}
	_ = p.tokenStore.Save(ctx, p.storeKey(), token)
}
