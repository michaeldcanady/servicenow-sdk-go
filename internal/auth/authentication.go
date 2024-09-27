package auth

import (
	"context"
	"fmt"
)

type AuthenticationProvider interface {
	GetCredential(context.Context) (string, error)
}

type AuthTypeProvider interface {
	GetAuthType() string
}

type authenticationProvider struct {
	authTypeProvider AuthTypeProvider
	strategy         AuthorizationStrategy
}

func NewAuthenticationProvider(authTypeProvider AuthTypeProvider, strategy AuthorizationStrategy) AuthenticationProvider {
	return &authenticationProvider{
		authTypeProvider: authTypeProvider,
		strategy:         strategy,
	}
}

func (a *authenticationProvider) GetCredential(ctx context.Context) (string, error) {
	auth, err := a.strategy.GetAuth(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", a.authTypeProvider.GetAuthType(), auth), nil
}
