package documentsapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	attachURLTemplate = "{+baseurl}/api/now/v1/documents/{provider_id}/attach"
)

// AttachRequestBuilder provides operations to manage the attach endpoint.
type AttachRequestBuilder struct {
	*documentPostRequestBuilder
}

// NewAttachRequestBuilderInternal instantiates a new AttachRequestBuilder.
func NewAttachRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AttachRequestBuilder {
	return &AttachRequestBuilder{
		newDocumentPostRequestBuilder(requestAdapter, attachURLTemplate, pathParameters),
	}
}

// Post attaches a document using the specified provider.
func (rB *AttachRequestBuilder) Post(ctx context.Context, requestConfiguration *AttachRequestBuilderPostRequestConfiguration) (*core.BaseServiceNowItemResponse[Document], error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.post(ctx, (*documentPostRequestConfiguration)(requestConfiguration))
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *AttachRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AttachRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.toPostRequestInformation(ctx, (*documentPostRequestConfiguration)(requestConfiguration))
}
