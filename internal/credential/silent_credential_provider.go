package credential

type SilentCredentialsProvider struct {
	username     string
	password     string
	clientID     string
	clientSecret string
}

func NewSilentCredentialsProvider(clientID, clientSecret, username, password string) *SilentCredentialsProvider {
	return &SilentCredentialsProvider{
		username:     username,
		password:     password,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (cP *SilentCredentialsProvider) GetClientID() string {
	return cP.clientID
}
func (cP *SilentCredentialsProvider) GetClientSecret() string {
	return cP.clientSecret
}
func (cP *SilentCredentialsProvider) GetUsername() string {
	return cP.username
}
func (cP *SilentCredentialsProvider) GetPassword() string {
	return cP.password
}
