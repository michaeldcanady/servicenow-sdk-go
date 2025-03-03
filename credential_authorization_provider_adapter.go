package servicenowsdkgo

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

const (
	authorizationHeader = "Authorization"
)

var _ authentication.AuthenticationProvider = (*credentialAuthenticationProviderAdapter)(nil)

// credentialAuthenticationProviderAdapter adapter for core.Credential to kiota AuthenticationProvider
type credentialAuthenticationProviderAdapter struct {
	cred core.Credential //nolint: staticcheck
}

// newCredentialAuthenticationProviderAdapter adapts provided credential to an AuthenticationProvider
func newCredentialAuthenticationProviderAdapter(credential core.Credential) (*credentialAuthenticationProviderAdapter, error) {
	if internal.IsNil(credential) {
		return nil, errors.New("credential is nil")
	}

	return &credentialAuthenticationProviderAdapter{
		cred: credential,
	}, nil
}

// AuthenticateRequest authenticates the provided RequestInformation.
func (provider *credentialAuthenticationProviderAdapter) AuthenticateRequest(ctx context.Context, request *abstractions.RequestInformation, _ map[string]interface{}) error {
	if request == nil {
		return errors.New("request is nil")
	}
	if request.Headers == nil {
		request.Headers = abstractions.NewRequestHeaders()
	}
	if !request.Headers.ContainsKey(authorizationHeader) {
		authString, err := provider.cred.GetAuthentication()
		if err != nil {
			return err
		}
		if authString != "" {
			request.Headers.Add(authorizationHeader, authString)
		}
	}

	return nil
}
