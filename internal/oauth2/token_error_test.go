// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package oauth2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *TokenError
		expected string
	}{
		{
			name:     "error code only",
			err:      &TokenError{Err: "invalid_request"},
			expected: "oauth2 error: invalid_request",
		},
		{
			name:     "error code with description",
			err:      &TokenError{Err: "invalid_grant", ErrorDescription: "refresh token expired"},
			expected: "oauth2 error: invalid_grant (refresh token expired)",
		},
		{
			name:     "error code with status code",
			err:      &TokenError{Err: "invalid_client", StatusCode: 401},
			expected: "oauth2 error: invalid_client (status code: 401)",
		},
		{
			name:     "error code with description and status code",
			err:      &TokenError{Err: "invalid_scope", ErrorDescription: "unknown scope", StatusCode: 400},
			expected: "oauth2 error: invalid_scope (unknown scope) (status code: 400)",
		},
		{
			name:     "empty error code still renders the prefix",
			err:      &TokenError{},
			expected: "oauth2 error: ",
		},
		{
			name:     "error URI and raw body are not part of the message",
			err:      &TokenError{Err: "server_error", ErrorURI: "http://docs", RawBody: `{"error":"server_error"}`},
			expected: "oauth2 error: server_error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.err.Error())
		})
	}
}

func TestTokenError_ErrorsAs(t *testing.T) {
	var err error = &TokenError{Err: "invalid_grant", StatusCode: 400}

	var tokenErr *TokenError
	require.ErrorAs(t, err, &tokenErr)
	assert.Equal(t, "invalid_grant", tokenErr.Err)
	assert.Equal(t, 400, tokenErr.StatusCode)
}

// A typed-nil *TokenError must return a safe string instead of panicking,
// matching the nil-receiver guard on NilPointerError (#591).
func TestTokenError_Error_NilReceiver(t *testing.T) {
	var e *TokenError

	var got string
	assert.NotPanics(t, func() { got = e.Error() })
	assert.Equal(t, "nil oauth2 token error", got)
}
