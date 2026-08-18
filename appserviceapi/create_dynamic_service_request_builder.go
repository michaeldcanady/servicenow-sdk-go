package appserviceapi

import (
	"context"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const createDynamicServiceURLTemplate = "{+baseurl}/api/now/cmdb/app_service/createDynamicService"

// CreateDynamicServiceRequestBuilder provides operations to create a dynamic application service.
type CreateDynamicServiceRequestBuilder struct {
	*servicePostRequestBuilder[*AppServiceActionRequest, AppServiceActionResponse]
}

// NewCreateDynamicServiceRequestBuilderInternal instantiates a new CreateDynamicServiceRequestBuilder.
func NewCreateDynamicServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CreateDynamicServiceRequestBuilder {
	return &CreateDynamicServiceRequestBuilder{
		newServicePostRequestBuilder[*AppServiceActionRequest, AppServiceActionResponse](requestAdapter, createDynamicServiceURLTemplate, pathParameters, CreateAppServiceActionResponseFromDiscriminatorValue),
	}
}

// NewCreateDynamicServiceRequestBuilder instantiates a new [CreateDynamicServiceRequestBuilder].
func NewCreateDynamicServiceRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *CreateDynamicServiceRequestBuilder {
	return NewCreateDynamicServiceRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to create a dynamic application service.
func (rB *CreateDynamicServiceRequestBuilder) Post(ctx context.Context, body *AppServiceActionRequest, config *CreateDynamicServiceRequestConfiguration) (AppServiceActionResponse, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.post(ctx, body, config)
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *CreateDynamicServiceRequestBuilder) ToPostRequestInformation(ctx context.Context, body *AppServiceActionRequest, config *CreateDynamicServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.toPostRequestInformation(ctx, body, config)
}
