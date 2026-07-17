package appserviceapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// servicePostRequestBuilder is the shared implementation behind every appserviceapi
// request builder whose only operation is a POST with a Parsable body returning a
// single response (Create, RegisterService) - they differ only in URL template,
// body type, response type, and response factory.
type servicePostRequestBuilder[TBody serialization.Parsable, TResponse any] struct {
	core.RequestBuilder
	responseFactory serialization.ParsableFactory
}

// newServicePostRequestBuilder instantiates a new servicePostRequestBuilder.
func newServicePostRequestBuilder[TBody serialization.Parsable, TResponse any](requestAdapter abstractions.RequestAdapter, urlTemplate string, pathParameters map[string]string, responseFactory serialization.ParsableFactory) *servicePostRequestBuilder[TBody, TResponse] {
	return &servicePostRequestBuilder[TBody, TResponse]{
		RequestBuilder:  core.NewBaseRequestBuilder(requestAdapter, urlTemplate, pathParameters),
		responseFactory: responseFactory,
	}
}

// post sends a POST request with the given body and returns the decoded response.
func (rB *servicePostRequestBuilder[TBody, TResponse]) post(ctx context.Context, body TBody, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (TResponse, error) {
	var zero TResponse

	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return zero, nil
	}

	requestInfo, err := rB.toPostRequestInformation(ctx, body, config)
	if err != nil {
		return zero, err
	}
	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, rB.responseFactory, errorMapping)
	if err != nil {
		return zero, err
	}
	if res == nil {
		return zero, nil
	}
	return res.(TResponse), nil
}

// toPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *servicePostRequestBuilder[TBody, TResponse]) toPostRequestInformation(ctx context.Context, body TBody, config *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	return kiotaRequestInfo.RequestInformation, nil
}
