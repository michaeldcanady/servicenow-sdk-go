package credentials

import (
	"context"
	"errors"
	"net/url"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type CredentialAdapter struct {
	authenticationProvider authentication.AuthenticationProvider
	baseURL                string
}

func NewCredentialAdapter(
	authenticationProvider authentication.AuthenticationProvider,
	baseURL string,
) *CredentialAdapter {
	return &CredentialAdapter{
		authenticationProvider: authenticationProvider,
		baseURL:                baseURL,
	}
}

func (cA *CredentialAdapter) GetAuthentication() (string, error) {
	requestInformation := abstractions.NewRequestInformation()

	if _, ok := cA.authenticationProvider.(*authentication.BaseBearerTokenAuthenticationProvider); ok && cA.baseURL == "" {
		return "", errors.New("baseURL is required for oauth2 authentication")
	}

	if cA.baseURL != "" {
		uri, err := url.Parse(cA.baseURL)
		if err != nil {
			return "", err
		}
		requestInformation.SetUri(*uri)
	}
	if err := cA.authenticationProvider.AuthenticateRequest(context.Background(), requestInformation, nil); err != nil {
		return "", err
	}

	headers := requestInformation.Headers.Get("Authorization")

	return headers[0], nil
}
