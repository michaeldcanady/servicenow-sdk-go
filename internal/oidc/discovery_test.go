// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package oidc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockHTTPClient implements oauth2.HTTPClient for testing.
type mockHTTPClient struct {
	handler func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.handler != nil {
		return m.handler(req)
	}
	// Default: make a real HTTP request
	return http.DefaultClient.Do(req)
}

func TestDiscover_Success(t *testing.T) {
	var server *httptest.Server //nolint:staticcheck
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/.well-known/openid-configuration", r.URL.Path)

		doc := map[string]interface{}{
			"issuer":                 server.URL,
			"authorization_endpoint": server.URL + "/authorize",
			"token_endpoint":         server.URL + "/token",
			"revocation_endpoint":    server.URL + "/revoke",
			"jwks_uri":               server.URL + "/jwks",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(doc)
	}))
	defer server.Close()

	client := &mockHTTPClient{}
	endpoints, err := Discover(context.Background(), server.URL, client)

	require.NoError(t, err)
	assert.Equal(t, server.URL+"/authorize", endpoints.AuthURL)
	assert.Equal(t, server.URL+"/token", endpoints.TokenURL)
	assert.Equal(t, server.URL+"/revoke", endpoints.RevocationURL)
}

func TestDiscover_SuccessWithoutRevocationEndpoint(t *testing.T) {
	var server *httptest.Server //nolint:staticcheck
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		doc := map[string]interface{}{
			"issuer":                 server.URL,
			"authorization_endpoint": server.URL + "/authorize",
			"token_endpoint":         server.URL + "/token",
		}
		json.NewEncoder(w).Encode(doc)
	}))
	defer server.Close()

	client := &mockHTTPClient{}
	endpoints, err := Discover(context.Background(), server.URL, client)

	require.NoError(t, err)
	assert.Empty(t, endpoints.RevocationURL)
}

func TestDiscover_MissingIssuer(t *testing.T) {
	var server *httptest.Server //nolint:staticcheck
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		doc := map[string]interface{}{
			"authorization_endpoint": server.URL + "/authorize",
			"token_endpoint":         server.URL + "/token",
		}
		json.NewEncoder(w).Encode(doc)
	}))
	defer server.Close()

	client := &mockHTTPClient{}
	_, err := Discover(context.Background(), server.URL, client)

	require.Error(t, err)
	require.ErrorIs(t, err, errors.ErrOIDCDiscoveryMissingEndpoint)
	assert.Contains(t, err.Error(), "issuer")
}

func TestDiscover_MissingAuthorizationEndpoint(t *testing.T) {
	var server *httptest.Server //nolint:staticcheck
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		doc := map[string]interface{}{
			"issuer":         server.URL,
			"token_endpoint": server.URL + "/token",
		}
		json.NewEncoder(w).Encode(doc)
	}))
	defer server.Close()

	client := &mockHTTPClient{}
	_, err := Discover(context.Background(), server.URL, client)

	require.Error(t, err)
	require.ErrorIs(t, err, errors.ErrOIDCDiscoveryMissingEndpoint)
	assert.Contains(t, err.Error(), "authorization_endpoint")
}

func TestDiscover_MissingTokenEndpoint(t *testing.T) {
	var server *httptest.Server //nolint:staticcheck
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		doc := map[string]interface{}{
			"issuer":                 server.URL,
			"authorization_endpoint": server.URL + "/authorize",
		}
		json.NewEncoder(w).Encode(doc)
	}))
	defer server.Close()

	client := &mockHTTPClient{}
	_, err := Discover(context.Background(), server.URL, client)

	require.Error(t, err)
	require.ErrorIs(t, err, errors.ErrOIDCDiscoveryMissingEndpoint)
	assert.Contains(t, err.Error(), "token_endpoint")
}

func TestDiscover_IssuerMismatch(t *testing.T) {
	var server *httptest.Server //nolint:staticcheck
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		doc := map[string]interface{}{
			"issuer":                 "https://different-idp.example.com",
			"authorization_endpoint": server.URL + "/authorize",
			"token_endpoint":         server.URL + "/token",
		}
		json.NewEncoder(w).Encode(doc)
	}))
	defer server.Close()

	client := &mockHTTPClient{}
	_, err := Discover(context.Background(), server.URL, client)

	require.Error(t, err)
	assert.ErrorIs(t, err, errors.ErrOIDCDiscoveryIssuerMismatch)
}

func TestDiscover_HTTPError(t *testing.T) {
	var server *httptest.Server //nolint:staticcheck
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := &mockHTTPClient{}
	_, err := Discover(context.Background(), server.URL, client)

	require.Error(t, err)
	require.ErrorIs(t, err, errors.ErrOIDCDiscoveryFailed)
	assert.Contains(t, err.Error(), "HTTP 404")
}

func TestDiscover_NetworkError(t *testing.T) {
	client := &mockHTTPClient{
		handler: func(req *http.Request) (*http.Response, error) {
			return nil, fmt.Errorf("network error")
		},
	}
	_, err := Discover(context.Background(), "https://idp.example.com", client)

	require.Error(t, err)
	require.ErrorIs(t, err, errors.ErrOIDCDiscoveryFailed)
}

func TestDiscover_InvalidJSON(t *testing.T) {
	var server *httptest.Server //nolint:staticcheck
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{invalid json"))
	}))
	defer server.Close()

	client := &mockHTTPClient{}
	_, err := Discover(context.Background(), server.URL, client)

	require.Error(t, err)
	require.ErrorIs(t, err, errors.ErrOIDCDiscoveryFailed)
}

func TestDiscover_FullMetadataURL(t *testing.T) {
	var server *httptest.Server //nolint:staticcheck
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/.well-known/openid-configuration", r.URL.Path)

		doc := map[string]interface{}{
			"issuer":                 server.URL,
			"authorization_endpoint": server.URL + "/authorize",
			"token_endpoint":         server.URL + "/token",
		}
		json.NewEncoder(w).Encode(doc)
	}))
	defer server.Close()

	client := &mockHTTPClient{}
	endpoints, err := Discover(context.Background(), server.URL, client)

	require.NoError(t, err)
	assert.Equal(t, server.URL+"/authorize", endpoints.AuthURL)
}

func TestDiscover_IssuerWithTrailingSlash(t *testing.T) {
	var server *httptest.Server //nolint:staticcheck
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		doc := map[string]interface{}{
			"issuer":                 server.URL + "/",
			"authorization_endpoint": server.URL + "/authorize",
			"token_endpoint":         server.URL + "/token",
		}
		json.NewEncoder(w).Encode(doc)
	}))
	defer server.Close()

	client := &mockHTTPClient{}
	_, err := Discover(context.Background(), server.URL, client)

	require.NoError(t, err)
}

func TestDiscover_IssuerCaseInsensitiveHost(t *testing.T) {
	var server *httptest.Server //nolint:staticcheck
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		doc := map[string]interface{}{
			"issuer":                 server.URL,
			"authorization_endpoint": server.URL + "/authorize",
			"token_endpoint":         server.URL + "/token",
		}
		json.NewEncoder(w).Encode(doc)
	}))
	defer server.Close()

	client := &mockHTTPClient{}
	_, err := Discover(context.Background(), server.URL, client)

	require.NoError(t, err)
}

func TestResolveMetadataURL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"https://idp.example.com", "https://idp.example.com/.well-known/openid-configuration"},
		{"https://idp.example.com/", "https://idp.example.com/.well-known/openid-configuration"},
		{"https://idp.example.com/.well-known/openid-configuration", "https://idp.example.com/.well-known/openid-configuration"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, resolveMetadataURL(tt.input))
		})
	}
}
