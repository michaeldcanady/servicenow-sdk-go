package credentials

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type BasicAuthenticationProvider struct {
	username string
	password string
}

// NewBasicProvider instantiates a new BasicAuthenticationProvider.
func NewBasicProvider(username, password string) *BasicAuthenticationProvider {
	return &BasicAuthenticationProvider{
		username: username,
		password: password,
	}
}

// AuthenticateRequest authenticates the provided RequestInformation.
func (b *BasicAuthenticationProvider) AuthenticateRequest(context context.Context, request *abstractions.RequestInformation, _ map[string]interface{}) error {
	if internal.IsNil(b) {
		return errors.New("provider is nil")
	}

	if err := context.Err(); err != nil {
		return err
	}

	if strings.TrimSpace(b.username) == "" {
		return errors.New("username is empty")
	}

	if strings.TrimSpace(b.password) == "" {
		return errors.New("password is empty")
	}

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", b.username, b.password)))

	request.Headers.Add("Authorization", fmt.Sprintf("Basic %s", auth))

	return nil
}
