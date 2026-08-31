// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package caseapi

import (
	"context"
	"fmt"
	"maps"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const caseURLTemplate = "{+baseurl}/api/sn_customerservice/v1/case{?sysparm_query}"

// CaseRequestBuilder provides operations to manage cases.
type CaseRequestBuilder struct {
	core.RequestBuilder
}

// NewCaseRequestBuilderInternal instantiates a new [CaseRequestBuilder].
func NewCaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CaseRequestBuilder {
	return &CaseRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, caseURLTemplate, pathParameters),
	}
}

// NewDefaultCaseRequestBuilder instantiates a new [CaseRequestBuilder].
func NewDefaultCaseRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *CaseRequestBuilder {
	return NewCaseRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// ByID returns a [CaseItemRequestBuilder] for the specified case ID.
func (rB *CaseRequestBuilder) ByID(id string) *CaseItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["id"] = id
	return NewCaseItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// FieldValues returns a [CaseFieldValuesRequestBuilder] for the specified field name.
func (rB *CaseRequestBuilder) FieldValues(fieldName string) *CaseFieldValuesRequestBuilder { // V
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters[fieldNameKey] = fieldName
	return NewCaseFieldValuesRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Get sends a GET request to search cases.
func (rB *CaseRequestBuilder) Get(ctx context.Context, config *CaseRequestBuilderGetRequestConfiguration) (CaseCollectionResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCaseCollectionResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(CaseCollectionResponse)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*CaseCollectionResponse)(nil))
	}

	return typedResp, nil
}

// ToGetRequestInformation creates a [abstractions.RequestInformation] object for a GET request.
func (rB *CaseRequestBuilder) ToGetRequestInformation(_ context.Context, config *CaseRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	// TODO: check for nil/empty template and path parameters
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}

// Post sends a POST request to create a case.
func (rB *CaseRequestBuilder) Post(ctx context.Context, body CaseResult, config *CaseRequestBuilderPostRequestConfiguration) (CaseItemResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCaseItemResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(CaseItemResponse)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*CaseItemResponse)(nil))
	}

	return typedResp, nil
}

// ToPostRequestInformation creates a [abstractions.RequestInformation] object for a POST request.
func (rB *CaseRequestBuilder) ToPostRequestInformation(ctx context.Context, body CaseResult, config *CaseRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	// TODO: check for nil/empty template and path parameters
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}

	return requestInfo, nil
}
