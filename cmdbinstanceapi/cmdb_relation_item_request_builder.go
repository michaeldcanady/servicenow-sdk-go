package cmdbinstanceapi

import (
	"context"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
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
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return snerrors.ErrNilRequestBuilder
	}

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
	if !conversion.IsNil(config) {
		if headers := config.Headers; !conversion.IsNil(headers) {
			requestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !conversion.IsNil(options) {
			requestInfo.AddRequestOptions(options)
		}
	}

	return requestInfo, nil
}
