package cmdbinstanceapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cmdbRelationItemURLTemplate = "{+baseurl}/api/now/v1/cmdb/instance/{className}/{sys_id}/relation/{rel_sys_id}"
)

// CmdbRelationItemRequestBuilder provides operations to manage a specific CI relationship.
type CmdbRelationItemRequestBuilder struct {
	core.RequestBuilder
}

// NewCmdbRelationItemRequestBuilderInternal instantiates a new CmdbRelationItemRequestBuilder.
func NewCmdbRelationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbRelationItemRequestBuilder {
	return &CmdbRelationItemRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, cmdbRelationItemURLTemplate, pathParameters),
	}
}

// Delete deletes a specific Relation for the CI.
func (rB *CmdbRelationItemRequestBuilder) Delete(ctx context.Context, config *CmdbRelationItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := rB.ToDeleteRequestInformation(ctx, config)
	if err != nil {
		return err
	}

	errorMapping := core.DefaultErrorMapping()
	err = rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return nil
}

// ToDeleteRequestInformation converts request configurations to Delete request information.
func (rB *CmdbRelationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, config *CmdbRelationItemRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if headers := config.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}

	return requestInfo, nil
}
