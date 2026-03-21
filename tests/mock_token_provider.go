package tests

import (
	"context"
	"net/url"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type staticTokenProvider struct {
	token string
}

func (p *staticTokenProvider) GetAuthorizationToken(ctx context.Context, url *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	return p.token, nil
}

func (p *staticTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}

func NewStaticTokenProvider(token string) authentication.AccessTokenProvider {
	return &staticTokenProvider{token: token}
}
