// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import "context"

// TokenKey identifies a stored token by the authority that issued it and the
// client ID that requested it, so tokens for two instances or two client
// registrations never collide in one store.
type TokenKey struct {
	// Authority is the token issuer: the ServiceNow instance URL for native
	// flows, or the OIDC issuer for external identity providers.
	Authority string
	// ClientID is the OAuth2 client ID the token was issued to.
	ClientID string
}

// KeyedTokenStore is an interface for persisting and retrieving access tokens
// keyed by both the issuing authority and the client ID.
type KeyedTokenStore interface {
	// Save stores the token for the given key.
	Save(ctx context.Context, key TokenKey, token *AccessToken) error
	// Load retrieves the token for the given key. A missing key loads as
	// (nil, nil) rather than an error.
	Load(ctx context.Context, key TokenKey) (*AccessToken, error)
	// Delete removes the token for the given key.
	Delete(ctx context.Context, key TokenKey) error
}

// TokenStore is an interface for persisting and retrieving access tokens by
// client ID.
//
// Deprecated: implement KeyedTokenStore instead. A client-ID-only key cannot
// tell tokens for two instances or two client registrations apart. Set
// CommonOptions.TokenStore to a KeyedTokenStore to migrate.
type TokenStore interface {
	// Save stores the token for the given client ID.
	Save(ctx context.Context, clientID string, token *AccessToken) error
	// Load retrieves the token for the given client ID.
	Load(ctx context.Context, clientID string) (*AccessToken, error)
	// Delete removes the token for the given client ID.
	Delete(ctx context.Context, clientID string) error
}

// clientIDTokenStore adapts a client-ID-keyed TokenStore to KeyedTokenStore by
// projecting the key onto its client ID. It backs the deprecated path where a
// caller supplies a TokenStore through WithTokenStore or SetTokenStore.
type clientIDTokenStore struct {
	store TokenStore
}

// Save stores the token under the key's client ID.
func (a clientIDTokenStore) Save(ctx context.Context, key TokenKey, token *AccessToken) error {
	return a.store.Save(ctx, key.ClientID, token)
}

// Load retrieves the token stored under the key's client ID.
func (a clientIDTokenStore) Load(ctx context.Context, key TokenKey) (*AccessToken, error) {
	return a.store.Load(ctx, key.ClientID)
}

// Delete removes the token stored under the key's client ID.
func (a clientIDTokenStore) Delete(ctx context.Context, key TokenKey) error {
	return a.store.Delete(ctx, key.ClientID)
}
