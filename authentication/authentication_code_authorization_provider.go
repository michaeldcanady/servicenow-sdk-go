package authentication

import (
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/pkg/browser"
)

func NewAuthorizationCodeAuthorizationProvider(uriCallback func(string) error, port int, clientID, clientSecret string) *bearerAuthorizationProvider {
	return newBearerAuthorizationProvider(newAuthorizationCodeTokenProvider(uriCallback, port, clientID, clientSecret))
}

func NewBrowserAuthCodeAuthorizationProvider(port int, clientID, clientSecret string) *bearerAuthorizationProvider {
	return NewAuthorizationCodeAuthorizationProvider(browser.OpenURL, port, clientID, clientSecret)
}

func NewExternalJWTAuthorizationProvider(tokenProvider authentication.AccessTokenProvider, clientID, clientSecret string) *bearerAuthorizationProvider {
	return newBearerAuthorizationProvider(newExternalJWTTokenProvider(tokenProvider, clientID, clientSecret))
}

func NewImplicitAuthorizationProvider(uriCallback func(string) error, port int, clientID string) *bearerAuthorizationProvider {
	return newBearerAuthorizationProvider(newImplicitTokenProvider(uriCallback, port, clientID))
}

func NewBrowserImplicitAuthorizationProvider(port int, clientID string) *bearerAuthorizationProvider {
	return NewImplicitAuthorizationProvider(browser.OpenURL, port, clientID)
}

func NewROPCAuthorizationProvider(clientID, clientSecret, username, password string) *bearerAuthorizationProvider {
	return newBearerAuthorizationProvider(newResourceOwnerPasswordTokenProvider(clientID, clientSecret, username, password))
}

func NewNonInteractiveBasicAuthorizationProvider(username, password string) *basicAuthorizationProvider {
	return newBasicAuthorizationProvider(newNonInteractiveUserPassProvider(username, password))
}
