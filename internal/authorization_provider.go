package internal

const (
	AuthorizationHeader = "Authorization"
)

type Credential interface {
	GetAuthentication() (string, error)
}

type AuthorizationProvider interface {
	AuthorizeRequest(request RequestInformation) error
}

type BasicAuthorizationProvider struct {
	credential Credential
}

func (b *BasicAuthorizationProvider) AuthorizeRequest(request RequestInformation) error {
	header, err := b.credential.GetAuthentication()
	if err != nil {
		return err
	}
}
