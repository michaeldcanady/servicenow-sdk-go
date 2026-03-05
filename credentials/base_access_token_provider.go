package credentials

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// baseAccessTokenProvider provides common functionality for OAuth2 token providers.
type baseAccessTokenProvider struct {
	token                 *AccessToken
	mutex                 sync.RWMutex
	allowedHostsValidator authentication.AllowedHostsValidator
	// retrieveInitialToken is a function to get the first token.
	retrieveInitialToken func(ctx context.Context) (*AccessToken, error)
	// refreshToken is a function to refresh the token.
	refreshToken func(ctx context.Context, refreshToken string) (*AccessToken, error)
}

func newBaseAccessTokenProvider(allowedHosts []string) *baseAccessTokenProvider {
	return &baseAccessTokenProvider{
		//nolint: staticcheck // while being deprecated there doesn't seem to be a replacement
		allowedHostsValidator: authentication.NewAllowedHostsValidator(allowedHosts),
	}
}

// GetAllowedHostsValidator returns the allowed hosts validator.
func (p *baseAccessTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return &p.allowedHostsValidator
}

// GetAuthorizationToken gets the authorization token.
func (p *baseAccessTokenProvider) GetAuthorizationToken(ctx context.Context, uri *url.URL, _ map[string]interface{}) (string, error) {
	if uri != nil && !p.allowedHostsValidator.IsUrlHostValid(uri) {
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

	newToken, err := p.retrieveInitialToken(ctx)
	if err != nil {
		return "", fmt.Errorf("initial token acquisition failed: %w", err)
	}

	p.token = newToken
	return p.token.AccessToken, nil
}
