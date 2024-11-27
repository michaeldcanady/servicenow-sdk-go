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

func (provider *authorizationCodeTokenProvider) sendRequest(ctx context.Context, pathParameters map[string]string, code, redirectURI string) (string, error) {
	if internal.IsNil(provider) {
		return "", errors.New("provider is nil")
	}

	grantType := "authorization_code"

	body := newAuthorizationCodeRequest()
	if err := body.SetClientID(&provider.clientID); err != nil {
		return "", err
	}
	if err := body.setGrantType(&grantType); err != nil {
		return "", err
	}
	if err := body.SetClientSecret(&provider.clientSecret); err != nil {
		return "", err
	}
	if err := body.SetCode(&code); err != nil {
		return "", err
	}
	if err := body.SetRedirectURI(&redirectURI); err != nil {
		return "", err
	}

	builder := NewOauthTokenRequestBuilderInternal(pathParameters, provider.requestAdapter)
	resp, err := builder.Post(ctx, body, nil)
	if err != nil {
		return "", err
	}

	token, err := resp.GetAccessToken()
	if err != nil {
		return "", err
	}

	if token == nil || *token == "" {
		return "", errors.New("access token is empty")
	}

	return *token, nil
}

// GetAuthorizationToken returns the access token for the provided url.
func (provider *authorizationCodeTokenProvider) GetAuthorizationToken(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if internal.IsNil(provider) {
		return "", errors.New("provider is nil")
	}

	// Clear path and fragment
	uri.Path, uri.Fragment = "", ""
	stringURL := uri.String()

	pathParameters := map[string]string{
		"baseurl": stringURL,
	}

	state, err := randomString(5)
	if err != nil {
		return "", err
	}

	server, err := NewAuthenticationCodeRedirectServer(state, provider.port)
	if err != nil {
		return "", err
	}
	defer server.Shutdown()

	params := &authorizationCodeQueryParameters{
		responseType: "token",
		redirectURI:  server.Addr,
		clientID:     provider.clientID,
		state:        state,
	}

	stringURI, err := buildOauthURL(pathParameters, params)
	if err != nil {
		return "", err
	}

	if err := provider.uriCallback(stringURI); err != nil {
		return "", err
	}

	result := server.Result(ctx)
	if result.Err != nil {
		return "", result.Err
	}

	return provider.sendRequest(ctx, pathParameters, result.Code, server.Addr)
}

// GetAllowedHostsValidator returns the hosts validator.
func (provider *authorizationCodeTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}
