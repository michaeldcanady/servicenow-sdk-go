// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

// Package credentials implements authentication providers for ServiceNow OAuth2
// and external OIDC flows. It provides a single public register of
// NewXProvider constructors that return authentication.AuthenticationProvider,
// with the ServiceNow instance URL configured once on the client and passed to
// the provider through Initialize.
//
// The ROPC (Resource Owner Password Credentials) flow is available via
// NewROPCProvider but is strongly discouraged. Per ADR-016, ROPC sends user
// credentials directly to the client application, which expands the attack
// surface, defeats MFA, and violates modern OAuth2 best practices. Prefer
// Authorization Code flow with PKCE (NewPublicAuthorizationCodeProvider) or
// Client Credentials (NewClientCredentialsProvider) for machine-to-machine
// scenarios. External OIDC (NewExternalOIDCProvider) is the recommended path
// for workforce identity.
package credentials
