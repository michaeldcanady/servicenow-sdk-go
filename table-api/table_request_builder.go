package tableapi

import (
	"context"
	"errors"
	"fmt"

	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	// tableURLTemplate the url template for Service-Now batch API
	tableURLTemplate = "{+baseurl}/api/now/v1/table/{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

// TableRequestBuilder provides operations to manage Service-Now table collections.
type TableRequestBuilder[T model.ServiceNowItem] struct {
	kiota.RequestBuilder
	factory serialization.ParsableFactory
}

// NewTableRequestBuilderInternal instantiates a new TableRequestBuilder with custom parsable for table entries.
func NewTableRequestBuilderInternal[T model.ServiceNowItem](
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder[T] {
	m := &TableRequestBuilder[T]{
		RequestBuilder: kiota.NewBaseRequestBuilder(requestAdapter, tableURLTemplate, pathParameters),
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
	urlParams[utils.RawURLKey] = rawURL
	return NewTableRequestBuilderInternal[T](urlParams, requestAdapter, factory)
}

// Get sends an HTTP GET request and returns a collection of table records.
func (rB *TableRequestBuilder[T]) Get(ctx context.Context, requestConfiguration *TableRequestBuilderGetRequestConfiguration) (newInternal.ServiceNowCollectionResponse[T], error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, rB.factory, errorMapping)
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

	return typedResp, nil
}

// Post sends an HTTP POST request to create a new table record and returns the created record.
func (rB *TableRequestBuilder[T]) Post(ctx context.Context, body T, requestConfiguration *TableRequestBuilderPostRequestConfiguration) (newInternal.ServiceNowItemResponse[T], error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if utils.IsNil(body) {
		return nil, errors.New("body is nil")
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, rB.factory, errorMapping)
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

// ById returns a TableItemRequestBuilder for the specified sysId.
func (rB *TableRequestBuilder[T]) ById(sysId string) *TableItemRequestBuilder[T] {
	pathParameters := make(map[string]string)
	for k, v := range rB.GetPathParameters() {
		pathParameters[k] = v
	}
	pathParameters["sysId"] = sysId
	return NewTableItemRequestBuilderInternal[T](pathParameters, rB.GetRequestAdapter(), rB.factory)
}

// TODO: Add Head method which returns headers

// ToGetRequestInformation converts provided parameters into request information
func (rB *TableRequestBuilder[T]) ToGetRequestInformation(_ context.Context, requestConfiguration *TableRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &kiota.KiotaRequestInformation{RequestInformation: requestInfo}
	if !utils.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !utils.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !utils.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !utils.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), utils.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}

// ToPostRequestInformation converts provided parameters into request information
func (rB *TableRequestBuilder[T]) ToPostRequestInformation(ctx context.Context, body T, requestConfiguration *TableRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &kiota.KiotaRequestInformation{RequestInformation: requestInfo}
	if !utils.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !utils.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !utils.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !utils.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), utils.ContentTypeApplicationJSON)

	if !utils.IsNil(body) {
		if err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), utils.ContentTypeApplicationJSON, body); err != nil {
			return nil, err
		}
	}

	return kiotaRequestInfo.RequestInformation, nil
}

// ToHeadRequestInformation converts provided parameters into request information
func (rB *TableRequestBuilder[T]) ToHeadRequestInformation(_ context.Context, requestConfiguration *TableRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.HEAD, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &kiota.KiotaRequestInformation{RequestInformation: requestInfo}
	if !utils.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !utils.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !utils.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), utils.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}
