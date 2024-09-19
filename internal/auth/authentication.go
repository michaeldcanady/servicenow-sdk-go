package auth

import (
	"context"
	"fmt"
)

type Authentication interface {
	GetCredential(context.Context) (string, error)
}

type authenticationImpl struct {
	authType string
	strategy AuthorizationStrategy
}

func NewAuthentication(authType string, strategy AuthorizationStrategy) Authentication {
	return &authenticationImpl{
		authType: authType,
		strategy: strategy,
	}
}

func (a *authenticationImpl) GetCredential(ctx context.Context) (string, error) {
	auth, err := a.strategy.GetAuth(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", a.authType, auth), nil
}
