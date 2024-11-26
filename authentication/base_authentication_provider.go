package authentication

import (
	"context"
	"errors"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

const (
	authorizationHeader = "Authorization"
	claimsKey           = "claims"
)

var _ authentication.AuthenticationProvider = (*baseAuthenticationProvider)(nil)

type baseAuthenticationProvider struct {
	authorizationProvider AuthorizationProvider
}

func (provider *baseAuthenticationProvider) AuthenticateRequest(ctx context.Context, request *abstractions.RequestInformation, additionalAuthenticationContext map[string]interface{}) error {
	if internal.IsNil(provider) {
		return errors.New("provider is nil")
	}

	if request == nil {
		return errors.New("request is nil")
	}

	if request.Headers == nil {
		request.Headers = abstractions.NewRequestHeaders()
	}

	// Handle additional authentication context
	if len(additionalAuthenticationContext) > 0 {
		if _, ok := additionalAuthenticationContext[claimsKey]; ok && request.Headers.ContainsKey(authorizationHeader) {
			request.Headers.Remove(authorizationHeader)
		}
	}

	if !request.Headers.ContainsKey(authorizationHeader) {
		uri, err := request.GetUri()
		if err != nil {
			return fmt.Errorf("failed to get URI: %w", err)
		}

		authorization, err := provider.authorizationProvider.GetAuthorization(ctx, uri, additionalAuthenticationContext)
		if err != nil {
			return fmt.Errorf("failed to get authorization: %w", err)
		}

		if authorization != "" {
			request.Headers.Add(authorizationHeader, authorization)
		}
	}

	return nil
}
