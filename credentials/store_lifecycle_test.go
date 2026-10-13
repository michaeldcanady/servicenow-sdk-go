// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	revokeTestClientID  = "test-client"
	revokeTestInstance  = "https://dev12345.service-now.com"
	revokeTestTokenBody = `{"access_token":"acquired-token","token_type":"Bearer","expires_in":3600}`
)

// recordingStore records store operations for the revoke lifecycle tests.
type recordingStore struct {
	mu      sync.Mutex
	tokens  map[TokenKey]*AccessToken
	deleted []TokenKey
}

func newRecordingStore() *recordingStore {
	return &recordingStore{tokens: make(map[TokenKey]*AccessToken)}
}

func (s *recordingStore) Save(_ context.Context, key TokenKey, token *AccessToken) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens[key] = token
	return nil
}

func (s *recordingStore) Load(_ context.Context, key TokenKey) (*AccessToken, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.tokens[key], nil
}

func (s *recordingStore) Delete(_ context.Context, key TokenKey) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.tokens, key)
	s.deleted = append(s.deleted, key)
	return nil
}

func (s *recordingStore) deleteCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.deleted)
}

func (s *recordingStore) deletedKeys() []TokenKey {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]TokenKey, len(s.deleted))
	copy(out, s.deleted)
	return out
}

// okJSONRoundTripper serves the same JSON body for every request.
type okJSONRoundTripper struct {
	body []byte
}

func (rt okJSONRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       io.NopCloser(bytes.NewReader(rt.body)),
		Header:     make(http.Header),
		Request:    req,
	}, nil
}

func mustInPackageURL(t *testing.T, raw string) *url.URL {
	t.Helper()
	u, err := url.Parse(raw)
	require.NoError(t, err)
	return u
}

// newInPackageProvider builds a client-credentials provider and returns the
// underlying credential so its Revoke can be exercised.
func newInPackageProvider(t *testing.T, store KeyedTokenStore) (authentication.AuthenticationProvider, *clientCredentialsCredential) {
	t.Helper()
	provider, err := NewClientCredentialsProvider(revokeTestClientID, "test-secret",
		ClientCredentialsOptions{
			CommonOptions: CommonOptions{
				HTTPClient: &http.Client{Transport: okJSONRoundTripper{body: []byte(revokeTestTokenBody)}},
				TokenStore: store,
			},
		},
	)
	require.NoError(t, err)

	// Initialize the provider with the instance URL
	if prep, ok := provider.(interface{ Initialize(string) }); ok {
		prep.Initialize(revokeTestInstance)
	}

	bearer, ok := provider.(*BearerTokenAuthenticationProvider)
	require.True(t, ok)
	credential, ok := bearer.GetAuthorizationTokenProvider().(*clientCredentialsCredential)
	require.True(t, ok)
	return provider, credential
}

func TestTokenLifecycle_RevokeDeletesStoredToken(t *testing.T) {
	store := newRecordingStore()
	provider, credential := newInPackageProvider(t, store)
	key := TokenKey{Authority: revokeTestInstance, ClientID: revokeTestClientID}

	request := abstractions.NewRequestInformation()
	request.SetUri(*mustInPackageURL(t, revokeTestInstance+"/now/table/incident"))
	require.NoError(t, provider.AuthenticateRequest(context.Background(), request, nil))
	require.NotNil(t, store.tokens[key], "acquisition must persist the token before revoke")

	require.NoError(t, credential.Revoke(context.Background()))

	require.Equal(t, 1, store.deleteCount(), "revoke must delete the stored token")
	assert.Equal(t, key, store.deletedKeys()[0])
	assert.Nil(t, store.tokens[key])
}

func TestTokenLifecycle_RevokeDeletesStoredTokenWithoutInMemoryToken(t *testing.T) {
	store := newRecordingStore()
	key := TokenKey{Authority: revokeTestInstance, ClientID: revokeTestClientID}
	store.tokens[key] = &AccessToken{AccessToken: "stored-token", ExpiresAt: time.Now().Add(time.Hour)}

	// The provider is built against the store but never acquires, modeling a
	// process that restarted and only has the persisted token.
	_, credential := newInPackageProvider(t, store)

	require.NoError(t, credential.Revoke(context.Background()))

	require.Equal(t, 1, store.deleteCount(), "revoke must delete the stored token even without an in-memory token")
	assert.Nil(t, store.tokens[key])
}
