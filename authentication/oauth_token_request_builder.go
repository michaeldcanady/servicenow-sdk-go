package authentication

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	oauthTokenURLTemplate = "{+baseurl}/oauth_token.do{?}" //nolint:gosec
	rawURLKey             = "request-raw-url"
)

type OauthTokenRequestBuilderPostRequestConfiguration struct {
	// Headers Request headers
	Headers *abstractions.RequestHeaders
	// Options Request options
	Options []abstractions.RequestOption
}

type oauthTokenRequestBuilder struct {
	abstractions.BaseRequestBuilder
}

// NewOauthTokenRequestBuilderInternal instantiates a new OauthTokenRequestBuilder with custom parsable for table entries.
func NewOauthTokenRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *oauthTokenRequestBuilder {
	m := &oauthTokenRequestBuilder{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, oauthTokenURLTemplate, pathParameters),
	}
	return m
}

// NewOauthTokenRequestBuilder instantiates a new OauthTokenRequestBuilder with custom parsable for table entries.
func NewOauthTokenRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *oauthTokenRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[rawURLKey] = rawURL
	return NewOauthTokenRequestBuilderInternal(urlParams, requestAdapter)
}

func (rB *oauthTokenRequestBuilder) Post(ctx context.Context, body grantTypeRequestable, requestConfiguration *OauthTokenRequestBuilderPostRequestConfiguration) (authenticationTokenResponsable, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	resp, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuthenticationTokenResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	typedResp, ok := resp.(authenticationTokenResponsable)
	if !ok {
		return nil, errors.New("resp is not authenticationTokenResponsable")
	}

	return typedResp, nil
}

func (rB *oauthTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body grantTypeRequestable, requestConfiguration *OauthTokenRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := intHttp.NewRequestInformationWithMethodAndURLTemplateAndPathParameters(abstractions.POST, rB.UrlTemplate, rB.PathParameters)
	if !internal.IsNil(requestConfiguration) {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	if err := requestInfo.SetContentFromParsable(ctx, rB.BaseRequestBuilder.RequestAdapter, "multipart/form-data", body); err != nil {
		return nil, err
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")

	return &requestInfo.RequestInformation, nil
}
