package authentication

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
)

type BaseUserPassProvider struct {
	username string
	password string
}

func NewBaseUserPassProvider(username, password string) *BaseUserPassProvider {
	return &BaseUserPassProvider{
		username: username,
		password: password,
	}
}

func (provider *BaseUserPassProvider) GetAuthorizationToken(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if provider == nil {
		return "", nil
	}

	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", provider.username, provider.password))), nil
}
