package authentication

import (
	"context"
	"fmt"
	"net/url"

	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/yosida95/uritemplate/v3"
	"golang.org/x/oauth2"
)

type AuthorizationCodeTokenProvider struct {
	config      *oauth2.Config
	tokenSource oauth2.TokenSource
	uriCallback func(string) error
}

func NewAuthorizationCodeTokenProvider(
	clientID,
	clientSecret string,
	port int,
	scopes []string,
	uriCallback func(string) error,
) *AuthorizationCodeTokenProvider {
	return &AuthorizationCodeTokenProvider{
		config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint: oauth2.Endpoint{
				AuthStyle: oauth2.AuthStyleInParams,
			},
			RedirectURL: fmt.Sprintf("localhost:%v", port),
			Scopes:      scopes,
		},
		uriCallback: uriCallback,
	}
}

func (provider *AuthorizationCodeTokenProvider) GetAuthorizationToken(ctx context.Context, uri *url.URL, _ map[string]interface{}) (string, error) {
	if provider == nil {
		return "", nil
	}

	if provider.tokenSource != nil {
		token, err := provider.tokenSource.Token()
		if err != nil {
			return "", err
		}

		return token.AccessToken, nil
	}

	uri.Path, uri.Fragment, uri.RawQuery = "", "", ""
	baseURL := uri.String()

	values := uritemplate.Values{}
	values.Set("baseurl", uritemplate.String(baseURL))

	authURL, err := authURLTemplate.Expand(values)
	if err != nil {
		return "", err
	}
	tokenURL, err := tokenURLTemplate.Expand(values)
	if err != nil {
		return "", err
	}
	provider.config.Endpoint.AuthURL = authURL
	provider.config.Endpoint.TokenURL = tokenURL

	state := ""

	// Token is either nil or expired; fetch a new one
	codeURL := provider.config.AuthCodeURL(state)

	server, err := NewAuthenticationCodeRedirectServer(state, 5000)
	if err != nil {
		return "", err
	}
	defer server.Shutdown()

	if err != provider.uriCallback(codeURL) {
		return "", err
	}

	result := server.Result(ctx)
	if result.Err != nil {
		return "", result.Err
	}

	token, err := provider.config.Exchange(ctx, result.Code)
	if err != nil {
		return "", err
	}

	provider.tokenSource = provider.config.TokenSource(ctx, token)

	return token.AccessToken, nil
}

func (provider *AuthorizationCodeTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}
