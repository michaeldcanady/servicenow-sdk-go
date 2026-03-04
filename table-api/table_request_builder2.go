package tableapi

import (
	"context"
	"errors"
	"fmt"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	// batchURLTemplate the url template for Service-Now batch API
	tableURLTemplate2 = "{+baseurl}/api/now/v1/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

// TableRequestBuilder2 provides operations to manage Service-Now table collections.
type TableRequestBuilder2[T model.ServiceNowItem] struct {
	newInternal.RequestBuilder
	factory serialization.ParsableFactory
}

// NewTableRequestBuilder2Internal instantiates a new TableRequestBuilder2 with custom parsable for table entries.
func NewTableRequestBuilder2Internal[T model.ServiceNowItem](
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder2[T] {
	m := &TableRequestBuilder2[T]{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, tableURLTemplate2, pathParameters),
		factory:        factory,
	}
	return m
}

// NewDefaultTableRequestBuilder2Internal instantiates a new TableRequestBuilder2 with default table record parsable.
func NewDefaultTableRequestBuilder2Internal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *TableRequestBuilder2[*TableRecord] {
	return NewTableRequestBuilder2Internal[*TableRecord](pathParameters, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// NewDefaultTableRequestBuilder2 instantiates a new TableRequestBuilder2 with a raw URL and default table record parsable.
func NewDefaultTableRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *TableRequestBuilder2[*TableRecord] {
	return NewTableRequestBuilder2[*TableRecord](rawURL, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// NewTableRequestBuilder2 instantiates a new TableRequestBuilder2 with a raw URL and custom parsable.
func NewTableRequestBuilder2[T model.ServiceNowItem](
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder2[T] {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewTableRequestBuilder2Internal[T](urlParams, requestAdapter, factory)
}

// Get sends an HTTP GET request and returns a collection of table records.
func (rB *TableRequestBuilder2[T]) Get(ctx context.Context, requestConfiguration *TableRequestBuilder2GetRequestConfiguration) (newInternal.ServiceNowCollectionResponse[T], error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &TableRequestBuilder2GetRequestConfiguration{}
	}

	headerOpt := nethttplibrary.NewHeadersInspectionOptions()
	headerOpt.InspectResponseHeaders = true

	//if existingOpts := requestConfiguration.Options; len(requestConfiguration.Options) > 0 {

	//}

	requestConfiguration.Options = append(requestConfiguration.Options, headerOpt)

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowCollectionResponseFromDiscriminatorValue[T](rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("response is nil")
	}

	typedResp, ok := resp.(newInternal.ServiceNowCollectionResponse[T])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*newInternal.ServiceNowCollectionResponse[T])(nil))
	}

	typedResp.ParseHeaders(headerOpt.GetResponseHeaders())

	return typedResp, nil
}

// Post sends an HTTP POST request to create a new table record and returns the created record.
func (rB *TableRequestBuilder2[T]) Post(ctx context.Context, body T, requestConfiguration *TableRequestBuilder2PostRequestConfiguration) (newInternal.ServiceNowItemResponse[T], error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if internal.IsNil(body) {
		return nil, errors.New("body is nil")
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowItemResponseFromDiscriminatorValue[T](rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("response is nil")
	}

	typedResp, ok := resp.(newInternal.ServiceNowItemResponse[T])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*newInternal.ServiceNowItemResponse[T])(nil))
	}

	return typedResp, nil
}

// ById returns a TableItemRequestBuilder2 for the specified sysId.
func (rB *TableRequestBuilder2[T]) ById(sysId string) *TableItemRequestBuilder2[T] {
	pathParameters := make(map[string]string)
	for k, v := range rB.GetPathParameters() {
		pathParameters[k] = v
	}
	pathParameters["sysId"] = sysId
	return NewTableItemRequestBuilder2Internal[T](pathParameters, rB.GetRequestAdapter(), rB.factory)
}

// TODO: Add Head method which returns headers

// ToGetRequestInformation converts provided parameters into request information
func (rB *TableRequestBuilder2[T]) ToGetRequestInformation(_ context.Context, requestConfiguration *TableRequestBuilder2GetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !internal.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}

// ToPostRequestInformation converts provided parameters into request information
func (rB *TableRequestBuilder2[T]) ToPostRequestInformation(ctx context.Context, body T, requestConfiguration *TableRequestBuilder2PostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !internal.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	if !internal.IsNil(body) {
		if err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, body); err != nil {
			return nil, err
		}
	}

	return kiotaRequestInfo.RequestInformation, nil
}

// ToHeadRequestInformation converts provided parameters into request information
func (rB *TableRequestBuilder2[T]) ToHeadRequestInformation(_ context.Context, requestConfiguration *TableRequestBuilder2GetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.HEAD, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
