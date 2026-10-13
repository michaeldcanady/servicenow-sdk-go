// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"

	credentials "github.com/michaeldcanady/servicenow-sdk-go/v2/credentials"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockTokenEndpoint is a fake token endpoint: it serves canned responses and
// counts request hits, with an optional delay to widen race windows.
type mockTokenEndpoint struct {
	mu     sync.Mutex
	hits   int
	status int
	body   []byte
	delay  time.Duration
}

func (m *mockTokenEndpoint) RoundTrip(req *http.Request) (*http.Response, error) {
	m.mu.Lock()
	m.hits++
	m.mu.Unlock()

	if m.delay > 0 {
		time.Sleep(m.delay)
	}

	status := m.status
	if status == 0 {
		status = http.StatusOK
	}

	return &http.Response{
		StatusCode: status,
		Status:     http.StatusText(status),
		Body:       io.NopCloser(bytes.NewReader(m.body)),
		Header:     make(http.Header),
		Request:    req,
	}, nil
}

func (m *mockTokenEndpoint) count() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.hits
}

func newMockTokenEndpoint() *mockTokenEndpoint {
	return &mockTokenEndpoint{
		status: http.StatusOK,
		body:   []byte(`{"access_token":"mock-token","token_type":"Bearer","expires_in":3600}`),
	}
}

const (
	testInstanceURL  = "https://dev12345.service-now.com"
	testClientID     = "test-client"
	testClientSecret = "test-secret"
)

// lifecycleStore is a recording fake TokenStore implementing the exported
// contract. A preset token is served for a key; absent keys load as (nil, nil).
type lifecycleStore struct {
	mu      sync.Mutex
	tokens  map[credentials.TokenKey]*credentials.AccessToken
	saved   []credentials.TokenKey
	deleted []credentials.TokenKey
	loadErr error
	saveErr error
}

func (s *lifecycleStore) Save(_ context.Context, key credentials.TokenKey, token *credentials.AccessToken) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.saveErr != nil {
		return s.saveErr
	}
	if s.tokens != nil {
		s.tokens[key] = token
	}
	s.saved = append(s.saved, key)
	return nil
}

func (s *lifecycleStore) Load(_ context.Context, key credentials.TokenKey) (*credentials.AccessToken, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.loadErr != nil {
		return nil, s.loadErr
	}
	return s.tokens[key], nil
}

func (s *lifecycleStore) Delete(_ context.Context, key credentials.TokenKey) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.tokens != nil {
		delete(s.tokens, key)
	}
	s.deleted = append(s.deleted, key)
	return nil
}

func (s *lifecycleStore) saveCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.saved)
}

func newLifecycleStore() *lifecycleStore {
	return &lifecycleStore{tokens: make(map[credentials.TokenKey]*credentials.AccessToken)}
}

func mustURL(t *testing.T, raw string) *url.URL {
	t.Helper()
	u, err := url.Parse(raw)
	require.NoError(t, err)
	return u
}

func storedValidToken() *credentials.AccessToken {
	return &credentials.AccessToken{
		AccessToken: "stored-token",
		ExpiresAt:   time.Now().Add(time.Hour),
	}
}

func buildProvider(t *testing.T, endpoint *mockTokenEndpoint, store credentials.TokenStore, opts credentials.ClientCredentialsOptions) authentication.AuthenticationProvider {
	t.Helper()
	opts.HTTPClient = &http.Client{Transport: endpoint}
	opts.TokenStore = store
	provider, err := credentials.NewClientCredentialsProvider(testClientID, testClientSecret, opts)
	require.NoError(t, err)
	// The provider now expects the client to call Initialize with the base URL.
	// In production the client does this; in tests we call it directly.
	if prep, ok := provider.(interface{ Initialize(string) }); ok {
		prep.Initialize(testInstanceURL)
	}
	return provider
}

func authenticate(t *testing.T, provider authentication.AuthenticationProvider, uri *url.URL) (string, error) {
	t.Helper()
	request := abstractions.NewRequestInformation()
	if uri != nil {
		request.SetUri(*uri)
	}
	err := provider.AuthenticateRequest(context.Background(), request, nil)
	if err != nil {
		return "", err
	}
	header := request.Headers.Get("Authorization")
	if len(header) == 0 {
		return "", nil
	}
	return header[0], nil
}

func TestTokenLifecycle_ValidStoredTokenServedWithoutRequest(t *testing.T) {
	store := newLifecycleStore()
	store.tokens[credentials.TokenKey{Authority: testInstanceURL, ClientID: testClientID}] = storedValidToken()
	endpoint := newMockTokenEndpoint()

	provider := buildProvider(t, endpoint, store, credentials.ClientCredentialsOptions{
		CommonOptions: credentials.CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
		},
	})

	header, err := authenticate(t, provider, mustURL(t, testInstanceURL+"/now/table/incident"))
	require.NoError(t, err)
	assert.Equal(t, "Bearer stored-token", header)
	assert.Zero(t, endpoint.count(), "a valid stored token must be served without hitting the token endpoint")
}

func TestTokenLifecycle_AbsentKeyLoadsNilAndAcquires(t *testing.T) {
	store := newLifecycleStore() // no token under the key
	endpoint := newMockTokenEndpoint()

	provider := buildProvider(t, endpoint, store, credentials.ClientCredentialsOptions{
		CommonOptions: credentials.CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
		},
	})

	header, err := authenticate(t, provider, mustURL(t, testInstanceURL+"/now/table/incident"))
	require.NoError(t, err)
	assert.Equal(t, "Bearer mock-token", header)
	assert.Equal(t, 1, endpoint.count(), "an absent key must proceed to a normal acquire")
}

func TestTokenLifecycle_AcquireSavesToStore(t *testing.T) {
	store := newLifecycleStore()
	endpoint := newMockTokenEndpoint()
	key := credentials.TokenKey{Authority: testInstanceURL, ClientID: testClientID}

	provider := buildProvider(t, endpoint, store, credentials.ClientCredentialsOptions{
		CommonOptions: credentials.CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
		},
	})

	header, err := authenticate(t, provider, mustURL(t, testInstanceURL+"/now/table/incident"))
	require.NoError(t, err)
	assert.Equal(t, "Bearer mock-token", header)

	require.Equal(t, 1, store.saveCount(), "a successful acquire must persist the token")
	stored, err := store.Load(context.Background(), key)
	require.NoError(t, err)
	require.NotNil(t, stored)
	assert.Equal(t, "mock-token", stored.AccessToken)
}

func TestTokenLifecycle_ExpiredStoredTokenRefreshesAndSaves(t *testing.T) {
	key := credentials.TokenKey{Authority: testInstanceURL, ClientID: testClientID}
	store := newLifecycleStore()
	store.tokens[key] = &credentials.AccessToken{
		AccessToken:  "stale-token",
		RefreshToken: "refresh-me",
		ExpiresAt:    time.Now().Add(-time.Minute),
	}
	endpoint := newMockTokenEndpoint()
	endpoint.body = []byte(`{"access_token":"refreshed-token","refresh_token":"new-refresh","token_type":"Bearer","expires_in":3600}`)

	provider := buildProvider(t, endpoint, store, credentials.ClientCredentialsOptions{
		CommonOptions: credentials.CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
		},
	})

	header, err := authenticate(t, provider, mustURL(t, testInstanceURL+"/now/table/incident"))
	require.NoError(t, err)
	assert.Equal(t, "Bearer refreshed-token", header)
	assert.Equal(t, 1, endpoint.count(), "an expired stored token must refresh exactly once")
	require.Equal(t, 1, store.saveCount(), "a successful refresh must persist the new token")

	stored, err := store.Load(context.Background(), key)
	require.NoError(t, err)
	require.NotNil(t, stored)
	assert.Equal(t, "refreshed-token", stored.AccessToken)
}

func TestTokenLifecycle_ConcurrentRequestsCoalesceToOneRequest(t *testing.T) {
	key := credentials.TokenKey{Authority: testInstanceURL, ClientID: testClientID}
	store := newLifecycleStore()
	store.tokens[key] = &credentials.AccessToken{
		AccessToken:  "stale-token",
		RefreshToken: "refresh-me",
		ExpiresAt:    time.Now().Add(-time.Minute),
	}
	endpoint := newMockTokenEndpoint()
	endpoint.delay = 50 * time.Millisecond
	provider := buildProvider(t, endpoint, store, credentials.ClientCredentialsOptions{
		CommonOptions: credentials.CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
		},
	})
	uri := mustURL(t, testInstanceURL+"/now/table/incident")

	const n = 8
	headers := make([]string, n)
	errs := make([]error, n)

	var wg sync.WaitGroup
	start := make(chan struct{})
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-start
			headers[i], errs[i] = authenticate(t, provider, uri)
		}(i)
	}
	close(start)
	wg.Wait()

	for i := 0; i < n; i++ {
		require.NoError(t, errs[i])
		assert.Equal(t, "Bearer mock-token", headers[i])
	}
	assert.Equal(t, 1, endpoint.count(), "concurrent requests must coalesce into exactly one token request")
}

func TestTokenLifecycle_FailedStampedeCoalescesToOneRequest(t *testing.T) {
	store := newLifecycleStore()
	endpoint := newMockTokenEndpoint()
	endpoint.status = http.StatusInternalServerError
	endpoint.body = []byte(`{"error":"server_error"}`)
	endpoint.delay = 50 * time.Millisecond
	provider := buildProvider(t, endpoint, store, credentials.ClientCredentialsOptions{
		CommonOptions: credentials.CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
		},
	})
	uri := mustURL(t, testInstanceURL+"/now/table/incident")

	const n = 8
	errs := make([]error, n)

	var wg sync.WaitGroup
	start := make(chan struct{})
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-start
			_, errs[i] = authenticate(t, provider, uri)
		}(i)
	}
	close(start)
	wg.Wait()

	for i := 0; i < n; i++ {
		require.Error(t, errs[i])
	}
	assert.Equal(t, 1, endpoint.count(), "a failing token endpoint must be hit once, not once per concurrent caller")
}

func TestTokenLifecycle_LoadErrorIsPropagated(t *testing.T) {
	store := newLifecycleStore()
	store.loadErr = errors.New("store unavailable")
	endpoint := newMockTokenEndpoint()
	provider := buildProvider(t, endpoint, store, credentials.ClientCredentialsOptions{
		CommonOptions: credentials.CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
		},
	})

	_, err := authenticate(t, provider, mustURL(t, testInstanceURL+"/now/table/incident"))
	require.Error(t, err)
	require.ErrorIs(t, err, store.loadErr)
	assert.Zero(t, endpoint.count(), "a store load failure must not trigger a token request")
}

func TestTokenLifecycle_SaveFailureStillServesToken(t *testing.T) {
	store := newLifecycleStore()
	store.saveErr = errors.New("store unavailable")
	endpoint := newMockTokenEndpoint()
	provider := buildProvider(t, endpoint, store, credentials.ClientCredentialsOptions{
		CommonOptions: credentials.CommonOptions{
			AllowedHosts: []string{"dev12345.service-now.com"},
		},
	})

	header, err := authenticate(t, provider, mustURL(t, testInstanceURL+"/now/table/incident"))
	require.NoError(t, err, "a persistence failure must not fail an otherwise successful request")
	assert.Equal(t, "Bearer mock-token", header)
	assert.Equal(t, 1, endpoint.count())
}

func TestTokenLifecycle_DisallowedHostIsNotAuthenticated(t *testing.T) {
	store := newLifecycleStore()
	endpoint := newMockTokenEndpoint()
	provider := buildProvider(t, endpoint, store, credentials.ClientCredentialsOptions{
		CommonOptions: credentials.CommonOptions{
			AllowedHosts: []string{"allowed.example.com"},
		},
	})

	header, err := authenticate(t, provider, mustURL(t, "https://not-allowed.example.com/now/table/incident"))
	require.NoError(t, err)
	assert.Empty(t, header, "a disallowed host must not receive an Authorization header")
	assert.Zero(t, endpoint.count(), "a disallowed host must not trigger a token request")
}
