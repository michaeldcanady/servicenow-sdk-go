package credentials

// DEPRECATED: deprecated since v{unreleased}. Use [authentication.AuthenticationProvider] implementation instead.
//
// Credential
type Credential interface {
	GetAuthentication() (string, error)
}
