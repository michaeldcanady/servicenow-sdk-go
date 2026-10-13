// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	"github.com/stretchr/testify/assert"
)

func TestConvertToken(t *testing.T) {
	tests := []struct {
		name     string
		token    *oauth2.Token
		expected *AccessToken
	}{
		{
			name:     "nil token returns nil",
			token:    nil,
			expected: nil,
		},
		{
			name: "populated token converts fields",
			token: &oauth2.Token{
				AccessToken:  "access",
				TokenType:    "Bearer",
				ExpiresIn:    3600,
				RefreshToken: "refresh",
				Scope:        "read",
			},
			expected: &AccessToken{
				AccessToken:  "access",
				TokenType:    "Bearer",
				ExpiresIn:    3600,
				RefreshToken: "refresh",
				Scope:        "read",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertToken(tt.token)
			if tt.expected == nil {
				assert.Nil(t, got)
				return
			}
			assert.Equal(t, tt.expected.AccessToken, got.AccessToken)
			assert.Equal(t, tt.expected.TokenType, got.TokenType)
			assert.Equal(t, tt.expected.ExpiresIn, got.ExpiresIn)
			assert.Equal(t, tt.expected.RefreshToken, got.RefreshToken)
			assert.Equal(t, tt.expected.Scope, got.Scope)
			assert.False(t, got.ExpiresAt.IsZero())
		})
	}
}
