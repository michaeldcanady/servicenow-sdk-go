// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithOIDCIssuer(t *testing.T) {
	config := &AuthConfig{}
	WithOIDCIssuer("https://idp.example.com")(config)
	assert.Equal(t, "https://idp.example.com", config.oidcIssuer)
}

func TestExternalOIDCCredential_GetToken(t *testing.T) {
	client := &authCodeTestClient{authURL: "https://auth.example.com", acquireToken: &AccessToken{AccessToken: "final"}}
	server := &mockAuthorizationCodeServer{}
	server.On("GetAddr").Return("http://localhost:5001")
	server.On("Result", mockAnyContext).Return("code", "state", nil)
	server.On("Shutdown", mockAnyContext).Return(nil)

	credential := newExternalOIDCCredential(client, nil, 0, func() string { return "state" }, func(string) error { return nil }, authCodeTestServerFactory(server))

	token, err := credential.GetToken(context.Background(), nil, nil)
	require.NoError(t, err)
	assert.Equal(t, &AccessToken{AccessToken: "final"}, token)
}

func TestExternalOIDCCredential_Initialize(t *testing.T) {
	client := &authCodeTestClient{}
	credential := newExternalOIDCCredential(client, nil, 0, nil, nil, nil)

	credential.Initialize("https://dev12345.service-now.com")

	assert.Equal(t, "https://dev12345.service-now.com", client.initializedWith)
	assert.Equal(t, "https://dev12345.service-now.com", credential.baseURL)
}

func TestNewExternalOIDCCredential_Defaults(t *testing.T) {
	credential := newExternalOIDCCredential(&authCodeTestClient{}, nil, 0, nil, nil, nil)

	assert.Equal(t, defaultAuthorizationCodePort, credential.port)
	assert.NotNil(t, credential.stateGenerator)
	assert.NotNil(t, credential.urlOpener)
	assert.NotNil(t, credential.serverFactory)
}

func TestNewExternalOIDCProvider_Validation(t *testing.T) {
	t.Run("empty client id fails", func(t *testing.T) {
		provider, err := NewExternalOIDCProvider("")
		require.ErrorIs(t, err, ErrEmptyClientID)
		assert.Nil(t, provider)
	})

	t.Run("empty issuer fails", func(t *testing.T) {
		provider, err := NewExternalOIDCProvider("client-id")
		require.ErrorIs(t, err, ErrEmptyOIDCIssuer)
		assert.Nil(t, provider)
	})

	t.Run("malformed discovery document fails", func(t *testing.T) {
		server := newOIDCDocumentServer(t, `{"issuer":"https://idp.example.com"}`)

		provider, err := NewExternalOIDCProvider("client-id", WithOIDCIssuer(server.URL), WithHTTPClient(server.Client()))
		require.ErrorContains(t, err, "authorization_endpoint")
		assert.Nil(t, provider)
	})

	t.Run("valid discovery returns a bearer provider", func(t *testing.T) {
		idp := newMockIDP(t, func(string) string {
			return idTokenResponse(fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix()}))
		})

		provider, err := NewExternalOIDCProvider("client-id", WithOIDCIssuer(idp.server.URL), WithHTTPClient(idp.server.Client()))
		require.NoError(t, err)
		assert.IsType(t, &BearerTokenAuthenticationProvider{}, provider)
	})
}

func TestExternalOIDCProvider_Contract(t *testing.T) {
	idToken := fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix(), "aud": "client-id"})
	idp := newMockIDP(t, func(string) string { return idTokenResponse(idToken) })

	var captured string
	provider := mustTestOIDCProvider(t, idp, &captured)

	request := newTestRequest()
	require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
	assert.Equal(t, []string{"Bearer " + idToken}, request.Headers.Get("Authorization"))

	// The authorization URL carried the PKCE challenge, state, and default scope.
	parsed, err := url.Parse(captured)
	require.NoError(t, err)
	query := parsed.Query()
	assert.Equal(t, "S256", query.Get("code_challenge_method"))
	assert.NotEmpty(t, query.Get("code_challenge"))
	assert.Equal(t, "test-state", query.Get("state"))
	assert.Equal(t, "openid", query.Get("scope"))

	// The token is cached and reused within its TTL.
	codeCalls, _ := idp.tokenCalls()
	assert.Equal(t, 1, codeCalls)

	second := newTestRequest()
	require.NoError(t, provider.AuthenticateRequest(context.Background(), second, nil))
	assert.Equal(t, []string{"Bearer " + idToken}, second.Headers.Get("Authorization"))

	codeCalls, _ = idp.tokenCalls()
	assert.Equal(t, 1, codeCalls)
}

func TestExternalOIDCProvider_PresentsAccessTokenWithScopes(t *testing.T) {
	idToken := fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix()})
	accessToken := fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix()})
	idp := newMockIDP(t, func(string) string {
		return `{"access_token":"` + accessToken + `","id_token":"` + idToken + `","token_type":"Bearer","expires_in":3600}`
	})

	var captured string
	provider := mustTestOIDCProvider(t, idp, &captured, WithScopes("api://client-id/.default"))

	request := newTestRequest()
	require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
	assert.Equal(t, []string{"Bearer " + accessToken}, request.Headers.Get("Authorization"))

	parsed, err := url.Parse(captured)
	require.NoError(t, err)
	assert.Equal(t, "api://client-id/.default", parsed.Query().Get("scope"))
}

func TestExternalOIDCProvider_RefreshesExpiredToken(t *testing.T) {
	expired := fakeJWT(t, map[string]any{"exp": time.Now().Add(-time.Minute).Unix()})
	fresh := fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix()})
	idp := newMockIDP(t, func(grant string) string {
		if grant == "refresh_token" {
			return `{"access_token":"access-token","id_token":"` + fresh + `","token_type":"Bearer","expires_in":3600}`
		}
		return `{"access_token":"access-token","id_token":"` + expired + `","token_type":"Bearer","expires_in":3600,"refresh_token":"refresh-token"}`
	})

	provider := mustTestOIDCProvider(t, idp, nil)

	first := newTestRequest()
	require.NoError(t, provider.AuthenticateRequest(context.Background(), first, nil))
	assert.Equal(t, []string{"Bearer " + expired}, first.Headers.Get("Authorization"))

	second := newTestRequest()
	require.NoError(t, provider.AuthenticateRequest(context.Background(), second, nil))
	assert.Equal(t, []string{"Bearer " + fresh}, second.Headers.Get("Authorization"))

	codeCalls, refreshCalls := idp.tokenCalls()
	assert.Equal(t, 1, codeCalls)
	assert.Equal(t, 1, refreshCalls)
}

func TestExternalOIDCProvider_PersistsTokensAcrossRestarts(t *testing.T) {
	idToken := fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix()})
	idp := newMockIDP(t, func(string) string { return idTokenResponse(idToken) })
	store := newMemoryTokenStore()

	first := mustTestOIDCProvider(t, idp, nil, WithTokenStore(store))
	request := newTestRequest()
	require.NoError(t, first.AuthenticateRequest(context.Background(), request, nil))
	assert.Equal(t, []string{"Bearer " + idToken}, request.Headers.Get("Authorization"))
	require.Contains(t, store.tokens, "client-id")

	// A fresh provider sharing the store (simulating a restart) reuses the token
	// without contacting the identity provider again.
	second := mustTestOIDCProvider(t, idp, nil, WithTokenStore(store))
	restarted := newTestRequest()
	require.NoError(t, second.AuthenticateRequest(context.Background(), restarted, nil))
	assert.Equal(t, []string{"Bearer " + idToken}, restarted.Headers.Get("Authorization"))

	codeCalls, _ := idp.tokenCalls()
	assert.Equal(t, 1, codeCalls, "the restarted provider should reuse the persisted token")
}

func TestExternalOIDCProvider_RevokeClearsTokenWithoutEndpoint(t *testing.T) {
	idToken := fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix()})
	idp := newMockIDP(t, func(string) string { return idTokenResponse(idToken) })
	store := newMemoryTokenStore()
	provider := mustTestOIDCProvider(t, idp, nil, WithTokenStore(store))

	request := newTestRequest()
	require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
	require.Contains(t, store.tokens, "client-id")

	require.NoError(t, provider.Revoke(context.Background()))
	assert.NotContains(t, store.tokens, "client-id")
	assert.Equal(t, 0, idp.revocationCalls(), "no revocation endpoint should be contacted")

	again := newTestRequest()
	require.NoError(t, provider.AuthenticateRequest(context.Background(), again, nil))

	codeCalls, _ := idp.tokenCalls()
	assert.Equal(t, 2, codeCalls, "revoke should clear the cached token")
}

func TestExternalOIDCProvider_RevokesAtProviderWithEndpoint(t *testing.T) {
	idToken := fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix()})
	idp := newMockIDPWithRevocation(t, func(string) string {
		return `{"access_token":"` + idToken + `","id_token":"` + idToken + `","token_type":"Bearer","expires_in":3600}`
	}, true)
	provider := mustTestOIDCProvider(t, idp, nil)

	request := newTestRequest()
	require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))

	require.NoError(t, provider.Revoke(context.Background()))
	assert.Equal(t, 1, idp.revocationCalls())
}

// mustTestOIDCProvider builds an external OIDC provider wired to a mock IdP and a
// mock redirect server, so no browser or local listener is required.
func mustTestOIDCProvider(t *testing.T, idp *mockIDP, captured *string, extra ...AuthOption) *BearerTokenAuthenticationProvider {
	t.Helper()

	server := &mockAuthorizationCodeServer{}
	server.On("GetAddr").Return("http://localhost:5001/callback")
	server.On("Result", mockAnyContext).Return("auth-code", "test-state", nil)
	server.On("Shutdown", mockAnyContext).Return(nil)

	opts := append([]AuthOption{
		WithOIDCIssuer(idp.server.URL),
		WithHTTPClient(idp.server.Client()),
		WithURL("https://dev12345.service-now.com"),
		WithStateGenerator(func() string { return "test-state" }),
		WithURLOpener(func(authorizationURL string) error {
			if captured != nil {
				*captured = authorizationURL
			}
			return nil
		}),
		WithServerFactory(authCodeTestServerFactory(server)),
	}, extra...)

	provider, err := NewExternalOIDCProvider("client-id", opts...)
	require.NoError(t, err)

	wrapper, ok := provider.(*BearerTokenAuthenticationProvider)
	require.True(t, ok)

	return wrapper
}

// mockIDP serves an OIDC discovery document, a token endpoint, and optionally a
// revocation endpoint. It counts the requests it receives.
type mockIDP struct {
	server       *httptest.Server
	tokenFunc    func(grantType string) string
	revoke       bool
	mu           sync.Mutex
	codeCalls    int
	refreshCalls int
	revokeCalls  int
}

func newMockIDP(t *testing.T, tokenFunc func(grantType string) string) *mockIDP {
	t.Helper()
	return newMockIDPWithRevocation(t, tokenFunc, false)
}

func newMockIDPWithRevocation(t *testing.T, tokenFunc func(grantType string) string, revoke bool) *mockIDP {
	t.Helper()
	idp := &mockIDP{tokenFunc: tokenFunc, revoke: revoke}
	idp.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/.well-known/openid-configuration":
			document := `{"issuer":"` + idp.server.URL + `","authorization_endpoint":"` + idp.server.URL + `/authorize","token_endpoint":"` + idp.server.URL + `/token"`
			if idp.revoke {
				document += `,"revocation_endpoint":"` + idp.server.URL + `/revoke"`
			}
			document += `}`
			_, _ = io.WriteString(w, document)
		case "/token":
			_ = r.ParseForm()
			grant := r.PostForm.Get("grant_type")
			idp.mu.Lock()
			if grant == "refresh_token" {
				idp.refreshCalls++
			} else {
				idp.codeCalls++
			}
			idp.mu.Unlock()
			_, _ = io.WriteString(w, tokenFunc(grant))
		case "/revoke":
			idp.mu.Lock()
			idp.revokeCalls++
			idp.mu.Unlock()
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(idp.server.Close)
	return idp
}

func (m *mockIDP) tokenCalls() (code, refresh int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.codeCalls, m.refreshCalls
}

func (m *mockIDP) revocationCalls() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.revokeCalls
}
