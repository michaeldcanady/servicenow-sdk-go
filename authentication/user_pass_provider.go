package authentication

import (
	"context"
	"net/url"
)

type UserPassProvider interface {
	GetUserPass(context.Context, *url.URL, map[string]interface{}) (string, error)
}
