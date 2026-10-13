// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The tests in this file prove that the deprecated v2 surface and the new
// provider surface coexist: both call styles build the same constructors.

func TestCompat_LegacyFunctionalOptions(t *testing.T) {
	provider, err := NewClientCredentialsProvider("client-id", "client-secret",
		WithInstance("dev12345"),
		WithScopes("useraccount"),
		WithHTTPClient(http.DefaultClient),
	)
	require.NoError(t, err)
	require.NotNil(t, provider)
}

func TestCompat_NewStructOptions(t *testing.T) {
	provider, err := NewClientCredentialsProvider("client-id", "client-secret",
		ClientCredentialsOptions{
			CommonOptions: CommonOptions{Scopes: []string{"useraccount"}},
		},
	)
	require.NoError(t, err)
	require.NotNil(t, provider)
}

func TestCompat_LegacyROPCFunctionalOptions(t *testing.T) {
	provider, err := NewROPCProvider("client-id", "client-secret", "user", "pass",
		WithURL("https://dev12345.service-now.com"),
	)
	require.NoError(t, err)
	require.NotNil(t, provider)
}

func TestCompat_Authority(t *testing.T) {
	instance := NewInstanceAuthority("dev12345")
	assert.Equal(t, "https://dev12345.service-now.com/oauth_token.do", instance.TokenURL())
	assert.Equal(t, "https://dev12345.service-now.com/oauth_auth.do", instance.AuthURL())
	assert.Equal(t, "https://dev12345.service-now.com/oauth_revoke.do", instance.RevocationURL())

	custom := NewCustomAuthority("example.com")
	assert.Equal(t, "https://example.com/oauth_token.do", custom.TokenURL())

	// The alias still allows a direct conversion, as v2 callers did.
	assert.Equal(t, Authority("https://example.com"), custom)
}

func TestCompat_LegacyCredentialConstructors(t *testing.T) {
	client, err := newConfidentialClient("client-id", "client-secret", authority(""))
	require.NoError(t, err)

	cc, err := NewClientCredentialsCredential(client, []string{"useraccount"}, []string{"dev12345.service-now.com"})
	require.NoError(t, err)
	require.NotNil(t, cc)
	require.Equal(t, cc, any(cc).(*ClientCredentialsCredential))

	rc, err := NewROPCCredential(client, "user", "pass", nil)
	require.NoError(t, err)
	require.NotNil(t, rc)

	// GetAuthentication is retained on the deprecated credential.
	_ = rc.GetAuthentication

	ac, err := NewAuthorizationCodeCredential(client, nil, 5001, nil, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, ac)

	jw, err := NewJWTCredential(client, staticTokenProvider{}, nil)
	require.NoError(t, err)
	require.NotNil(t, jw)
}

func TestCompat_DeprecatedErrors(t *testing.T) {
	assert.Equal(t, "boom", NewOauth2Error("boom").Error())
	assert.Equal(t, "boom", NewCredentialError("boom").Error())

	var nilErr *CredentialError
	assert.Equal(t, "nil credential error", nilErr.Error())

	// The deprecated sentinels are the shared sentinels, so errors.Is matches.
	require.ErrorIs(t, ErrEmptyClientID, snerrors.ErrEmptyClientID)
	require.ErrorIs(t, ErrEmptyClientSecret, snerrors.ErrEmptyClientSecret)
}

// staticTokenProvider is a minimal authentication.AccessTokenProvider for the
// deprecated JWT constructor.
type staticTokenProvider struct{}

func (staticTokenProvider) GetAuthorizationToken(_ context.Context, _ *url.URL, _ map[string]interface{}) (string, error) {
	return "token", nil
}

func (staticTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return createAllowedHostsValidator(nil)
}
