package oauth2

import "context"

type ClientCredentialsExchanger interface {
	// ExchangeClientCredentials performs a Client Credentials Grant exchange.
	ExchangeClientCredentials(ctx context.Context, scopes []string) (*Token, error)
}
