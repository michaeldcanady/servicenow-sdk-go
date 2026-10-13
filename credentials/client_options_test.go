// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2/pkce"
	"github.com/stretchr/testify/assert"
)

func TestWithPKCEChallenge(t *testing.T) {
	tests := []struct {
		name   string
		method pkce.Method
	}{
		{
			name:   "sets S256 method",
			method: pkce.MethodS256,
		},
		{
			name:   "sets plain method",
			method: pkce.MethodPlain,
		},
		{
			name:   "sets unset method",
			method: pkce.MethodUnset,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts := defaultOptions()
			withPKCEChallenge(test.method)(&opts)

			assert.Equal(t, test.method, opts.method)
		})
	}
}

func TestDefaultOptions(t *testing.T) {
	opts := defaultOptions()

	assert.NotNil(t, opts.httpClient)
	assert.Equal(t, pkce.MethodUnset, opts.method)
}
