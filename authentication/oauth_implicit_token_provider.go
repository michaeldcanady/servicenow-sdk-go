package authentication

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	stduritemplate "github.com/std-uritemplate/std-uritemplate/go"
)

const (
	oauthAuthURLTemplate = "{+baseurl}/oauth_auth.do{?response_type,redirect_uri,client_id,state}"
)

func buildOauthURL(pathParameters map[string]string, queryParameters *authorizationCodeQueryParameters) (string, error) {
	substitutions := make(map[string]any)

	for key, value := range pathParameters {
		substitutions[key] = any(value)
	}

	if !internal.IsNil(queryParameters) {
		queryValues, err := query.Values(queryParameters)
		if err != nil {
			return "", err
		}
		for key, values := range queryValues {
			substitutions[key] = strings.Join(values, ",")
		}
	}

	return stduritemplate.Expand(oauthAuthURLTemplate, substitutions)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) (string, error) {
	b := make([]byte, n)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		b[i] = letterBytes[num.Int64()]
	}
	return string(b), nil
}

var _ authentication.AccessTokenProvider = (*implicitTokenProvider)(nil)

type implicitTokenProvider struct {
	uriCallback func(string) error
	port        int
	clientID    string
}

func newImplicitTokenProvider(clientID string, port int, uriCallback func(string) error) *implicitTokenProvider {
	return &implicitTokenProvider{
		uriCallback: uriCallback,
		port:        port,
		clientID:    clientID,
	}
}

// GetAuthorizationToken returns the access token for the provided url.
func (provider *implicitTokenProvider) GetAuthorizationToken(context context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
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

	state, err := randomString(5)
	if err != nil {
		return "", err
	}

	server, err := NewAuthenticationTokenRedirectServer(state, provider.port)
	if err != nil {
		server.Shutdown()
		return "", err
	}

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

	result := server.Result(context)
	if result.Err != nil {
		return "", result.Err
	}

	return result.AccessToken, nil
}

// GetAllowedHostsValidator returns the hosts validator.
func (provider *implicitTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}
