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

// TokenStore is an interface for persisting and retrieving access tokens.
type TokenStore interface {
	// Save stores the token for the given key.
	Save(ctx context.Context, key TokenKey, token *AccessToken) error
	// Load retrieves the token for the given key. A missing key loads as
	// (nil, nil) rather than an error.
	Load(ctx context.Context, key TokenKey) (*AccessToken, error)
	// Delete removes the token for the given key.
	Delete(ctx context.Context, key TokenKey) error
}
