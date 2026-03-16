package credentials

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2/pkce"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/pkg/browser"
)

type authorizationCodeStrategy interface {
	getAuthorizationURL(redirectURI, state string, scopes []string) (string, error)
	acquireTokenByCode(ctx context.Context, code, redirectURI, state string) (*AccessToken, error)
	acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// AuthorizationCodeCredential implements the OAuth2 Authorization Code flow.
type AuthorizationCodeCredential struct {
	*baseAccessTokenProvider
	client authorizationCodeStrategy
}

// NewAuthorizationCodeCredential creates a new AuthorizationCodeCredential.
// This credential uses the provided authorization code to acquire an access token.
// The code is one-time use. Subsequent token acquisitions will use the refresh token.
func NewAuthorizationCodeCredential(client authorizationCodeStrategy, allowedHosts []string) (*AuthorizationCodeCredential, error) {
	port := 5001

	initialFunc := func(ctx context.Context) (token *AccessToken, err error) {
		state := uuid.NewString()

		server, err := oauth2.NewServer(state, port)
		if err != nil {
			return nil, err
		}

		defer func() {
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if shutdownErr := server.Shutdown(shutdownCtx); shutdownErr != nil && err == nil {
				err = shutdownErr
			}
		}()

		authURL, err := client.getAuthorizationURL(server.Addr, state, nil)
		if err != nil {
			return nil, err
		}

		if err := browser.OpenURL(authURL); err != nil {
			return nil, err
		}

		result := server.Result(ctx)

		if err := result.Err; err != nil {
			return nil, err
		}

		token, err = client.acquireTokenByCode(ctx, result.Code, server.Addr, state)
		if err != nil {
			return nil, err
		}
		return token, nil
	}

	refreshFunc := func(ctx context.Context, refreshToken string) (*AccessToken, error) {
		return client.acquireTokenByRefreshToken(ctx, refreshToken)
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = initialFunc
	base.refreshToken = refreshFunc
	base.revokeToken = client.revokeToken

	return &AuthorizationCodeCredential{
		baseAccessTokenProvider: base,
		client:                  client,
	}, nil
}

// NewAuthorizationCodeAuthenticationProvider creates a new AuthenticationProvider for the Authorization Code flow.
func NewAuthorizationCodeAuthenticationProvider(clientID, clientSecret string, authority Authority, allowedHosts []string) (authentication.AuthenticationProvider, error) {
	var (
		client authorizationCodeStrategy
		err    error
	)

	if strings.TrimSpace(clientSecret) != "" {
		client, err = newConfidentialClient(clientID, clientSecret, authority)
	} else {
		client, err = newPublicClient(clientID, authority, withPKCEChallenge(pkce.MethodS256))
	}
	if err != nil {
		return nil, err
	}

	tokenProvider, err := NewAuthorizationCodeCredential(client, allowedHosts)
	if err != nil {
		return nil, err
	}
	return authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider), nil
}
