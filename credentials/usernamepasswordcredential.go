package credentials

import "encoding/base64"

type UsernamePasswordCredential struct {
	username string
	password string
}

// NewUsernamePasswordCredential creates a new instance of UsernamePasswordCredential.
// It accepts the username and password as parameters and returns a pointer to the created UsernamePasswordCredential.
func NewUsernamePasswordCredential(username, password string) *UsernamePasswordCredential {
	return &UsernamePasswordCredential{
		username: username,
		password: password,
	}
}

func (C *UsernamePasswordCredential) getAuthoization() string {
	return C.BasicAuth(C.username, C.password)
}

func (C *UsernamePasswordCredential) getAuthoizationType() string {
	return "Basic"
}

// GetAuthentication returns the authentication string for UsernamePasswordCredential.
// It uses the BasicAuth method from the underlying Credential to generate the encoded authentication string.
// It returns the authentication string in the format "Basic <encoded-auth-string>".
func (C *UsernamePasswordCredential) GetAuthentication() (string, error) {
	return C.getAuthoizationType() + " " + C.getAuthoization(), nil
}

// BasicAuth returns a Basic Authentication string based on the provided username and password.
// The function combines the username and password with a colon and encodes the result using base64.
// It returns the encoded Basic Authentication string.
func (C *UsernamePasswordCredential) BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
