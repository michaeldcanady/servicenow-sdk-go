package oauth2

import (
	"errors"
	"strings"
)

// Endpoints holds the URLs for various OAuth2 service endpoints.
type Endpoints struct {
	// AuthURL is the URL of the authorization server's authorization endpoint.
	AuthURL string
	// TokenURL is the URL of the authorization server's token endpoint.
	TokenURL string
	// DeviceURL is the URL of the authorization server's device authorization endpoint.
	DeviceURL string
	// RevocationURL is the URL of the authorization server's token revocation endpoint.
	RevocationURL string
	// IntrospectionURL is the URL of the authorization server's token introspection endpoint.
	IntrospectionURL string
}

// Validate checks if the required endpoints for a given grant type are set.
func (e *Endpoints) Validate(grantType string) error {
	if e == nil {
		return errors.New("endpoints are not set")
	}

	if strings.TrimSpace(e.TokenURL) == "" {
		return errors.New("token endpoint is not set")
	}

	switch grantType {
	case GrantTypeAuthCode:
		if strings.TrimSpace(e.AuthURL) == "" {
			return errors.New("authorization endpoint is not set")
		}
	case GrantTypeDeviceCode:
		if strings.TrimSpace(e.DeviceURL) == "" {
			return errors.New("device authorization endpoint is not set")
		}
	}

	return nil
}
