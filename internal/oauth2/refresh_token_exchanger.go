package oauth2

import "context"

type RefreshTokenExchanger interface {
	// ExchangeRefreshToken performs a Refresh Token Grant exchange.
	ExchangeRefreshToken(ctx context.Context, refresh string) (*Token, error)
}
