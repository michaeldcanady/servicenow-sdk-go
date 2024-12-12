package authentication

import "github.com/michaeldcanady/servicenow-sdk-go/internal/authentication"

func NewSilentBasicCredential(username, password string) *authentication.BaseBasicTokenAuthenticationProvider {
	return authentication.NewBaseBasicAuthenticationProvider(authentication.NewBaseUserPassProvider(username, password))
}
