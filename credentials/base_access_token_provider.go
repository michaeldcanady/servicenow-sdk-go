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

// Revoke revokes the current access token and refresh token.
func (p *BaseAccessTokenProvider) Revoke(ctx context.Context) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.token == nil {
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

	if p.token != nil && !p.token.IsExpired() {
		return p.token.AccessToken, nil
	}

	if p.token != nil && p.token.RefreshToken != "" && p.refreshToken != nil {
		if newToken, err := p.refreshToken(ctx, p.token.RefreshToken); err == nil {
			p.token = newToken
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
	return p.token.AccessToken, nil
}
