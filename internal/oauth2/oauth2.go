package oauth2

import "github.com/michaeldcanady/servicenow-sdk-go/internal/auth"

func NewOauth2TokenStrategy[T, C any](
	strategy GrantStrategy[T],
	cache auth.Cache[Oauth2Token],
	config C,
	generator tokenOptionGenerator[T, C],
) *AuthorizationStrategy[T, C] {
	return &AuthorizationStrategy[T, C]{
		grantStrategy:        strategy,
		cache:                cache,
		config:               config,
		generateTokenOptions: generator,
	}
}
