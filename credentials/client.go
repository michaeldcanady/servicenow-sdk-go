package credentials

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

// client is an interface for OAuth2 clients.
type client interface {
	acquireTokenByUsernamePassword(ctx context.Context, username, password string) (*AccessToken, error)
	acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)
	acquireTokenByCode(ctx context.Context, code, redirectURI, state string) (*AccessToken, error)
	acquireTokenByJWT(ctx context.Context, assertion string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
	getAuthorizationURL(redirectURI, state string, scopes []string) (string, error)
}

// baseClient provides a base implementation for an OAuth2 client.
type baseClient struct {
	oauthClient *oauth2.Client
}

// acquireTokenByUsernamePassword acquires a token using the ROPC flow.
func (c *baseClient) acquireTokenByUsernamePassword(ctx context.Context, username, password string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangePassword(ctx, username, password, nil)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// acquireTokenByRefreshToken acquires a new token using a refresh token.
func (c *baseClient) acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangeRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// acquireTokenByJWT acquires a token using the JWT bearer grant.
func (c *baseClient) acquireTokenByJWT(ctx context.Context, assertion string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangeJWT(ctx, assertion)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// revokeToken revokes a token.
func (c *baseClient) revokeToken(ctx context.Context, token, tokenTypeHint string) error {
	return c.oauthClient.Revoke(ctx, token, tokenTypeHint)
}
