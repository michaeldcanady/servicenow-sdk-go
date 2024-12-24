package core

// Deprecated: deprecated since v{unreleased}.
//
// Credential ...
type Credential interface {
	GetAuthentication() (string, error)
}
