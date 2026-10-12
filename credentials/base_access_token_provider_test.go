// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"errors"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBaseAccessTokenProvider(t *testing.T) {
	tests := []struct {
		name             string
		allowedHosts     []string
		wantNilValidator bool
	}{
		{
			name:             "no allowed hosts leaves validator nil",
			allowedHosts:     nil,
			wantNilValidator: true,
		},
		{
			name:             "allowed hosts builds validator",
			allowedHosts:     []string{"example.com"},
			wantNilValidator: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newBaseAccessTokenProvider(tt.allowedHosts)
			if tt.wantNilValidator {
				assert.Nil(t, p.GetAllowedHostsValidator())
			} else {
				assert.NotNil(t, p.GetAllowedHostsValidator())
			}
		})
	}
}

func TestBaseAccessTokenProvider_Initialize(t *testing.T) {
	tests := []struct {
		name             string
		allowedHosts     []string
		baseURL          string
		wantNilValidator bool
	}{
		{
			name:             "empty base url keeps validator nil",
			allowedHosts:     nil,
			baseURL:          "",
			wantNilValidator: true,
		},
		{
			name:             "base url derives validator from host when none set",
			allowedHosts:     nil,
			baseURL:          "https://dev12345.service-now.com",
			wantNilValidator: false,
		},
		{
			name:             "existing validator is preserved",
			allowedHosts:     []string{"example.com"},
			baseURL:          "https://dev12345.service-now.com",
			wantNilValidator: false,
		},
		{
			name:             "invalid base url keeps validator nil",
			allowedHosts:     nil,
			baseURL:          "://not-a-url",
			wantNilValidator: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newBaseAccessTokenProvider(tt.allowedHosts)
			p.Initialize(tt.baseURL)

			assert.Equal(t, tt.baseURL, p.baseURL)
			if tt.wantNilValidator {
				assert.Nil(t, p.GetAllowedHostsValidator())
			} else {
				assert.NotNil(t, p.GetAllowedHostsValidator())
			}
		})
	}
}

func TestBaseAccessTokenProvider_SetTokenStore(t *testing.T) {
	p := newBaseAccessTokenProvider(nil)
	store := &fakeTokenStore{}
	p.SetTokenStore(store)
	assert.Equal(t, store, p.tokenStore)
}

func TestBaseAccessTokenProvider_Revoke(t *testing.T) {
	revokeErr := errors.New("revoke failed")

	tests := []struct {
		name        string
		token       *AccessToken
		revokeToken func(ctx context.Context, token, tokenTypeHint string) error
		wantErr     error
		wantNilTok  bool
	}{
		{
			name:       "nil token is a no-op",
			token:      nil,
			wantErr:    nil,
			wantNilTok: true,
		},
		{
			name:    "no revocation function returns error",
			token:   &AccessToken{AccessToken: "access"},
			wantErr: errors.New("token revocation failed: no revocation function available"),
		},
		{
			name:  "refresh token revocation error propagates and token retained",
			token: &AccessToken{AccessToken: "access", RefreshToken: "refresh"},
			revokeToken: func(_ context.Context, _, tokenTypeHint string) error {
				if tokenTypeHint == "refresh_token" {
					return revokeErr
				}
				return nil
			},
			wantErr:    revokeErr,
			wantNilTok: false,
		},
		{
			name:  "access token revocation error propagates",
			token: &AccessToken{AccessToken: "access"},
			revokeToken: func(_ context.Context, _, tokenTypeHint string) error {
				if tokenTypeHint == "access_token" {
					return revokeErr
				}
				return nil
			},
			wantErr:    revokeErr,
			wantNilTok: false,
		},
		{
			name:  "successful revoke clears token",
			token: &AccessToken{AccessToken: "access", RefreshToken: "refresh"},
			revokeToken: func(_ context.Context, _, _ string) error {
				return nil
			},
			wantErr:    nil,
			wantNilTok: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newBaseAccessTokenProvider(nil)
			p.token = tt.token
			p.revokeToken = tt.revokeToken

			err := p.Revoke(context.Background())

			if tt.wantErr != nil {
				require.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}

			if tt.wantNilTok {
				assert.Nil(t, p.token)
			} else {
				assert.NotNil(t, p.token)
			}
		})
	}
}

func TestBaseAccessTokenProvider_GetAuthorizationToken(t *testing.T) {
	acquisitionErr := errors.New("acquire failed")

	tests := []struct {
		name         string
		allowedHosts []string
		uri          *url.URL
		setup        func(p *BaseAccessTokenProvider)
		wantToken    string
		wantErr      string
	}{
		{
			name:         "url host not allowed returns empty token",
			allowedHosts: []string{"allowed.example.com"},
			uri:          mustParseURL(t, "https://not-allowed.example.com/path"),
			setup:        func(_ *BaseAccessTokenProvider) {},
			wantToken:    "",
		},
		{
			name: "valid cached token is returned without refresh",
			uri:  nil,
			setup: func(p *BaseAccessTokenProvider) {
				p.token = &AccessToken{AccessToken: "cached", ExpiresAt: time.Now().Add(time.Hour)}
			},
			wantToken: "cached",
		},
		{
			name: "expired token with successful refresh returns new token",
			uri:  nil,
			setup: func(p *BaseAccessTokenProvider) {
				p.token = &AccessToken{AccessToken: "old", RefreshToken: "refresh", ExpiresAt: time.Now().Add(-time.Minute)}
				p.refreshToken = func(_ context.Context, refreshToken string) (*AccessToken, error) {
					return &AccessToken{AccessToken: "refreshed-" + refreshToken, ExpiresAt: time.Now().Add(time.Hour)}, nil
				}
			},
			wantToken: "refreshed-refresh",
		},
		{
			name: "expired token with failing refresh falls back to initial retrieval",
			uri:  nil,
			setup: func(p *BaseAccessTokenProvider) {
				p.token = &AccessToken{AccessToken: "old", RefreshToken: "refresh", ExpiresAt: time.Now().Add(-time.Minute)}
				p.refreshToken = func(_ context.Context, _ string) (*AccessToken, error) {
					return nil, errors.New("refresh failed")
				}
				p.retrieveInitialToken = func(_ context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
					return &AccessToken{AccessToken: "fresh", ExpiresAt: time.Now().Add(time.Hour)}, nil
				}
			},
			wantToken: "fresh",
		},
		{
			name:    "no token and no retrieval function errors",
			uri:     nil,
			setup:   func(_ *BaseAccessTokenProvider) {},
			wantErr: "token acquisition failed: no initial retrieval function available",
		},
		{
			name: "retrieval function error is wrapped",
			uri:  nil,
			setup: func(p *BaseAccessTokenProvider) {
				p.retrieveInitialToken = func(_ context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
					return nil, acquisitionErr
				}
			},
			wantErr: "initial token acquisition failed: acquire failed",
		},
		{
			name: "successful initial retrieval sets token",
			uri:  nil,
			setup: func(p *BaseAccessTokenProvider) {
				p.retrieveInitialToken = func(_ context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
					return &AccessToken{AccessToken: "brand-new", ExpiresAt: time.Now().Add(time.Hour)}, nil
				}
			},
			wantToken: "brand-new",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := newBaseAccessTokenProvider(tt.allowedHosts)
			tt.setup(p)

			token, err := p.GetAuthorizationToken(context.Background(), tt.uri, nil)

			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantToken, token)
		})
	}
}

func mustParseURL(t *testing.T, raw string) *url.URL {
	t.Helper()
	u, err := url.Parse(raw)
	require.NoError(t, err)
	return u
}

type fakeTokenStore struct{}

func (f *fakeTokenStore) Save(_ context.Context, _ string, _ *AccessToken) error { return nil }
func (f *fakeTokenStore) Load(_ context.Context, _ string) (*AccessToken, error) { return nil, nil }
func (f *fakeTokenStore) Delete(_ context.Context, _ string) error               { return nil }

// memoryTokenStore is an in-memory TokenStore used to exercise persistence.
type memoryTokenStore struct {
	mu     sync.Mutex
	tokens map[string]*AccessToken
}

func newMemoryTokenStore() *memoryTokenStore {
	return &memoryTokenStore{tokens: map[string]*AccessToken{}}
}

func (m *memoryTokenStore) Save(_ context.Context, clientID string, token *AccessToken) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tokens[clientID] = token
	return nil
}

func (m *memoryTokenStore) Load(_ context.Context, clientID string) (*AccessToken, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.tokens[clientID], nil
}

func (m *memoryTokenStore) Delete(_ context.Context, clientID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.tokens, clientID)
	return nil
}

func TestBaseAccessTokenProvider_Persistence(t *testing.T) {
	t.Run("valid persisted token is loaded without retrieval", func(t *testing.T) {
		store := newMemoryTokenStore()
		store.tokens["client-id"] = &AccessToken{AccessToken: "persisted", ExpiresAt: time.Now().Add(time.Hour)}

		p := newBaseAccessTokenProvider(nil)
		p.setClientID("client-id")
		p.SetTokenStore(store)
		p.retrieveInitialToken = func(_ context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
			return nil, errors.New("retrieval should not be called")
		}

		token, err := p.GetAuthorizationToken(context.Background(), nil, nil)
		require.NoError(t, err)
		assert.Equal(t, "persisted", token)
	})

	t.Run("expired persisted token is refreshed and re-saved", func(t *testing.T) {
		store := newMemoryTokenStore()
		store.tokens["client-id"] = &AccessToken{AccessToken: "old", RefreshToken: "refresh", ExpiresAt: time.Now().Add(-time.Minute)}

		p := newBaseAccessTokenProvider(nil)
		p.setClientID("client-id")
		p.SetTokenStore(store)
		p.refreshToken = func(_ context.Context, _ string) (*AccessToken, error) {
			return &AccessToken{AccessToken: "refreshed", RefreshToken: "refresh-2", ExpiresAt: time.Now().Add(time.Hour)}, nil
		}

		token, err := p.GetAuthorizationToken(context.Background(), nil, nil)
		require.NoError(t, err)
		assert.Equal(t, "refreshed", token)
		assert.Equal(t, "refreshed", store.tokens["client-id"].AccessToken)
	})

	t.Run("initial retrieval is persisted", func(t *testing.T) {
		store := newMemoryTokenStore()

		p := newBaseAccessTokenProvider(nil)
		p.setClientID("client-id")
		p.SetTokenStore(store)
		p.retrieveInitialToken = func(_ context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
			return &AccessToken{AccessToken: "new", ExpiresAt: time.Now().Add(time.Hour)}, nil
		}

		token, err := p.GetAuthorizationToken(context.Background(), nil, nil)
		require.NoError(t, err)
		assert.Equal(t, "new", token)
		assert.Equal(t, "new", store.tokens["client-id"].AccessToken)
	})

	t.Run("revoke deletes the persisted token", func(t *testing.T) {
		store := newMemoryTokenStore()
		store.tokens["client-id"] = &AccessToken{AccessToken: "token", ExpiresAt: time.Now().Add(time.Hour)}

		p := newBaseAccessTokenProvider(nil)
		p.setClientID("client-id")
		p.SetTokenStore(store)
		p.token = &AccessToken{AccessToken: "token"}
		p.revokeToken = func(_ context.Context, _, _ string) error { return nil }

		require.NoError(t, p.Revoke(context.Background()))
		assert.NotContains(t, store.tokens, "client-id")
	})

	t.Run("no client id skips persistence", func(t *testing.T) {
		store := newMemoryTokenStore()

		p := newBaseAccessTokenProvider(nil)
		p.SetTokenStore(store)
		p.retrieveInitialToken = func(_ context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
			return &AccessToken{AccessToken: "new", ExpiresAt: time.Now().Add(time.Hour)}, nil
		}

		token, err := p.GetAuthorizationToken(context.Background(), nil, nil)
		require.NoError(t, err)
		assert.Equal(t, "new", token)
		assert.Empty(t, store.tokens)
	})
}
