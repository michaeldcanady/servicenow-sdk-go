package appserviceapi

import (
	"context"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// RegisterServiceRequestBuilder provides operations to register a CSDM service.
type RegisterServiceRequestBuilder struct {
	*servicePostRequestBuilder[*RegisterServiceRequest, RegisterServiceResponse]
}

// NewRegisterServiceRequestBuilderInternal instantiates a new RegisterServiceRequestBuilder.
func NewRegisterServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *RegisterServiceRequestBuilder {
	return &RegisterServiceRequestBuilder{
		newServicePostRequestBuilder[*RegisterServiceRequest, RegisterServiceResponse](requestAdapter, registerServiceURLTemplate, pathParameters, CreateRegisterServiceResponseFromDiscriminatorValue),
	}
}

// Post sends a POST request to register a service.
func (rB *RegisterServiceRequestBuilder) Post(ctx context.Context, body *RegisterServiceRequest, config *RegisterServiceRequestConfiguration) (RegisterServiceResponse, error) {
	return rB.post(ctx, body, config)
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *RegisterServiceRequestBuilder) ToPostRequestInformation(ctx context.Context, body *RegisterServiceRequest, config *RegisterServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	return rB.toPostRequestInformation(ctx, body, config)
}
