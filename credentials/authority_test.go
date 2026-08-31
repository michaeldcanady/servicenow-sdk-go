// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthority_Constructors(t *testing.T) {
	tests := []struct {
		name     string
		build    func() Authority
		expected Authority
	}{
		{
			name:     "instance authority",
			build:    func() Authority { return NewInstanceAuthority("dev12345") },
			expected: Authority("https://dev12345.service-now.com"),
		},
		{
			name:     "custom authority",
			build:    func() Authority { return NewCustomAuthority("mycustomerservicenowurl.com") },
			expected: Authority("https://mycustomerservicenowurl.com"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.build())
		})
	}
}

// TODO: only tests happy path, should also test invalid authority strings (e.g., empty, missing https, etc.)
func TestAuthority_URLs(t *testing.T) {
	authority := Authority("https://dev12345.service-now.com")

	tests := []struct {
		name     string
		build    func() string
		expected string
	}{
		{
			name:     "token url",
			build:    authority.TokenURL,
			expected: "https://dev12345.service-now.com/oauth_token.do",
		},
		{
			name:     "auth url",
			build:    authority.AuthURL,
			expected: "https://dev12345.service-now.com/oauth_auth.do",
		},
		{
			name:     "revocation url",
			build:    authority.RevocationURL,
			expected: "https://dev12345.service-now.com/oauth_revoke.do",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.build())
		})
	}
}
