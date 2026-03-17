package credentials

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

// client is an interface for OAuth2 clients.
type client interface {
	acquireTokenByClientCredentials(ctx context.Context, scopes []string) (*AccessToken, error)
	acquireTokenByUsernamePassword(ctx context.Context, username, password string) (*AccessToken, error)
	acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)
	acquireTokenByCode(ctx context.Context, code, redirectURI, state string) (*AccessToken, error)
	acquireTokenByJWT(ctx context.Context, assertion string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
	getAuthorizationURL(redirectURI, state string, scopes []string) (string, error)
}

// baseClient provides a base implementation for an OAuth2 client.
type baseClient struct {
	oauthClient  *oauth2.Client
	clientID     string
	clientSecret string
	httpClient   *http.Client
	mutex        sync.RWMutex
}

func (c *baseClient) Initialize(instance, baseURL string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.oauthClient != nil {
		return
	}

	authority := Authority(baseURL)

	c.oauthClient = &oauth2.Client{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		Endpoints: &oauth2.Endpoints{
			TokenURL:         authority.TokenURL(),
			AuthURL:          authority.AuthURL(),
			DeviceURL:        "",
			RevocationURL:    authority.RevocationURL(),
			IntrospectionURL: "",
		},
		AuthMethod: oauth2.AuthMethodClientSecretPost, // Default
		HTTPClient: c.httpClient,
	}
}

func (c *baseClient) getOAuthClient() (*oauth2.Client, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if c.oauthClient == nil {
		return nil, fmt.Errorf("OAuth2 client not initialized. Ensure instance or URL is provided")
	}
	return c.oauthClient, nil
}

// acquireTokenByUsernamePassword acquires a token using the ROPC flow.
func (c *baseClient) acquireTokenByUsernamePassword(ctx context.Context, username, password string) (*AccessToken, error) {
	client, err := c.getOAuthClient()
	if err != nil {
		return nil, err
	}
	token, err := client.ExchangePassword(ctx, username, password, nil)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// acquireTokenByRefreshToken acquires a new token using a refresh token.
func (c *baseClient) acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error) {
	client, err := c.getOAuthClient()
	if err != nil {
		return nil, err
	}
	token, err := client.ExchangeRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// acquireTokenByJWT acquires a token using the JWT bearer grant.
func (c *baseClient) acquireTokenByJWT(ctx context.Context, assertion string) (*AccessToken, error) {
	client, err := c.getOAuthClient()
	if err != nil {
		return nil, err
	}
	token, err := client.ExchangeJWT(ctx, assertion)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// revokeToken revokes a token.
func (c *baseClient) revokeToken(ctx context.Context, token, tokenTypeHint string) error {
	client, err := c.getOAuthClient()
	if err != nil {
		return err
	}
	return client.Revoke(ctx, token, tokenTypeHint)
}
