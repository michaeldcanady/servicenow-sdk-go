package documentsapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	versionStateURLTemplate = "{+baseurl}/api/now/v1/documents/versionstate/{version_sys_id}"
)

// VersionStateRequestBuilder provides operations to manage the versionstate endpoint.
type VersionStateRequestBuilder struct {
	*documentGetRequestBuilder
}

// NewVersionStateRequestBuilderInternal instantiates a new VersionStateRequestBuilder.
func NewVersionStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *VersionStateRequestBuilder {
	return &VersionStateRequestBuilder{
		newDocumentGetRequestBuilder(requestAdapter, versionStateURLTemplate, pathParameters),
	}
}

// Get retrieves the state of the specified document version.
func (rB *VersionStateRequestBuilder) Get(ctx context.Context, requestConfiguration *VersionStateRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowItemResponse[Document], error) {
	return rB.get(ctx, (*documentGetRequestConfiguration)(requestConfiguration))
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *VersionStateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VersionStateRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	return rB.toGetRequestInformation(ctx, (*documentGetRequestConfiguration)(requestConfiguration))
}
