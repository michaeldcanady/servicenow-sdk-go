// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// This file keeps the v2 credentials types and constructors working while the
// provider surface from NewClientCredentialsProvider and its siblings becomes
// the supported path. Every symbol here is deprecated and is scheduled for
// removal in the next major release.

// ClientCredentialsCredential implements the OAuth2 Client Credentials flow.
//
// Deprecated: use the provider returned by NewClientCredentialsProvider.
type ClientCredentialsCredential = clientCredentialsCredential

// NewClientCredentialsCredential creates a new ClientCredentialsCredential.
//
// Deprecated: use NewClientCredentialsProvider.
func NewClientCredentialsCredential(client clientCredentialsClient, scopes []string, allowedHosts []string) (*ClientCredentialsCredential, error) {
	c := newClientCredentialsCredential(client, scopes)
	c.SetAllowedHosts(allowedHosts)
	return c, nil
}

// ROPCCredential implements the OAuth2 ROPC flow.
//
// Deprecated: use the provider returned by NewROPCProvider.
type ROPCCredential = ropcCredential

// NewROPCCredential creates a new ROPCCredential.
//
// Deprecated: use NewROPCProvider.
func NewROPCCredential(client ropcClient, username, password string, allowedHosts []string) (*ROPCCredential, error) {
	c := newROPCCredential(client, username, password)
	c.SetAllowedHosts(allowedHosts)
	return c, nil
}

// JWTCredential implements the OAuth2 JWT Bearer Token flow.
//
// Deprecated: use the provider returned by NewJWTProvider.
type JWTCredential = jwtCredential

// NewJWTCredential creates a new JWTCredential.
//
// Deprecated: use NewJWTProvider.
func NewJWTCredential(client jwtClient, tokenProvider authentication.AccessTokenProvider, allowedHosts []string) (*JWTCredential, error) {
	c := newJWTCredential(client, tokenProvider)
	c.SetAllowedHosts(allowedHosts)
	return c, nil
}

// AuthorizationCodeCredential implements the OAuth2 Authorization Code flow.
//
// Deprecated: use the provider returned by NewAuthorizationCodeProvider.
type AuthorizationCodeCredential = authorizationCodeCredential

// NewAuthorizationCodeCredential creates a new AuthorizationCodeCredential.
//
// Deprecated: use NewAuthorizationCodeProvider.
func NewAuthorizationCodeCredential(client authorizationCodeClient, allowedHosts []string, port int, stateGenerator func() string, urlOpener func(string) error, serverFactory ServerFactory) (*AuthorizationCodeCredential, error) {
	c := newAuthorizationCodeCredential(client, port, stateGenerator, urlOpener, serverFactory)
	c.SetAllowedHosts(allowedHosts)
	return c, nil
}

// BaseAccessTokenProvider provides common functionality for OAuth2 token providers.
//
// Deprecated: use the provider returned by NewClientCredentialsProvider or its
// sibling constructors.
type BaseAccessTokenProvider = baseAccessTokenProvider
