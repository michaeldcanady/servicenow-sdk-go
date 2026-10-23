// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlowValues(t *testing.T) {
	assert.Equal(t, FlowUnknown, Flow(0))
	assert.Equal(t, "unknown", FlowUnknown.String())
	assert.Equal(t, "authorization_code", FlowAuthorizationCode.String())
	assert.Equal(t, "client_credentials", FlowClientCredentials.String())
	assert.Equal(t, "ropc", FlowROPC.String())
	assert.Equal(t, "jwt_bearer", FlowJWTBearer.String())
	assert.Equal(t, "external_oidc", FlowExternalOIDC.String())
	assert.Equal(t, "unknown", Flow(99).String())
	assert.Equal(t, FlowAuthorizationCode, Flow(1))
	assert.Equal(t, FlowExternalOIDC, Flow(5))
}

func TestClientAuthMethodString(t *testing.T) {
	tests := []struct {
		name     string
		method   ClientAuthMethod
		expected string
	}{
		{name: "default", method: ClientAuthMethodDefault, expected: "client_secret_post"},
		{name: "post", method: ClientAuthMethodClientSecretPost, expected: "client_secret_post"},
		{name: "basic", method: ClientAuthMethodClientSecretBasic, expected: "client_secret_basic"},
		{name: "none", method: ClientAuthMethodNone, expected: "none"},
		{name: "unknown", method: ClientAuthMethodUnknown, expected: "unknown"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.method.String())
		})
	}
}

func TestAuthLocationString(t *testing.T) {
	assert.Equal(t, "header", AuthLocationHeader.String())
	assert.Equal(t, "query", AuthLocationQuery.String())
	assert.Equal(t, "unknown", AuthLocationUnknown.String())
}
