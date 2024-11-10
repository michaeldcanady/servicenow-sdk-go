package auth

import (
	"context"
	"errors"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	oauthTokenURITemplate = "{+baseurl}/oauth_token.do" //nolint:gosec
)

// OauthTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OauthTokenRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *abstractions.RequestHeaders
	// Request options
	Options []abstractions.RequestOption
}

type OauthTokenRequestBuilder struct {
	abstractions.BaseRequestBuilder
}

func NewOauthTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *OauthTokenRequestBuilder {
	return &OauthTokenRequestBuilder{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, oauthTokenURITemplate, pathParameters),
	}
}

func NewOauthTokenRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *OauthTokenRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawURL
	return NewOauthTokenRequestBuilderInternal(urlParams, requestAdapter)
}

func (rB *OauthTokenRequestBuilder) Post(ctx context.Context, body Authenticatable, requestConfiguration *OauthTokenRequestBuilderPostRequestConfiguration) (AccessTokenable, error) {
	requestInfo, err := rB.toPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := abstractions.ErrorMappings{}
	res, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAccessTokenFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	typedResp, ok := res.(AccessTokenable)
	if !ok {
		return nil, errors.New("res is not AccessTokenable")
	}
	return typedResp, nil
}

func (rB *OauthTokenRequestBuilder) toPostRequestInformation(ctx context.Context, body Authenticatable, requestConfiguration *OauthTokenRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.BaseRequestBuilder.UrlTemplate, rB.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestInfo.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
		err := requestInfo.SetContentFromParsable(ctx, rB.RequestAdapter, "application/x-www-form-urlencoded", body)
		if err != nil {
			return nil, err
		}
	}
	return requestInfo, nil
}
