package authentication

import (
	"context"
	"net/url"
)

type UserPassProvider interface {
	GetAuthorizationToken(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error)
}
