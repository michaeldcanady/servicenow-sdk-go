package authentication

import (
	"context"
	"errors"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

var _ authentication.AccessTokenProvider = (*externalJWTTokenProvider)(nil)

type externalJWTTokenProvider struct {
	tokenProvider  authentication.AccessTokenProvider
	requestAdapter abstractions.RequestAdapter
	clientID       *string
	clientSecret   *string
}

func newExternalJWTTokenProvider(tokenProvider authentication.AccessTokenProvider, clientID, clientSecret string) *externalJWTTokenProvider {
	return &externalJWTTokenProvider{
		tokenProvider:  tokenProvider,
		requestAdapter: newRequestAdapter(),
		clientID:       &clientID,
		clientSecret:   &clientSecret,
	}
}

func (provider *externalJWTTokenProvider) buildRequest(token string) (externalJWTRequestable, error) {
	if internal.IsNil(provider) {
		return nil, errors.New("provider is nil")
	}

	request := newExternalJWTRequest()
	grantType := "urn:ietf:params:oauth:grant-type:jwt-bearer"
	if err := request.setGrantType(&grantType); err != nil {
		return nil, err
	}
	if err := request.SetAssertion(&token); err != nil {
		return nil, err
	}
	if provider.clientID != nil || *provider.clientID != "" {
		if err := request.SetClientID(provider.clientID); err != nil {
			return nil, err
		}
	}
	if provider.clientSecret != nil || *provider.clientSecret != "" {
		if err := request.SetClientSecret(provider.clientSecret); err != nil {
			return nil, err
		}
	}

	return request, nil
}

func (provider *externalJWTTokenProvider) sendRequest(context context.Context, request externalJWTRequestable, uri *url.URL) (string, error) {
	if internal.IsNil(provider) {
		return "", errors.New("provider is nil")
	}

	uri.Path = ""
	uri.Fragment = ""
	uri.Path = ""
	stringURL := uri.String()

	pathParameters := map[string]string{
		"baseurl": stringURL,
	}

	requestBuilder := NewOauthTokenRequestBuilderInternal(pathParameters, provider.requestAdapter)
	resp, err := requestBuilder.Post(context, request, nil)
	if err != nil {
		return "", err
	}

	accessToken, err := resp.GetAccessToken()
	if err != nil {
		return "", err
	}

	if accessToken == nil || *accessToken == "" {
		return "", nil
	}

	return *accessToken, nil
}

// GetAuthorizationToken returns the access token for the provided url.
func (provider *externalJWTTokenProvider) GetAuthorizationToken(context context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if internal.IsNil(provider) {
		return "", errors.New("provider is nil")
	}

	token, err := provider.tokenProvider.GetAuthorizationToken(context, uri, additionalAuthenticationContext)
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", nil
	}

	request, err := provider.buildRequest(token)
	if err != nil {
		return "", err
	}

	return provider.sendRequest(context, request, uri)
}

// GetAllowedHostsValidator returns the hosts validator.
func (provider *externalJWTTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}
