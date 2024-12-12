package authentication

import (
	"context"
	"net/url"

	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/yosida95/uritemplate/v3"
	"golang.org/x/oauth2"
)

type ResourceOwnerPasswordTokenProvider struct {
	config      *oauth2.Config
	username    string
	password    string
	tokenSource oauth2.TokenSource
}

func NewResourceOwnerPasswordTokenProvider(
	clientID,
	clientSecret,
	username,
	password string,
	scopes []string,
) *ResourceOwnerPasswordTokenProvider {
	return &ResourceOwnerPasswordTokenProvider{
		config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint: oauth2.Endpoint{
				AuthStyle: oauth2.AuthStyleInParams,
			},
			Scopes: scopes,
		},
		username: username,
		password: password,
	}
}

func (provider *ResourceOwnerPasswordTokenProvider) GetAuthorizationToken(ctx context.Context, uri *url.URL, _ map[string]interface{}) (string, error) {
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

	// Token is either nil or expired; fetch a new one
	token, err := provider.config.PasswordCredentialsToken(ctx, provider.username, provider.password)
	if err != nil {
		return "", err
	}

	provider.tokenSource = provider.config.TokenSource(ctx, token)

	return token.AccessToken, nil
}

func (provider *ResourceOwnerPasswordTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}
