package credentials

import (
	"context"
	"crypto/rand"
	"math/big"

	u "net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/auth"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	nethttplibrary "github.com/microsoft/kiota-http-go"
	"github.com/pkg/browser"
)

var _ authentication.AccessTokenProvider = (*ImplicitGrantFlowTokenProvider)(nil)

type ImplicitGrantFlowTokenProvider struct {
	redirectURI          string
	clientID             string
	allowedHostValidator *authentication.AllowedHostsValidator
	requestAdapter       abstractions.RequestAdapter
}

func NewImplicitGrantFlowTokenProvider(
	redirectURI string,
	clientID string,
) *ImplicitGrantFlowTokenProvider {

	requestAdapter, _ := nethttplibrary.NewNetHttpRequestAdapter(&authentication.AnonymousAuthenticationProvider{})

	return &ImplicitGrantFlowTokenProvider{
		redirectURI:          redirectURI,
		clientID:             clientID,
		allowedHostValidator: nil,
		requestAdapter:       requestAdapter,
	}
}

func generateState() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 4
	b := make([]byte, length)
	for i := range b {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[index.Int64()]
	}
	return string(b), nil
}

// GetAuthorizationToken returns the access token for the provided url.
func (tP *ImplicitGrantFlowTokenProvider) GetAuthorizationToken(context context.Context, url *u.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	host := url.Host

	pathParameters := map[string]string{
		"baseurl": host,
	}

	state, err := generateState()
	if err != nil {
		return "", err
	}

	builder := auth.NewOauthAuthRequestBuilderInternal(pathParameters, tP.requestAdapter)
	config := &auth.OauthAuthRequestBuilderGetRequestConfiguration{
		QueryParameters: &auth.OauthAuthRequestBuilderGetQueryParameters{
			ResponseType: auth.ResponseTypeToken,
			RedirectURI:  tP.redirectURI,
			ClientID:     tP.clientID,
			State:        state,
		},
	}

	info, err := builder.ToGetRequestInformation(context, config)
	if err != nil {
		return "", err
	}

	uri, err := info.GetUri()
	if err != nil {
		return "", err
	}

	server, _ := auth.New(state, 8000)

	if err := browser.OpenURL(uri.String()); err != nil {
		return "", err
	}

	params := server.Result(context)

	return params.Code, nil
}

// GetAllowedHostsValidator returns the hosts validator.
func (tP *ImplicitGrantFlowTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return &authentication.AllowedHostsValidator{}
}
