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
		build    func() authority
		expected authority
	}{
		{
			name:     "instance authority",
			build:    func() authority { return newInstanceAuthority("dev12345") },
			expected: authority("https://dev12345.service-now.com"),
		},
		{
			name:     "custom authority",
			build:    func() authority { return newCustomAuthority("mycustomerservicenowurl.com") },
			expected: authority("https://mycustomerservicenowurl.com"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.build())
		})
	}
}

func TestAuthority_URLs(t *testing.T) {
	authority := authority("https://dev12345.service-now.com")

	tests := []struct {
		name     string
		build    func() string
		expected string
	}{
		{
			name:     "token url",
			build:    authority.tokenURL,
			expected: "https://dev12345.service-now.com/oauth_token.do",
		},
		{
			name:     "auth url",
			build:    authority.authURL,
			expected: "https://dev12345.service-now.com/oauth_auth.do",
		},
		{
			name:     "revocation url",
			build:    authority.revocationURL,
			expected: "https://dev12345.service-now.com/oauth_revoke.do",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.build())
		})
	}
}
