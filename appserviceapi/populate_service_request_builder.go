package appserviceapi

import (
	"context"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const populateServiceURLTemplate = "{+baseurl}/api/now/cmdb/csdm/app_service/{sys_id}/populate_service"

// PopulateServiceRequestBuilder provides operations to populate a CSDM service.
type PopulateServiceRequestBuilder struct {
	core.RequestBuilder
}

// NewPopulateServiceRequestBuilderInternal instantiates a new [PopulateServiceRequestBuilder].
func NewPopulateServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *PopulateServiceRequestBuilder {
	return &PopulateServiceRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, populateServiceURLTemplate, pathParameters),
	}
}

// NewPopulateServiceRequestBuilder instantiates a new [PopulateServiceRequestBuilder] with a raw URL.
func NewPopulateServiceRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *PopulateServiceRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewPopulateServiceRequestBuilderInternal(urlParams, requestAdapter)
}

// Put sends a PUT request to populate a service.
func (rB *PopulateServiceRequestBuilder) Put(ctx context.Context, body *PopulateServiceRequest, config *PopulateServiceRequestConfiguration) (PopulateServiceResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToPutRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreatePopulateServiceResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(PopulateServiceResponse)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*PopulateServiceResponse)(nil))
	}

	return typedResp, nil
}

// ToPutRequestInformation creates a [abstractions.RequestInformation] object for a PUT request.
func (rB *PopulateServiceRequestBuilder) ToPutRequestInformation(ctx context.Context, body *PopulateServiceRequest, config *PopulateServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	// TODO: check for nil/empty template and path parameters
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}
