package auth

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	oauthAuthURITemplate = "https://{+baseurl}/oauth_auth.do{?response_type,redirect_uri,client_id,state}" //nolint:gosec
)

type OauthAuthRequestBuilderGetQueryParameters struct {
	ResponseType ResponseType `url:"response_type"`
	RedirectURI  string       `url:"redirect_uri"`
	ClientID     string       `url:"client_id"`
	State        string       `url:"state"`
}

// OauthAuthRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OauthAuthRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *abstractions.RequestHeaders
	// Request options
	Options []abstractions.RequestOption

	QueryParameters *OauthAuthRequestBuilderGetQueryParameters
}

type OauthAuthRequestBuilder struct {
	abstractions.BaseRequestBuilder
}

func NewOauthAuthRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *OauthAuthRequestBuilder {
	return &OauthAuthRequestBuilder{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, oauthAuthURITemplate, pathParameters),
	}
}

func NewOauthAuthRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *OauthAuthRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawURL
	return NewOauthAuthRequestBuilderInternal(urlParams, requestAdapter)
}

func (rB *OauthAuthRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OauthAuthRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.BaseRequestBuilder.UrlTemplate, rB.BaseRequestBuilder.PathParameters)
	kRequestInfo := http.KiotaRequestInformation{*requestInfo}
	if requestConfiguration != nil {
		kRequestInfo.Headers.AddAll(kRequestInfo.Headers)
		kRequestInfo.AddRequestOptions(requestConfiguration.Options)
		kRequestInfo.AddQueryParameters(requestConfiguration.QueryParameters)
	}
	return &kRequestInfo.RequestInformation, nil
}
