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
	tableItemURLTemplate2 = "{+baseurl}/api/now/v1/table/{/table}/{sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view}"
)

// TableItemRequestBuilder provides operations to manage a single Service-Now table record.
type TableItemRequestBuilder[T model.ServiceNowItem] struct {
	kiota.RequestBuilder
	factory serialization.ParsableFactory
}

// NewTableItemRequestBuilderInternal instantiates a new TableItemRequestBuilder.
func NewTableItemRequestBuilderInternal[T model.ServiceNowItem](
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableItemRequestBuilder[T] {
	m := &TableItemRequestBuilder[T]{
		RequestBuilder: kiota.NewBaseRequestBuilder(requestAdapter, tableItemURLTemplate2, pathParameters),
		factory:        factory,
	}
	return m
}

// NewDefaultTableItemRequestBuilderInternal instantiates a new TableItemRequestBuilder with default table record parsable.
func NewDefaultTableItemRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *TableItemRequestBuilder[*TableRecord] {
	return NewTableItemRequestBuilderInternal[*TableRecord](pathParameters, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

func NewDefaultTableItemRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *TableItemRequestBuilder[*TableRecord] {
	return NewTableItemRequestBuilder[*TableRecord](rawURL, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// NewTableItemRequestBuilder instantiates a new TableItemRequestBuilder with a raw URL and custom parsable.
func NewTableItemRequestBuilder[T model.ServiceNowItem](
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableItemRequestBuilder[T] {
	urlParams := make(map[string]string)
	urlParams[utils.RawURLKey] = rawURL
	return NewTableItemRequestBuilderInternal[T](urlParams, requestAdapter, factory)
}

// Get sends an HTTP GET request and returns the single table record.
func (rB *TableItemRequestBuilder[T]) Get(ctx context.Context, requestConfiguration *TableItemRequestBuilderGetRequestConfiguration) (newInternal.ServiceNowItemResponse[T], error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": model.CreateServiceNowErrorFromDiscriminatorValue,
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

// Delete sends an HTTP DELETE request to remove the single table record.
func (rB *TableItemRequestBuilder[T]) Delete(ctx context.Context, requestConfiguration *TableItemRequestBuilderDeleteRequestConfiguration) error {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

// Put updates a table item using an HTTP PUT request, replacing the entire record.
func (rB *TableItemRequestBuilder[T]) Put(ctx context.Context, body T, requestConfiguration *TableItemRequestBuilderPutRequestConfiguration) (newInternal.ServiceNowItemResponse[T], error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if utils.IsNil(body) {
		return nil, errors.New("body is nil")
	}

	requestInfo, err := rB.ToPutRequestInformation(ctx, body, requestConfiguration)
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

// Patch updates a table item using an HTTP PATCH request, applying partial updates.
func (rB *TableItemRequestBuilder[T]) Patch(ctx context.Context, body T, requestConfiguration *TableItemRequestBuilderPatchRequestConfiguration) (newInternal.ServiceNowItemResponse[T], error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if utils.IsNil(body) {
		return nil, errors.New("body is nil")
	}

	requestInfo, err := rB.ToPatchRequestInformation(ctx, body, requestConfiguration)
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

// ToGetRequestInformation converts provided parameters into request information
func (rB *TableItemRequestBuilder[T]) ToGetRequestInformation(_ context.Context, requestConfiguration *TableItemRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
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

// ToDeleteRequestInformation converts provided parameters into request information
func (rB *TableItemRequestBuilder[T]) ToDeleteRequestInformation(_ context.Context, requestConfiguration *TableItemRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
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

// ToPutRequestInformation converts provided parameters into request information
func (rB *TableItemRequestBuilder[T]) ToPutRequestInformation(ctx context.Context, body T, requestConfiguration *TableItemRequestBuilderPutRequestConfiguration) (*abstractions.RequestInformation, error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
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

// ToPatchRequestInformation converts provided parameters into request information
func (rB *TableItemRequestBuilder[T]) ToPatchRequestInformation(ctx context.Context, body T, requestConfiguration *TableItemRequestBuilderPatchRequestConfiguration) (*abstractions.RequestInformation, error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PATCH, rB.GetURLTemplate(), rB.GetPathParameters())
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
