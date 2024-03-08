package internal

import "net/http"

type BaseAuthorizationProvider struct {
	credential Credential
}

func NewBaseAuthorizationProvider(credential Credential) (*BaseAuthorizationProvider, error) {
	if IsNil(credential) {
		return nil, ErrNilCredential
	}

	return &BaseAuthorizationProvider{
		credential: credential,
	}, nil
}

func (b *BaseAuthorizationProvider) AuthorizeRequest(request RequestInformation) error {
	if IsNil(request) {
		return ErrNilRequest
	}

	headerString, err := b.credential.GetAuthentication()
	if err != nil {
		return err
	}

	newHeaders := http.Header{}
	newHeaders.Set(authorizationHeader, headerString)

	err = request.AddHeaders(newHeaders)
	if err != nil {
		return err
	}

	return nil
}
