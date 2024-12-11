package authentication

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/authentication"
	kauthentication "github.com/microsoft/kiota-abstractions-go/authentication"
)

func NewSilentROPCCredential(
	clientID,
	clientSecret,
	username,
	password string,
	scopes []string,
) *kauthentication.BaseBearerTokenAuthenticationProvider {
	return kauthentication.NewBaseBearerTokenAuthenticationProvider(authentication.NewResourceOwnerPasswordTokenProvider(clientID,
		clientSecret,
		username,
		password,
		scopes))
}
