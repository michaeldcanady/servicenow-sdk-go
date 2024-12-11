package authentication

import (
	"context"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type resourceOwnerPasswordFlow struct {
	clientID       string
	requestAdapter abstractions.RequestAdapter
	clientSecret   string
	username       string
	password       string
}

func newResourceOwnerPasswordFlow(clientID string, requestAdapter abstractions.RequestAdapter, clientSecret string, username string, password string) *resourceOwnerPasswordFlow {
	return &resourceOwnerPasswordFlow{
		clientID:       clientID,
		requestAdapter: requestAdapter,
		clientSecret:   clientSecret,
		username:       username,
		password:       password,
	}

}

func (flow *resourceOwnerPasswordFlow) AcquireAuthRecord(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (authenticationTokenResponsable, error) {
	if internal.IsNil(flow) {
		return nil, nil
	}

	pathParameters := map[string]string{}

	builder := NewOauthTokenRequestBuilderInternal(pathParameters, flow.requestAdapter)

	body := newResourceOwnerRequest()
	grantType := grantTypePassword
	if err := body.setGrantType(&grantType); err != nil {
		return nil, err
	}
	if err := body.SetClientID(&flow.clientID); err != nil {
		return nil, err
	}
	if err := body.SetClientSecret(&flow.clientSecret); err != nil {
		return nil, err
	}
	if err := body.SetPassword(&flow.password); err != nil {
		return nil, err
	}
	if err := body.SetUsername(&flow.username); err != nil {
		return nil, err
	}

	return builder.Post(ctx, body, nil)
}
