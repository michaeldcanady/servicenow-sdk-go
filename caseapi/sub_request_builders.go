package caseapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	caseItemURLTemplate        = "{+baseurl}/api/sn_customerservice/v1/case/{id}"
	caseActivitiesURLTemplate  = "{+baseurl}/api/sn_customerservice/v1/case/{id}/activities"
	fieldValuesURLTemplate     = "{+baseurl}/api/sn_customerservice/v1/case/field_values/{field_name}"
	itemFieldValuesURLTemplate = "{+baseurl}/api/sn_customerservice/v1/case/{id}/field_values/{field_name}"
)

// CaseItemRequestBuilder provides operations to manage a single case.
type CaseItemRequestBuilder struct {
	core.RequestBuilder
}

func NewCaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CaseItemRequestBuilder {
	return &CaseItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, caseItemURLTemplate, pathParameters),
	}
}

// Activities returns a CaseActivitiesRequestBuilder.
func (rB *CaseItemRequestBuilder) Activities() *CaseActivitiesRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewCaseActivitiesRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// FieldValues returns a CaseFieldValuesRequestBuilder for the specified field name.
func (rB *CaseItemRequestBuilder) FieldValues(fieldName string) *CaseFieldValuesRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["field_name"] = fieldName
	return NewItemFieldValuesRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// Get sends a GET request to retrieve a single case.
func (rB *CaseItemRequestBuilder) Get(ctx context.Context, config *CaseItemRequestBuilderGetRequestConfiguration) (CaseItemResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
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

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *CaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, config *CaseItemRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}

// Put sends a PUT request to update an existing case.
func (rB *CaseItemRequestBuilder) Put(ctx context.Context, body CaseResult, config *CaseItemRequestBuilderPutRequestConfiguration) (CaseItemResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}
	requestInfo, err := rB.ToPutRequestInformation(ctx, body, config)
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

// ToPutRequestInformation creates a RequestInformation object for a PUT request.
func (rB *CaseItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body CaseResult, config *CaseItemRequestBuilderPutRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body)
	if err != nil {
		return nil, err
	}

	return kiotaRequestInfo.RequestInformation, nil
}

// CaseActivitiesRequestBuilder provides operations to manage case activities.
type CaseActivitiesRequestBuilder struct {
	core.RequestBuilder
}

func NewCaseActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CaseActivitiesRequestBuilder {
	return &CaseActivitiesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, caseActivitiesURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve case activities.
func (rB *CaseActivitiesRequestBuilder) Get(ctx context.Context, config *CaseActivitiesRequestBuilderGetRequestConfiguration) (ActivitiesResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateActivitiesResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(ActivitiesResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *CaseActivitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, config *CaseActivitiesRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}

// CaseFieldValuesRequestBuilder provides operations to manage field values.
type CaseFieldValuesRequestBuilder struct {
	core.RequestBuilder
}

func NewCaseFieldValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CaseFieldValuesRequestBuilder {
	return &CaseFieldValuesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, fieldValuesURLTemplate, pathParameters),
	}
}

func NewItemFieldValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CaseFieldValuesRequestBuilder {
	return &CaseFieldValuesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, itemFieldValuesURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve field values.
func (rB *CaseFieldValuesRequestBuilder) Get(ctx context.Context, config *CaseFieldValuesRequestBuilderGetRequestConfiguration) (FieldValuesResponse, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateFieldValuesResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(FieldValuesResponse), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *CaseFieldValuesRequestBuilder) ToGetRequestInformation(ctx context.Context, config *CaseFieldValuesRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
