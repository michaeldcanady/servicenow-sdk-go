package appserviceapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// FindServiceRequestBuilder provides operations to find an application service.
type FindServiceRequestBuilder struct {
	core.RequestBuilder
}

// NewFindServiceRequestBuilderInternal instantiates a new FindServiceRequestBuilder.
func NewFindServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *FindServiceRequestBuilder {
	return &FindServiceRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, findServiceURLTemplate, pathParameters),
	}
}

// Get sends a GET request to find an application service.
func (rB *FindServiceRequestBuilder) Get(ctx context.Context, config *FindServiceRequestConfiguration) (FindServiceResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}
	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateFindServiceResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(FindServiceResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *FindServiceRequestBuilder) ToGetRequestInformation(ctx context.Context, config *FindServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	return kiotaRequestInfo.RequestInformation, nil
}
