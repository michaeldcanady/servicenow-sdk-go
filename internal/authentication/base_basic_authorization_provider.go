package authentication

import (
	"context"
	"errors"
	"fmt"

	abs "github.com/microsoft/kiota-abstractions-go"
)

const (
	authorizationHeader    = "Authorization"
	claimsKey              = "claims"
	basicAuthorizationType = "Basic"
)

// BaseBasicTokenAuthenticationProvider provides a base class implementing AuthenticationProvider for Basic token scheme.
type BaseBasicTokenAuthenticationProvider struct {
	// userPassProvider is called by the BaseBasicTokenAuthenticationProvider class to authenticate the request via the returned access token.
	userPassProvider UserPassProvider
}

// NewBaseBasicAuthenticationProvider creates a new instance of the BaseBasicTokenAuthenticationProvider class.
func NewBaseBasicAuthenticationProvider(userPassProvider UserPassProvider) *BaseBasicTokenAuthenticationProvider {
	return &BaseBasicTokenAuthenticationProvider{userPassProvider}
}

// AuthenticateRequest authenticates the provided RequestInformation instance using the provided authorization token callback.
func (provider *BaseBasicTokenAuthenticationProvider) AuthenticateRequest(ctx context.Context, request *abs.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	if request == nil {
		return errors.New("request is nil")
	}
	if request.Headers == nil {
		request.Headers = abs.NewRequestHeaders()
	}
	if provider.userPassProvider == nil {
		return errors.New("this class needs to be initialized with an access token provider")
	}
	if len(additionalAuthenticationContext) > 0 &&
		additionalAuthenticationContext[claimsKey] != nil &&
		request.Headers.ContainsKey(authorizationHeader) {
		request.Headers.Remove(authorizationHeader)
	}
	if !request.Headers.ContainsKey(authorizationHeader) {
		uri, err := request.GetUri()
		if err != nil {
			return err
		}
		token, err := provider.userPassProvider.GetAuthorizationToken(ctx, uri, additionalAuthenticationContext)
		if err != nil {
			return err
		}
		if token != "" {
			request.Headers.Add(authorizationHeader, fmt.Sprintf("%s %s", basicAuthorizationType, token))
		}
	}

	return nil
}

// GetAuthorizationTokenProvider returns the access token provider the BaseBasicTokenAuthenticationProvider class uses to authenticate the request.
func (provider *BaseBasicTokenAuthenticationProvider) GetAuthorizationTokenProvider() UserPassProvider {
	return provider.userPassProvider
}
