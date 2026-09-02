// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultServerFactory(t *testing.T) {
	server, err := defaultServerFactory("state", 0)
	require.NoError(t, err)
	require.NotNil(t, server)

	addr := server.GetAddr()
	assert.NotEmpty(t, addr)

	parsed, err := url.Parse(addr)
	require.NoError(t, err)
	assert.Equal(t, "http", parsed.Scheme)
	assert.NotEmpty(t, parsed.Port())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	require.NoError(t, server.Shutdown(ctx))
}

func TestRedirectServer_Result(t *testing.T) {
	t.Run("cancelled context yields an error result", func(t *testing.T) {
		server, err := defaultServerFactory("state", 0)
		require.NoError(t, err)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		code, state, resultErr := server.Result(ctx)

		assert.Empty(t, code)
		assert.Empty(t, state)
		require.Error(t, resultErr)

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer shutdownCancel()
		require.NoError(t, server.Shutdown(shutdownCtx))
	})
}
