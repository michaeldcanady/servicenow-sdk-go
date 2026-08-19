package appserviceapi

import (
	"context"
	"fmt"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const serviceDetailsURLTemplate = "{+baseurl}/api/now/cmdb/csdm/app_service/{sys_id}/service_details"

// ServiceDetailsRequestBuilder provides operations to update details of a CSDM service.
type ServiceDetailsRequestBuilder struct {
	core.RequestBuilder
}

// NewServiceDetailsRequestBuilderInternal instantiates a new [ServiceDetailsRequestBuilder].
func NewServiceDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ServiceDetailsRequestBuilder {
	return &ServiceDetailsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, serviceDetailsURLTemplate, pathParameters),
	}
}

// NewServiceDetailsRequestBuilder instantiates a new [ServiceDetailsRequestBuilder] with a raw URL.
func NewServiceDetailsRequestBuilder(rawURL string, requestAdapter abstractions.RequestAdapter) *ServiceDetailsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewServiceDetailsRequestBuilderInternal(urlParams, requestAdapter)
}

// Put sends a PUT request to update service details.
func (rB *ServiceDetailsRequestBuilder) Put(ctx context.Context, body *ServiceDetailsRequest, config *ServiceDetailsRequestConfiguration) (ServiceDetailsResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateServiceDetailsResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(ServiceDetailsResponse)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*ServiceDetailsResponse)(nil))
	}

	return typedResp, nil
}

// ToPutRequestInformation creates a RequestInformation object for a PUT request.
func (rB *ServiceDetailsRequestBuilder) ToPutRequestInformation(ctx context.Context, body *ServiceDetailsRequest, config *ServiceDetailsRequestConfiguration) (*abstractions.RequestInformation, error) {
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
