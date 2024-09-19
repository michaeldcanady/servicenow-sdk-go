package oauth2

import "context"

type TokenStrategy interface {
	FetchToken(context.Context) (Oauth2Token, error)
	RefreshToken(string, context.Context) (Oauth2Token, error)
}
