package tableapi

import (
	"context"
	"fmt"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	sysIDKey = "sysId"

	// batchURLTemplate the url template for Service-Now batch API
	tableURLTemplate = "{+baseurl}/api/now/v1/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

// TableRequestBuilder provides operations to manage Service-Now table collections.
type TableRequestBuilder[T model.ServiceNowItem] struct {
	core.RequestBuilder
	factory serialization.ParsableFactory
}

// NewTableRequestBuilderInternal instantiates a new TableRequestBuilder with custom parsable for table entries.
func NewTableRequestBuilderInternal[T model.ServiceNowItem](
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder[T] {
	m := &TableRequestBuilder[T]{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, tableURLTemplate, pathParameters),
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
func (rB *TableRequestBuilder[T]) Get(ctx context.Context, requestConfiguration *TableRequestBuilderGetRequestConfiguration) (core.ServiceNowCollectionResponse[T], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
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

	errorMapping := core.DefaultErrorMapping()
	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowCollectionResponseFromDiscriminatorValue[T](rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := resp.(core.ServiceNowCollectionResponse[T])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*core.ServiceNowCollectionResponse[T])(nil))
	}

	core.ParseHeaders(typedResp, headerOpt.GetResponseHeaders())

	return typedResp, nil
}

// Post sends an HTTP POST request to create a new table record and returns the created record.
func (rB *TableRequestBuilder[T]) Post(ctx context.Context, body T, requestConfiguration *TableRequestBuilderPostRequestConfiguration) (core.ServiceNowItemResponse[T], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	if conversion.IsNil(body) {
		return nil, snerrors.ErrNilBody
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[T](rB.factory), errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := resp.(core.ServiceNowItemResponse[T])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*core.ServiceNowItemResponse[T])(nil))
	}

	return typedResp, nil
}

// ByID returns a TableItemRequestBuilder for the specified sysId.
func (rB *TableRequestBuilder[T]) ByID(sysId string) *TableItemRequestBuilder[T] {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters[sysIDKey] = sysId
	return NewTableItemRequestBuilderInternal[T](pathParameters, rB.GetRequestAdapter(), rB.factory)
}

// Head sends an HTTP HEAD request and returns the response headers.
func (rB *TableRequestBuilder[T]) Head(ctx context.Context, requestConfiguration *TableRequestBuilderGetRequestConfiguration) (*abstractions.ResponseHeaders, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
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

	errorMapping := core.DefaultErrorMapping()
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

	internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return kiotaRequestInfo.RequestInformation, nil
}

// ToPostRequestInformation converts provided parameters into request information
func (rB *TableRequestBuilder[T]) ToPostRequestInformation(ctx context.Context, body T, requestConfiguration *TableRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	if !conversion.IsNil(body) {
		if err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body); err != nil {
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

	internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return kiotaRequestInfo.RequestInformation, nil
}
