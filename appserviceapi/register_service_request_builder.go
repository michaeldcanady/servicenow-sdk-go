package appserviceapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
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
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.post(ctx, body, config)
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *RegisterServiceRequestBuilder) ToPostRequestInformation(ctx context.Context, body *RegisterServiceRequest, config *RegisterServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.toPostRequestInformation(ctx, body, config)
}
