package authentication

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

const (
	bearerAuthorizationKey = "bearer"
)

var _ AuthorizationProvider = (*bearerAuthorizationProvider)(nil)

type bearerAuthorizationProvider struct {
	accessTokenProvider authentication.AccessTokenProvider
}

func (provider *bearerAuthorizationProvider) GetAuthorization(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if internal.IsNil(provider) {
		return "", errors.New("provider is nil")
	}

	token, err := provider.accessTokenProvider.GetAuthorizationToken(ctx, uri, additionalAuthenticationContext)
	if err != nil {
		return "", fmt.Errorf("failed to get authorization token: %w", err)
	}

	if token != "" {
		return fmt.Sprintf("%s %s", bearerAuthorizationKey, token), nil
	}
	return "", nil
}
