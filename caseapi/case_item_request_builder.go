package caseapi

import (
	"context"
	"fmt"
	"maps"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const caseItemURLTemplate = "{+baseurl}/api/sn_customerservice/v1/case/{id}{?sysparm_display_value}"

// CaseItemRequestBuilder provides operations to manage a single case.
type CaseItemRequestBuilder struct {
	core.RequestBuilder
}

// NewCaseItemRequestBuilderInternal instantiates a new CaseItemRequestBuilder.
func NewCaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CaseItemRequestBuilder {
	return &CaseItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, caseItemURLTemplate, pathParameters),
	}
}

// Activities returns a CaseActivitiesRequestBuilder.
func (rB *CaseItemRequestBuilder) Activities() *CaseActivitiesRequestBuilder { // V
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewCaseActivitiesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// FieldValues returns a CaseFieldValuesRequestBuilder for the specified field name.
func (rB *CaseItemRequestBuilder) FieldValues(fieldName string) *CaseFieldValuesRequestBuilder { // V
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters[fieldNameKey] = fieldName
	return NewItemFieldValuesRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Get sends a GET request to retrieve a single case.
func (rB *CaseItemRequestBuilder) Get(ctx context.Context, config *CaseItemRequestBuilderGetRequestConfiguration) (CaseItemResponse, error) {
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

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *CaseItemRequestBuilder) ToGetRequestInformation(_ context.Context, config *CaseItemRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	// TODO: check for nil/empty template and path parameters
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}

// Put sends a PUT request to update an existing case.
func (rB *CaseItemRequestBuilder) Put(ctx context.Context, body CaseResult, config *CaseItemRequestBuilderPutRequestConfiguration) (CaseItemResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}
	requestInfo, err := rB.ToPutRequestInformation(ctx, body, config)
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

// ToPutRequestInformation creates a RequestInformation object for a PUT request.
func (rB *CaseItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body CaseResult, config *CaseItemRequestBuilderPutRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	// TODO: check for nil/empty template and path parameters
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}

	return requestInfo, nil
}
