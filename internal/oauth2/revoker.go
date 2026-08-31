// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package oauth2

import "context"

type Revoker interface {
	// Revoke invalidates the provided token (RFC 7009).
	Revoke(ctx context.Context, token, tokenTypeHint string) error
}
