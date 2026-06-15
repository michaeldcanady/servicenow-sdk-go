package tableapi

import (
	"context"
	"errors"
	"fmt"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	tableItemURLTemplate2 = "{+baseurl}/api/now/v1/table/{table}/{sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view}"
)

// TableItemRequestBuilder2 provides operations to manage a single Service-Now table record.
type TableItemRequestBuilder2[T model.ServiceNowItem] struct {
	internal.RequestBuilder
	factory serialization.ParsableFactory
}

// NewTableItemRequestBuilder2Internal instantiates a new TableItemRequestBuilder2.
func NewTableItemRequestBuilder2Internal[T model.ServiceNowItem](
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableItemRequestBuilder2[T] {
	m := &TableItemRequestBuilder2[T]{
		RequestBuilder: internal.NewBaseRequestBuilder(requestAdapter, tableItemURLTemplate2, pathParameters),
		factory:        factory,
	}
	return m
}

// NewDefaultTableItemRequestBuilderInternal instantiates a new TableItemRequestBuilder2 with default table record parsable.
func NewDefaultTableItemRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *TableItemRequestBuilder2[*TableRecord] {
	return NewTableItemRequestBuilder2Internal[*TableRecord](pathParameters, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

func NewDefaultTableItemRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *TableItemRequestBuilder2[*TableRecord] {
	return NewTableItemRequestBuilder3[*TableRecord](rawURL, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// NewTableItemRequestBuilder3 instantiates a new TableItemRequestBuilder2 with a raw URL and custom parsable.
func NewTableItemRequestBuilder3[T model.ServiceNowItem](
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableItemRequestBuilder2[T] {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewTableItemRequestBuilder2Internal[T](urlParams, requestAdapter, factory)
}

// Get sends an HTTP GET request and returns the single table record.
func (rB *TableItemRequestBuilder2[T]) Get(ctx context.Context, requestConfiguration *TableItemRequestBuilder2GetRequestConfiguration) (internal.ServiceNowItemResponse[T], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
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

// Delete sends an HTTP DELETE request to remove the single table record.
func (rB *TableItemRequestBuilder2[T]) Delete(ctx context.Context, requestConfiguration *TableItemRequestBuilder2DeleteRequestConfiguration) error {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": internal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	return rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, errorMapping)
}

// Put updates a table item using an HTTP PUT request, replacing the entire record.
func (rB *TableItemRequestBuilder2[T]) Put(ctx context.Context, body T, requestConfiguration *TableItemRequestBuilder2PutRequestConfiguration) (internal.ServiceNowItemResponse[T], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(body) {
		return nil, errors.New("body is nil")
	}

	requestInfo, err := rB.ToPutRequestInformation(ctx, body, requestConfiguration)
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

// Patch updates a table item using an HTTP PATCH request, applying partial updates.
func (rB *TableItemRequestBuilder2[T]) Patch(ctx context.Context, body T, requestConfiguration *TableItemRequestBuilder2PatchRequestConfiguration) (internal.ServiceNowItemResponse[T], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(body) {
		return nil, errors.New("body is nil")
	}

	requestInfo, err := rB.ToPatchRequestInformation(ctx, body, requestConfiguration)
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

// ToGetRequestInformation converts provided parameters into request information
func (rB *TableItemRequestBuilder2[T]) ToGetRequestInformation(_ context.Context, requestConfiguration *TableItemRequestBuilder2GetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !conversion.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}

// ToDeleteRequestInformation converts provided parameters into request information
func (rB *TableItemRequestBuilder2[T]) ToDeleteRequestInformation(_ context.Context, requestConfiguration *TableItemRequestBuilder2DeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !conversion.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}

// ToPutRequestInformation converts provided parameters into request information
func (rB *TableItemRequestBuilder2[T]) ToPutRequestInformation(ctx context.Context, body T, requestConfiguration *TableItemRequestBuilder2PutRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !conversion.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	if !conversion.IsNil(body) {
		if err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body); err != nil {
			return nil, err
		}
	}

	return kiotaRequestInfo.RequestInformation, nil
}

// ToPatchRequestInformation converts provided parameters into request information
func (rB *TableItemRequestBuilder2[T]) ToPatchRequestInformation(ctx context.Context, body T, requestConfiguration *TableItemRequestBuilder2PatchRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PATCH, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParams := requestConfiguration.QueryParameters; !conversion.IsNil(queryParams) {
			kiotaRequestInfo.AddQueryParameters(queryParams)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	if !conversion.IsNil(body) {
		if err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body); err != nil {
			return nil, err
		}
	}

	return kiotaRequestInfo.RequestInformation, nil
}
