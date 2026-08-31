// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package pkce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethod_String(t *testing.T) {
	tests := []struct {
		name   string
		method Method
		want   string
	}{
		{
			name:   "plain returns correct string",
			method: MethodPlain,
			want:   "plain",
		},
		{
			name:   "S256 returns correct string",
			method: MethodS256,
			want:   "S256",
		},
		{
			name:   "unset returns correct string",
			method: MethodUnset,
			want:   "unset",
		},
		{
			name:   "unknown returns correct string",
			method: MethodUnknown,
			want:   "unknown",
		},
		{
			name:   "unsupported value falls back to unknown",
			method: Method(999),
			want:   "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.method.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
