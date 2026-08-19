package appserviceapi

import (
	"context"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const convertToManualServiceURLTemplate = "{+baseurl}/api/now/cmdb/app_service/convertToManualService"

// ConvertToManualServiceRequestBuilder provides operations to convert an application service to a manual service.
type ConvertToManualServiceRequestBuilder struct {
	*servicePostRequestBuilder[*AppServiceActionRequest, AppServiceActionResponse]
}

// NewConvertToManualServiceRequestBuilderInternal instantiates a new ConvertToManualServiceRequestBuilder.
func NewConvertToManualServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ConvertToManualServiceRequestBuilder {
	return &ConvertToManualServiceRequestBuilder{
		newServicePostRequestBuilder[*AppServiceActionRequest, AppServiceActionResponse](requestAdapter, convertToManualServiceURLTemplate, pathParameters, CreateAppServiceActionResponseFromDiscriminatorValue),
	}
}

// NewConvertToManualServiceRequestBuilder instantiates a new [ConvertToManualServiceRequestBuilder].
func NewConvertToManualServiceRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *ConvertToManualServiceRequestBuilder {
	return NewConvertToManualServiceRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to convert an application service to a manual service.
func (rB *ConvertToManualServiceRequestBuilder) Post(ctx context.Context, body *AppServiceActionRequest, config *ConvertToManualServiceRequestConfiguration) (AppServiceActionResponse, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.post(ctx, body, config)
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *ConvertToManualServiceRequestBuilder) ToPostRequestInformation(ctx context.Context, body *AppServiceActionRequest, config *ConvertToManualServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.toPostRequestInformation(ctx, body, config)
}
