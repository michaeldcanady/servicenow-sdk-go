package credentials

type CredentialError struct {
	Message string
}

type Oauth2Error struct {
	CredentialError
}

func NewOauth2Error(message string) *Oauth2Error {
	return &Oauth2Error{
		CredentialError: *(NewCredentialError(message)),
	}
}

func NewCredentialError(message string) *CredentialError {
	return &CredentialError{Message: message}
}

func (e *CredentialError) Error() string {
	return e.Message
}
