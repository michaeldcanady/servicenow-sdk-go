package tableapi

import (
	"context"
	"errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	tableURLTemplate = "{+baseurl}/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

type TableRequestBuilder2 struct {
	abstractions.BaseRequestBuilder
	factory serialization.ParsableFactory
}

// NewRequestBuilder2Internal instantiates a new TableRequestBuilderKiota and sets the default values.
func NewRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder2 {
	m := &TableRequestBuilder2{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, tableURLTemplate, pathParameters),
		factory:            factory,
	}
	return m
}

// NewRequestBuilderBuilder2 instantiates a new TableRequestBuilderKiota and sets the default values.
func NewRequestBuilderBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawURL
	return NewRequestBuilder2Internal(urlParams, requestAdapter, factory)
}

func (rB *TableRequestBuilder2) ByID(sysID string) *TableItemRequestBuilder2 {
	if internal.IsNil(rB) {
		return nil
	}

	pathParameters := maps.Clone(rB.BaseRequestBuilder.PathParameters)
	pathParameters["sysid"] = sysID

	return NewTableItemRequestBuilder2Internal(pathParameters, rB.BaseRequestBuilder.RequestAdapter, rB.factory)
}

func (rB *TableRequestBuilder2) Get(ctx context.Context, requestConfiguration *TableRequestBuilder2GetRequestConfiguration) ([]TableRecord, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.toGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	res, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceNowCollectionResponseFromDiscriminatorValue(rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	snRes, ok := res.(ServiceNowCollectionResponse)
	if !ok {
		return nil, errors.New("res is not ServiceNowResponse")
	}

	result, err := snRes.GetResult()
	if err != nil {
		return nil, err
	}

	records, ok := interface{}(result).([]TableRecord)
	if !ok {
		return nil, errors.New("result is not TableRecord")
	}

	return records, nil
}

func (rB *TableRequestBuilder2) Post(ctx context.Context, body TableRecord, requestConfiguration *TableRequestBuilder2PostRequestConfiguration) (TableRecord, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.toPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	res, err := rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceNowResponseFromDiscriminatorValue(rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	snRes, ok := res.(ServiceNowResponse)
	if !ok {
		return nil, errors.New("res is not ServiceNowResponse")
	}

	result, err := snRes.GetResult()
	if err != nil {
		return nil, err
	}

	record, ok := result.(TableRecord)
	if !ok {
		return nil, errors.New("result is not TableRecord")
	}

	return record, nil
}

func (rB *TableRequestBuilder2) toGetRequestInformation(ctx context.Context, requestConfiguration *TableRequestBuilder2GetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			kiotaRequestInfo.AddQueryParameters(*params)
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.AddAll(requestConfiguration.Headers)
	kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	kiotaRequestInfo.Headers.TryAdd("Accept", "application/json")

	return &kiotaRequestInfo.RequestInformation, nil
}

func (rB *TableRequestBuilder2) toPostRequestInformation(ctx context.Context, body TableRecord, requestConfiguration *TableRequestBuilder2PostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			kiotaRequestInfo.AddQueryParameters(*params)
		}
		kiotaRequestInfo.Headers.AddAll(requestConfiguration.Headers)
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.AddAll(requestConfiguration.Headers)
	kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	kiotaRequestInfo.Headers.TryAdd("Accept", "application/json")

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}

	return &kiotaRequestInfo.RequestInformation, nil
}
