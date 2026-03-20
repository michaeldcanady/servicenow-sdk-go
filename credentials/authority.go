package credentials

// Authority represents the base URL for the ServiceNow instance.
type Authority string

// NewInstanceAuthority creates an authority from a ServiceNow instance name.
// e.g., "dev12345" becomes "https://dev12345.service-now.com"
func NewInstanceAuthority(instance string) Authority {
	return Authority("https://" + instance + ".service-now.com")
}

// NewCustomAuthority creates an authority from a custom domain.
// e.g., "mycustomerservicenowurl.com" becomes "https://mycustomerservicenowurl.com"
func NewCustomAuthority(domain string) Authority {
	return Authority("https://" + domain)
}

// TokenURL returns the OAuth2 token endpoint for the authority.
func (a Authority) TokenURL() string {
	return string(a) + "/oauth_token.do"
}

// AuthURL returns the OAuth2 authorization endpoint for the authority.
func (a Authority) AuthURL() string {
	return string(a) + "/oauth_auth.do"
}

// RevocationURL returns the OAuth2 token revocation endpoint for the authority.
func (a Authority) RevocationURL() string {
	return string(a) + "/oauth_revoke.do"
}
