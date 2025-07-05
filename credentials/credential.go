package credentials

type Credential interface {
	GetAuthentication() (string, error)
}
