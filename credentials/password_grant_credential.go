package credentials

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

// PasswordProvider is a user-supplied callback that returns username/password.
// This allows credentials to be injected from env vars, config, secret managers, etc.
type PasswordProvider func(ctx context.Context) (string, string, error)

// Oauth2PasswordClient defines the minimal client contract for password + refresh flows.
type Oauth2PasswordClient interface {
	ExchangePassword(ctx context.Context, user string, pass string) (*oauth2.Token, error)
	ExchangeRefreshToken(ctx context.Context, refresh string) (*oauth2.Token, error)
}

// ResourceOwnerPasswordCredential implements the Credential interface using the
// Resource Owner Password Credentials (ROPC) grant type.
type ResourceOwnerPasswordCredential struct {
	client           Oauth2PasswordClient
	passwordProvider PasswordProvider
	token            *oauth2.Token
}

// NewResourceOwnerPasswordCredential constructs a credential with a required password provider.
func NewResourceOwnerPasswordCredential(clientID, clientSecret, baseURL string, provider PasswordProvider) (*ResourceOwnerPasswordCredential, error) {
	if provider == nil {
		return nil, errors.New("password provider cannot be nil")
	}

	return &ResourceOwnerPasswordCredential{
		client: &oauth2.Client{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoints: &oauth2.Endpoints{
				AuthURL:  fmt.Sprintf("%s/oauth_auth.do", baseURL),
				TokenURL: fmt.Sprintf("%s/oauth_token.do", baseURL),
			},
			HTTPClient: http.DefaultClient,
		},
		passwordProvider: provider,
	}, nil
}

// GetAuthentication returns a valid access token string.
// It will reuse the current token if valid, refresh if expired, or fetch a new one.
func (c *ResourceOwnerPasswordCredential) GetAuthentication() (string, error) {
	ctx := context.Background()

	if c.token != nil && c.token.IsValid() {
		return fmt.Sprintf("Bearer %s", c.token.AccessToken), nil
	}

	token, err := c.resolveToken(ctx)
	if err != nil {
		return "", err
	}

	c.token = token
	return fmt.Sprintf("Bearer %s", c.token.AccessToken), nil
}

// resolveToken decides whether to refresh or fetch a new token.
func (c *ResourceOwnerPasswordCredential) resolveToken(ctx context.Context) (*oauth2.Token, error) {
	if c.token != nil && c.token.IsExpired() && strings.TrimSpace(c.token.RefreshToken) != "" {
		return c.refreshToken(ctx)
	}
	return c.fetchToken(ctx)
}

// fetchToken retrieves a new token using username/password.
func (c *ResourceOwnerPasswordCredential) fetchToken(ctx context.Context) (*oauth2.Token, error) {
	user, pass, err := c.passwordProvider(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve username/password: %w", err)
	}

	token, err := c.client.ExchangePassword(ctx, user, pass)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve token: %w", err)
	}
	return token, nil
}

// refreshToken retrieves a new token using the refresh token.
func (c *ResourceOwnerPasswordCredential) refreshToken(ctx context.Context) (*oauth2.Token, error) {
	token, err := c.client.ExchangeRefreshToken(ctx, c.token.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}
	return token, nil
}
