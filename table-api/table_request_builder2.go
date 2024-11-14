package tableapi

import (
	"context"
	"errors"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/RecoLabs/servicenow-sdk-go/internal"
	intCore "github.com/RecoLabs/servicenow-sdk-go/internal/core"
)

const (
	tableURLTemplate = "{+baseurl}/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

// TableRequestBuilder2 defines the operations available for the table as a whole.
type TableRequestBuilder2[T TableRecord] interface {
	Get(context.Context, *TableRequestBuilderGetQueryParameters) (TableCollectionResponse3[T], error)
	ByID(string) (TableItemRequestBuilder2[T], error)
	Post(context.Context, interface{}, *TableRequestBuilderPostQueryParameters) (TableItemResponse3[T], error)
	Count(context.Context, *TableRequestBuilderGetQueryParameters) (int, error)
}

// TableRequestBuilder is responsible for building requests for table operations.
type tableRequestBuilder2[T TableRecord] struct {
	intCore.RequestBuilder2
}

func NewDefaultTableRequestBuilder2(client intCore.ClientSendable, pathParameters map[string]string) (TableRequestBuilder2[*TableRecordImpl], error) {
	return newTableRequestBuilder2[*TableRecordImpl](client, pathParameters)
}

// NewTableRequestBuilder initializes a new TableRequestBuilder with the given client and path parameters.
func newTableRequestBuilder2[T TableRecord](client intCore.ClientSendable, pathParameters map[string]string) (TableRequestBuilder2[T], error) {
	if internal.IsNil(client) {
		return nil, ErrNilClient
	}

	_, basePathOk := pathParameters[internal.BasePathParameter]
	if !basePathOk {
		return nil, core.ErrMissingBasePathParam
	}
	_, tableOk := pathParameters["table"]
	if !tableOk {
		return nil, errors.New("missing \"table\" parameter")
	}

	return &tableRequestBuilder2[T]{
		intCore.NewRequestBuilder2(client, tableURLTemplate, pathParameters),
	}, nil
}

// ByID creates a TableItemRequestBuilder for a specific record in the table identified by sysID.
func (rB *tableRequestBuilder2[T]) ByID(sysID string) (TableItemRequestBuilder2[T], error) {
	pathParameters := rB.GetPathParameters()
	client := rB.GetClient()
	pathParameters["sysId"] = sysID

	builder, err := newTableItemRequestBuilder2[T](client, pathParameters)

	return builder, err
}

// Get retrieves a collection of table items based on the provided query parameters.
func (rB *tableRequestBuilder2[T]) Get(ctx context.Context, params *TableRequestBuilderGetQueryParameters) (TableCollectionResponse3[T], error) {
	config := &intCore.RequestConfigurationImpl{
		Header:          nil,
		QueryParameters: interface{}(params),
		Data:            nil,
		ErrorMapping:    nil,
		Response:        &tableCollectionResponse3[T]{},
	}

	resp, err := rB.Send(
		ctx,
		intCore.MethodGet,
		config,
	)
	if err != nil {
		return nil, err
	}

	return resp.(TableCollectionResponse3[T]), nil
}

// Post creates a new table item with the provided data and query parameters.
func (rB *tableRequestBuilder2[T]) Post(ctx context.Context, data interface{}, params *TableRequestBuilderPostQueryParameters) (TableItemResponse3[T], error) {
	data, err := convertFromTableEntry(data)
	if err != nil {
		return nil, err
	}

	config := &intCore.RequestConfigurationImpl{
		Header:          nil,
		QueryParameters: interface{}(params),
		Data:            data.(map[string]string),
		ErrorMapping:    nil,
		Response:        &tableItemResponse3[T]{},
	}

	resp, err := rB.Send(ctx, intCore.MethodPost, config)
	if err != nil {
		return nil, err
	}

	return resp.(TableItemResponse3[T]), nil
}

// Count retrieves the total count of items in the table.
func (rB *tableRequestBuilder2[T]) Count(ctx context.Context, params *TableRequestBuilderGetQueryParameters) (int, error) {
	config := &intCore.RequestConfigurationImpl{
		Header:          nil,
		QueryParameters: interface{}(params),
		Data:            nil,
		ErrorMapping:    nil,
		Response:        &tableCollectionResponse3[T]{},
	}

	resp, err := rB.Send(ctx, intCore.MethodHead, config)
	if err != nil {
		return -1, err
	}

	return resp.(TableCollectionResponse3[T]).GetCount(), nil
}
