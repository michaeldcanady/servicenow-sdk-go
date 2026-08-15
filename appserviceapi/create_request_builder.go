package appserviceapi

import (
	"context"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const createURLTemplate = "{+baseurl}/api/now/v1/cmdb/app_service/create"

// CreateRequestBuilder provides operations to create an application service.
type CreateRequestBuilder struct {
	*servicePostRequestBuilder[*CreateServiceRequest, CreateServiceResponse]
}

// NewCreateRequestBuilderInternal instantiates a new CreateRequestBuilder.
func NewCreateRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		newServicePostRequestBuilder[*CreateServiceRequest, CreateServiceResponse](requestAdapter, createURLTemplate, pathParameters, CreateCreateServiceResponseFromDiscriminatorValue),
	}
}

// NewCreateRequestBuilder instantiates a new [CreateRequestBuilder].
func NewCreateRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *CreateRequestBuilder {
	return NewCreateRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to create an application service.
func (rB *CreateRequestBuilder) Post(ctx context.Context, body *CreateServiceRequest, config *CreateRequestConfiguration) (CreateServiceResponse, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.post(ctx, body, config)
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *CreateRequestBuilder) ToPostRequestInformation(ctx context.Context, body *CreateServiceRequest, config *CreateRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.toPostRequestInformation(ctx, body, config)
}
