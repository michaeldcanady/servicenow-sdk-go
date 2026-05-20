package cdmeditorapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	editorURLTemplate         = "{+baseurl}/api/sn_cdm/editor"
	editorNodesURLTemplate    = "{+baseurl}/api/sn_cdm/editor/v1/nodes{?sys_id,parent_id,type}"
	editorNodeItemURLTemplate = "{+baseurl}/api/sn_cdm/editor/v1/nodes/{node_sys_id}"
	editorValidateURLTemplate = "{+baseurl}/api/sn_cdm/editor/v1/validation{?cdm_id}"
)

// CdmEditorRequestBuilder provides operations to manage CDM Editor.
type CdmEditorRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewCdmEditorRequestBuilderInternal instantiates a new CdmEditorRequestBuilder.
func NewCdmEditorRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CdmEditorRequestBuilder {
	return &CdmEditorRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, editorURLTemplate, pathParameters),
	}
}

// Nodes returns a NodesRequestBuilder.
func (rB *CdmEditorRequestBuilder) Nodes() *NodesRequestBuilder {
	return NewNodesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Validation returns a ValidationRequestBuilder.
func (rB *CdmEditorRequestBuilder) Validation() *ValidationRequestBuilder {
	return NewValidationRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// NodesRequestBuilder handles /v1/nodes endpoint.
type NodesRequestBuilder struct {
	newInternal.RequestBuilder
}

func NewNodesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *NodesRequestBuilder {
	return &NodesRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, editorNodesURLTemplate, pathParameters),
	}
}

func (rB *NodesRequestBuilder) ByID(id string) *NodeItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["node_sys_id"] = id
	return NewNodeItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

func (rB *NodesRequestBuilder) Get(ctx context.Context, config *NodesRequestBuilderGetRequestConfiguration) (NodesResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			kiotaRequestInfo.AddQueryParameters(config.QueryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateNodesResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(NodesResponse), nil
}

func (rB *NodesRequestBuilder) Post(ctx context.Context, body NodeCreateRequest, config *NodesRequestBuilderPostRequestConfiguration) (NodeResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateNodeResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(NodeResponse), nil
}

// NodeItemRequestBuilder handles /v1/nodes/{id} endpoint.
type NodeItemRequestBuilder struct {
	newInternal.RequestBuilder
}

func NewNodeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *NodeItemRequestBuilder {
	return &NodeItemRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, editorNodeItemURLTemplate, pathParameters),
	}
}

func (rB *NodeItemRequestBuilder) Put(ctx context.Context, body NodeUpdateRequest, config *NodeItemRequestBuilderPutRequestConfiguration) (NodeResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateNodeResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(NodeResponse), nil
}

func (rB *NodeItemRequestBuilder) Delete(ctx context.Context, config *NodeItemRequestBuilderDeleteRequestConfiguration) (NodeDeleteResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateNodeDeleteResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(NodeDeleteResponse), nil
}

// ValidationRequestBuilder handles /v1/validation endpoint.
type ValidationRequestBuilder struct {
	newInternal.RequestBuilder
}

func NewValidationRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ValidationRequestBuilder {
	return &ValidationRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, editorValidateURLTemplate, pathParameters),
	}
}

func (rB *ValidationRequestBuilder) Get(ctx context.Context, config *ValidationRequestBuilderGetRequestConfiguration) (ValidationResponse, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if config.Headers != nil {
			kiotaRequestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			kiotaRequestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			kiotaRequestInfo.AddQueryParameters(config.QueryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateValidationResponseFromDiscriminatorValue, nil)
	if err != nil {
		return nil, err
	}
	return res.(ValidationResponse), nil
}
