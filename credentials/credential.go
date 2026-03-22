package credentials

// DEPRECATED: deprecated since v1.11.0. Use [authentication.AuthenticationProvider] implementation instead.
//
// Credential
type Credential interface {
	GetAuthentication() (string, error)
}
