// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"net/url"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRequest(t *testing.T, rawURL string) *abstractions.RequestInformation {
	t.Helper()
	uri, err := url.Parse(rawURL)
	require.NoError(t, err)
	request := abstractions.NewRequestInformation()
	request.SetUri(*uri)
	return request
}

func TestProviderConfigValidate(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.NoError(t, (ProviderConfig{}).Validate())
	})
	t.Run("valid host", func(t *testing.T) {
		assert.NoError(t, (ProviderConfig{AllowedHosts: []string{"example.com", ".example.org"}}).Validate())
	})
	t.Run("invalid host", func(t *testing.T) {
		assert.Error(t, (ProviderConfig{AllowedHosts: []string{"https://example.com"}}).Validate())
	})
}

func TestBasicProvider(t *testing.T) {
	config := BasicConfig{
		ProviderConfig: ProviderConfig{AllowedHosts: []string{"example.com"}},
		Username:       "user",
		Password:       "password",
	}

	t.Run("validates", func(t *testing.T) {
		_, err := NewBasicProvider(BasicConfig{})
		require.ErrorIs(t, err, snerrors.ErrEmptyUsername)

		_, err = NewBasicProvider(BasicConfig{Username: "user"})
		require.ErrorIs(t, err, snerrors.ErrEmptyPassword)

		_, err = NewBasicProvider(BasicConfig{
			ProviderConfig: ProviderConfig{AllowedHosts: []string{"https://example.com"}},
			Username:       "user",
			Password:       "password",
		})
		assert.Error(t, err)
	})

	t.Run("adds the header", func(t *testing.T) {
		provider, err := NewBasicProvider(config)
		require.NoError(t, err)

		request := testRequest(t, "https://example.com/api/now")
		require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
		assert.Equal(t, []string{"Basic dXNlcjpwYXNzd29yZA=="}, request.Headers.Get("Authorization"))
	})

	t.Run("does not override an existing header", func(t *testing.T) {
		provider, err := NewBasicProvider(config)
		require.NoError(t, err)

		request := testRequest(t, "https://example.com/api/now")
		request.Headers.Add("Authorization", "Bearer existing")
		require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
		assert.Equal(t, []string{"Bearer existing"}, request.Headers.Get("Authorization"))
	})

	t.Run("skips a host that is not allowed", func(t *testing.T) {
		provider, err := NewBasicProvider(config)
		require.NoError(t, err)

		request := testRequest(t, "https://other.example.com/api/now")
		require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
		assert.Empty(t, request.Headers.Get("Authorization"))
	})

	t.Run("requires HTTPS", func(t *testing.T) {
		provider, err := NewBasicProvider(BasicConfig{Username: "user", Password: "password"})
		require.NoError(t, err)

		request := testRequest(t, "http://example.com/api/now")
		require.ErrorIs(t, provider.AuthenticateRequest(context.Background(), request, nil), snerrors.ErrInsecureURL)
	})

	t.Run("rejects nil inputs", func(t *testing.T) {
		var provider *BasicProvider
		require.ErrorIs(t, provider.Initialize("https://example.com"), snerrors.ErrNilProvider)
		require.ErrorIs(t, provider.AuthenticateRequest(context.Background(), testRequest(t, "https://example.com"), nil), snerrors.ErrNilProvider)

		valid, err := NewBasicProvider(config)
		require.NoError(t, err)
		//nolint:staticcheck // Verify the explicit nil guard.
		require.ErrorIs(t, valid.AuthenticateRequest(nil, testRequest(t, "https://example.com"), nil), snerrors.ErrNilContext)
		require.ErrorIs(t, valid.AuthenticateRequest(context.Background(), nil, nil), snerrors.ErrNilRequestInfo)
		assert.NoError(t, valid.Initialize("https://example.com"))
	})
}

func TestBearerProvider(t *testing.T) {
	config := BearerConfig{
		ProviderConfig: ProviderConfig{AllowedHosts: []string{"example.com"}},
		Token:          "token",
	}

	t.Run("validates", func(t *testing.T) {
		_, err := NewBearerProvider(BearerConfig{})
		require.ErrorIs(t, err, snerrors.ErrEmptyToken)

		_, err = NewBearerProvider(BearerConfig{
			ProviderConfig: ProviderConfig{AllowedHosts: []string{"https://example.com"}},
			Token:          "token",
		})
		assert.Error(t, err)
	})

	t.Run("adds the header", func(t *testing.T) {
		provider, err := NewBearerProvider(config)
		require.NoError(t, err)

		request := testRequest(t, "https://example.com/api/now")
		require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
		assert.Equal(t, []string{"Bearer token"}, request.Headers.Get("Authorization"))
	})

	t.Run("skips a host that is not allowed", func(t *testing.T) {
		provider, err := NewBearerProvider(config)
		require.NoError(t, err)

		request := testRequest(t, "https://other.example.com/api/now")
		require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
		assert.Empty(t, request.Headers.Get("Authorization"))
	})

	t.Run("rejects nil inputs", func(t *testing.T) {
		var provider *BearerProvider
		require.ErrorIs(t, provider.Initialize("https://example.com"), snerrors.ErrNilProvider)
		require.ErrorIs(t, provider.AuthenticateRequest(context.Background(), testRequest(t, "https://example.com"), nil), snerrors.ErrNilProvider)

		valid, err := NewBearerProvider(config)
		require.NoError(t, err)
		//nolint:staticcheck // Verify the explicit nil guard.
		require.ErrorIs(t, valid.AuthenticateRequest(nil, testRequest(t, "https://example.com"), nil), snerrors.ErrNilContext)
		require.ErrorIs(t, valid.AuthenticateRequest(context.Background(), nil, nil), snerrors.ErrNilRequestInfo)
	})
}

func TestAPIKeyProvider(t *testing.T) {
	t.Run("validates", func(t *testing.T) {
		_, err := NewAPIKeyProvider(APIKeyConfig{})
		require.ErrorIs(t, err, snerrors.ErrEmptyAPIKey)

		_, err = NewAPIKeyProvider(APIKeyConfig{
			Key:      "key",
			Location: AuthLocationUnknown,
		})
		require.ErrorIs(t, err, snerrors.ErrInvalidAuthLocation)
	})

	t.Run("uses the default header", func(t *testing.T) {
		provider, err := NewAPIKeyProvider(APIKeyConfig{Key: "key"})
		require.NoError(t, err)

		request := testRequest(t, "https://example.com/api/now")
		require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
		assert.Equal(t, []string{"key"}, request.Headers.Get(DefaultAPIKeyParameter))
	})

	t.Run("uses a custom header and prefix", func(t *testing.T) {
		provider, err := NewAPIKeyProvider(APIKeyConfig{
			Key:           "key",
			ParameterName: "x-custom-key",
			Prefix:        "Token ",
		})
		require.NoError(t, err)

		request := testRequest(t, "https://example.com/api/now")
		require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
		assert.Equal(t, []string{"Token key"}, request.Headers.Get("x-custom-key"))
	})

	t.Run("adds a query parameter", func(t *testing.T) {
		provider, err := NewAPIKeyProvider(APIKeyConfig{
			Key:           "key",
			ParameterName: "api_key",
			Location:      AuthLocationQuery,
		})
		require.NoError(t, err)

		request := testRequest(t, "https://example.com/api/now?existing=value")
		require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))

		uri, err := request.GetUri()
		require.NoError(t, err)
		assert.Equal(t, "value", uri.Query().Get("existing"))
		assert.Equal(t, "key", uri.Query().Get("api_key"))
	})

	t.Run("skips a host that is not allowed", func(t *testing.T) {
		provider, err := NewAPIKeyProvider(APIKeyConfig{
			ProviderConfig: ProviderConfig{AllowedHosts: []string{"example.com"}},
			Key:            "key",
		})
		require.NoError(t, err)

		request := testRequest(t, "https://other.example.com/api/now")
		require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
		assert.Empty(t, request.Headers.Get(DefaultAPIKeyParameter))
	})

	t.Run("requires HTTPS", func(t *testing.T) {
		provider, err := NewAPIKeyProvider(APIKeyConfig{Key: "key"})
		require.NoError(t, err)

		request := testRequest(t, "http://example.com/api/now")
		require.ErrorIs(t, provider.AuthenticateRequest(context.Background(), request, nil), snerrors.ErrInsecureURL)
	})

	t.Run("rejects nil inputs", func(t *testing.T) {
		var provider *APIKeyProvider
		require.ErrorIs(t, provider.Initialize("https://example.com"), snerrors.ErrNilProvider)
		require.ErrorIs(t, provider.AuthenticateRequest(context.Background(), testRequest(t, "https://example.com"), nil), snerrors.ErrNilProvider)

		valid, err := NewAPIKeyProvider(APIKeyConfig{Key: "key"})
		require.NoError(t, err)
		//nolint:staticcheck // Verify the explicit nil guard.
		require.ErrorIs(t, valid.AuthenticateRequest(nil, testRequest(t, "https://example.com"), nil), snerrors.ErrNilContext)
		require.ErrorIs(t, valid.AuthenticateRequest(context.Background(), nil, nil), snerrors.ErrNilRequestInfo)
	})
}
