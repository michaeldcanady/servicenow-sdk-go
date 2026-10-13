// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	oidcTestIssuer = "https://idp.example.com"
	oidcTestDoc    = `{
		"issuer": "https://idp.example.com",
		"authorization_endpoint": "https://idp.example.com/authorize",
		"token_endpoint": "https://idp.example.com/token",
		"revocation_endpoint": "https://idp.example.com/revoke"
	}`
)

// statusRoundTripper returns a fixed status code and body for every request.
type statusRoundTripper struct {
	status int
	body   string
}

func (rt statusRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	header := make(http.Header)
	header.Set("Content-Type", "application/json")
	return &http.Response{
		StatusCode: rt.status,
		Status:     http.StatusText(rt.status),
		Body:       io.NopCloser(strings.NewReader(rt.body)),
		Header:     header,
		Request:    req,
	}, nil
}

func TestExtractIDToken(t *testing.T) {
	tests := []struct {
		name  string
		token *AccessToken
		want  string
	}{
		{name: "nil token", token: nil, want: ""},
		{name: "no ID token", token: &AccessToken{AccessToken: "access"}, want: ""},
		{name: "ID token present", token: &AccessToken{AccessToken: "access", IDToken: "id-token"}, want: "id-token"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, extractIDToken(tt.token))
		})
	}
}

func TestNewExternalOIDCProvider_Validation(t *testing.T) {
	tests := []struct {
		name     string
		clientID string
		issuer   string
		wantErr  error
	}{
		{
			name:     "empty client ID",
			clientID: "",
			issuer:   oidcTestIssuer,
			wantErr:  snerrors.ErrEmptyClientID,
		},
		{
			name:     "empty issuer",
			clientID: "client-id",
			issuer:   "   ",
			wantErr:  snerrors.ErrEmptyIssuer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider, err := NewExternalOIDCProvider(tt.clientID, ExternalOIDCOptions{Issuer: tt.issuer})
			require.ErrorIs(t, err, tt.wantErr)
			assert.Nil(t, provider)
		})
	}
}

func TestNewExternalOIDCProvider_DiscoveryFailure(t *testing.T) {
	provider, err := NewExternalOIDCProvider("client-id", ExternalOIDCOptions{
		CommonOptions: CommonOptions{
			HTTPClient: &http.Client{Transport: statusRoundTripper{status: http.StatusInternalServerError}},
		},
		Issuer: oidcTestIssuer,
	})
	require.ErrorIs(t, err, snerrors.ErrOIDCDiscoveryFailed)
	assert.Nil(t, provider)
}

func TestNewExternalOIDCProvider_DefaultsToOpenIDScope(t *testing.T) {
	provider, err := NewExternalOIDCProvider("client-id", ExternalOIDCOptions{
		CommonOptions: CommonOptions{
			HTTPClient: &http.Client{Transport: okJSONRoundTripper{body: []byte(oidcTestDoc)}},
		},
		Issuer: oidcTestIssuer,
	})
	require.NoError(t, err)
	require.NotNil(t, provider)

	cred := externalOIDCCredentialFrom(t, provider)
	assert.Equal(t, []string{"openid"}, cred.scopes)
	assert.Equal(t, ExternalOIDCTokenSourceIDToken, cred.tokenSource)
}

func TestNewExternalOIDCProvider_CustomScopes(t *testing.T) {
	provider, err := NewExternalOIDCProvider("client-id", ExternalOIDCOptions{
		CommonOptions: CommonOptions{
			HTTPClient: &http.Client{Transport: okJSONRoundTripper{body: []byte(oidcTestDoc)}},
			Scopes:     []string{"openid", "profile"},
		},
		Issuer: oidcTestIssuer,
	})
	require.NoError(t, err)

	cred := externalOIDCCredentialFrom(t, provider)
	assert.Equal(t, []string{"openid", "profile"}, cred.scopes)
}

func TestExternalOIDCCredential_Refresh_UsesIDToken(t *testing.T) {
	client := &fakeAuthCodeClient{
		token: &AccessToken{AccessToken: "access", IDToken: "id-token", RefreshToken: "refresh"},
	}
	cred := &externalOIDCCredential{
		client:      client,
		tokenSource: ExternalOIDCTokenSourceIDToken,
	}

	token, err := cred.refreshTokenFunc(context.Background(), "refresh")
	require.NoError(t, err)
	require.NotNil(t, token)
	assert.Equal(t, "id-token", token.AccessToken)
	assert.Equal(t, "id-token", cred.idToken)
}

func TestExternalOIDCCredential_Refresh_UsesAccessToken(t *testing.T) {
	client := &fakeAuthCodeClient{
		token: &AccessToken{AccessToken: "access", IDToken: "id-token", RefreshToken: "refresh"},
	}
	cred := &externalOIDCCredential{
		client:      client,
		tokenSource: ExternalOIDCTokenSourceAccessToken,
	}

	token, err := cred.refreshTokenFunc(context.Background(), "refresh")
	require.NoError(t, err)
	require.NotNil(t, token)
	assert.Equal(t, "access", token.AccessToken)
}

func TestExternalOIDCCredential_Revoke_MissingEndpointSucceeds(t *testing.T) {
	client := &fakeAuthCodeClient{revokeErr: snerrors.ErrRevocationEndpointNotSet}
	cred := &externalOIDCCredential{
		client:             client,
		idToken:            "id-token",
		storedRefreshToken: "refresh",
	}

	err := cred.revokeTokenFunc(context.Background(), "id-token", "access_token")
	require.NoError(t, err)
	assert.Empty(t, cred.idToken)
	assert.Empty(t, cred.storedRefreshToken)
}

func TestExternalOIDCCredential_Revoke_PropagatesOtherErrors(t *testing.T) {
	wantErr := errors.New("revocation unavailable")
	client := &fakeAuthCodeClient{revokeErr: wantErr}
	cred := &externalOIDCCredential{
		client:             client,
		idToken:            "id-token",
		storedRefreshToken: "refresh",
	}

	err := cred.revokeTokenFunc(context.Background(), "id-token", "access_token")
	require.ErrorIs(t, err, wantErr)
	// Local state stays put when revocation fails for a reason other than a
	// missing endpoint, so a retry can still use the token.
	assert.Equal(t, "id-token", cred.idToken)
	assert.Equal(t, "refresh", cred.storedRefreshToken)
}

// externalOIDCCredentialFrom digs the underlying credential out of the provider
// so tests can assert on its configuration.
func externalOIDCCredentialFrom(t *testing.T, provider authentication.AuthenticationProvider) *externalOIDCCredential {
	t.Helper()
	bearer, ok := provider.(*BearerTokenAuthenticationProvider)
	require.True(t, ok)
	cred, ok := bearer.GetAuthorizationTokenProvider().(*externalOIDCCredential)
	require.True(t, ok)
	return cred
}

// fakeAuthCodeClient is an authorizationCodeClient test double.
type fakeAuthCodeClient struct {
	token      *AccessToken
	refreshErr error
	revokeErr  error
}

func (*fakeAuthCodeClient) Initialize(string) {}

func (*fakeAuthCodeClient) getAuthorizationURL(string, string, []string) (string, string, error) {
	return "", "", nil
}

func (*fakeAuthCodeClient) acquireTokenByCode(context.Context, string, string, string, string) (*AccessToken, error) {
	return nil, nil
}

func (f *fakeAuthCodeClient) acquireTokenByRefreshToken(context.Context, string) (*AccessToken, error) {
	return f.token, f.refreshErr
}

func (f *fakeAuthCodeClient) revokeToken(context.Context, string, string) error {
	return f.revokeErr
}
