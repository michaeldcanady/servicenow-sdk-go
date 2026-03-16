package credentials

import (
	"context"
	"crypto"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type jwtFlow interface {
	acquireTokenByJWT(ctx context.Context, assertion string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// JWTCredential implements the OAuth2 JWT Bearer Token flow.
type JWTCredential struct {
	*baseAccessTokenProvider
	client jwtFlow
}

// NewJWTCredential creates a new JWTCredential.
func NewJWTCredential(client jwtFlow, privateKey crypto.PrivateKey, issuer, subject, audience string, allowedHosts []string) (*JWTCredential, error) {
	initialFunc := func(ctx context.Context) (*AccessToken, error) {
		now := time.Now()
		claims := jwt.MapClaims{
			"iss": issuer,
			"sub": subject,
			"aud": audience,
			"jti": uuid.NewString(),
			"iat": now.Unix(),
			"exp": now.Add(time.Hour).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
		assertion, err := token.SignedString(privateKey)
		if err != nil {
			return nil, err
		}

		return client.acquireTokenByJWT(ctx, assertion)
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = initialFunc
	base.revokeToken = client.revokeToken

	return &JWTCredential{
		baseAccessTokenProvider: base,
		client:                  client,
	}, nil
}

// NewJWTAuthenticationProvider creates a new AuthenticationProvider for the JWT Bearer Token flow.
func NewJWTAuthenticationProvider(clientID, clientSecret string, privateKey crypto.PrivateKey, issuer, subject, audience string, authority Authority, allowedHosts []string) (authentication.AuthenticationProvider, error) {
	client, err := newConfidentialClient(clientID, clientSecret, authority)
	if err != nil {
		return nil, err
	}

	tokenProvider, err := NewJWTCredential(client, privateKey, issuer, subject, audience, allowedHosts)
	if err != nil {
		return nil, err
	}
	return authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider), nil
}
