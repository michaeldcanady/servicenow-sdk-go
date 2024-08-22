package tableapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	tableURLTemplate = "{+baseurl}/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"
)

// TableRequestBuilder2 defines the operations available for the table as a whole.
type TableRequestBuilder2 interface {
	Get(context.Context, *TableRequestBuilderGetQueryParameters) (TableCollectionResponse3[*tableRecord], error)
	ByID(string) (TableItemRequestBuilder2, error)
	Post(context.Context, interface{}, *TableRequestBuilderPostQueryParameters) (TableItemResponse3[*tableRecord], error)
	Count(context.Context, *TableRequestBuilderGetQueryParameters) (int, error)
}

// TableRequestBuilder is responsible for building requests for table operations.
type tableRequestBuilder2 struct {
	intCore.RequestBuilder2
}

// NewTableRequestBuilder initializes a new TableRequestBuilder with the given client and path parameters.
func NewTableRequestBuilder2(client intCore.Client2, pathParameters map[string]string) (TableRequestBuilder2, error) {
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

	return &tableRequestBuilder2{
		intCore.NewRequestBuilder2(client, tableURLTemplate, pathParameters),
	}, nil
}

// ByID creates a TableItemRequestBuilder for a specific record in the table identified by sysID.
func (rB *tableRequestBuilder2) ByID(sysID string) (TableItemRequestBuilder2, error) {
	pathParameters := rB.GetPathParameters()
	client := rB.GetClient()
	pathParameters["sysId"] = sysID
	return NewTableItemRequestBuilder2(client, pathParameters)
}

// Get retrieves a collection of table items based on the provided query parameters.
func (rB *tableRequestBuilder2) Get(ctx context.Context, params *TableRequestBuilderGetQueryParameters) (TableCollectionResponse3[*tableRecord], error) {
	config := &tableGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		mapping:  nil,
		response: TableCollectionResponse3[TableEntry]{},
	}

	resp, err := rB.Send(ctx, intCore.MethodGet, config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return resp.(TableCollectionResponse3[*tableRecord]), nil
}

// Post creates a new table item with the provided data and query parameters.
func (rB *tableRequestBuilder2) Post(ctx context.Context, data interface{}, params *TableRequestBuilderPostQueryParameters) (TableItemResponse3[*tableRecord], error) {
	data, err := convertFromTableEntry(data)
	if err != nil {
		return nil, err
	}

	config := &tablePostRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     data.(map[string]string),
		mapping:  nil,
		response: TableItemResponse3[*tableRecord]{},
	}

	resp, err := rB.Send(ctx, intCore.MethodPost, config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return resp.(TableItemResponse3[*tableRecord]), nil
}

// Count retrieves the total count of items in the table.
func (rB *tableRequestBuilder2) Count(ctx context.Context, params *TableRequestBuilderGetQueryParameters) (int, error) {

	config := &tableGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		mapping:  nil,
		response: TableCollectionResponse3[TableEntry]{},
	}

	resp, err := rB.Send(ctx, intCore.MethodHead, config.toConfiguration())
	if err != nil {
		return -1, err
	}

	return resp.(TableCollectionResponse3[*tableRecord]).GetCount(), nil
}
