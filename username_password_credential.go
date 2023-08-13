package servicenowsdkgo

type UsernamePasswordCredential struct {
	Credential Credential
	username   string
	password   string
}

// NewUsernamePasswordCredential creates a new instance of UsernamePasswordCredential.
// It accepts the username and password as parameters and returns a pointer to the created UsernamePasswordCredential.
func NewUsernamePasswordCredential(username, password string) *UsernamePasswordCredential {
	return &UsernamePasswordCredential{
		Credential: *NewCredential(),
		username:   username,
		password:   password,
	}
}

// GetAuthentication returns the authentication string for UsernamePasswordCredential.
// It uses the BasicAuth method from the underlying Credential to generate the encoded authentication string.
// It returns the authentication string in the format "Basic <encoded-auth-string>".
func (U *UsernamePasswordCredential) GetAuthentication() string {
	authString := U.Credential.BasicAuth(U.username, U.password)
	return "Basic " + authString
}
