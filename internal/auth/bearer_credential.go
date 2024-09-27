package auth

import (
	"context"
	"errors"
)

const bearerAuthenticationType = "Bearer"

type BearerAuthStrategy struct {
	tokenStrategy TokenStrategy
}

func NewBearerAuthenticationStrategy(strategy TokenStrategy) AuthenticationProvider {
	return NewAuthenticationProvider(
		NewStaticAuthTypeProvider(bearerAuthenticationType),
		&BearerAuthStrategy{
			tokenStrategy: strategy,
		},
	)
}

func (b *BearerAuthStrategy) GetAuth(ctx context.Context) (string, error) {
	token, err := b.tokenStrategy.GetToken(ctx)
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", errors.New("token is empty")
	}
	return token, nil
}
