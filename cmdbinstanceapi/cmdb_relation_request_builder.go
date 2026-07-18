package cmdbinstanceapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	relSysIDKey             = "rel_sys_id"
	cmdbRelationURLTemplate = "{+baseurl}/api/now/v1/cmdb/instance/{className}/{sys_id}/relation"
)

// CmdbRelationRequestBuilder provides operations to manage CI relationships.
type CmdbRelationRequestBuilder struct {
	core.RequestBuilder
}

var _ core.ItemPostRequestBuilder[CmdbInstance, abstractions.DefaultQueryParameters, abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]] = (*CmdbRelationRequestBuilder)(nil)

// NewCmdbRelationRequestBuilderInternal instantiates a new CmdbRelationRequestBuilder.
func NewCmdbRelationRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CmdbRelationRequestBuilder {
	return &CmdbRelationRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, cmdbRelationURLTemplate, pathParameters),
	}
}

// Post creates a Relation for the CI.
func (rB *CmdbRelationRequestBuilder) Post(ctx context.Context, body CmdbInstance, config *CmdbRelationRequestBuilderPostRequestConfiguration) (*core.BaseServiceNowItemResponse[CmdbInstance], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[CmdbInstance](CreateCmdbInstanceFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*core.BaseServiceNowItemResponse[CmdbInstance]), nil
}

// ToPostRequestInformation converts request configurations to Post request information.
func (rB *CmdbRelationRequestBuilder) ToPostRequestInformation(ctx context.Context, body CmdbInstance, config *CmdbRelationRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if headers := config.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	if !conversion.IsNil(body) {
		err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
		if err != nil {
			return nil, err
		}
	}

	return requestInfo, nil
}

// ByID provides operations to manage a specific CI relationship.
func (rB *CmdbRelationRequestBuilder) ByID(relSysID string) *CmdbRelationItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters[relSysIDKey] = relSysID
	return NewCmdbRelationItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}
