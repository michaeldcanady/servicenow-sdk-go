// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func clientCredentialsConfig(server *httptest.Server) ClientCredentialsConfig {
	parsed, err := url.Parse(server.URL)
	if err != nil {
		panic(err)
	}
	return ClientCredentialsConfig{
		ProviderConfig: ProviderConfig{
			HTTPClient:   server.Client(),
			AllowedHosts: []string{parsed.Hostname()},
		},
		ClientID:     "client",
		ClientSecret: "secret",
		Scopes:       []string{"scope:b", "scope:a"},
	}
}

func TestClientCredentialsConfigValidate(t *testing.T) {
	valid := ClientCredentialsConfig{
		ClientID:     "client",
		ClientSecret: "secret",
	}

	tests := []struct {
		name   string
		mutate func(*ClientCredentialsConfig)
		want   error
	}{
		{name: "valid", mutate: func(*ClientCredentialsConfig) {}},
		{name: "empty client ID", mutate: func(config *ClientCredentialsConfig) { config.ClientID = " " }, want: snerrors.ErrEmptyClientID},
		{name: "empty client secret", mutate: func(config *ClientCredentialsConfig) { config.ClientSecret = " " }, want: snerrors.ErrEmptyClientSecret},
		{name: "none", mutate: func(config *ClientCredentialsConfig) { config.ClientAuthMethod = ClientAuthMethodNone }, want: snerrors.ErrInvalidAuthMethod},
		{name: "unknown", mutate: func(config *ClientCredentialsConfig) { config.ClientAuthMethod = ClientAuthMethodUnknown }, want: snerrors.ErrInvalidAuthMethod},
		{name: "invalid host", mutate: func(config *ClientCredentialsConfig) {
			config.AllowedHosts = []string{"https://example.com"}
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := valid
			test.mutate(&config)
			err := config.Validate()
			if test.want != nil {
				require.ErrorIs(t, err, test.want)
			} else if test.name == "invalid host" {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewClientCredentialsProviderDoesNotMakeRequests(t *testing.T) {
	var calls atomic.Int32
	client := &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		calls.Add(1)
		return nil, errors.New("unexpected request")
	})}

	_, err := NewClientCredentialsProvider(ClientCredentialsConfig{
		ProviderConfig: ProviderConfig{HTTPClient: client},
		ClientID:       "client",
		ClientSecret:   "secret",
	})
	require.NoError(t, err)
	assert.Equal(t, int32(0), calls.Load())
}

func TestClientCredentialsProviderInitialize(t *testing.T) {
	provider, err := NewClientCredentialsProvider(ClientCredentialsConfig{
		ClientID:     "client",
		ClientSecret: "secret",
	})
	require.NoError(t, err)

	require.ErrorIs(t, provider.Initialize(""), snerrors.ErrEmptyURI)
	require.ErrorIs(t, provider.Initialize("http://example.com"), snerrors.ErrInsecureURL)
	require.ErrorIs(t, provider.Initialize("://bad"), snerrors.ErrInvalidURL)
	require.NoError(t, provider.Initialize("https://example.com/"))
	require.NoError(t, provider.Initialize("https://EXAMPLE.com"))
	require.ErrorIs(t, provider.Initialize("https://other.example.com"), snerrors.ErrAlreadyInitialized)

	var nilProvider *ClientCredentialsProvider
	require.ErrorIs(t, nilProvider.Initialize("https://example.com"), snerrors.ErrNilProvider)
}

func TestClientCredentialsProviderRequiresInitialization(t *testing.T) {
	provider, err := NewClientCredentialsProvider(ClientCredentialsConfig{
		ClientID:     "client",
		ClientSecret: "secret",
	})
	require.NoError(t, err)

	request := testRequest(t, "https://example.com/api/now")
	require.ErrorIs(t, provider.AuthenticateRequest(context.Background(), request, nil), snerrors.ErrNotInitialized)
	require.ErrorIs(t, provider.Revoke(context.Background()), snerrors.ErrNotInitialized)

	var nilProvider *ClientCredentialsProvider
	require.ErrorIs(t, nilProvider.AuthenticateRequest(context.Background(), request, nil), snerrors.ErrNilProvider)
	require.ErrorIs(t, nilProvider.Revoke(context.Background()), snerrors.ErrNilProvider)
}

func TestClientCredentialsProviderAuthenticates(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		calls.Add(1)
		assert.NoError(t, request.ParseForm())
		assert.Equal(t, "/oauth_token.do", request.URL.Path)
		assert.Equal(t, "client_credentials", request.Form.Get("grant_type"))
		assert.Equal(t, "client", request.Form.Get("client_id"))
		assert.Equal(t, "secret", request.Form.Get("client_secret"))
		assert.Equal(t, "scope:a scope:b", request.Form.Get("scope"))

		writer.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(writer, `{"access_token":"token","token_type":"Bearer","expires_in":3600,"scope":"scope:a scope:b"}`)
	}))
	defer server.Close()

	config := clientCredentialsConfig(server)
	provider, err := NewClientCredentialsProvider(config)
	require.NoError(t, err)
	require.NoError(t, provider.Initialize(server.URL))

	first := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, provider.AuthenticateRequest(context.Background(), first, nil))
	assert.Equal(t, []string{"Bearer token"}, first.Headers.Get("Authorization"))

	second := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, provider.AuthenticateRequest(context.Background(), second, nil))
	assert.Equal(t, []string{"Bearer token"}, second.Headers.Get("Authorization"))
	assert.Equal(t, int32(1), calls.Load())
}

func TestClientCredentialsProviderClientSecretBasic(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		calls.Add(1)
		assert.NoError(t, request.ParseForm())
		assert.Empty(t, request.Form.Get("client_secret"))
		username, password, ok := request.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, "client", username)
		assert.Equal(t, "secret", password)

		writer.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(writer, `{"access_token":"token","token_type":"Bearer","expires_in":3600}`)
	}))
	defer server.Close()

	config := clientCredentialsConfig(server)
	config.ClientAuthMethod = ClientAuthMethodClientSecretBasic
	provider, err := NewClientCredentialsProvider(config)
	require.NoError(t, err)
	require.NoError(t, provider.Initialize(server.URL))

	request := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
	assert.Equal(t, []string{"Bearer token"}, request.Headers.Get("Authorization"))
	assert.Equal(t, int32(1), calls.Load())
}

func TestClientCredentialsProviderRefreshes(t *testing.T) {
	clock := newFakeClock(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	var calls atomic.Int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		calls.Add(1)
		assert.NoError(t, request.ParseForm())

		writer.Header().Set("Content-Type", "application/json")
		switch request.Form.Get("grant_type") {
		case "client_credentials":
			_, _ = io.WriteString(writer, `{"access_token":"old","refresh_token":"refresh","expires_in":60}`)
		case "refresh_token":
			assert.Equal(t, "refresh", request.Form.Get("refresh_token"))
			_, _ = io.WriteString(writer, `{"access_token":"new","expires_in":3600}`)
		default:
			t.Errorf("unexpected grant type %q", request.Form.Get("grant_type"))
		}
	}))
	defer server.Close()

	provider, err := NewClientCredentialsProvider(clientCredentialsConfig(server))
	require.NoError(t, err)
	provider.lifecycle.now = clock.Now
	require.NoError(t, provider.Initialize(server.URL))

	first := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, provider.AuthenticateRequest(context.Background(), first, nil))
	assert.Equal(t, []string{"Bearer old"}, first.Headers.Get("Authorization"))

	clock.Advance(31 * time.Second)
	second := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, provider.AuthenticateRequest(context.Background(), second, nil))
	assert.Equal(t, []string{"Bearer new"}, second.Headers.Get("Authorization"))
	assert.Equal(t, int32(2), calls.Load())
}

func TestClientCredentialsProviderReacquiresAfterInvalidGrant(t *testing.T) {
	clock := newFakeClock(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	var calls atomic.Int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		call := calls.Add(1)
		assert.NoError(t, request.ParseForm())

		writer.Header().Set("Content-Type", "application/json")
		switch request.Form.Get("grant_type") {
		case "client_credentials":
			if call == 1 {
				_, _ = io.WriteString(writer, `{"access_token":"old","refresh_token":"refresh","expires_in":60}`)
				return
			}
			_, _ = io.WriteString(writer, `{"access_token":"new","expires_in":3600}`)
		case "refresh_token":
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = io.WriteString(writer, `{"error":"invalid_grant","error_description":"refresh expired"}`)
		default:
			t.Errorf("unexpected grant type %q", request.Form.Get("grant_type"))
		}
	}))
	defer server.Close()

	provider, err := NewClientCredentialsProvider(clientCredentialsConfig(server))
	require.NoError(t, err)
	provider.lifecycle.now = clock.Now
	require.NoError(t, provider.Initialize(server.URL))

	first := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, provider.AuthenticateRequest(context.Background(), first, nil))
	assert.Equal(t, []string{"Bearer old"}, first.Headers.Get("Authorization"))

	clock.Advance(31 * time.Second)
	second := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, provider.AuthenticateRequest(context.Background(), second, nil))
	assert.Equal(t, []string{"Bearer new"}, second.Headers.Get("Authorization"))
	assert.Equal(t, int32(3), calls.Load())
}

func TestClientCredentialsProviderTokenStore(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		calls.Add(1)
		writer.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(writer, `{"access_token":"token","expires_in":3600}`)
	}))
	defer server.Close()

	store := NewMemoryTokenStore()
	config := clientCredentialsConfig(server)
	config.TokenStore = store

	firstProvider, err := NewClientCredentialsProvider(config)
	require.NoError(t, err)
	require.NoError(t, firstProvider.Initialize(server.URL))
	firstRequest := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, firstProvider.AuthenticateRequest(context.Background(), firstRequest, nil))
	assert.Equal(t, int32(1), calls.Load())

	secondProvider, err := NewClientCredentialsProvider(config)
	require.NoError(t, err)
	require.NoError(t, secondProvider.Initialize(server.URL))
	secondRequest := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, secondProvider.AuthenticateRequest(context.Background(), secondRequest, nil))
	assert.Equal(t, []string{"Bearer token"}, secondRequest.Headers.Get("Authorization"))
	assert.Equal(t, int32(1), calls.Load())
}

func TestClientCredentialsProviderRevokes(t *testing.T) {
	var tokenCalls atomic.Int32
	var revokeCalls atomic.Int32
	var revoked []string
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		assert.NoError(t, request.ParseForm())
		switch request.URL.Path {
		case "/oauth_token.do":
			tokenCalls.Add(1)
			writer.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(writer, `{"access_token":"token","refresh_token":"refresh","expires_in":3600}`)
		case "/oauth_revoke.do":
			revokeCalls.Add(1)
			revoked = append(revoked, request.Form.Get("token"), request.Form.Get("token_type_hint"))
			writer.WriteHeader(http.StatusOK)
		default:
			t.Errorf("unexpected path %q", request.URL.Path)
		}
	}))
	defer server.Close()

	provider, err := NewClientCredentialsProvider(clientCredentialsConfig(server))
	require.NoError(t, err)
	require.NoError(t, provider.Initialize(server.URL))

	request := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
	require.NoError(t, provider.Revoke(context.Background()))
	assert.Equal(t, int32(2), revokeCalls.Load())
	assert.Equal(t, []string{"refresh", "refresh_token", "token", "access_token"}, revoked)

	again := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, provider.AuthenticateRequest(context.Background(), again, nil))
	assert.Equal(t, int32(2), tokenCalls.Load())
}

func TestClientCredentialsProviderReportsOAuthErrors(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		_, _ = io.WriteString(writer, `{"error":"invalid_client","error_description":"bad secret"}`)
	}))
	defer server.Close()

	provider, err := NewClientCredentialsProvider(clientCredentialsConfig(server))
	require.NoError(t, err)
	require.NoError(t, provider.Initialize(server.URL))

	request := testRequest(t, server.URL+"/api/now/table/incident")
	err = provider.AuthenticateRequest(context.Background(), request, nil)
	require.Error(t, err)
	require.ErrorIs(t, err, snerrors.ErrTokenAcquisition)

	var oauthError *OAuthError
	require.ErrorAs(t, err, &oauthError)
	assert.Equal(t, "acquire", oauthError.Operation)
	assert.Equal(t, FlowClientCredentials, oauthError.Flow)
	assert.Equal(t, "invalid_client", oauthError.Code)
	assert.Equal(t, http.StatusUnauthorized, oauthError.StatusCode)
	assert.NotContains(t, err.Error(), "bad secret")
	assert.NotContains(t, err.Error(), "secret")
}

func TestClientCredentialsProviderHonorsExistingAuthorization(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		calls.Add(1)
		writer.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(writer, `{"access_token":"token","expires_in":3600}`)
	}))
	defer server.Close()

	provider, err := NewClientCredentialsProvider(clientCredentialsConfig(server))
	require.NoError(t, err)
	require.NoError(t, provider.Initialize(server.URL))

	request := testRequest(t, server.URL+"/api/now/table/incident")
	request.Headers.Add("Authorization", "Bearer caller")
	require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
	assert.Equal(t, []string{"Bearer caller"}, request.Headers.Get("Authorization"))
	assert.Equal(t, int32(0), calls.Load())
}

func TestClientCredentialsProviderHostIsolation(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		calls.Add(1)
		writer.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(writer, `{"access_token":"token","expires_in":3600}`)
	}))
	defer server.Close()

	config := clientCredentialsConfig(server)
	config.AllowedHosts = []string{"example.com"}
	provider, err := NewClientCredentialsProvider(config)
	require.NoError(t, err)
	require.NoError(t, provider.Initialize(server.URL))

	request := testRequest(t, server.URL+"/api/now/table/incident")
	require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
	assert.Empty(t, request.Headers.Get("Authorization"))
	assert.Equal(t, int32(0), calls.Load())
}

func TestTokenFromOAuth2(t *testing.T) {
	now := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)

	token, err := tokenFromOAuth2(&oauth2.Token{
		AccessToken:  "access",
		RefreshToken: "refresh",
		TokenType:    "Bearer",
		Scope:        "scope",
		ExpiresIn:    60,
		Raw:          map[string]any{"id_token": "id"},
	}, now)
	require.NoError(t, err)
	assert.Equal(t, "access", token.AccessToken)
	assert.Equal(t, "refresh", token.RefreshToken)
	assert.Equal(t, "id", token.IDToken)
	assert.Equal(t, now.Add(time.Minute), token.ExpiresAt)

	token, err = tokenFromOAuth2(&oauth2.Token{AccessToken: "access"}, now)
	require.NoError(t, err)
	assert.True(t, token.ExpiresAt.IsZero())

	token, err = tokenFromOAuth2(nil, now)
	assert.Nil(t, token)
	require.ErrorIs(t, err, snerrors.ErrMissingAccessToken)
}

func TestClientCredentialsProviderErrorFormat(t *testing.T) {
	provider, err := NewClientCredentialsProvider(ClientCredentialsConfig{
		ClientID:     "client",
		ClientSecret: "super-secret",
	})
	require.NoError(t, err)

	request := testRequest(t, "https://example.com/api/now")
	err = provider.AuthenticateRequest(context.Background(), request, nil)
	require.Error(t, err)
	assert.NotContains(t, err.Error(), "super-secret")
}
