// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"errors"
	"net/url"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeClock struct {
	mu  sync.Mutex
	now time.Time
}

func newFakeClock(now time.Time) *fakeClock {
	return &fakeClock{now: now}
}

func (c *fakeClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *fakeClock) Advance(duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.now = c.now.Add(duration)
}

func newTestLifecycle(t *testing.T, clock *fakeClock, scopes ...string) *tokenLifecycle {
	t.Helper()
	lifecycle := newTokenLifecycle(FlowClientCredentials, "client", scopes, "", NewMemoryTokenStore(), nil)
	require.NoError(t, lifecycle.initialize("https://example.com"))
	lifecycle.now = clock.Now
	return lifecycle
}

func TestTokenLifecycleCachesAndExpires(t *testing.T) {
	ctx := context.Background()
	clock := newFakeClock(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	lifecycle := newTestLifecycle(t, clock)

	var calls atomic.Int32
	lifecycle.acquire = func(context.Context, string) (*Token, error) {
		calls.Add(1)
		return &Token{AccessToken: "first", ExpiresAt: clock.Now().Add(time.Minute)}, nil
	}

	token, err := lifecycle.tokenFor(ctx)
	require.NoError(t, err)
	assert.Equal(t, "first", token.AccessToken)

	token, err = lifecycle.tokenFor(ctx)
	require.NoError(t, err)
	assert.Equal(t, "first", token.AccessToken)
	assert.Equal(t, int32(1), calls.Load())

	clock.Advance(31 * time.Second)
	token, err = lifecycle.tokenFor(ctx)
	require.NoError(t, err)
	assert.Equal(t, "first", token.AccessToken)
	assert.Equal(t, int32(2), calls.Load())
}

func TestTokenLifecycleRefreshes(t *testing.T) {
	ctx := context.Background()
	clock := newFakeClock(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	lifecycle := newTestLifecycle(t, clock)

	var refreshCalls atomic.Int32
	lifecycle.acquire = func(context.Context, string) (*Token, error) {
		return &Token{
			AccessToken:  "old",
			RefreshToken: "refresh",
			TokenType:    "Bearer",
			Scope:        "scope",
			ExpiresAt:    clock.Now().Add(time.Minute),
		}, nil
	}
	lifecycle.refresh = func(_ context.Context, _ string, refreshToken string) (*Token, error) {
		refreshCalls.Add(1)
		assert.Equal(t, "refresh", refreshToken)
		return &Token{AccessToken: "new", ExpiresAt: clock.Now().Add(time.Hour)}, nil
	}

	_, err := lifecycle.tokenFor(ctx)
	require.NoError(t, err)

	clock.Advance(time.Minute)
	token, err := lifecycle.tokenFor(ctx)
	require.NoError(t, err)
	assert.Equal(t, "new", token.AccessToken)
	assert.Equal(t, "refresh", token.RefreshToken)
	assert.Equal(t, "Bearer", token.TokenType)
	assert.Equal(t, "scope", token.Scope)
	assert.Equal(t, int32(1), refreshCalls.Load())
}

func TestTokenLifecycleReacquiresAfterInvalidGrant(t *testing.T) {
	ctx := context.Background()
	clock := newFakeClock(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	lifecycle := newTestLifecycle(t, clock)
	lifecycle.reacquireOnInvalidGrant = true

	var acquireCalls atomic.Int32
	lifecycle.acquire = func(context.Context, string) (*Token, error) {
		if acquireCalls.Add(1) == 1 {
			return &Token{AccessToken: "old", RefreshToken: "refresh", ExpiresAt: clock.Now().Add(time.Minute)}, nil
		}
		return &Token{AccessToken: "new", ExpiresAt: clock.Now().Add(time.Hour)}, nil
	}
	lifecycle.refresh = func(context.Context, string, string) (*Token, error) {
		return nil, &oauth2.TokenError{Err: "invalid_grant", StatusCode: 400}
	}

	_, err := lifecycle.tokenFor(ctx)
	require.NoError(t, err)

	clock.Advance(time.Minute)
	token, err := lifecycle.tokenFor(ctx)
	require.NoError(t, err)
	assert.Equal(t, "new", token.AccessToken)
	assert.Equal(t, int32(2), acquireCalls.Load())
}

func TestTokenLifecycleRequiresReauthentication(t *testing.T) {
	ctx := context.Background()
	clock := newFakeClock(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	lifecycle := newTestLifecycle(t, clock)

	lifecycle.acquire = func(context.Context, string) (*Token, error) {
		return &Token{AccessToken: "old", RefreshToken: "refresh", ExpiresAt: clock.Now().Add(time.Minute)}, nil
	}
	lifecycle.refresh = func(context.Context, string, string) (*Token, error) {
		return nil, &oauth2.TokenError{Err: "invalid_grant", StatusCode: 400}
	}

	_, err := lifecycle.tokenFor(ctx)
	require.NoError(t, err)

	clock.Advance(time.Minute)
	token, err := lifecycle.tokenFor(ctx)
	assert.Nil(t, token)
	require.ErrorIs(t, err, snerrors.ErrReauthRequired)

	var oauthError *OAuthError
	require.ErrorAs(t, err, &oauthError)
	assert.Equal(t, "refresh", oauthError.Operation)
	assert.Equal(t, "invalid_grant", oauthError.Code)
	assert.Nil(t, lifecycle.token)
}

func TestTokenLifecycleKeepsAValidTokenOnTransientRefreshFailure(t *testing.T) {
	ctx := context.Background()
	clock := newFakeClock(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	lifecycle := newTestLifecycle(t, clock)

	lifecycle.acquire = func(context.Context, string) (*Token, error) {
		return &Token{AccessToken: "old", RefreshToken: "refresh", ExpiresAt: clock.Now().Add(40 * time.Second)}, nil
	}
	lifecycle.refresh = func(context.Context, string, string) (*Token, error) {
		return nil, &oauth2.TokenError{Err: "server_error", StatusCode: 500}
	}

	_, err := lifecycle.tokenFor(ctx)
	require.NoError(t, err)

	clock.Advance(15 * time.Second)
	token, err := lifecycle.tokenFor(ctx)
	require.NoError(t, err)
	assert.Equal(t, "old", token.AccessToken)

	clock.Advance(26 * time.Second)
	token, err = lifecycle.tokenFor(ctx)
	assert.Nil(t, token)
	require.ErrorIs(t, err, snerrors.ErrTokenAcquisition)

	var oauthError *OAuthError
	require.ErrorAs(t, err, &oauthError)
	assert.Equal(t, "refresh", oauthError.Operation)
	assert.Equal(t, "server_error", oauthError.Code)
}

func TestTokenLifecycleSingleFlight(t *testing.T) {
	clock := newFakeClock(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	lifecycle := newTestLifecycle(t, clock)

	started := make(chan struct{})
	release := make(chan struct{})
	var calls atomic.Int32
	lifecycle.acquire = func(context.Context, string) (*Token, error) {
		calls.Add(1)
		close(started)
		<-release
		return &Token{AccessToken: "token", ExpiresAt: clock.Now().Add(time.Hour)}, nil
	}

	const goroutines = 8
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for range goroutines {
		go func() {
			defer wg.Done()
			token, err := lifecycle.tokenFor(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, "token", token.AccessToken)
		}()
	}

	<-started
	close(release)
	wg.Wait()
	assert.Equal(t, int32(1), calls.Load())
}

func TestTokenLifecycleLoadsFromStore(t *testing.T) {
	ctx := context.Background()
	clock := newFakeClock(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	lifecycle := newTestLifecycle(t, clock, "z", "a", "a")

	stored := &Token{AccessToken: "stored", ExpiresAt: clock.Now().Add(time.Hour)}
	require.NoError(t, lifecycle.store.Save(ctx, lifecycle.keyFor("https://example.com"), stored))

	calls := 0
	lifecycle.acquire = func(context.Context, string) (*Token, error) {
		calls++
		return nil, errors.New("should not acquire")
	}

	token, err := lifecycle.tokenFor(ctx)
	require.NoError(t, err)
	assert.Equal(t, "stored", token.AccessToken)
	assert.Equal(t, 0, calls)

	key := lifecycle.keyFor("https://example.com")
	assert.Equal(t, "a z", key.Scopes)
}

func TestTokenLifecycleAuthorizationToken(t *testing.T) {
	ctx := context.Background()
	clock := newFakeClock(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	lifecycle := newTokenLifecycle(FlowClientCredentials, "client", nil, "", nil, nil)
	require.NoError(t, lifecycle.initialize("https://example.com"))
	lifecycle.now = clock.Now
	lifecycle.acquire = func(context.Context, string) (*Token, error) {
		return &Token{AccessToken: "token", ExpiresAt: clock.Now().Add(time.Hour)}, nil
	}

	token, err := lifecycle.authorizationToken(ctx, &url.URL{Scheme: "https", Host: "example.com"})
	require.NoError(t, err)
	assert.Equal(t, "token", token)

	token, err = lifecycle.authorizationToken(ctx, &url.URL{Scheme: "http", Host: "example.com"})
	assert.Empty(t, token)
	require.ErrorIs(t, err, snerrors.ErrInsecureURL)

	token, err = lifecycle.authorizationToken(ctx, nil)
	assert.Empty(t, token)
	require.ErrorIs(t, err, snerrors.ErrInvalidURL)

	validator, err := authentication.NewAllowedHostsValidatorErrorCheck([]string{"example.com"})
	require.NoError(t, err)
	lifecycle.validator = validator

	token, err = lifecycle.authorizationToken(ctx, &url.URL{Scheme: "https", Host: "other.example.com"})
	require.NoError(t, err)
	assert.Empty(t, token)
}

func TestTokenLifecycleRequiresInitialization(t *testing.T) {
	lifecycle := newTokenLifecycle(FlowClientCredentials, "client", nil, "", nil, nil)

	token, err := lifecycle.tokenFor(context.Background())
	assert.Nil(t, token)
	require.ErrorIs(t, err, snerrors.ErrNotInitialized)

	_, err = lifecycle.authorizationToken(context.Background(), &url.URL{Scheme: "https", Host: "example.com"})
	require.ErrorIs(t, err, snerrors.ErrNotInitialized)
}

func TestTokenLifecycleInitialize(t *testing.T) {
	lifecycle := newTokenLifecycle(FlowClientCredentials, "client", nil, "", nil, nil)

	require.ErrorIs(t, lifecycle.initialize(""), snerrors.ErrEmptyURI)
	require.ErrorIs(t, lifecycle.initialize("http://example.com"), snerrors.ErrInsecureURL)
	require.ErrorIs(t, lifecycle.initialize("://bad"), snerrors.ErrInvalidURL)
	require.NoError(t, lifecycle.initialize("https://example.com/"))
	require.NoError(t, lifecycle.initialize("https://EXAMPLE.com"))
	require.ErrorIs(t, lifecycle.initialize("https://other.example.com"), snerrors.ErrAlreadyInitialized)

	var nilLifecycle *tokenLifecycle
	require.ErrorIs(t, nilLifecycle.initialize("https://example.com"), snerrors.ErrNilProvider)
}

func TestTokenLifecycleRevoke(t *testing.T) {
	ctx := context.Background()
	clock := newFakeClock(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))

	t.Run("revokes and clears", func(t *testing.T) {
		lifecycle := newTestLifecycle(t, clock)
		lifecycle.acquire = func(context.Context, string) (*Token, error) {
			return &Token{AccessToken: "token", RefreshToken: "refresh", ExpiresAt: clock.Now().Add(time.Hour)}, nil
		}

		var revoked []string
		lifecycle.revokeHook = func(_ context.Context, _ string, token *Token) error {
			revoked = append(revoked, token.RefreshToken, token.AccessToken)
			return nil
		}

		_, err := lifecycle.tokenFor(ctx)
		require.NoError(t, err)
		require.NoError(t, lifecycle.revoke(ctx))
		assert.Equal(t, []string{"refresh", "token"}, revoked)
		assert.Nil(t, lifecycle.token)

		stored, err := lifecycle.store.Load(ctx, lifecycle.keyFor("https://example.com"))
		require.NoError(t, err)
		assert.Nil(t, stored)
	})

	t.Run("keeps the token when revocation fails", func(t *testing.T) {
		lifecycle := newTestLifecycle(t, clock)
		lifecycle.acquire = func(context.Context, string) (*Token, error) {
			return &Token{AccessToken: "token", ExpiresAt: clock.Now().Add(time.Hour)}, nil
		}
		lifecycle.revokeHook = func(context.Context, string, *Token) error {
			return errors.New("revocation failed")
		}

		_, err := lifecycle.tokenFor(ctx)
		require.NoError(t, err)
		err = lifecycle.revoke(ctx)
		require.ErrorIs(t, err, snerrors.ErrTokenAcquisition)
		assert.NotNil(t, lifecycle.token)
	})

	t.Run("no token", func(t *testing.T) {
		lifecycle := newTestLifecycle(t, clock)
		require.NoError(t, lifecycle.revoke(ctx))
	})

	t.Run("nil and uninitialized", func(t *testing.T) {
		var nilLifecycle *tokenLifecycle
		require.ErrorIs(t, nilLifecycle.revoke(ctx), snerrors.ErrNilProvider)

		lifecycle := newTokenLifecycle(FlowClientCredentials, "client", nil, "", nil, nil)
		require.ErrorIs(t, lifecycle.revoke(ctx), snerrors.ErrNotInitialized)
		var nilContext context.Context
		require.ErrorIs(t, lifecycle.revoke(nilContext), snerrors.ErrNilContext)
	})
}
