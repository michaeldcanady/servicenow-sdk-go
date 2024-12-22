package core

// Deprecated: deprecated since v{unreleased}.
type Credential interface {
	GetAuthentication() (string, error)
}
