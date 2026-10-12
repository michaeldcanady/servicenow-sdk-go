// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeJWT builds an unsigned JWT carrying the provided claims. The signature is
// never verified, so the value is irrelevant to the tests.
func fakeJWT(t *testing.T, claims map[string]any) string {
	t.Helper()
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"none","typ":"JWT"}`))
	payload, err := json.Marshal(claims)
	require.NoError(t, err)
	return header + "." + base64.RawURLEncoding.EncodeToString(payload) + ".signature"
}

// idTokenResponse builds a token endpoint response that includes the given ID token.
func idTokenResponse(idToken string) string {
	return `{"access_token":"access-token","id_token":"` + idToken + `","token_type":"Bearer","expires_in":3600,"refresh_token":"refresh-token"}`
}

// recordingTokenServer is a token endpoint that records the form of every request.
type recordingTokenServer struct {
	mu      sync.Mutex
	forms   []url.Values
	server  *httptest.Server
	respond func(form url.Values) (status int, body string)
}

func newRecordingTokenServer(t *testing.T, respond func(form url.Values) (int, string)) *recordingTokenServer {
	t.Helper()
	r := &recordingTokenServer{respond: respond}
	r.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_ = req.ParseForm()
		r.mu.Lock()
		r.forms = append(r.forms, req.PostForm)
		r.mu.Unlock()

		status, body := respond(req.PostForm)
		w.WriteHeader(status)
		_, _ = io.WriteString(w, body)
	}))
	t.Cleanup(r.server.Close)
	return r
}

func (r *recordingTokenServer) endpoints() *oauth2.Endpoints {
	return &oauth2.Endpoints{
		AuthURL:  r.server.URL + "/authorize",
		TokenURL: r.server.URL + "/token",
	}
}

func (r *recordingTokenServer) lastForm(t *testing.T) url.Values {
	t.Helper()
	r.mu.Lock()
	defer r.mu.Unlock()
	require.NotEmpty(t, r.forms)
	return r.forms[len(r.forms)-1]
}

// newClientForServer builds an external OIDC client pointed at the recording server.
func newClientForServer(t *testing.T, server *recordingTokenServer, scopes []string, useAccessToken bool) *externalOIDCClient {
	t.Helper()
	client, err := newExternalOIDCClient("client-id", server.endpoints(), scopes, useAccessToken, server.server.Client())
	require.NoError(t, err)
	return client
}

func TestNewExternalOIDCClient(t *testing.T) {
	t.Run("empty client id fails", func(t *testing.T) {
		client, err := newExternalOIDCClient("", &oauth2.Endpoints{}, nil, false, nil)
		require.ErrorIs(t, err, ErrEmptyClientID)
		assert.Nil(t, client)
	})

	t.Run("valid client builds the oauth client", func(t *testing.T) {
		client, err := newExternalOIDCClient("id", &oauth2.Endpoints{TokenURL: "https://idp/token", AuthURL: "https://idp/auth"}, []string{"openid"}, false, nil)
		require.NoError(t, err)
		assert.NotNil(t, client.oauthClient)
	})
}

func TestExternalOIDCClient_GetAuthorizationURL(t *testing.T) {
	client, err := newExternalOIDCClient("client-id",
		&oauth2.Endpoints{AuthURL: "https://idp.example.com/authorize", TokenURL: "https://idp.example.com/token"},
		[]string{"openid"}, false, nil)
	require.NoError(t, err)

	raw, err := client.getAuthorizationURL("http://localhost:5001/callback", "state-123", nil)
	require.NoError(t, err)

	parsed, err := url.Parse(raw)
	require.NoError(t, err)
	query := parsed.Query()

	assert.Equal(t, "https://idp.example.com/authorize", parsed.Scheme+"://"+parsed.Host+parsed.Path)
	assert.Equal(t, "code", query.Get("response_type"))
	assert.Equal(t, "client-id", query.Get("client_id"))
	assert.Equal(t, "http://localhost:5001/callback", query.Get("redirect_uri"))
	assert.Equal(t, "state-123", query.Get("state"))
	assert.Equal(t, "openid", query.Get("scope"))
	assert.Equal(t, "S256", query.Get("code_challenge_method"))
	assert.NotEmpty(t, query.Get("code_challenge"))
	assert.Empty(t, query.Get("client_secret"))
	assert.NotEmpty(t, client.verifier)
}

func TestExternalOIDCClient_AcquireTokenByCode(t *testing.T) {
	idToken := fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix()})

	t.Run("presents the id token by default", func(t *testing.T) {
		server := newRecordingTokenServer(t, func(url.Values) (int, string) {
			return http.StatusOK, idTokenResponse(idToken)
		})
		client := newClientForServer(t, server, []string{"openid"}, false)
		client.verifier = "verifier-123"

		token, err := client.acquireTokenByCode(context.Background(), "auth-code", "http://localhost:5001/callback", "state-123")
		require.NoError(t, err)
		assert.Equal(t, idToken, token.AccessToken)
		assert.Equal(t, "refresh-token", token.RefreshToken)
		assert.False(t, token.IsExpired())

		form := server.lastForm(t)
		assert.Equal(t, "authorization_code", form.Get("grant_type"))
		assert.Equal(t, "auth-code", form.Get("code"))
		assert.Equal(t, "verifier-123", form.Get("code_verifier"))
		assert.Equal(t, "http://localhost:5001/callback", form.Get("redirect_uri"))
		assert.Equal(t, "client-id", form.Get("client_id"))
	})

	t.Run("presents the access token when scopes are configured", func(t *testing.T) {
		accessToken := fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix()})
		server := newRecordingTokenServer(t, func(url.Values) (int, string) {
			return http.StatusOK, `{"access_token":"` + accessToken + `","id_token":"` + idToken + `","token_type":"Bearer","expires_in":3600}`
		})
		client := newClientForServer(t, server, []string{"api://client-id/.default"}, true)
		client.verifier = "verifier-123"

		token, err := client.acquireTokenByCode(context.Background(), "auth-code", "http://localhost:5001/callback", "state-123")
		require.NoError(t, err)
		assert.Equal(t, accessToken, token.AccessToken)
	})

	t.Run("non-success token response propagates", func(t *testing.T) {
		server := newRecordingTokenServer(t, func(url.Values) (int, string) {
			return http.StatusBadRequest, `{"error":"invalid_grant"}`
		})
		client := newClientForServer(t, server, []string{"openid"}, false)
		client.verifier = "verifier-123"

		_, err := client.acquireTokenByCode(context.Background(), "auth-code", "http://localhost:5001/callback", "state-123")
		require.Error(t, err)
	})
}

func TestExternalOIDCClient_PresentedTokenSelection(t *testing.T) {
	valid := fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix()})

	tests := []struct {
		name           string
		useAccessToken bool
		response       string
		want           string
		wantErr        error
	}{
		{
			name:     "id token is presented by default",
			response: idTokenResponse(valid),
			want:     valid,
		},
		{
			name:     "missing id token fails",
			response: `{"access_token":"access-token","token_type":"Bearer"}`,
			wantErr:  ErrMissingIDToken,
		},
		{
			name:           "access token required but missing fails",
			useAccessToken: true,
			response:       `{"id_token":"` + valid + `","token_type":"Bearer"}`,
			wantErr:        ErrMissingAccessToken,
		},
		{
			name:     "unparseable token fails",
			response: `{"id_token":"not-a-jwt","token_type":"Bearer"}`,
			wantErr:  ErrUnparseableToken,
		},
		{
			name:     "token without exp fails",
			response: `{"id_token":"` + fakeJWT(t, map[string]any{"sub": "user"}) + `","token_type":"Bearer"}`,
			wantErr:  ErrMissingTokenExpiry,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := newRecordingTokenServer(t, func(url.Values) (int, string) {
				return http.StatusOK, test.response
			})
			client := newClientForServer(t, server, []string{"openid"}, test.useAccessToken)
			client.verifier = "verifier-123"

			token, err := client.acquireTokenByCode(context.Background(), "auth-code", "http://localhost:5001/callback", "state-123")

			if test.wantErr != nil {
				require.ErrorIs(t, err, test.wantErr)
				require.ErrorIs(t, err, ErrTokenAcquisition)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.want, token.AccessToken)
		})
	}
}

func TestExternalOIDCClient_AcquireTokenByCode_ProviderError(t *testing.T) {
	server := newRecordingTokenServer(t, func(url.Values) (int, string) {
		return http.StatusBadRequest, `{"error":"invalid_grant","error_description":"bad code"}`
	})
	client := newClientForServer(t, server, []string{"openid"}, false)
	client.verifier = "verifier-123"

	_, err := client.acquireTokenByCode(context.Background(), "auth-code", "http://localhost:5001/callback", "state-123")
	require.ErrorIs(t, err, ErrTokenAcquisition)
}

func TestExternalOIDCClient_AcquireTokenByRefreshToken_ProviderError(t *testing.T) {
	server := newRecordingTokenServer(t, func(url.Values) (int, string) {
		return http.StatusBadRequest, `{"error":"invalid_grant"}`
	})
	client := newClientForServer(t, server, []string{"openid"}, false)

	_, err := client.acquireTokenByRefreshToken(context.Background(), "refresh-token")
	require.ErrorIs(t, err, ErrTokenAcquisition)
}

func TestExternalOIDCClient_AcquireTokenByRefreshToken(t *testing.T) {
	idToken := fakeJWT(t, map[string]any{"exp": time.Now().Add(time.Hour).Unix()})
	server := newRecordingTokenServer(t, func(url.Values) (int, string) {
		return http.StatusOK, idTokenResponse(idToken)
	})
	client := newClientForServer(t, server, []string{"openid"}, false)

	token, err := client.acquireTokenByRefreshToken(context.Background(), "refresh-token")
	require.NoError(t, err)
	assert.Equal(t, idToken, token.AccessToken)

	form := server.lastForm(t)
	assert.Equal(t, "refresh_token", form.Get("grant_type"))
	assert.Equal(t, "refresh-token", form.Get("refresh_token"))
}

func TestExternalOIDCClient_RevokeToken(t *testing.T) {
	t.Run("without a revocation endpoint succeeds as a no-op", func(t *testing.T) {
		client, err := newExternalOIDCClient("client-id",
			&oauth2.Endpoints{AuthURL: "https://idp.example.com/authorize", TokenURL: "https://idp.example.com/token"},
			[]string{"openid"}, false, nil)
		require.NoError(t, err)

		assert.False(t, client.revocationSupported())
		require.NoError(t, client.revokeToken(context.Background(), "token", "access_token"))
	})

	t.Run("with a revocation endpoint succeeds", func(t *testing.T) {
		server := newRecordingTokenServer(t, func(url.Values) (int, string) {
			return http.StatusOK, ""
		})
		endpoints := server.endpoints()
		endpoints.RevocationURL = server.server.URL + "/revoke"

		client, err := newExternalOIDCClient("client-id", endpoints, []string{"openid"}, false, server.server.Client())
		require.NoError(t, err)

		assert.True(t, client.revocationSupported())
		require.NoError(t, client.revokeToken(context.Background(), "token", "access_token"))
	})
}
