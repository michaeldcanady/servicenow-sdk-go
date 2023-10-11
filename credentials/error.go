package credentials

type CredentialError struct {
	Message string
}

func NewCredentialError(message string) *CredentialError {
	return &CredentialError{Message: message}
}

func (e *CredentialError) Error() string {
	return e.Message
}
