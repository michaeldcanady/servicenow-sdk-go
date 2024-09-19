package credentials

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/auth"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/credential"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

type Authentication = auth.Authentication

func NewSilentOauth2Credential(clientID, clientSecret, username, password, baseURL string) auth.Authentication {
	strategy, _ := credential.NewServiceNowOauth2Strategy(
		credential.NewSilentCredentialsProvider(
			clientID,
			clientSecret,
			username,
			password,
		),
		baseURL,
	)

	return credential.NewBearerAuthentication(oauth2.NewAuthorizationStrategy(strategy))
}
