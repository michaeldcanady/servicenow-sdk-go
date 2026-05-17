package cmdbinstanceapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	cmdbRelationURLTemplate = "{+baseurl}/api/now/v1/cmdb/instance/{className}/{sys_id}/relation"
)

// CmdbRelationRequestBuilder provides operations to manage CI relationships.
type CmdbRelationRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewCmdbRelationRequestBuilderInternal instantiates a new CmdbRelationRequestBuilder.
func NewCmdbRelationRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbRelationRequestBuilder {
	return &CmdbRelationRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, cmdbRelationURLTemplate, pathParameters),
	}
}

// Post creates a Relation for the CI.
func (rB *CmdbRelationRequestBuilder) Post(ctx context.Context, requestConfiguration *CmdbRelationRequestBuilderPostRequestConfiguration) (*newInternal.BaseServiceNowItemResponse[CmdbInstance], error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowItemResponseFromDiscriminatorValue[CmdbInstance](CreateCmdbInstanceFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	return res.(*newInternal.BaseServiceNowItemResponse[CmdbInstance]), nil
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *CmdbRelationRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CmdbRelationRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if data := requestConfiguration.Data; !internal.IsNil(data) {
			err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, data)
			if err != nil {
				return nil, err
			}
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
