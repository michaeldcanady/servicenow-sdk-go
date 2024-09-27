package oauth2

import "context"

type GrantStrategy[T any] interface {
	AcquireToken(context.Context, ...TokenOption[T]) (Oauth2Token, error)
}
