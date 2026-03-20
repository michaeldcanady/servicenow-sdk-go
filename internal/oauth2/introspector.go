package oauth2

import "context"

type Introspector interface {
	// Introspect checks the status and metadata of a token (RFC 7662).
	Introspect(ctx context.Context, token, tokenTypeHint string) (*IntrospectionResponse, error)
}
