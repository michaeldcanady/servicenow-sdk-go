package authentication

import (
	"context"
	"errors"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

var _ authentication.AccessTokenProvider = (*authorizationCodeTokenProvider)(nil)

type authorizationCodeTokenProvider struct {
	requestAdapter abstractions.RequestAdapter
	uriCallback    func(string) error
	port           int
	clientID       string
	clientSecret   string
	authCodeFlow   *authorizationCodeFlow
}

func newAuthorizationCodeTokenProvider(uriCallback func(string) error, port int, clientID string, clientSecret string) *authorizationCodeTokenProvider {
	return &authorizationCodeTokenProvider{
		requestAdapter: newRequestAdapter(),
		uriCallback:    uriCallback,
		port:           port,
		clientID:       clientID,
		clientSecret:   clientSecret,
	}
}

// GetAuthorizationToken returns the access token for the provided url.
func (provider *authorizationCodeTokenProvider) GetAuthorizationToken(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if internal.IsNil(provider) {
		return "", errors.New("provider is nil")
	}

	record, err := provider.authCodeFlow.AcquireAuthRecord(ctx, uri, additionalAuthenticationContext)
	if err != nil {
		return "", err
	}
	//TODO: cache access token, expiration time frame, and refresh token
	accessToken, err := record.GetAccessToken()
	if err != nil {
		return "", nil
	}

	return *accessToken, nil
}

// GetAllowedHostsValidator returns the hosts validator.
func (provider *authorizationCodeTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}
