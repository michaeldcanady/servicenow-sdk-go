package credential

import "github.com/michaeldcanady/servicenow-sdk-go/internal/auth"

const bearerAuthenticationType = "Bearer"

func NewBearerAuthentication(strategy auth.AuthorizationStrategy) auth.Authentication {
	return auth.NewAuthentication(bearerAuthenticationType, strategy)
}
