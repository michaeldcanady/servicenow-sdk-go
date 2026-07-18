package caseapi

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

const caseURLTemplate = "{+baseurl}/api/sn_customerservice/v1/case{?sysparm_query}"

// CaseRequestBuilder provides operations to manage cases.
type CaseRequestBuilder struct {
	core.RequestBuilder
}

// NewCaseRequestBuilderInternal instantiates a new CaseRequestBuilder with the provided request parameters.
func NewCaseRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CaseRequestBuilder {
	return &CaseRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, caseURLTemplate, pathParameters),
	}
}

// ByID returns a CaseItemRequestBuilder for the specified case ID.
func (rB *CaseRequestBuilder) ByID(id string) *CaseItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["id"] = id
	return NewCaseItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// FieldValues returns a CaseFieldValuesRequestBuilder for the specified field name.
func (rB *CaseRequestBuilder) FieldValues(fieldName string) *CaseFieldValuesRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["field_name"] = fieldName
	return NewCaseFieldValuesRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Get sends a GET request to search cases.
func (rB *CaseRequestBuilder) Get(ctx context.Context, config *CaseRequestBuilderGetRequestConfiguration) (CaseCollectionResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCaseCollectionResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(CaseCollectionResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *CaseRequestBuilder) ToGetRequestInformation(ctx context.Context, config *CaseRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return kiotaRequestInfo.RequestInformation, nil
}

// Post sends a POST request to create a case.
func (rB *CaseRequestBuilder) Post(ctx context.Context, body CaseResult, config *CaseRequestBuilderPostRequestConfiguration) (CaseItemResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateCaseItemResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(CaseItemResponse), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *CaseRequestBuilder) ToPostRequestInformation(ctx context.Context, body CaseResult, config *CaseRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)
	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}

	return kiotaRequestInfo.RequestInformation, nil
}
