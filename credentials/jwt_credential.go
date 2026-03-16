package credentials

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type jwtFlow interface {
	acquireTokenByJWT(ctx context.Context, assertion string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// JWTCredential implements the OAuth2 JWT Bearer Token flow.
type JWTCredential struct {
	*baseAccessTokenProvider
	tokenProvider authentication.AccessTokenProvider
	client        jwtFlow
}

func validateJWT(rawToken string) error {
	// Parse without verifying the signature
	token, _, err := jwt.NewParser(
		jwt.WithoutClaimsValidation(),
	).ParseUnverified(rawToken, jwt.MapClaims{})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid JWT claims format")
	}

	// Validate required claims exist
	required := []string{"iss", "sub", "aud", "exp", "iat", "jti"}
	for _, key := range required {
		if _, exists := claims[key]; !exists {
			return errors.New("missing required claim: " + key)
		}
	}

	// Validate issued-at
	if iat, ok := claims["iat"].(float64); ok {
		if int64(iat) > time.Now().Unix() {
			return errors.New("JWT issued in the future")
		}
	} else {
		return errors.New("iat claim is not numeric")
	}

	// Validate algorithm
	if token.Method.Alg() != jwt.SigningMethodRS256.Alg() {
		return errors.New("unexpected signing algorithm: " + token.Method.Alg())
	}

	return nil
}

// NewJWTCredential creates a new JWTCredential.
func NewJWTCredential(client jwtFlow, tokenProvider authentication.AccessTokenProvider, allowedHosts []string) (*JWTCredential, error) {
	initialFunc := func(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (*AccessToken, error) {

		assertion, err := tokenProvider.GetAuthorizationToken(ctx, uri, additionalAuthenticationContext)
		if err != nil {
			return nil, err
		}

		if err := validateJWT(assertion); err != nil {
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
func NewJWTAuthenticationProvider(clientID, clientSecret string, tokenProvider authentication.AccessTokenProvider, authority Authority, allowedHosts []string) (authentication.AuthenticationProvider, error) {
	client, err := newConfidentialClient(clientID, clientSecret, authority)
	if err != nil {
		return nil, err
	}

	snTokenProvider, err := NewJWTCredential(client, tokenProvider, allowedHosts)
	if err != nil {
		return nil, err
	}
	return authentication.NewBaseBearerTokenAuthenticationProvider(snTokenProvider), nil
}
