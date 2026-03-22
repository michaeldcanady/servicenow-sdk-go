package credentials

import "encoding/base64"

// Deprecated: deprecated since v1.11.0. Use [BasicAuthenticationProvider] instead.
//
// UsernamePasswordCredential
type UsernamePasswordCredential struct {
	Username string
	Password string //nolint:gosec // G117: Needed for flow, not secret
}

// Deprecated: deprecated since v1.11.0. Use [NewBasicProvider] instead.
//
// NewUsernamePasswordCredential creates a new instance of UsernamePasswordCredential.
// It accepts the username and password as parameters and returns a pointer to the created UsernamePasswordCredential.
func NewUsernamePasswordCredential(username, password string) *UsernamePasswordCredential {
	return &UsernamePasswordCredential{
		Username: username,
		Password: password,
	}
}

// GetAuthentication returns the authentication string for UsernamePasswordCredential.
// It uses the BasicAuth method from the underlying Credential to generate the encoded authentication string.
// It returns the authentication string in the format "Basic <encoded-auth-string>".
func (c *UsernamePasswordCredential) GetAuthentication() (string, error) {
	return c.getAuthoizationType() + " " + c.getAuthoization(), nil
}

// BasicAuth returns a Basic Authentication string based on the provided username and password.
// The function combines the username and password with a colon and encodes the result using base64.
// It returns the encoded Basic Authentication string.
func (c *UsernamePasswordCredential) BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c *UsernamePasswordCredential) getAuthoization() string {
	return c.BasicAuth(c.Username, c.Password)
}

func (c *UsernamePasswordCredential) getAuthoizationType() string {
	return "Basic"
}
