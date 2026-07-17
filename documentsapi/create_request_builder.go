package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	createURLTemplate = "{+baseurl}/api/now/v1/documents/create"
)

// CreateRequestBuilder provides operations to manage the create endpoint.
type CreateRequestBuilder struct {
	*documentPostRequestBuilder
}

// NewCreateRequestBuilderInternal instantiates a new CreateRequestBuilder.
func NewCreateRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CreateRequestBuilder {
	return &CreateRequestBuilder{
		newDocumentPostRequestBuilder(requestAdapter, createURLTemplate, pathParameters),
	}
}

// Post creates a new document.
func (rB *CreateRequestBuilder) Post(ctx context.Context, requestConfiguration *CreateRequestBuilderPostRequestConfiguration) (*core.BaseServiceNowItemResponse[Document], error) {
	return rB.post(ctx, (*documentPostRequestConfiguration)(requestConfiguration))
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *CreateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CreateRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	return rB.toPostRequestInformation(ctx, (*documentPostRequestConfiguration)(requestConfiguration))
}
