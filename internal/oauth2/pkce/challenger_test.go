// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package pkce

import (
	"crypto/sha256"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChallengerFor(t *testing.T) {
	tests := []struct {
		name        string
		method      Method
		wantChaller Challenger
		wantErr     bool
		wantErrMsg  string
	}{
		{
			name:        "S256 returns S256Challenger",
			method:      MethodS256,
			wantChaller: S256Challenger{},
		},
		{
			name:        "Plain returns PlainChallenger",
			method:      MethodPlain,
			wantChaller: PlainChallenger{},
		},
		{
			name:       "unsupported method returns error",
			method:     MethodUnknown,
			wantErr:    true,
			wantErrMsg: "unsupported PKCE method",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ChallengerFor(tt.method)

			if tt.wantErr {
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErrMsg)
				assert.Nil(t, got)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantChaller, got)
		})
	}
}

func TestNewPKCEChallenge(t *testing.T) {
	verifier := "test-verifier"
	sum := sha256.Sum256([]byte(verifier))
	wantS256 := base64.RawURLEncoding.EncodeToString(sum[:])

	tests := []struct {
		name       string
		method     Method
		verifier   string
		want       string
		wantErr    bool
		wantErrMsg string
	}{
		{
			name:     "plain method returns verifier unchanged",
			method:   MethodPlain,
			verifier: verifier,
			want:     verifier,
		},
		{
			name:     "S256 method returns hashed verifier",
			method:   MethodS256,
			verifier: verifier,
			want:     wantS256,
		},
		{
			name:       "empty verifier returns error",
			method:     MethodPlain,
			verifier:   "",
			wantErr:    true,
			wantErrMsg: "verifier cannot be empty",
		},
		{
			name:       "unsupported method returns error",
			method:     MethodUnknown,
			verifier:   verifier,
			wantErr:    true,
			wantErrMsg: "unsupported PKCE method",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPKCEChallenge(tt.method, tt.verifier)

			if tt.wantErr {
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErrMsg)
				assert.Empty(t, got)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPlainChallenger(t *testing.T) {
	tests := []struct {
		name     string
		verifier string
	}{
		{
			name:     "returns verifier unchanged",
			verifier: "abc123",
		},
		{
			name:     "empty verifier is passed through",
			verifier: "",
		},
	}

	c := PlainChallenger{}

	assert.Equal(t, MethodPlain, c.Name())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Challenge(tt.verifier)
			require.NoError(t, err)
			assert.Equal(t, tt.verifier, got)
		})
	}
}

func TestS256Challenger(t *testing.T) {
	tests := []struct {
		name     string
		verifier string
	}{
		{
			name:     "hashes non-empty verifier",
			verifier: "abc123",
		},
		{
			name:     "hashes empty verifier",
			verifier: "",
		},
	}

	c := S256Challenger{}

	assert.Equal(t, MethodS256, c.Name())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := sha256.Sum256([]byte(tt.verifier))
			want := base64.RawURLEncoding.EncodeToString(sum[:])

			got, err := c.Challenge(tt.verifier)
			require.NoError(t, err)
			assert.Equal(t, want, got)
		})
	}
}
