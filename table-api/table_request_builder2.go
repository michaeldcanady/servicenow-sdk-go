//go:build preview.tableApiV2

package tableapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

const (
	// batchURLTemplate the url template for Service-Now batch API
	batchURLTemplate = "{+baseurl}/api/now/v1/batch"
)

type TableRequestBuilder2[T serialization.Parsable, V any] struct {
	newInternal.RequestBuilder
	constructor serialization.ParsableFactory
}

// NewTableRequestBuilder2Internal instantiates a new TableRequestBuilder2[T] with custom parsable for table entries.
func NewTableRequestBuilder2Internal[T serialization.Parsable, V any](
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
	constructor serialization.ParsableFactory,
) *TableRequestBuilder2[T, V] {
	m := &TableRequestBuilder2[T, V]{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, batchURLTemplate, pathParameters),
		constructor:    constructor,
	}
	return m
}

func (builder *TableRequestBuilder2[T, V]) Get(ctx context.Context, requestConfiguration *abstractions.RequestConfiguration[V]) (*newInternal.BaseServiceNowCollectionResponse[T], error) {
	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &abstractions.RequestConfiguration[V]{}
	}

	opts := nethttplibrary.NewHeadersInspectionOptions()
	opts.InspectResponseHeaders = true

	requestConfiguration.Options = append(requestConfiguration.Options, opts)

	requestInfo, err := builder.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	res, err := builder.GetRequestAdapter().Send(ctx, requestInfo, builder.constructor, errorMapping)
	if err != nil {
		return nil, err
	}

	if internal.IsNil(res) {
		return nil, nil
	}

	snRes, ok := res.(*newInternal.BaseServiceNowCollectionResponse[T])
	if !ok {
		return nil, errors.New("res is not *AttachmentCollectionResponse2Model")
	}

	return snRes, nil
}

func (builder *TableRequestBuilder2[T, V]) ToGetRequestInformation(_ context.Context, requestConfiguration *abstractions.RequestConfiguration[V]) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, builder.GetURLTemplate(), builder.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if parameter := requestConfiguration.QueryParameters; !internal.IsNil(parameter) {
			kiotaRequestInfo.AddQueryParameters(parameter)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return kiotaRequestInfo.RequestInformation, nil
}
