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
)

const (
	// batchURLTemplate the url template for Service-Now batch API
	tableURLTemplate2 = "{+baseurl}/api/now/v1/table/{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

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
		newInternal.NewBaseRequestBuilder(requestAdapter, tableURLTemplate2, pathParameters),
		factory,
	}
	return m
}

func NewDefaultTableRequestBuilder2Internal[T model.ServiceNowItem](
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *TableRequestBuilder2[*TableRecord] {
	return NewTableRequestBuilder2Internal[*TableRecord](pathParameters, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// NewTableRequestBuilder2 instantiates a new BatchRequestBuilder with custom parsable for table entries.
func NewDefaultTableRequestBuilder2(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *TableRequestBuilder2[*TableRecord] {
	return NewTableRequestBuilder2[*TableRecord](rawURL, requestAdapter, CreateTableRecordFromDiscriminatorValue)
}

// NewTableRequestBuilder2 instantiates a new BatchRequestBuilder with custom parsable for table entries.
func NewTableRequestBuilder2[T model.ServiceNowItem](
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
	factory serialization.ParsableFactory,
) *TableRequestBuilder2[T] {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewTableRequestBuilder2Internal[T](urlParams, requestAdapter, factory)
}

func (rB *TableRequestBuilder2[T]) Post(ctx context.Context, body T, requestConfiguration *TableRequestBuilder2GetRequestConfiguration) (newInternal.ServiceNowItemResponse[T], error) {
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

// TODO: Add Head method which returns headers

// ToPostRequestInformation converts provided parameters into request information
func (rB *TableRequestBuilder2[T]) ToPostRequestInformation(ctx context.Context, body T, requestConfiguration *TableRequestBuilder2GetRequestConfiguration) (*abstractions.RequestInformation, error) {
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
