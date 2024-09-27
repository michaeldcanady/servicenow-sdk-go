package credentials

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/auth"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/credential"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

type AuthenticationProvider = auth.AuthenticationProvider

func NewROPCCrdential(clientID, clientSecret, username, password, baseURL string) auth.AuthenticationProvider {
	strategy, _ := credential.NewServiceNowOauth2Strategy(
		credential.NewROPCCredentialProvider(
			clientID,
			clientSecret,
			username,
			password,
		),
		baseURL,
	)

	return auth.NewBearerAuthenticationStrategy(oauth2.NewAuthorizationStrategy(strategy))
}
