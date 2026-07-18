package documentsapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	syncDownURLTemplate = "{+baseurl}/api/now/v1/documents/{documentSysId}/syncDown"
)

// SyncDownRequestBuilder provides operations to manage the syncDown endpoint.
type SyncDownRequestBuilder struct {
	*documentPostRequestBuilder
}

// NewSyncDownRequestBuilderInternal instantiates a new SyncDownRequestBuilder.
func NewSyncDownRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SyncDownRequestBuilder {
	return &SyncDownRequestBuilder{
		newDocumentPostRequestBuilder(requestAdapter, syncDownURLTemplate, pathParameters),
	}
}

// Post synchronizes the specified document.
func (rB *SyncDownRequestBuilder) Post(ctx context.Context, requestConfiguration *SyncDownRequestBuilderPostRequestConfiguration) (*core.BaseServiceNowItemResponse[Document], error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.post(ctx, (*documentPostRequestConfiguration)(requestConfiguration))
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *SyncDownRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SyncDownRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.toPostRequestInformation(ctx, (*documentPostRequestConfiguration)(requestConfiguration))
}
