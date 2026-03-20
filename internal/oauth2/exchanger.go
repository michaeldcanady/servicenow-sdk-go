package oauth2

import (
	"context"
	"net/url"
)

type Exchanger interface {
	// Exchange performs a generic token exchange by sending the provided parameters to the token endpoint.
	Exchange(ctx context.Context, params url.Values) (*Token, error)
}
