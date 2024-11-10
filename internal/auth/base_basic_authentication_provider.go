package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

var _ authentication.AuthenticationProvider = (*BaseBasicAuthenticationProvider)(nil)

const (
	authorizationHeader = "Authorization"
	basicKey            = "Basic"
)

type BaseBasicAuthenticationProvider struct {
	userPassProvider UserPassProvider
}

// NewBaseBasicAuthenticationProvider creates a new instance of the BaseBasicAuthenticationProvider class.
func NewBaseBasicAuthenticationProvider(userPassProvider UserPassProvider) *BaseBasicAuthenticationProvider {
	return &BaseBasicAuthenticationProvider{userPassProvider: userPassProvider}
}

// AuthenticateRequest authenticates the provided RequestInformation instance using the provided authorization token callback.
func (provider *BaseBasicAuthenticationProvider) AuthenticateRequest(ctx context.Context, request *abstractions.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	if request == nil {
		return errors.New("request is nil")
	}
	if request.Headers == nil {
		request.Headers = abstractions.NewRequestHeaders()
	}
	if provider.userPassProvider == nil {
		return errors.New("this class needs to be initialized with a user-pass provider")
	}
	if !request.Headers.ContainsKey(authorizationHeader) {
		uri, err := request.GetUri()
		if err != nil {
			return err
		}
		userPass, err := provider.userPassProvider.GetAuthorizationToken(ctx, uri, additionalAuthenticationContext)
		if err != nil {
			return err
		}
		if userPass != "" {
			request.Headers.Add(authorizationHeader, fmt.Sprintf("%s %s", basicKey, userPass))
		}
	}

	return nil
}

// GetAuthorizationTokenProvider returns the user-pass provider the BaseBasicAuthenticationProvider class uses to authenticate the request.
func (provider *BaseBasicAuthenticationProvider) GetAuthorizationTokenProvider() UserPassProvider {
	if core.IsNil(provider) {
		return nil
	}

	return provider.userPassProvider
}
