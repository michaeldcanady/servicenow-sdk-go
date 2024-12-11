package authentication

import (
	"context"
	"net/url"
)

type AuthorizationProvider interface {
	GetAuthorization(context.Context, *url.URL, map[string]interface{}) (string, error)
}
