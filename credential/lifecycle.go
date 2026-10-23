// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

const defaultExpirySkew = 30 * time.Second

type tokenFlight struct {
	done  chan struct{}
	token *Token
	err   error
}

type tokenLifecycle struct {
	flow       Flow
	clientID   string
	scopes     []string
	sessionKey string
	store      TokenStore
	validator  *authentication.AllowedHostsValidator

	acquire                 func(ctx context.Context, baseURL string) (*Token, error)
	refresh                 func(ctx context.Context, baseURL, refreshToken string) (*Token, error)
	revokeHook              func(ctx context.Context, baseURL string, token *Token) error
	reacquireOnInvalidGrant bool

	now  func() time.Time
	skew time.Duration

	mu          sync.Mutex
	baseURL     string
	initialized bool
	token       *Token
	flight      *tokenFlight
}

//nolint:unparam // The shared lifecycle serves flows added in later phases.
func newTokenLifecycle(
	flow Flow,
	clientID string,
	scopes []string,
	sessionKey string,
	store TokenStore,
	validator *authentication.AllowedHostsValidator,
) *tokenLifecycle {
	if isNilTokenStore(store) {
		store = NewMemoryTokenStore()
	}
	return &tokenLifecycle{
		flow:       flow,
		clientID:   clientID,
		scopes:     normalizeScopes(scopes),
		sessionKey: sessionKey,
		store:      store,
		validator:  validator,
		now:        time.Now,
		skew:       defaultExpirySkew,
	}
}

func (l *tokenLifecycle) initialize(baseURL string) error {
	normalized, err := normalizeBaseURL(baseURL)
	if err != nil {
		return err
	}
	if l == nil {
		return snerrors.ErrNilProvider
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.initialized {
		if l.baseURL == normalized {
			return nil
		}
		return snerrors.ErrAlreadyInitialized
	}

	l.baseURL = normalized
	l.initialized = true
	return nil
}

func (l *tokenLifecycle) authorizationToken(ctx context.Context, uri *url.URL) (string, error) {
	if l == nil {
		return "", snerrors.ErrNilProvider
	}
	if ctx == nil {
		return "", snerrors.ErrNilContext
	}
	if uri == nil {
		// TODO: why?
		return "", fmt.Errorf("%w: request URI is empty", snerrors.ErrInvalidURL)
	}
	if l.validator != nil && !l.validator.IsUrlHostValid(uri) {
		return "", nil
	}
	if !strings.EqualFold(uri.Scheme, "https") {
		return "", snerrors.ErrInsecureURL
	}

	token, err := l.tokenFor(ctx)
	if err != nil {
		return "", err
	}
	if token == nil || token.AccessToken == "" {
		return "", snerrors.ErrMissingAccessToken
	}
	return token.AccessToken, nil
}

func (l *tokenLifecycle) tokenFor(ctx context.Context) (*Token, error) {
	if ctx == nil {
		return nil, snerrors.ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	l.mu.Lock()
	if !l.initialized {
		l.mu.Unlock()
		return nil, snerrors.ErrNotInitialized
	}

	if l.token == nil {
		key := l.keyLocked()
		stored, err := l.store.Load(ctx, key)
		if err == nil && stored != nil {
			l.token = stored
		}
	}

	if l.token != nil && !l.expired(l.token, l.now(), l.skew) {
		token := cloneToken(l.token)
		l.mu.Unlock()
		return token, nil
	}

	if l.flight != nil {
		flight := l.flight
		l.mu.Unlock()
		return waitForToken(ctx, flight)
	}

	cached := cloneToken(l.token)
	baseURL := l.baseURL
	flight := &tokenFlight{done: make(chan struct{})}
	l.flight = flight
	l.mu.Unlock()

	// The shared attempt outlives the caller that started it. A caller that
	// cancels stops waiting, but other callers still receive the result.
	go func() {
		token, err := l.fetch(context.WithoutCancel(ctx), baseURL, cached)
		l.complete(flight, baseURL, token, err)
	}()

	return waitForToken(ctx, flight)
}

func waitForToken(ctx context.Context, flight *tokenFlight) (*Token, error) {
	select {
	case <-flight.done:
		return cloneToken(flight.token), flight.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (l *tokenLifecycle) complete(flight *tokenFlight, baseURL string, token *Token, err error) {
	l.mu.Lock()
	if err == nil && token != nil {
		l.token = token
		_ = l.store.Save(context.Background(), l.keyFor(baseURL), token)
	}
	l.flight = nil
	flight.token = token
	flight.err = err
	l.mu.Unlock()

	close(flight.done)
}

func (l *tokenLifecycle) fetch(ctx context.Context, baseURL string, cached *Token) (*Token, error) {
	if cached != nil && cached.RefreshToken != "" && l.refresh != nil {
		refreshed, err := l.refresh(ctx, baseURL, cached.RefreshToken)
		if err == nil {
			if refreshed == nil || refreshed.AccessToken == "" {
				return nil, &OAuthError{
					Operation: "refresh",
					Flow:      l.flow,
					Err:       snerrors.ErrMissingAccessToken,
				}
			}
			return mergeRefreshedToken(cached, refreshed), nil
		}
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		if isInvalidGrant(err) {
			l.clear(ctx, baseURL)
			if !l.reacquireOnInvalidGrant {
				return nil, wrapOAuthError("refresh", l.flow, err, snerrors.ErrReauthRequired)
			}
			acquired, acquireErr := l.acquireToken(ctx, baseURL)
			if acquireErr != nil {
				return nil, acquireErr
			}
			return acquired, nil
		}
		if isTransient(err) && l.validNow(cached, l.now()) {
			return cloneToken(cached), nil
		}
		return nil, wrapOAuthError("refresh", l.flow, err, snerrors.ErrTokenAcquisition)
	}

	return l.acquireToken(ctx, baseURL)
}

func (l *tokenLifecycle) acquireToken(ctx context.Context, baseURL string) (*Token, error) {
	if l.acquire == nil {
		return nil, &OAuthError{
			Operation: "acquire",
			Flow:      l.flow,
			Err:       snerrors.ErrTokenAcquisition,
		}
	}

	token, err := l.acquire(ctx, baseURL)
	if err != nil {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		return nil, wrapOAuthError("acquire", l.flow, err, snerrors.ErrTokenAcquisition)
	}
	if token == nil || token.AccessToken == "" {
		return nil, &OAuthError{
			Operation: "acquire",
			Flow:      l.flow,
			Err:       snerrors.ErrMissingAccessToken,
		}
	}
	return token, nil
}

func (l *tokenLifecycle) revoke(ctx context.Context) error {
	if l == nil {
		return snerrors.ErrNilProvider
	}
	if ctx == nil {
		return snerrors.ErrNilContext
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	l.mu.Lock()
	if !l.initialized {
		l.mu.Unlock()
		return snerrors.ErrNotInitialized
	}
	token := cloneToken(l.token)
	baseURL := l.baseURL
	key := l.keyLocked()
	l.mu.Unlock()

	if token == nil {
		return l.store.Delete(ctx, key)
	}
	if l.revokeHook != nil {
		if err := l.revokeHook(ctx, baseURL, token); err != nil {
			return wrapOAuthError("revoke", l.flow, err, snerrors.ErrTokenAcquisition)
		}
	}

	l.mu.Lock()
	l.token = nil
	l.mu.Unlock()
	return l.store.Delete(ctx, key)
}

func (l *tokenLifecycle) clear(ctx context.Context, baseURL string) {
	l.mu.Lock()
	l.token = nil
	l.mu.Unlock()
	_ = l.store.Delete(ctx, l.keyFor(baseURL))
}

func (l *tokenLifecycle) expired(token *Token, now time.Time, skew time.Duration) bool {
	if token == nil || token.ExpiresAt.IsZero() {
		return false
	}
	return !now.Add(skew).Before(token.ExpiresAt)
}

func (l *tokenLifecycle) validNow(token *Token, now time.Time) bool {
	if token == nil || token.AccessToken == "" {
		return false
	}
	return token.ExpiresAt.IsZero() || now.Before(token.ExpiresAt)
}

func (l *tokenLifecycle) keyLocked() TokenKey {
	return l.keyFor(l.baseURL)
}

func (l *tokenLifecycle) keyFor(baseURL string) TokenKey {
	return TokenKey{
		BaseURL:    baseURL,
		ClientID:   l.clientID,
		Flow:       l.flow,
		Scopes:     strings.Join(l.scopes, " "),
		SessionKey: l.sessionKey,
	}
}

func cloneToken(token *Token) *Token {
	if token == nil {
		return nil
	}
	cloned := *token
	return &cloned
}

func mergeRefreshedToken(previous, next *Token) *Token {
	merged := cloneToken(next)
	if merged == nil {
		return nil
	}
	if merged.RefreshToken == "" {
		merged.RefreshToken = previous.RefreshToken
	}
	if merged.IDToken == "" {
		merged.IDToken = previous.IDToken
	}
	if merged.TokenType == "" {
		merged.TokenType = previous.TokenType
	}
	if merged.Scope == "" {
		merged.Scope = previous.Scope
	}
	return merged
}

func isInvalidGrant(err error) bool {
	var tokenError *oauth2.TokenError
	if errors.As(err, &tokenError) {
		return tokenError.Err == "invalid_grant"
	}
	return false
}

func isTransient(err error) bool {
	var tokenError *oauth2.TokenError
	if errors.As(err, &tokenError) {
		return tokenError.StatusCode == http.StatusTooManyRequests || tokenError.StatusCode >= http.StatusInternalServerError
	}
	return true
}

func wrapOAuthError(operation string, flow Flow, err error, fallback error) error {
	if err == nil {
		return nil
	}

	var existing *OAuthError
	if errors.As(err, &existing) {
		return existing
	}

	var tokenError *oauth2.TokenError
	if errors.As(err, &tokenError) {
		return &OAuthError{
			Operation:   operation,
			Flow:        flow,
			Code:        tokenError.Err,
			Description: tokenError.ErrorDescription,
			URI:         tokenError.ErrorURI,
			StatusCode:  tokenError.StatusCode,
			Err:         fallback,
		}
	}

	return &OAuthError{
		Operation: operation,
		Flow:      flow,
		Err:       fmt.Errorf("%w: %w", fallback, err),
	}
}
