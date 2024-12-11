package authentication

import (
	"context"
	"errors"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type jwtBearerFlow struct {
	clientID       string
	tokenProvider  authentication.AccessTokenProvider
	requestAdapter abstractions.RequestAdapter
	clientSecret   string
}

func newJWTBearerFlow(tokenProvider authentication.AccessTokenProvider, requestAdapter abstractions.RequestAdapter, clientID, clientSecret string) *jwtBearerFlow {
	return &jwtBearerFlow{
		clientID:       clientID,
		tokenProvider:  tokenProvider,
		requestAdapter: requestAdapter,
		clientSecret:   clientSecret,
	}
}

func (flow *jwtBearerFlow) acquireTokenFromToken(ctx context.Context, uri *url.URL, _ map[string]interface{}, token string) (authenticationTokenResponsable, error) {
	if internal.IsNil(flow) {
		return nil, errors.New("provider is nil")
	}

	// Clear path and fragment
	uri.Path, uri.Fragment = "", ""
	baseURL := uri.String()

	request := newJWTBearerRequest()
	grantType := grantTypeJWTBearer
	if err := request.setGrantType(&grantType); err != nil {
		return nil, err
	}
	if err := request.SetAssertion(&token); err != nil {
		return nil, err
	}
	if flow.clientID != "" {
		if err := request.SetClientID(&flow.clientID); err != nil {
			return nil, err
		}
	}
	if flow.clientSecret != "" {
		if err := request.SetClientSecret(&flow.clientSecret); err != nil {
			return nil, err
		}
	}

	pathParameters := map[string]string{
		"baseurl": baseURL,
	}

	builder := NewOauthTokenRequestBuilderInternal(pathParameters, flow.requestAdapter)
	resp, err := builder.Post(ctx, request, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (flow *jwtBearerFlow) acquireExternalJWTToken(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (string, error) {
	if internal.IsNil(flow) {
		return "", errors.New("provider is nil")
	}

	if flow.tokenProvider == nil {
		return "", errors.New("tokenProvider is nil")
	}

	return flow.tokenProvider.GetAuthorizationToken(ctx, uri, additionalAuthenticationContext)
}

func (flow *jwtBearerFlow) AcquireAuthRecord(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (authenticationTokenResponsable, error) {
	if internal.IsNil(flow) {
		return nil, errors.New("provider is nil")
	}

	token, err := flow.acquireExternalJWTToken(ctx, uri, additionalAuthenticationContext)
	if err != nil {
		return nil, err
	}

	return flow.acquireTokenFromToken(ctx, uri, additionalAuthenticationContext, token)
}
