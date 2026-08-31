// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import "context"

// TokenStore is an interface for persisting and retrieving access tokens.
type TokenStore interface {
	// Save stores the token for the given client ID.
	Save(ctx context.Context, clientID string, token *AccessToken) error
	// Load retrieves the token for the given client ID.
	Load(ctx context.Context, clientID string) (*AccessToken, error)
	// Delete removes the token for the given client ID.
	Delete(ctx context.Context, clientID string) error
}
