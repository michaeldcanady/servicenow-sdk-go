package oauth2

import "context"

type PasswordExchanger interface {
	// ExchangePassword performs a Resource Owner Password Credentials Grant exchange.
	ExchangePassword(ctx context.Context, user, pass string, scopes []string) (*Token, error)
}
