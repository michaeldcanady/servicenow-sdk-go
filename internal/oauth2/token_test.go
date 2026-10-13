// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package oauth2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToken_DecodesIDToken(t *testing.T) {
	body := []byte(`{"access_token":"access","token_type":"Bearer","expires_in":3600,"id_token":"header.payload.signature"}`)

	var token Token
	require.NoError(t, json.Unmarshal(body, &token))

	assert.Equal(t, "access", token.AccessToken)
	assert.Equal(t, "header.payload.signature", token.IDToken)
}

func TestToken_IDTokenAbsentStaysEmpty(t *testing.T) {
	body := []byte(`{"access_token":"access","token_type":"Bearer","expires_in":3600}`)

	var token Token
	require.NoError(t, json.Unmarshal(body, &token))

	assert.Empty(t, token.IDToken)
}
