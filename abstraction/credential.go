package abstraction

import "encoding/base64"

type Credential struct {
}

// NewCredential creates a new instance of the Credential struct.
// It returns a pointer to the newly created Credential.
func NewCredential() *Credential {
	return &Credential{}
}

// BasicAuth returns a Basic Authentication string based on the provided username and password.
// The function combines the username and password with a colon and encodes the result using base64.
// It returns the encoded Basic Authentication string.
func (C *Credential) BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
