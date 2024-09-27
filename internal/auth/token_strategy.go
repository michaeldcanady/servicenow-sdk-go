package auth

import "context"

type TokenStrategy interface {
	GetToken(context.Context) (string, error)
}
