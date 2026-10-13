// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

// authority represents the base URL for the ServiceNow instance.
type authority string

// Authority represents the base URL for the ServiceNow instance.
//
// Deprecated: register the instance on the client once instead.
type Authority = authority

// NewInstanceAuthority creates an authority from a ServiceNow instance name.
// e.g., "dev12345" becomes "https://dev12345.service-now.com"
//
// Deprecated: register the instance on the client once instead.
func NewInstanceAuthority(instance string) Authority {
	return newInstanceAuthority(instance)
}

// NewCustomAuthority creates an authority from a custom domain.
// e.g., "mycustomerservicenowurl.com" becomes "https://mycustomerservicenowurl.com"
//
// Deprecated: register the instance on the client once instead.
func NewCustomAuthority(domain string) Authority {
	return newCustomAuthority(domain)
}

// newInstanceAuthority creates an authority from a ServiceNow instance name.
// e.g., "dev12345" becomes "https://dev12345.service-now.com"
func newInstanceAuthority(instance string) authority {
	return authority("https://" + instance + ".service-now.com")
}

// newCustomAuthority creates an authority from a custom domain.
// e.g., "mycustomerservicenowurl.com" becomes "https://mycustomerservicenowurl.com"
func newCustomAuthority(domain string) authority {
	return authority("https://" + domain)
}

// TokenURL returns the OAuth2 token endpoint for the authority.
//
// Deprecated: use the client's instance configuration instead.
func (a authority) TokenURL() string {
	return a.tokenURL()
}

// AuthURL returns the OAuth2 authorization endpoint for the authority.
//
// Deprecated: use the client's instance configuration instead.
func (a authority) AuthURL() string {
	return a.authURL()
}

// RevocationURL returns the OAuth2 token revocation endpoint for the authority.
//
// Deprecated: use the client's instance configuration instead.
func (a authority) RevocationURL() string {
	return a.revocationURL()
}

// tokenURL returns the OAuth2 token endpoint for the authority.
func (a authority) tokenURL() string {
	return string(a) + "/oauth_token.do"
}

// authURL returns the OAuth2 authorization endpoint for the authority.
func (a authority) authURL() string {
	return string(a) + "/oauth_auth.do"
}

// revocationURL returns the OAuth2 token revocation endpoint for the authority.
func (a authority) revocationURL() string {
	return string(a) + "/oauth_revoke.do"
}
