package credentials

import (
	"context"
	"crypto"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// JWTCredential implements the OAuth2 JWT Bearer Token flow.
type JWTCredential struct {
	*baseAccessTokenProvider
}

// NewJWTCredential creates a new JWTCredential.
func NewJWTCredential(clientID, clientSecret string, privateKey crypto.PrivateKey, issuer, subject, audience string, authority Authority, allowedHosts []string) (*JWTCredential, error) {
	client, err := newConfidentialClient(clientID, clientSecret, authority) // Secret is dummy here
	if err != nil {
		return nil, err
	}

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

		tok, err := client.oauthClient.ExchangeJWT(ctx, assertion)
		if err != nil {
			return nil, err
		}
		return convertToken(tok), nil
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = initialFunc
	base.revokeToken = client.revokeToken

	return &JWTCredential{
		baseAccessTokenProvider: base,
	}, nil
}

// NewJWTAuthenticationProvider creates a new AuthenticationProvider for the JWT Bearer Token flow.
func NewJWTAuthenticationProvider(clientID, clientSecret string, privateKey crypto.PrivateKey, issuer, subject, audience string, authority Authority, allowedHosts []string) (authentication.AuthenticationProvider, error) {
	tokenProvider, err := NewJWTCredential(clientID, clientSecret, privateKey, issuer, subject, audience, authority, allowedHosts)
	if err != nil {
		return nil, err
	}
	return authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider), nil
}
