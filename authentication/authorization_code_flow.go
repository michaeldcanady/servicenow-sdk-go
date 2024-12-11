package authentication

import (
	"context"
	"errors"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type authorizationCodeFlow struct {
	clientID       string
	redirectURI    string
	port           int
	uriCallback    func(string) error
	requestAdapter abstractions.RequestAdapter
	clientSecret   string
}

func (flow *authorizationCodeFlow) acquireToken(ctx context.Context, baseURL, code string) (authenticationTokenResponsable, error) {
	if internal.IsNil(flow) {
		return nil, errors.New("provider is nil")
	}

	grantType := grantTypeAuthorizationCode

	body := newAuthorizationCodeRequest()
	if err := body.SetClientID(&flow.clientID); err != nil {
		return nil, err
	}
	if err := body.setGrantType(&grantType); err != nil {
		return nil, err
	}
	if err := body.SetClientSecret(&flow.clientSecret); err != nil {
		return nil, err
	}
	if err := body.SetCode(&code); err != nil {
		return nil, err
	}
	if err := body.SetRedirectURI(&flow.redirectURI); err != nil {
		return nil, err
	}

	pathParameters := map[string]string{
		"baseurl": baseURL,
	}

	builder := NewOauthTokenRequestBuilderInternal(pathParameters, flow.requestAdapter)
	resp, err := builder.Post(ctx, body, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (flow *authorizationCodeFlow) acquireAuthorizationCode(ctx context.Context, baseURL string) (string, error) {
	if internal.IsNil(flow) {
		return "", errors.New("provider is nil")
	}

	pathParameters := map[string]string{
		"baseurl": baseURL,
	}

	state, err := randomString(5)
	if err != nil {
		return "", err
	}

	server, err := NewAuthenticationCodeRedirectServer(state, flow.port)
	if err != nil {
		return "", err
	}
	defer server.Shutdown()
	flow.redirectURI = server.Addr

	params := &oauthAuthQueryParameters{
		responseType: responseTypeCode,
		redirectURI:  flow.redirectURI,
		clientID:     flow.clientID,
		state:        state,
	}

	oauthURL, err := buildOauthURL(pathParameters, params)
	if err != nil {
		return "", err
	}

	if err := flow.uriCallback(oauthURL); err != nil {
		return "", err
	}

	result := server.Result(ctx)
	if result.Err != nil {
		return "", result.Err
	}

	return result.Code, nil
}

func (flow *authorizationCodeFlow) AcquireAuthRecord(ctx context.Context, uri *url.URL, _ map[string]interface{}) (authenticationTokenResponsable, error) {
	if internal.IsNil(flow) {
		return nil, errors.New("provider is nil")
	}

	// Clear path and fragment
	uri.Path, uri.Fragment = "", ""
	baseURL := uri.String()

	code, err := flow.acquireAuthorizationCode(ctx, baseURL)
	if err != nil {
		return nil, err
	}

	return flow.acquireToken(ctx, baseURL, code)
}
