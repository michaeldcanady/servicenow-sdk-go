// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

//nolint:dupl // per-verb request-builder methods share the mandatory nil-guard/send boilerplate by convention; each depends on its own outer type, response type, and discriminator factory, so it can't be extracted into a shared helper
package cdmeditorapi

import (
	"context"
	"fmt"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
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
	core.RequestBuilder
}

// NewCdmEditorRequestBuilderInternal instantiates a new CdmEditorRequestBuilder.
func NewCdmEditorRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CdmEditorRequestBuilder {
	return &CdmEditorRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, editorURLTemplate, pathParameters),
	}
}

// Nodes returns a NodesRequestBuilder.
func (rB *CdmEditorRequestBuilder) Nodes() *NodesRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewNodesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Validation returns a ValidationRequestBuilder.
func (rB *CdmEditorRequestBuilder) Validation() *ValidationRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewValidationRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// NodesRequestBuilder handles /v1/nodes endpoint.
type NodesRequestBuilder struct {
	core.RequestBuilder
}

// NewNodesRequestBuilderInternal instantiates a new NodesRequestBuilder.
func NewNodesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *NodesRequestBuilder {
	return &NodesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, editorNodesURLTemplate, pathParameters),
	}
}

// ByID returns the by id request builder.
func (rB *NodesRequestBuilder) ByID(id string) *NodeItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["node_sys_id"] = id
	return NewNodeItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Get sends a GET request.
func (rB *NodesRequestBuilder) Get(ctx context.Context, config *NodesRequestBuilderGetRequestConfiguration) (NodesResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			requestInfo.AddQueryParameters(*config.QueryParameters)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateNodesResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(NodesResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*NodesResponse)(nil))
	}
	return typedRes, nil
}

// Post sends a POST request.
func (rB *NodesRequestBuilder) Post(ctx context.Context, body NodeCreateRequest, config *NodesRequestBuilderPostRequestConfiguration) (NodeResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateNodeResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(NodeResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*NodeResponse)(nil))
	}
	return typedRes, nil
}

// NodeItemRequestBuilder handles /v1/nodes/{id} endpoint.
type NodeItemRequestBuilder struct {
	core.RequestBuilder
}

// NewNodeItemRequestBuilderInternal instantiates a new NodeItemRequestBuilder.
func NewNodeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *NodeItemRequestBuilder {
	return &NodeItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, editorNodeItemURLTemplate, pathParameters),
	}
}

// Put sends a PUT request.
func (rB *NodeItemRequestBuilder) Put(ctx context.Context, body NodeUpdateRequest, config *NodeItemRequestBuilderPutRequestConfiguration) (NodeResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateNodeResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(NodeResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*NodeResponse)(nil))
	}
	return typedRes, nil
}

// Delete removes a node. It returns only an error: the endpoint answers with no content, so
// there is no response body to hand back.
func (rB *NodeItemRequestBuilder) Delete(ctx context.Context, config *NodeItemRequestBuilderDeleteRequestConfiguration) error {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return snerrors.ErrNilRequestAdapter
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, core.DefaultErrorMapping())
}

// ValidationRequestBuilder handles /v1/validation endpoint.
type ValidationRequestBuilder struct {
	core.RequestBuilder
}

// NewValidationRequestBuilderInternal instantiates a new ValidationRequestBuilder.
func NewValidationRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ValidationRequestBuilder {
	return &ValidationRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, editorValidateURLTemplate, pathParameters),
	}
}

// Get sends a GET request.
func (rB *ValidationRequestBuilder) Get(ctx context.Context, config *ValidationRequestBuilderGetRequestConfiguration) (ValidationResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
		if config.QueryParameters != nil {
			requestInfo.AddQueryParameters(*config.QueryParameters)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateValidationResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(ValidationResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*ValidationResponse)(nil))
	}
	return typedRes, nil
}
