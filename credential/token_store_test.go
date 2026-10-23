// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"errors"
	"testing"
	"time"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemoryTokenStore(t *testing.T) {
	t.Run("saves and loads a copy", func(t *testing.T) {
		store := NewMemoryTokenStore()
		key := TokenKey{BaseURL: "https://example.com", ClientID: "client", Flow: FlowClientCredentials, Scopes: "a b"}
		token := &Token{AccessToken: "access", RefreshToken: "refresh", ExpiresAt: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)}

		require.NoError(t, store.Save(context.Background(), key, token))
		loaded, err := store.Load(context.Background(), key)
		require.NoError(t, err)
		require.NotNil(t, loaded)
		assert.Equal(t, token, loaded)
		assert.NotSame(t, token, loaded)

		loaded.AccessToken = "changed"
		again, err := store.Load(context.Background(), key)
		require.NoError(t, err)
		assert.Equal(t, "access", again.AccessToken)
	})

	t.Run("returns nil for a missing key", func(t *testing.T) {
		store := NewMemoryTokenStore()
		token, err := store.Load(context.Background(), TokenKey{})
		require.NoError(t, err)
		assert.Nil(t, token)
	})

	t.Run("deletes a token", func(t *testing.T) {
		store := NewMemoryTokenStore()
		key := TokenKey{BaseURL: "https://example.com", Flow: FlowClientCredentials}
		require.NoError(t, store.Save(context.Background(), key, &Token{AccessToken: "access"}))
		require.NoError(t, store.Delete(context.Background(), key))

		token, err := store.Load(context.Background(), key)
		require.NoError(t, err)
		assert.Nil(t, token)
	})

	t.Run("deleting a missing key succeeds", func(t *testing.T) {
		store := NewMemoryTokenStore()
		assert.NoError(t, store.Delete(context.Background(), TokenKey{}))
	})

	t.Run("rejects nil contexts", func(t *testing.T) {
		store := NewMemoryTokenStore()
		var nilContext context.Context
		token, err := store.Load(nilContext, TokenKey{})
		assert.Nil(t, token)
		require.ErrorIs(t, err, snerrors.ErrNilContext)

		require.ErrorIs(t, store.Save(nilContext, TokenKey{}, &Token{}), snerrors.ErrNilContext)
		require.ErrorIs(t, store.Delete(nilContext, TokenKey{}), snerrors.ErrNilContext)
	})

	t.Run("rejects a nil token", func(t *testing.T) {
		store := NewMemoryTokenStore()
		require.ErrorIs(t, store.Save(context.Background(), TokenKey{}, nil), snerrors.ErrNilModel)
	})

	t.Run("rejects a canceled context", func(t *testing.T) {
		store := NewMemoryTokenStore()
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		token, err := store.Load(ctx, TokenKey{})
		assert.Nil(t, token)
		require.ErrorIs(t, err, context.Canceled)
		require.ErrorIs(t, store.Save(ctx, TokenKey{}, &Token{}), context.Canceled)
		require.ErrorIs(t, store.Delete(ctx, TokenKey{}), context.Canceled)
	})

	t.Run("nil store", func(t *testing.T) {
		var store *MemoryTokenStore
		token, err := store.Load(context.Background(), TokenKey{})
		assert.Nil(t, token)
		require.ErrorIs(t, err, snerrors.ErrNilStore)
		require.ErrorIs(t, store.Save(context.Background(), TokenKey{}, &Token{}), snerrors.ErrNilStore)
		require.ErrorIs(t, store.Delete(context.Background(), TokenKey{}), snerrors.ErrNilStore)
	})
}

func TestOAuthError(t *testing.T) {
	t.Run("formats useful metadata", func(t *testing.T) {
		err := &OAuthError{
			Operation:   "refresh",
			Flow:        FlowClientCredentials,
			Code:        "invalid_client",
			Description: "secret should not be printed",
			StatusCode:  401,
			Err:         snerrors.ErrTokenAcquisition,
		}

		assert.Equal(t, "oauth2 refresh failed for flow client_credentials: invalid_client (status code: 401)", err.Error())
		assert.NotContains(t, err.Error(), "secret should not be printed")
		require.ErrorIs(t, err, snerrors.ErrTokenAcquisition)
	})

	t.Run("nil receiver", func(t *testing.T) {
		var err *OAuthError
		assert.Equal(t, "nil OAuth error", err.Error())
		assert.NoError(t, err.Unwrap())
	})

	t.Run("unwraps typed causes", func(t *testing.T) {
		cause := errors.New("cause")
		err := &OAuthError{Err: cause}
		require.ErrorIs(t, err, cause)
	})
}
