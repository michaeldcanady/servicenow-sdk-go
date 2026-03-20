package oauth2

import "context"

type JWTExchanger interface {
	// ExchangeJWT performs a JWT Bearer Token Grant exchange.
	ExchangeJWT(ctx context.Context, assertion string) (*Token, error)
}
