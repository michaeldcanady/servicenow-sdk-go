package cmdbinstanceapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cmdbRelationItemURLTemplate = "{+baseurl}/api/now/v1/cmdb/instance/{className}/{sys_id}/relation/{rel_sys_id}"
)

// CmdbRelationItemRequestBuilder provides operations to manage a specific CI relationship.
type CmdbRelationItemRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewCmdbRelationItemRequestBuilderInternal instantiates a new CmdbRelationItemRequestBuilder.
func NewCmdbRelationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbRelationItemRequestBuilder {
	return &CmdbRelationItemRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, cmdbRelationItemURLTemplate, pathParameters),
	}
}

// Delete deletes a specific Relation for the CI.
func (rB *CmdbRelationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CmdbRelationItemRequestBuilderDeleteRequestConfiguration) error {
	if internal.IsNil(rB) {
		return nil
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	err = rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return nil
}

// ToDeleteRequestInformation converts request configurations to Delete request information.
func (rB *CmdbRelationItemRequestBuilder) ToDeleteRequestInformation(_ context.Context, requestConfiguration *CmdbRelationItemRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}

	return kiotaRequestInfo.RequestInformation, nil
}
