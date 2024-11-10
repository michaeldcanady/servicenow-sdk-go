package auth

import (
	"context"
	"net/url"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

func pointer[T interface{}](val T) *T {
	return &val
}

var _ authentication.AccessTokenProvider = (*ROPCProvider)(nil)

type ROPCProvider struct {
	username       string
	password       string
	clientID       string
	clientSecret   string
	accessToken    string
	refreshToken   string
	expiresAt      time.Time
	RequestAdapter abstractions.RequestAdapter
}

func NewROPCProvider(clientID, clientSecret string) *ROPCProvider {
	return &ROPCProvider{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (tokenProvider *ROPCProvider) GetAuthorizationToken(ctx context.Context, url *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if tokenProvider == nil {
		return "", nil
	}

	if tokenProvider.accessToken == "" || tokenProvider.expiresAt.Before(time.Now()) {
		var err error
		if tokenProvider.accessToken == "" {
			err = tokenProvider.acquireTokenByPassword(ctx, url)
		} else {
			err = tokenProvider.acquireTokenByRefreshToken(ctx, url)
		}

		if err != nil {
			return "", err
		}
	}
	return tokenProvider.accessToken, nil
}

func (tokenProvider *ROPCProvider) acquireTokenByPassword(ctx context.Context, url *url.URL) error {
	body := NewPasswordAuthenticate()
	if err := body.SetGrantType(pointer(GrantTypePassword)); err != nil {
		return err
	}

	if err := body.SetUsername(&tokenProvider.username); err != nil {
		return err
	}

	if err := body.SetPassword(&tokenProvider.password); err != nil {
		return err
	}

	if err := body.SetClientID(&tokenProvider.clientID); err != nil {
		return err
	}

	if err := body.SetClientSecret(&tokenProvider.clientSecret); err != nil {
		return err
	}
	return tokenProvider.requestToken(ctx, url, body)
}

func (tokenProvider *ROPCProvider) acquireTokenByRefreshToken(ctx context.Context, url *url.URL) error {
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

	if err := body.SetClientSecret(&tokenProvider.clientSecret); err != nil {
		return err
	}
	return tokenProvider.requestToken(ctx, url, body)
}

func (tokenProvider *ROPCProvider) requestToken(ctx context.Context, url *url.URL, body Authenticatable) error {
	pathParameters := map[string]string{"baseurl": url.Host}
	builder := NewOauthTokenRequestBuilderInternal(pathParameters, tokenProvider.RequestAdapter)
	token, err := builder.Post(ctx, body, nil)
	if err != nil {
		return err
	}
	return tokenProvider.updateToken(token)
}

func (tokenProvider *ROPCProvider) updateToken(token AccessTokenable) error {
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

func (tokenProvider *ROPCProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	if internal.IsNil(tokenProvider) {
		return nil
	}
	return nil
}
