package authentication

import "github.com/michaeldcanady/servicenow-sdk-go/internal/authentication"

// NewSilentBasicCredential creates a new BaseBasicTokenAuthenticationProvider using the provided arguments
func NewSilentBasicCredential(username, password string) *authentication.BaseBasicTokenAuthenticationProvider {
	return authentication.NewBaseBasicAuthenticationProvider(authentication.NewBaseUserPassProvider(username, password))
}
