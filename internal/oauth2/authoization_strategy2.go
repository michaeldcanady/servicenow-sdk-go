package oauth2

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/auth"
)

const (
	accessTokenCacheKey  = "access_token"
	refreshTokenCacheKey = "refresh_token"
)

var _ auth.TokenStrategy = (*AuthorizationStrategy[any, any])(nil)

type tokenOptionGenerator[T, C any] func(C) []TokenOption[T]

type AuthorizationStrategy[T, C any] struct {
	grantStrategy        GrantStrategy[T]
	cache                auth.Cache[Oauth2Token]
	config               C
	generateTokenOptions tokenOptionGenerator[T, C]
}

func (aS *AuthorizationStrategy[T, C]) GetToken(ctx context.Context) (string, error) {
	token, err := aS.cache.Retrieve(accessTokenCacheKey)
	if err != nil {
		return "", err
	}

	if token != nil && !token.IsExpired() {
		return token.GetAccessToken(), nil
	}

	refreshToken, err := aS.cache.Retrieve(refreshTokenCacheKey)

	// TODO: not working
	if supportsRefresh, ok := aS.config.(SupportsRefresh); ok && refreshToken != nil {
		supportsRefresh.SetRefreshToken(refreshToken)
	}

	// Generate options for acquiring the token
	opts := aS.generateTokenOptions(aS.config)

	// Acquire a new token
	token, err = aS.grantStrategy.AcquireToken(ctx, opts...)
	if err != nil {
		return "", err
	}

	// Store the token in the cache
	if err := aS.cache.Store(accessTokenCacheKey, token); err != nil {
		return "", err
	}

	// Update the refresh token if available
	if refreshable, ok := token.(HasRefreshToken); ok {
		if err = aS.cache.Store(refreshTokenCacheKey, refreshable.GetRefreshToken()); err != nil {
			return "", err
		}
	}

	return token.GetAccessToken(), nil
}
