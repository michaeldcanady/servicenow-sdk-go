package authentication

import (
	"context"
	"errors"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type refreshTokenFlow struct {
	clientID       string
	requestAdapter abstractions.RequestAdapter
	clientSecret   string
}

func newRefreshTokenFlow(requestAdapter abstractions.RequestAdapter, clientID, clientSecret string) *refreshTokenFlow {
	return &refreshTokenFlow{
		clientID:       clientID,
		clientSecret:   clientSecret,
		requestAdapter: requestAdapter,
	}
}

func (flow *refreshTokenFlow) AcquireAuthRecord(ctx context.Context, uri *url.URL, _ map[string]interface{}, token string) (authenticationTokenResponsable, error) {
	if internal.IsNil(flow) {
		return nil, errors.New("provider is nil")
	}

	// Clear path and fragment
	uri.Path, uri.Fragment = "", ""
	baseURL := uri.String()

	grantType := grantTypeRefreshToken
	request := newRefreshTokenRequest()
	if err := request.setGrantType(&grantType); err != nil {
		return nil, err
	}
	if err := request.SetClientID(&flow.clientID); err != nil {
		return nil, err
	}
	if err := request.SetClientSecret(&flow.clientSecret); err != nil {
		return nil, err
	}
	if err := request.SetRefreshToken(&token); err != nil {
		return nil, err
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
