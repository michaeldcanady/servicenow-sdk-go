package auth

import (
	"context"
	"net/url"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type InteractiveBrowserProvider struct {
	clientID       string
	clientSecret   string
	accessToken    string
	refreshToken   string
	expiresAt      time.Time
	RequestAdapter abstractions.RequestAdapter
}

func (tokenProvider *InteractiveBrowserProvider) GetAuthorizationToken(ctx context.Context, url *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if tokenProvider == nil {
		return "", nil
	}

	if tokenProvider.accessToken == "" || tokenProvider.expiresAt.Before(time.Now()) {
		var err error
		if tokenProvider.accessToken == "" {

		} else {
			err = tokenProvider.acquireTokenByRefreshToken(ctx, url)
		}

		if err != nil {
			return "", err
		}
	}
	return tokenProvider.accessToken, nil
}

func (tokenProvider *InteractiveBrowserProvider) acquireTokenByInteractive(ctx context.Context, url *url.URL) error {

	// TODO: Start server with redirect URL

	pathParameters := map[string]string{"baseurl": url.Host}
	builder := NewOauthAuthRequestBuilderInternal(pathParameters, tokenProvider.RequestAdapter)
	builder.ToGetRequestInformation(ctx, &OauthAuthRequestBuilderGetRequestConfiguration{})

	// TODO: Open Browser

	return nil
}

func (tokenProvider *InteractiveBrowserProvider) acquireTokenByRefreshToken(ctx context.Context, url *url.URL) error {
	body := NewRefreshTokenAuthenticate()

	if err := body.SetGrantType(pointer(GrantTypeRefreshToken)); err != nil {
		return err
	}

	if err := body.SetRefreshToken(&tokenProvider.refreshToken); err != nil {
		return err
	}

	if err := body.SetClientID(&tokenProvider.clientID); err != nil {
		return err
	}

	return tokenProvider.requestToken(ctx, url, body)
}

func (tokenProvider *InteractiveBrowserProvider) requestToken(ctx context.Context, url *url.URL, body Authenticatable) error {
	pathParameters := map[string]string{"baseurl": url.Host}
	builder := NewOauthTokenRequestBuilderInternal(pathParameters, tokenProvider.RequestAdapter)
	token, err := builder.Post(ctx, body, nil)
	if err != nil {
		return err
	}
	return tokenProvider.updateToken(token)
}

func (tokenProvider *InteractiveBrowserProvider) updateToken(token AccessTokenable) error {
	startTime := time.Now()
	rawAccessToken, err := token.GetAccessToken()
	if err != nil {
		return err
	}
	expiresIn, err := token.GetExpiresIn()
	if err != nil {
		return err
	}
	refreshToken, err := token.GetRefreshToken()
	if err != nil {
		return err
	}
	tokenProvider.accessToken = *rawAccessToken
	tokenProvider.expiresAt = startTime.Add(time.Duration(*expiresIn) * time.Second)
	tokenProvider.refreshToken = *refreshToken
	return nil
}

func (tokenProvider *InteractiveBrowserProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	if internal.IsNil(tokenProvider) {
		return nil
	}
	return nil
}
