package oauth2

import "context"

type CodeExchanger interface {
	// ExchangeCode performs an Authorization Code Grant exchange.
	ExchangeCode(ctx context.Context, code, redirectURI, verifier, state string) (*Token, error)
}
