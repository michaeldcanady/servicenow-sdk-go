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
	tableItemURLTemplate = "{+baseurl}/api/now/v2/table{/table}{/sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_input_display_value,sysparm_query_no_domain,sysparm_view,sysparm_query_no_domain}"
)

// TableItemRequestBuilder2 provides operations to manage Service-Now table entries.
type TableItemRequestBuilder2 struct {
	abstractions.BaseRequestBuilder
	factory serialization.ParsableFactory
}

// NewDefaultTableItemRequestBuilder2Internal instantiates a new TableItemRequestBuilder2 and sets the default values.
func NewDefaultTableItemRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *TableItemRequestBuilder2 {
	return newTableItemRequestBuilder2Internal(pathParameters, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// newTableItemRequestBuilder2Internal instantiates a new TableItemRequestBuilder2 with custom parsable for table entry.
func newTableItemRequestBuilder2Internal(
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

// NewDefaultTableItemRequestBuilder2 instantiates a new TableItemRequestBuilder2 and sets the default values.
func NewDefaultTableItemRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *TableItemRequestBuilder2 {
	return newTableItemRequestBuilder2(rawURL, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// newTableItemRequestBuilder2 instantiates a new TableItemRequestBuilder2 with custom parsable for table entry.
func newTableItemRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableItemRequestBuilder2 {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawURL
	return newTableItemRequestBuilder2Internal(urlParams, requestAdapter, factory)
}

// Get Fetches a Table Record resource.
func (rB *TableItemRequestBuilder2) Get(ctx context.Context, requestConfiguration *TableItemRequestBuilder2GetRequestConfiguration) (TableRecord, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
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

// Delete Deletes a Table Record resource.
func (rB *TableItemRequestBuilder2) Delete(ctx context.Context, requestConfiguration *TableItemRequestBuilder2DeleteRequestConfiguration) error {
	if internal.IsNil(rB) {
		return nil
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	return rB.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
}

// Put Updates a Table Record resource.
func (rB *TableItemRequestBuilder2) Put(ctx context.Context, body TableRecord, requestConfiguration *TableItemRequestBuilder2PutRequestConfiguration) (TableRecord, error) { //nolint:dupl
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo, err := rB.ToPutRequestInformation(ctx, body, requestConfiguration)
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

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *TableItemRequestBuilder2) ToGetRequestInformation(_ context.Context, requestConfiguration *TableItemRequestBuilder2GetRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:unparam
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			requestInfo.AddQueryParameters(*params)
		}
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.TryAdd("Accept", "application/json")

	return &kiotaRequestInfo.RequestInformation, nil
}

// ToDeleteRequestInformation converts request configurations to Delete request information.
func (rB *TableItemRequestBuilder2) ToDeleteRequestInformation(_ context.Context, requestConfiguration *TableItemRequestBuilder2DeleteRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:unparam
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			kiotaRequestInfo.AddQueryParameters(*params)
		}
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.AddAll(requestConfiguration.Headers)
	kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	kiotaRequestInfo.Headers.TryAdd("Accept", "application/json")

	return &kiotaRequestInfo.RequestInformation, nil
}

// ToPutRequestInformation converts request configurations to Put request information.
func (rB *TableItemRequestBuilder2) ToPutRequestInformation(ctx context.Context, body TableRecord, requestConfiguration *TableItemRequestBuilder2PutRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:dupl
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.UrlTemplate, rB.PathParameters)
	kiotaRequestInfo := &intHttp.KiotaRequestInformation{RequestInformation: *requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if params := requestConfiguration.QueryParameters; !internal.IsNil(params) {
			kiotaRequestInfo.AddQueryParameters(*params)
		}
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		kiotaRequestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	kiotaRequestInfo.Headers.TryAdd("Accept", "application/json")

	err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}

	return &kiotaRequestInfo.RequestInformation, nil
}
