package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	createDocumentURLTemplate = "{+baseurl}/api/now/v1/documents/createDocument"
)

// CreateDocumentRequestBuilder provides operations to manage the createDocument endpoint.
type CreateDocumentRequestBuilder struct {
	*documentPostRequestBuilder
}

// NewCreateDocumentRequestBuilderInternal instantiates a new CreateDocumentRequestBuilder.
func NewCreateDocumentRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CreateDocumentRequestBuilder {
	return &CreateDocumentRequestBuilder{
		newDocumentPostRequestBuilder(requestAdapter, createDocumentURLTemplate, pathParameters),
	}
}

// Post creates or links a document from an attachment, DMS repo or cloud.
func (rB *CreateDocumentRequestBuilder) Post(ctx context.Context, requestConfiguration *CreateDocumentRequestBuilderPostRequestConfiguration) (*core.BaseServiceNowItemResponse[Document], error) {
	return rB.post(ctx, (*documentPostRequestConfiguration)(requestConfiguration))
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *CreateDocumentRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CreateDocumentRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	return rB.toPostRequestInformation(ctx, (*documentPostRequestConfiguration)(requestConfiguration))
}
