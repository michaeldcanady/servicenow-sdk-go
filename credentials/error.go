package credentials

// CredentialError represents an error related to credentials or authentication.
type CredentialError struct {
	// Message is the error message associated with the credential error.
	Message string
}

// Oauth2Error represents an error related to OAuth2 authentication and extends CredentialError.
type Oauth2Error struct {
	CredentialError
}

// NewOauth2Error creates a new Oauth2Error with the specified error message.
func NewOauth2Error(message string) *Oauth2Error {
	return &Oauth2Error{
		CredentialError: *NewCredentialError(message),
	}
}

// NewCredentialError creates a new CredentialError with the specified error message.
func NewCredentialError(message string) *CredentialError {
	return &CredentialError{Message: message}
}

// Error returns the error message associated with the CredentialError.
func (e *CredentialError) Error() string {
	return e.Message
}

var (
	ErrEmptyClientID     = NewOauth2Error("clientId is empty")
	ErrEmptyClientSecret = NewOauth2Error("clientSecret is empty")
	ErrEmptyBaseURL      = NewOauth2Error("baseURL is empty")
	ErrEmptyUsername     = NewOauth2Error("username is empty")
	ErrEmptyPassword     = NewOauth2Error("password is empty")
)
