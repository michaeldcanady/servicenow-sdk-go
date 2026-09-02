// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentialError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "credential error message",
			err:      NewCredentialError("boom"),
			expected: "boom",
		},
		{
			name:     "oauth2 error message",
			err:      NewOauth2Error("oauth boom"),
			expected: "oauth boom",
		},
		{
			name:     "empty client id sentinel",
			err:      ErrEmptyClientID,
			expected: "clientId is empty",
		},
		{
			name:     "empty client secret sentinel",
			err:      ErrEmptyClientSecret,
			expected: "clientSecret is empty",
		},
		{
			name:     "empty base url sentinel",
			err:      ErrEmptyBaseURL,
			expected: "baseURL is empty",
		},
		{
			name:     "empty username sentinel",
			err:      ErrEmptyUsername,
			expected: "username is empty",
		},
		{
			name:     "empty password sentinel",
			err:      ErrEmptyPassword,
			expected: "password is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.err.Error())
		})
	}
}

// A typed-nil *CredentialError must return a safe string instead of panicking,
// matching the nil-receiver guard on NilPointerError (#591).
func TestCredentialError_Error_NilReceiver(t *testing.T) {
	var e *CredentialError

	var got string
	assert.NotPanics(t, func() { got = e.Error() })
	assert.Equal(t, "nil credential error", got)
}
