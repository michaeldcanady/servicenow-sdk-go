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

// AuthorizationCodeCredential implements the OAuth2 Authorization Code flow.
type AuthorizationCodeCredential struct {
	*baseAccessTokenProvider
	client client
	public bool
}

// nolint: unused // needed for later
type revokeTokenClient interface {
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// NewAuthorizationCodeCredential creates a new AuthorizationCodeCredential.
// This credential uses the provided authorization code to acquire an access token.
// The code is one-time use. Subsequent token acquisitions will use the refresh token.
func NewAuthorizationCodeCredential(clientID, clientSecret, redirectURI string, authority Authority, allowedHosts []string) (*AuthorizationCodeCredential, error) {
	port := 5001

	var (
		client client
		err    error
		public bool
	)

	if strings.TrimSpace(clientSecret) != "" {
		client, err = newConfidentialClient(clientID, clientSecret, authority)
		public = false
	} else {
		client, err = newPublicClient(clientID, authority, withPKCEChallenge(pkce.MethodS256))
		public = true
	}
	if err != nil {
		return nil, err
	}

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

	return &AuthorizationCodeCredential{
		baseAccessTokenProvider: base,
		client:                  client,
		public:                  public,
	}, nil
}

// NewAuthorizationCodeAuthenticationProvider creates a new AuthenticationProvider for the Authorization Code flow.
func NewAuthorizationCodeAuthenticationProvider(clientID, clientSecret, redirectURI string, authority Authority, allowedHosts []string) (authentication.AuthenticationProvider, error) {
	tokenProvider, err := NewAuthorizationCodeCredential(clientID, clientSecret, redirectURI, authority, allowedHosts)
	if err != nil {
		return nil, err
	}
	return authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider), nil
}
