package auth

import "context"

type AuthorizationStrategy interface {
	GetAuth(context.Context) (string, error)
}
