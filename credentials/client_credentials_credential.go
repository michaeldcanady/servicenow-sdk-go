package credentials

import (
	"context"
	"net/url"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type clientCredentialsFlow interface {
	acquireTokenByClientCredentials(ctx context.Context, scopes []string) (*AccessToken, error)
	acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// ClientCredentialsCredential implements the OAuth2 Client Credentials flow.
type ClientCredentialsCredential struct {
	*BaseAccessTokenProvider
	client clientCredentialsFlow
	scopes []string
}

// NewClientCredentialsCredential creates a new ClientCredentialsCredential.
func NewClientCredentialsCredential(client clientCredentialsFlow, scopes []string, allowedHosts []string) (*ClientCredentialsCredential, error) {
	c := &ClientCredentialsCredential{
		client: client,
		scopes: scopes,
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = c.GetToken
	base.refreshToken = client.acquireTokenByRefreshToken
	base.revokeToken = client.revokeToken

	c.BaseAccessTokenProvider = base

	return c, nil
}

// GetToken acquires a token using the client credentials.
func (c *ClientCredentialsCredential) GetToken(ctx context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
	return c.client.acquireTokenByClientCredentials(ctx, c.scopes)
}

// NewClientCredentialsProvider creates a new AuthenticationProvider for the Client Credentials flow using functional options.
func NewClientCredentialsProvider(clientID, clientSecret string, opts ...AuthOption) (authentication.AuthenticationProvider, error) {
	config := defaultAuthConfig()
	for _, opt := range opts {
		opt(config)
	}

	authority := Authority(config.baseURL)
	if authority == "" && config.instance != "" {
		authority = NewInstanceAuthority(config.instance)
	}

	client, err := newConfidentialClient(clientID, clientSecret, authority, func(co *clientOptions) {
		co.httpClient = config.httpClient
	})
	if err != nil {
		return nil, err
	}

	tokenProvider, err := NewClientCredentialsCredential(client, config.scopes, config.allowedHosts)
	if err != nil {
		return nil, err
	}

	if config.tokenStore != nil {
		tokenProvider.SetTokenStore(config.tokenStore)
	}

	return NewBearerTokenAuthenticationProvider(tokenProvider), nil
}
