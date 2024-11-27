package authentication

import (
	"context"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

var _ authentication.AccessTokenProvider = (*resourceOwnerPasswordTokenProvider)(nil)

type resourceOwnerPasswordTokenProvider struct {
	requestAdapter abstractions.RequestAdapter
	clientID       string
	clientSecret   string
	username       string
	password       string
}

func newResourceOwnerPasswordTokenProvider(clientID, clientSecret, username, password string) *resourceOwnerPasswordTokenProvider {
	return &resourceOwnerPasswordTokenProvider{
		requestAdapter: newRequestAdapter(),
		clientID:       clientID,
		clientSecret:   clientSecret,
		username:       username,
		password:       password,
	}
}

// GetAuthorizationToken returns the access token for the provided url.
func (provider *resourceOwnerPasswordTokenProvider) GetAuthorizationToken(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if internal.IsNil(provider) {
		return "", nil
	}

	pathParameters := map[string]string{}

	builder := NewOauthTokenRequestBuilderInternal(pathParameters, provider.requestAdapter)

	body := newResourceOwnerRequest()
	grantType := "password"
	if err := body.setGrantType(&grantType); err != nil {
		return "", err
	}
	if err := body.SetClientID(&provider.clientID); err != nil {
		return "", err
	}
	if err := body.SetClientSecret(&provider.clientSecret); err != nil {
		return "", err
	}
	if err := body.SetPassword(&provider.password); err != nil {
		return "", err
	}
	if err := body.SetUsername(&provider.username); err != nil {
		return "", err
	}

	resp, err := builder.Post(ctx, body, nil)
	if err != nil {
		return "", err
	}

	// TODO: add caching for access token
	// TODO: add caching for refresh token
	token, err := resp.GetAccessToken()
	if err != nil {
		return "", err
	}

	if token == nil || *token == "" {
		return "", nil
	}

	return *token, nil
}

// GetAllowedHostsValidator returns the hosts validator.
func (provider *resourceOwnerPasswordTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}
