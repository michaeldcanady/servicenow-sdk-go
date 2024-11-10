package auth

import (
	"context"
	"net/url"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// UserPassProvider returns access tokens.
type UserPassProvider interface {
	// GetAuthorizationToken returns the access token for the provided url.
	GetAuthorizationToken(context context.Context, url *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error)
	// GetAllowedHostsValidator returns the hosts validator.
	GetAllowedHostsValidator() *authentication.AllowedHostsValidator
}
