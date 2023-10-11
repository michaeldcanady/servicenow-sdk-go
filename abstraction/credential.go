package abstraction

type Credential interface {
	GetAuthentication() (string, error)
}
