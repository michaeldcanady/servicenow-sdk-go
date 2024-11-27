package authentication

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

var _ UserPassProvider = (*nonInteractiveUserPassProvider)(nil)

type nonInteractiveUserPassProvider struct {
	username string
	password string
}

func newNonInteractiveUserPassProvider(username, password string) *nonInteractiveUserPassProvider {
	return &nonInteractiveUserPassProvider{
		username: username,
		password: password,
	}
}

func (provider *nonInteractiveUserPassProvider) GetUserPass(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if internal.IsNil(provider) {
		return "", nil
	}

	userPassString := fmt.Sprintf("%s:%s", provider.username, provider.password)

	return base64.StdEncoding.EncodeToString([]byte(userPassString)), nil
}
