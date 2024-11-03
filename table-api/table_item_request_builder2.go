package tableapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	tableItemURLTemplate = "{+baseurl}/table{/table}{/sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_input_display_value,sysparm_query_no_domain,sysparm_view,sysparm_query_no_domain}"
)

type TableItemRequestBuilder2 struct {
	abstractions.BaseRequestBuilder
	factory serialization.ParsableFactory
}

// NewTableItemRequestBuilder2Internal instantiates a new TableItemRequestBuilder2 and sets the default values.
func NewTableItemRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableItemRequestBuilder2 {
	m := &TableItemRequestBuilder2{
		factory:            factory,
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, tableItemURLTemplate, pathParameters),
	}
	return m
}

// NewTableItemRequestBuilder2 instantiates a new TableItemRequestBuilder2 and sets the default values.
func NewTableItemRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableItemRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawURL
	return NewTableItemRequestBuilder2Internal(urlParams, requestAdapter, factory)
}

func (rB *TableItemRequestBuilder2) Get(ctx context.Context, requestConfiguration *TableItemRequestBuilder2GetRequestConfiguration) (TableRecord, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.toGetRequestInformation(ctx, requestConfiguration)
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

func (rB *TableItemRequestBuilder2) Delete(ctx context.Context, requestConfiguration *TableItemRequestBuilder2DeleteRequestConfiguration) error {
	if internal.IsNil(rB) {
		return nil
	}

	requestInfo, err := rB.toDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	_, err = rB.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceNowResponseFromDiscriminatorValue(rB.factory), errorMapping)
	if err != nil {
		return err
	}

	return nil
}

func (rB *TableItemRequestBuilder2) Put(ctx context.Context, body TableRecord, requestConfiguration *TableItemRequestBuilder2PutRequestConfiguration) (TableRecord, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.toPutRequestInformation(ctx, body, requestConfiguration)
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

func (rB *TableItemRequestBuilder2) toGetRequestInformation(ctx context.Context, requestConfiguration *TableItemRequestBuilder2GetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			requestInfo.AddQueryParameters(*params)
		}
		kiotaRequestInfo.Headers.AddAll(requestConfiguration.Headers)
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.AddAll(requestConfiguration.Headers)
	kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	kiotaRequestInfo.Headers.TryAdd("Accept", "application/json")

	return &kiotaRequestInfo.RequestInformation, nil
}

func (rB *TableItemRequestBuilder2) toDeleteRequestInformation(ctx context.Context, requestConfiguration *TableItemRequestBuilder2DeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
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

	return &kiotaRequestInfo.RequestInformation, nil
}

func (rB *TableItemRequestBuilder2) toPutRequestInformation(ctx context.Context, body TableRecord, requestConfiguration *TableItemRequestBuilder2PutRequestConfiguration) (*abstractions.RequestInformation, error) {
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
