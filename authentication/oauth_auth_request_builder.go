package authentication

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

type OauthAuthRequestBuilderGetRequestConfiguration struct {
	// Headers Request headers
	Headers *abstractions.RequestHeaders
	// Options Request options
	Options []abstractions.RequestOption

	QueryParameters oauthAuthQueryParameters
}

type oauthAuthRequestBuilder struct {
	abstractions.BaseRequestBuilder
}

// NewOauthAuthRequestBuilderInternal instantiates a new OauthAuthRequestBuilder with custom parsable for table entries.
func NewOauthAuthRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *oauthAuthRequestBuilder {
	m := &oauthAuthRequestBuilder{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, oauthAuthURLTemplate, pathParameters),
	}
	return m
}

// NewOauthAuthRequestBuilder instantiates a new OauthAuthRequestBuilder with custom parsable for table entries.
func NewOauthAuthRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *oauthAuthRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[rawURLKey] = rawURL
	return NewOauthAuthRequestBuilderInternal(urlParams, requestAdapter)
}

func (rB *oauthAuthRequestBuilder) Get(ctx context.Context, requestConfiguration *OauthAuthRequestBuilderGetRequestConfiguration) (authenticationTokenResponsable, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}

func (rB *oauthAuthRequestBuilder) ToGetRequestInformation(ctx context.Context, body grantTypeRequestable) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	return nil, errors.New("not implemented")
}
