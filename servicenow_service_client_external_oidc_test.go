// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package servicenowsdkgo

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/credentials"
	"github.com/stretchr/testify/require"
)

// TestNewServiceNowServiceClient_WithExternalOIDCProvider verifies that the
// external OIDC provider satisfies Kiota's authentication provider contract and
// wires into the service client, which initializes it with the base URL.
func TestNewServiceNowServiceClient_WithExternalOIDCProvider(t *testing.T) {
	var idp *httptest.Server
	idp = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/.well-known/openid-configuration" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = io.WriteString(w, `{"issuer":"`+idp.URL+`","authorization_endpoint":"`+idp.URL+`/authorize","token_endpoint":"`+idp.URL+`/token"}`)
	}))
	t.Cleanup(idp.Close)

	provider, err := credentials.NewExternalOIDCProvider("client-id",
		credentials.WithOIDCIssuer(idp.URL),
		credentials.WithHTTPClient(idp.Client()),
	)
	require.NoError(t, err)
	require.NotNil(t, provider)

	client, err := NewServiceNowServiceClient(
		WithURL("https://dev12345.service-now.com"),
		WithAuthenticationProvider(provider),
	)
	require.NoError(t, err)
	require.NotNil(t, client)
}
