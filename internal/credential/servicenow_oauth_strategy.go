package credential

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

const (
	contentTypeHeaderKey      = "Content-Type"
	formURLEncodedContentType = "application/x-www-form-urlencoded"
	grantTypeKey              = "grant_type"
	grantTypePassword         = "password"
	grantTypeRefreshToken     = "refresh_token"
	clientIDKey               = "client_id"
	clientSecretKey           = "client_secret"
	refreshTokenKey           = "refresh_token"
	usernameKey               = "username"
	passwordKey               = "password"
)

type CredentialsProvider interface {
	GetClientID() string
	GetClientSecret() string
	GetUsername() string
	GetPassword() string
}

type ServiceNowOauth2Strategy struct {
	server      oauth2.Server
	client      *http.Client
	credentials CredentialsProvider
	baseURL     string
}

// ServiceNowOauth2Option function type
type ServiceNowOauth2Option func(*ServiceNowOauth2Strategy)

// WithClient sets a custom HTTP client
func WithClient(client *http.Client) ServiceNowOauth2Option {
	return func(s *ServiceNowOauth2Strategy) {
		s.client = client
	}
}

func NewServiceNowOauth2Strategy(creds CredentialsProvider, baseURL string, opts ...ServiceNowOauth2Option) (*ServiceNowOauth2Strategy, error) {
	// TODO: Make an option
	server, err := newServiceNowOauthServer("http://localhost")
	if err != nil {
		return nil, err
	}

	if core.IsNil(creds) {
		return nil, errors.New("creds can't be nil")
	}

	strategy := &ServiceNowOauth2Strategy{
		server:      server,
		client:      http.DefaultClient, // Default to http.DefaultClient
		credentials: creds,
		baseURL:     baseURL,
	}

	for _, opt := range opts {
		opt(strategy)
	}

	return strategy, nil
}

func (tS *ServiceNowOauth2Strategy) FetchToken(ctx context.Context) (oauth2.Oauth2Token, error) {
	if core.IsNil(tS) || core.IsNil(tS.credentials) {
		return nil, errors.New("CredentialsProvider or stategy is nil")
	}

	data := url.Values{}
	data.Set(grantTypeKey, grantTypePassword)
	data.Set(clientIDKey, tS.credentials.GetClientID())
	data.Set(clientSecretKey, tS.credentials.GetClientSecret())
	data.Set(usernameKey, tS.credentials.GetUsername())
	data.Set(passwordKey, tS.credentials.GetPassword())

	return tS.getToken(data, ctx)
}

func (tS *ServiceNowOauth2Strategy) RefreshToken(refreshToken string, ctx context.Context) (oauth2.Oauth2Token, error) {
	data := url.Values{}
	data.Set(grantTypeKey, grantTypeRefreshToken)
	data.Set(clientIDKey, tS.credentials.GetClientID())
	data.Set(clientSecretKey, tS.credentials.GetClientSecret())
	data.Set(refreshTokenKey, refreshToken)

	return tS.getToken(data, ctx)
}

func (tS *ServiceNowOauth2Strategy) getToken(values url.Values, ctx context.Context) (oauth2.Oauth2Token, error) {
	var err error
	go func() {
		// if directly made into 'err' can hide another error
		tempErr := tS.server.ListenAndServe()
		if tempErr != nil {
			err = tempErr
		}
	}()
	if err != nil {
		return nil, err
	}
	defer func() {
		// if directly made into 'err' can hide another error
		tempErr := tS.server.Shutdown(ctx)
		if tempErr != nil {
			err = tempErr
		}
	}()

	oauthURL := fmt.Sprintf("https://%s%s", tS.baseURL, snOauthTokenURITemplate)
	req, err := http.NewRequest("POST", oauthURL, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add(contentTypeHeaderKey, formURLEncodedContentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(bodyBytes))
	}
	defer resp.Body.Close()

	token, err := tS.decodeToken(resp)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (tS *ServiceNowOauth2Strategy) decodeToken(response *http.Response) (*snAccessToken, error) {
	defer response.Body.Close()
	var accessToken snAccessToken
	if err := json.NewDecoder(response.Body).Decode(&accessToken); err != nil {
		return nil, err
	}

	accessToken.ExpiresAt = time.Now().Add(time.Duration(accessToken.ExpiresIn) * time.Second)
	return &accessToken, nil
}
