package servicenowsdkgo

type UsernamePasswordCredential struct {
	Credential Credential
	username   string
	password   string
}

func NewUsernamePasswordCredential(username, password string) *UsernamePasswordCredential {
	return &UsernamePasswordCredential{
		Credential: *NewCredential(),
		username:   username,
		password:   password,
	}
}

func (U *UsernamePasswordCredential) GetAuthentication() string {
	authString := U.Credential.BasicAuth(U.username, U.password)
	return "Basic " + authString
}
