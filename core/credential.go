package core

type Credential interface {
	GetAuthentication() (string, error)
}
