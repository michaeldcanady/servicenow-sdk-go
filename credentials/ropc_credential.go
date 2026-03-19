package credentials

import (
	"context"
	"net/http"
	"net/url"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type ropcClient interface {
	acquireTokenByUsernamePassword(ctx context.Context, username, password string) (*AccessToken, error)
	acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// ROPCCredential implements the ROPC (Resource Owner Password Credentials) flow.
type ROPCCredential struct {
	*BaseAccessTokenProvider
	client   ropcClient
	username string
	password string
}

// NewROPCCredential creates a new ROPCCredential.
func NewROPCCredential(client ropcClient, username, password string, allowedHosts []string) (*ROPCCredential, error) {
	c := &ROPCCredential{
		client:   client,
		username: username,
		password: password,
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = c.GetToken
	base.refreshToken = client.acquireTokenByRefreshToken
	base.revokeToken = client.revokeToken

	c.BaseAccessTokenProvider = base

	return c, nil
}

// GetToken acquires a token using the username and password.
func (c *ROPCCredential) GetToken(ctx context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
	return c.client.acquireTokenByUsernamePassword(ctx, c.username, c.password)
}

// NewROPCProvider creates a new AuthenticationProvider for the ROPC flow using functional options.
func NewROPCProvider(clientID, clientSecret, username, password string, opts ...AuthOption) (authentication.AuthenticationProvider, error) {
	config := &AuthConfig{
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(config)
	}

	authority := Authority(config.baseURL)

	client, err := newConfidentialClient(clientID, clientSecret, authority, func(co *clientOptions) {
		co.httpClient = config.httpClient
	})
	if err != nil {
		return nil, err
	}

	tokenProvider, err := NewROPCCredential(client, username, password, config.allowedHosts)
	if err != nil {
		return nil, err
	}

	tokenProvider.Initialize(string(authority))

	if config.tokenStore != nil {
		tokenProvider.SetTokenStore(config.tokenStore)
	}

	return NewBearerTokenAuthenticationProvider(tokenProvider), nil
}
