package authentication

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

const (
	basicAuthorizationKey = "basic"
)

var _ AuthorizationProvider = (*basicAuthorizationProvider)(nil)

type basicAuthorizationProvider struct {
	userPassProvider UserPassProvider
}

func (provider *basicAuthorizationProvider) GetAuthorization(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if internal.IsNil(provider) {
		return "", errors.New("provider is nil")
	}

	userPass, err := provider.userPassProvider.GetUserPass(ctx, uri, additionalAuthenticationContext)
	if err != nil {
		return "", fmt.Errorf("failed to get user-pass: %w", err)
	}

	if userPass != "" {
		return fmt.Sprintf("%s %s", basicAuthorizationKey, userPass), nil
	}
	return "", nil
}
