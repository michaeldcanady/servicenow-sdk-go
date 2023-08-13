package servicenowsdkgo

import "encoding/base64"

type Credential struct {
}

func NewCredential() *Credential {
	return &Credential{}
}

func (C *Credential) BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
