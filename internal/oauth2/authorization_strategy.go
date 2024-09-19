package oauth2

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/auth"
)

var _ auth.AuthorizationStrategy = (*AuthorizationStrategy)(nil)

type AuthorizationStrategy struct {
	strategy     TokenStrategy
	token        Oauth2Token
	refreshToken string
}

func NewAuthorizationStrategy(tokenStrategy TokenStrategy) *AuthorizationStrategy {
	return &AuthorizationStrategy{
		strategy: tokenStrategy,
	}
}

func (aS *AuthorizationStrategy) GetAuth(ctx context.Context) (string, error) {
	if aS.token != nil && !aS.token.IsExpired() {
		return aS.token.GetAccessToken(), nil
	}

	var err error
	if aS.token == nil || aS.token.IsExpired() {
		if aS.refreshToken != "" {
			aS.token, err = aS.strategy.RefreshToken(aS.refreshToken, ctx)
		} else {
			aS.token, err = aS.strategy.FetchToken(ctx)
		}
		if err != nil {
			return "", err
		}
		if tokenWithRefresh, ok := aS.token.(HasRefreshToken); ok {
			aS.refreshToken = tokenWithRefresh.GetRefreshToken()
		}
	}

	return aS.token.GetAccessToken(), nil
}
