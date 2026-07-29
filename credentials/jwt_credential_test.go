package credentials

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"net/url"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func jwtTestSigningKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	return key
}

func jwtTestSignToken(t *testing.T, method jwt.SigningMethod, claims jwt.MapClaims, key any) string {
	t.Helper()
	token := jwt.NewWithClaims(method, claims)
	signed, err := token.SignedString(key)
	require.NoError(t, err)
	return signed
}

func jwtTestValidClaims() jwt.MapClaims {
	now := time.Now()
	return jwt.MapClaims{
		"iss": "issuer",
		"sub": "subject",
		"aud": "audience",
		"exp": now.Add(time.Hour).Unix(),
		"iat": float64(now.Unix()),
		"jti": "unique-id",
	}
}

func TestValidateJWT(t *testing.T) {
	key := jwtTestSigningKey(t)

	tests := []struct {
		name    string
		token   func() string
		wantErr string
	}{
		{
			name: "valid RS256 token with all required claims",
			token: func() string {
				return jwtTestSignToken(t, jwt.SigningMethodRS256, jwtTestValidClaims(), key)
			},
		},
		{
			name: "malformed token fails to parse",
			token: func() string {
				return "not-a-jwt"
			},
			wantErr: "token is malformed",
		},
		{
			name: "missing required claim",
			token: func() string {
				claims := jwtTestValidClaims()
				delete(claims, "sub")
				return jwtTestSignToken(t, jwt.SigningMethodRS256, claims, key)
			},
			wantErr: "missing required claim: sub",
		},
		{
			name: "iat in the future is rejected",
			token: func() string {
				claims := jwtTestValidClaims()
				claims["iat"] = float64(time.Now().Add(time.Hour).Unix())
				return jwtTestSignToken(t, jwt.SigningMethodRS256, claims, key)
			},
			wantErr: "JWT issued in the future",
		},
		{
			name: "iat claim is not numeric",
			token: func() string {
				claims := jwtTestValidClaims()
				claims["iat"] = "not-a-number"
				return jwtTestSignToken(t, jwt.SigningMethodRS256, claims, key)
			},
			wantErr: "iat claim is not numeric",
		},
		{
			name: "unexpected signing algorithm",
			token: func() string {
				return jwtTestSignToken(t, jwt.SigningMethodHS256, jwtTestValidClaims(), []byte("secret"))
			},
			wantErr: "unexpected signing algorithm: HS256",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateJWT(test.token())

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

type jwtTestAssertionProvider struct {
	assertion string
	err       error
}

func (p *jwtTestAssertionProvider) GetAuthorizationToken(_ context.Context, _ *url.URL, _ map[string]interface{}) (string, error) {
	return p.assertion, p.err
}

func (p *jwtTestAssertionProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}

type jwtTestClient struct {
	initializedWith string
	acquireToken    *AccessToken
	acquireErr      error
	revokeErr       error
}

func (c *jwtTestClient) Initialize(baseURL string) {
	c.initializedWith = baseURL
}

func (c *jwtTestClient) acquireTokenByJWT(_ context.Context, _ string) (*AccessToken, error) {
	return c.acquireToken, c.acquireErr
}

func (c *jwtTestClient) revokeToken(_ context.Context, _, _ string) error {
	return c.revokeErr
}

func TestJWTCredential_GetToken(t *testing.T) {
	key := jwtTestSigningKey(t)
	validAssertion := jwtTestSignToken(t, jwt.SigningMethodRS256, jwtTestValidClaims(), key)

	tests := []struct {
		name      string
		provider  *jwtTestAssertionProvider
		client    *jwtTestClient
		wantErr   string
		wantToken *AccessToken
	}{
		{
			name:     "token provider error propagates",
			provider: &jwtTestAssertionProvider{err: errors.New("assertion failed")},
			client:   &jwtTestClient{},
			wantErr:  "assertion failed",
		},
		{
			name:     "invalid assertion fails validation",
			provider: &jwtTestAssertionProvider{assertion: "not-a-jwt"},
			client:   &jwtTestClient{},
			wantErr:  "token is malformed",
		},
		{
			name:      "valid assertion acquires token via client",
			provider:  &jwtTestAssertionProvider{assertion: validAssertion},
			client:    &jwtTestClient{acquireToken: &AccessToken{AccessToken: "issued"}},
			wantToken: &AccessToken{AccessToken: "issued"},
		},
		{
			name:     "client acquisition error propagates",
			provider: &jwtTestAssertionProvider{assertion: validAssertion},
			client:   &jwtTestClient{acquireErr: errors.New("exchange failed")},
			wantErr:  "exchange failed",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			credential, err := NewJWTCredential(test.client, test.provider, nil)
			require.NoError(t, err)

			token, err := credential.GetToken(context.Background(), nil, nil)

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.wantToken, token)
		})
	}
}
