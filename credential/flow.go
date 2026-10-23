// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credential

// Flow identifies the OAuth flow that acquired a token. The numeric values are
// stable because they are part of the token store key. Append new flows; do not
// reorder or renumber existing values.
type Flow int

const (
	// FlowUnknown is the zero value and must not key a token.
	FlowUnknown Flow = iota
	// FlowAuthorizationCode identifies the authorization code flow.
	FlowAuthorizationCode
	// FlowClientCredentials identifies the client credentials flow.
	FlowClientCredentials
	// FlowROPC identifies the resource owner password credentials flow.
	FlowROPC
	// FlowJWTBearer identifies the JWT bearer flow.
	FlowJWTBearer
	// FlowExternalOIDC identifies the external OIDC flow.
	FlowExternalOIDC
)

// String returns the stored name of the flow.
func (f Flow) String() string {
	switch f {
	case FlowAuthorizationCode:
		return "authorization_code"
	case FlowClientCredentials:
		return "client_credentials"
	case FlowROPC:
		return "ropc"
	case FlowJWTBearer:
		return "jwt_bearer"
	case FlowExternalOIDC:
		return "external_oidc"
	default:
		return "unknown"
	}
}
