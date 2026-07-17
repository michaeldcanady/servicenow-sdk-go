package appserviceapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

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

// Post sends a POST request to create an application service.
func (rB *CreateRequestBuilder) Post(ctx context.Context, body *CreateServiceRequest, config *CreateRequestConfiguration) (CreateServiceResponse, error) {
	if conversion.IsNil(rB) {
		return nil, nil
	}
	return rB.post(ctx, body, config)
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *CreateRequestBuilder) ToPostRequestInformation(ctx context.Context, body *CreateServiceRequest, config *CreateRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, nil
	}
	return rB.toPostRequestInformation(ctx, body, config)
}
