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
	*baseAccessTokenProvider
	client clientCredentialsFlow
}

// NewClientCredentialsCredential creates a new ClientCredentialsCredential.
func NewClientCredentialsCredential(client clientCredentialsFlow, scopes []string, allowedHosts []string) (*ClientCredentialsCredential, error) {
	initialFunc := func(ctx context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
		return client.acquireTokenByClientCredentials(ctx, scopes)
	}

	refreshFunc := func(ctx context.Context, refreshToken string) (*AccessToken, error) {
		return client.acquireTokenByRefreshToken(ctx, refreshToken)
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = initialFunc
	base.refreshToken = refreshFunc
	base.revokeToken = client.revokeToken

	return &ClientCredentialsCredential{
		baseAccessTokenProvider: base,
		client:                  client,
	}, nil
}

// NewClientCredentialsAuthenticationProvider creates a new AuthenticationProvider for the Client Credentials flow.
func NewClientCredentialsAuthenticationProvider(clientID, clientSecret string, authority Authority, allowedHosts []string, scopes []string) (authentication.AuthenticationProvider, error) {
	client, err := newConfidentialClient(clientID, clientSecret, authority)
	if err != nil {
		return nil, err
	}

	tokenProvider, err := NewClientCredentialsCredential(client, scopes, allowedHosts)
	if err != nil {
		return nil, err
	}
	return authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider), nil
}
