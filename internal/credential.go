package internal

type Credential interface {
	GetAuthentication() (string, error)
}
