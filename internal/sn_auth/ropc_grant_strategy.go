package snauth

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

const (
	contentTypeHeaderKey      = "Content-Type"
	formURLEncodedContentType = "application/x-www-form-urlencoded"

	grantTypeKey          = "grant_type"
	grantTypePassword     = "password"
	grantTypeRefreshToken = "refresh_token"
	responseTypeCode      = "code"

	responseTypeKey = "response_type"
	stateKey        = "state"
	redirectURIKey  = "redirect_uri"

	clientIDKey     = "client_id"
	clientSecretKey = "client_secret"
	refreshTokenKey = "refresh_token"
	usernameKey     = "username"
	passwordKey     = "password"
)

type TokenDecoder interface {
	DecodeToken(resp *http.Response) (oauth2.Oauth2Token, error)
}

var _ oauth2.GrantStrategy[ROPCTokenConfig] = (*ROPCTokenStrategy)(nil)

type CredentialsProvider interface {
	GetClientID() string
	GetClientSecret() string
	GetUsername() string
	GetPassword() string
}

// ROPCTokenStrategy Resource Owner Password Credentials (ROPC) grant, which allows an application to sign in the user by directly handling their password.
type ROPCTokenStrategy struct {
	client   *http.Client
	decoder  TokenDecoder
	tokenURL string
}

func NewROPCTokenStrategy(client *http.Client, decoder TokenDecoder, tokenURL string) *ROPCTokenStrategy {
	return &ROPCTokenStrategy{
		client:   client,
		decoder:  decoder,
		tokenURL: tokenURL,
	}
}

func (tS *ROPCTokenStrategy) AcquireToken(ctx context.Context, opts ...oauth2.TokenOption[ROPCTokenConfig]) (oauth2.Oauth2Token, error) {

	var config ROPCTokenConfig

	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, err
		}
	}

	// TODO: Add check for if refresh token is empty if not, refresh
	token, err := tS.fetchToken(ctx, config)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (tS *ROPCTokenStrategy) fetchToken(ctx context.Context, config ROPCTokenConfig) (oauth2.Oauth2Token, error) {

	//TODO: if config.username || config.password is empty, retrieve them

	data := url.Values{}
	data.Set(grantTypeKey, grantTypePassword)
	data.Set(clientIDKey, config.clientID)
	data.Set(clientSecretKey, config.clientSecret)
	data.Set(usernameKey, config.username)
	data.Set(passwordKey, config.password)

	return tS.getToken(data, ctx)
}

func (tS *ROPCTokenStrategy) refreshToken(ctx context.Context, config ROPCTokenConfig) (oauth2.Oauth2Token, error) {
	data := url.Values{}
	data.Set(grantTypeKey, grantTypeRefreshToken)
	data.Set(clientIDKey, config.clientID)
	data.Set(clientSecretKey, config.clientSecret)
	data.Set(refreshTokenKey, config.refreshToken)

	return tS.getToken(data, ctx)
}

func (tS *ROPCTokenStrategy) getToken(values url.Values, ctx context.Context) (oauth2.Oauth2Token, error) {
	var err error

	req, err := http.NewRequestWithContext(ctx, "POST", tS.tokenURL, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add(contentTypeHeaderKey, formURLEncodedContentType)

	resp, err := tS.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(bodyBytes))
	}

	token, err := tS.decoder.DecodeToken(resp)
	if err != nil {
		return nil, err
	}

	return token, nil
}
