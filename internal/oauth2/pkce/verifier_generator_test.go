// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package pkce

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// errReader is an io.Reader that always returns an error, used to exercise
// the io.ReadFull failure branch of VerifierGenerator.Generate.
type errReader struct{}

func (errReader) Read(_ []byte) (int, error) {
	return 0, errors.New("boom")
}

// shortReader returns fewer bytes than requested without an error on the
// first call, forcing io.ReadFull to issue a second Read that fails.
type shortReader struct{ n int }

func (r *shortReader) Read(p []byte) (int, error) {
	if r.n == 0 {
		return 0, io.EOF
	}
	n := r.n
	if n > len(p) {
		n = len(p)
	}
	r.n = 0
	return n, nil
}

func TestNewVerifierGenerator(t *testing.T) {
	g := NewVerifierGenerator(DefaultVerifierEntropy)

	assert.Equal(t, DefaultVerifierEntropy, g.EntropyBytes)
	assert.Equal(t, rand.Reader, g.Rand)
}

func TestVerifierGenerator_Generate(t *testing.T) {
	tests := []struct {
		name       string
		generator  VerifierGenerator
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "happy path - generates url-safe base64 verifier",
			generator: VerifierGenerator{
				EntropyBytes: DefaultVerifierEntropy,
				Rand:         rand.Reader,
			},
		},
		{
			name: "zero entropy returns error",
			generator: VerifierGenerator{
				EntropyBytes: 0,
				Rand:         rand.Reader,
			},
			wantErr:    true,
			wantErrMsg: "entropy must be > 0",
		},
		{
			name: "negative entropy returns error",
			generator: VerifierGenerator{
				EntropyBytes: -1,
				Rand:         rand.Reader,
			},
			wantErr:    true,
			wantErrMsg: "entropy must be > 0",
		},
		{
			name: "nil rand reader returns error",
			generator: VerifierGenerator{
				EntropyBytes: DefaultVerifierEntropy,
				Rand:         nil,
			},
			wantErr:    true,
			wantErrMsg: "rand reader cannot be nil",
		},
		{
			name: "rand reader error propagates",
			generator: VerifierGenerator{
				EntropyBytes: DefaultVerifierEntropy,
				Rand:         errReader{},
			},
			wantErr:    true,
			wantErrMsg: "boom",
		},
		{
			name: "short read from rand reader propagates ReadFull error",
			generator: VerifierGenerator{
				EntropyBytes: DefaultVerifierEntropy,
				Rand:         &shortReader{n: 1},
			},
			wantErr:    true,
			wantErrMsg: io.ErrUnexpectedEOF.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.generator.Generate()

			if tt.wantErr {
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErrMsg)
				assert.Empty(t, got)
				return
			}

			require.NoError(t, err)
			assert.NotEmpty(t, got)

			decoded, decodeErr := base64.RawURLEncoding.DecodeString(got)
			require.NoError(t, decodeErr)
			assert.Len(t, decoded, tt.generator.EntropyBytes)
		})
	}
}
