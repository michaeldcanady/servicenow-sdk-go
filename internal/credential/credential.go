package credential

type Credential interface {
	GetAuthentication() string
	authType() string
	authorization() string
}
