package credential

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/auth"
)

var _ auth.AuthorizationStrategy = (*BasicAuthStrategy)(nil)

const basicAuthenticationType = "Basic"

type BasicAuthStrategy struct {
	username *string
	password *string
}

func NewBasicAuthenticationStrategy(strategy auth.AuthorizationStrategy) auth.AuthenticationProvider {
	return auth.NewAuthenticationProvider(
		auth.NewStaticAuthTypeProvider(basicAuthenticationType),
		strategy,
	)
}

func NewInteractiveBasicCredential() *BasicAuthStrategy {
	return &BasicAuthStrategy{
		username: nil,
		password: nil,
	}
}

func NewNonInteractiveCredential(username string, password string) *BasicAuthStrategy {
	return &BasicAuthStrategy{
		username: &username,
		password: &password,
	}
}

func (cP *BasicAuthStrategy) GetAuth(ctx context.Context) (string, error) {
	if cP.username == nil || cP.password == nil {
		return "", errors.New("username or password is nil")
	}

	auth := fmt.Sprintf("%s:%s", *cP.username, *cP.password)
	return base64.StdEncoding.EncodeToString([]byte(auth)), nil
}