// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessToken_IsExpired(t *testing.T) {
	const skewBuffer = 20 * time.Second

	tests := []struct {
		name     string
		token    *AccessToken
		expected bool
	}{
		{
			name:     "nil receiver is not expired",
			token:    nil,
			expected: false,
		},
		{
			name:     "expired token",
			token:    &AccessToken{ExpiresAt: time.Now().Add(-time.Minute)},
			expected: true,
		},
		{
			name:     "non-expired token",
			token:    &AccessToken{ExpiresAt: time.Now().Add(time.Minute)},
			expected: false,
		},
		{
			name:     "token within the clock-skew buffer is expired",
			token:    &AccessToken{ExpiresAt: time.Now().Add(skewBuffer / 2)},
			expected: true,
		},
		{
			name:     "token beyond the clock-skew buffer is valid",
			token:    &AccessToken{ExpiresAt: time.Now().Add(skewBuffer * 2)},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.token.IsExpired())
		})
	}
}
