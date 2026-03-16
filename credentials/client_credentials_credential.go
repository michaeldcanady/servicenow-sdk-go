package credentials

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// ClientCredentialsCredential implements the OAuth2 Client Credentials flow.
type ClientCredentialsCredential struct {
	*baseAccessTokenProvider
}

// NewClientCredentialsCredential creates a new ClientCredentialsCredential.
func NewClientCredentialsCredential(clientID, clientSecret string, authority Authority, allowedHosts []string, scopes []string) (*ClientCredentialsCredential, error) {
	client, err := newConfidentialClient(clientID, clientSecret, authority)
	if err != nil {
		return nil, err
	}

	initialFunc := func(ctx context.Context) (*AccessToken, error) {
		token, err := client.oauthClient.ExchangeClientCredentials(ctx, scopes)
		if err != nil {
			return nil, err
		}
		return convertToken(token), nil
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
	}, nil
}

// NewClientCredentialsAuthenticationProvider creates a new AuthenticationProvider for the Client Credentials flow.
func NewClientCredentialsAuthenticationProvider(clientID, clientSecret string, authority Authority, allowedHosts []string, scopes []string) (authentication.AuthenticationProvider, error) {
	tokenProvider, err := NewClientCredentialsCredential(clientID, clientSecret, authority, allowedHosts, scopes)
	if err != nil {
		return nil, err
	}
	return authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider), nil
}
