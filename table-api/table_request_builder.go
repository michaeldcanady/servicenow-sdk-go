package tableapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	// batchURLTemplate the url template for Service-Now batch API
	tableURLTemplate = "{+baseurl}/api/now/v1/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

// TableRequestBuilder provides operations to manage Service-Now table collections.
type TableRequestBuilder[T model.ServiceNowItem] struct {
	internal.RequestBuilder
	factory serialization.ParsableFactory
}

// NewTableRequestBuilderInternal instantiates a new TableRequestBuilder with custom parsable for table entries.
func NewTableRequestBuilderInternal[T model.ServiceNowItem](
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder[T] {
	m := &TableRequestBuilder[T]{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, tableURLTemplate, pathParameters),
		factory:        factory,
	}
	return m
}

// NewDefaultTableRequestBuilderInternal instantiates a new TableRequestBuilder with default table record parsable.
func NewDefaultTableRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *TableRequestBuilder[*TableRecord] {
	return NewTableRequestBuilderInternal[*TableRecord](pathParameters, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// NewDefaultTableRequestBuilder instantiates a new TableRequestBuilder with a raw URL and default table record parsable.
func NewDefaultTableRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *TableRequestBuilder[*TableRecord] {
	return NewTableRequestBuilder[*TableRecord](rawURL, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// NewTableRequestBuilder instantiates a new TableRequestBuilder with a raw URL and custom parsable.
func NewTableRequestBuilder[T model.ServiceNowItem](
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder[T] {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewTableRequestBuilderInternal[T](urlParams, requestAdapter, factory)
}

// Get sends an HTTP GET request and returns a collection of table records.
func (rB *TableRequestBuilder[T]) Get(ctx context.Context, requestConfiguration *TableRequestBuilderGetRequestConfiguration) (internal.ServiceNowCollectionResponse[T], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(requestConfiguration) {
		requestConfiguration = &TableRequestBuilderGetRequestConfiguration{}
	}

	headerOpt := nethttplibrary.NewHeadersInspectionOptions()
	headerOpt.InspectResponseHeaders = true

	requestConfiguration.Options = append(requestConfiguration.Options, headerOpt)

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowCollectionResponseFromDiscriminatorValue[T](rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("response is nil")
	}

	typedResp, ok := resp.(internal.ServiceNowCollectionResponse[T])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*internal.ServiceNowCollectionResponse[T])(nil))
	}

	internal.ParseHeaders(typedResp, headerOpt.GetResponseHeaders())

	return typedResp, nil
}

// Post sends an HTTP POST request to create a new table record and returns the created record.
func (rB *TableRequestBuilder[T]) Post(ctx context.Context, body T, requestConfiguration *TableRequestBuilderPostRequestConfiguration) (internal.ServiceNowItemResponse[T], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(body) {
		return nil, errors.New("body is nil")
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowItemResponseFromDiscriminatorValue[T](rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("response is nil")
	}

	typedResp, ok := resp.(internal.ServiceNowItemResponse[T])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*internal.ServiceNowItemResponse[T])(nil))
	}

	return typedResp, nil
}

// ByID returns a TableItemRequestBuilder for the specified sysId.
func (rB *TableRequestBuilder[T]) ByID(sysId string) *TableItemRequestBuilder[T] {
	pathParameters := make(map[string]string)
	for k, v := range rB.GetPathParameters() {
		pathParameters[k] = v
	}
	pathParameters["sysId"] = sysId
	return NewTableItemRequestBuilderInternal[T](pathParameters, rB.GetRequestAdapter(), rB.factory)
}

// Head sends an HTTP HEAD request and returns the response headers.
func (rB *TableRequestBuilder[T]) Head(ctx context.Context, requestConfiguration *TableRequestBuilderGetRequestConfiguration) (*abstractions.ResponseHeaders, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(requestConfiguration) {
		requestConfiguration = &TableRequestBuilderGetRequestConfiguration{}
	}

	headerOpt := nethttplibrary.NewHeadersInspectionOptions()
	headerOpt.InspectResponseHeaders = true
	requestConfiguration.Options = append(requestConfiguration.Options, headerOpt)

	requestInfo, err := rB.ToHeadRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	if err := rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping); err != nil {
		return nil, err
	}

	return headerOpt.GetResponseHeaders(), nil
}

// ToGetRequestInformation converts provided parameters into request information
func (rB *TableRequestBuilder[T]) ToGetRequestInformation(_ context.Context, requestConfiguration *TableRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}

// ToPostRequestInformation converts provided parameters into request information
func (rB *TableRequestBuilder[T]) ToPostRequestInformation(ctx context.Context, body T, requestConfiguration *TableRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	if !conversion.IsNil(body) {
		if err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body); err != nil {
			return nil, err
		}
	}

	return kiotaRequestInfo.RequestInformation, nil
}

// ToHeadRequestInformation converts provided parameters into request information
func (rB *TableRequestBuilder[T]) ToHeadRequestInformation(_ context.Context, requestConfiguration *TableRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.HEAD, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
