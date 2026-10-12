// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOIDCDiscoveryURL(t *testing.T) {
	tests := []struct {
		name    string
		issuer  string
		want    string
		wantErr error
		wantMsg string
	}{
		{
			name:   "plain issuer appends the well-known path",
			issuer: "https://idp.example.com",
			want:   "https://idp.example.com/.well-known/openid-configuration",
		},
		{
			name:   "issuer with a trailing slash",
			issuer: "https://idp.example.com/",
			want:   "https://idp.example.com/.well-known/openid-configuration",
		},
		{
			name:   "issuer with a path segment",
			issuer: "https://idp.example.com/tenant",
			want:   "https://idp.example.com/tenant/.well-known/openid-configuration",
		},
		{
			name:   "full metadata url is used as-is",
			issuer: "https://idp.example.com/.well-known/openid-configuration",
			want:   "https://idp.example.com/.well-known/openid-configuration",
		},
		{
			name:    "empty issuer",
			issuer:  "",
			wantErr: ErrEmptyOIDCIssuer,
		},
		{
			name:    "relative issuer is rejected",
			issuer:  "idp.example.com",
			wantMsg: "must be an absolute URL",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := oidcDiscoveryURL(test.issuer)

			if test.wantErr != nil {
				require.ErrorIs(t, err, test.wantErr)
				return
			}
			if test.wantMsg != "" {
				require.ErrorContains(t, err, test.wantMsg)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestFetchOIDCDiscovery(t *testing.T) {
	t.Run("valid document", func(t *testing.T) {
		var server *httptest.Server
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/.well-known/openid-configuration" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			_, _ = io.WriteString(w, `{"issuer":"`+server.URL+`","authorization_endpoint":"`+server.URL+`/authorize","token_endpoint":"`+server.URL+`/token","revocation_endpoint":"`+server.URL+`/revoke"}`)
		}))
		t.Cleanup(server.Close)

		document, err := fetchOIDCDiscovery(context.Background(), server.URL, server.Client())
		require.NoError(t, err)
		assert.Equal(t, server.URL+"/authorize", document.AuthorizationEndpoint)
		assert.Equal(t, server.URL+"/token", document.TokenEndpoint)
		assert.Equal(t, server.URL+"/revoke", document.RevocationEndpoint)
		assert.Equal(t, server.URL+"/authorize", document.endpoints().AuthURL)
	})

	t.Run("full metadata url is used as-is", func(t *testing.T) {
		var server *httptest.Server
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/.well-known/openid-configuration" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			_, _ = io.WriteString(w, `{"authorization_endpoint":"`+server.URL+`/authorize","token_endpoint":"`+server.URL+`/token"}`)
		}))
		t.Cleanup(server.Close)

		_, err := fetchOIDCDiscovery(context.Background(), server.URL+"/.well-known/openid-configuration", server.Client())
		require.NoError(t, err)
	})

	t.Run("missing authorization endpoint fails", func(t *testing.T) {
		server := newOIDCDocumentServer(t, `{"issuer":"https://idp.example.com","token_endpoint":"https://idp.example.com/token"}`)

		_, err := fetchOIDCDiscovery(context.Background(), server.URL, server.Client())
		require.ErrorContains(t, err, "authorization_endpoint")
	})

	t.Run("missing token endpoint fails", func(t *testing.T) {
		server := newOIDCDocumentServer(t, `{"issuer":"https://idp.example.com","authorization_endpoint":"https://idp.example.com/authorize"}`)

		_, err := fetchOIDCDiscovery(context.Background(), server.URL, server.Client())
		require.ErrorContains(t, err, "token_endpoint")
	})

	t.Run("non-success status fails", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		t.Cleanup(server.Close)

		_, err := fetchOIDCDiscovery(context.Background(), server.URL, server.Client())
		require.ErrorContains(t, err, "status 500")
	})

	t.Run("malformed json fails", func(t *testing.T) {
		server := newOIDCDocumentServer(t, `not json`)

		_, err := fetchOIDCDiscovery(context.Background(), server.URL, server.Client())
		require.ErrorContains(t, err, "failed to parse")
	})
}

// newOIDCDocumentServer serves the given body at the well-known discovery path
// and returns 404 for any other path.
func newOIDCDocumentServer(t *testing.T, body string) *httptest.Server {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/.well-known/openid-configuration" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = io.WriteString(w, body)
	}))
	t.Cleanup(server.Close)
	return server
}
