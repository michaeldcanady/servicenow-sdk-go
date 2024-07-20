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

// TableRequestBuilder is responsible for building requests for table operations.
type TableRequestBuilder2 struct {
	intCore.RequestBuilder2
}

// NewTableRequestBuilder initializes a new TableRequestBuilder with the given client and path parameters.
func NewTableRequestBuilder2(client intCore.Client2, pathParameters map[string]string) (*TableRequestBuilder2, error) {
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

	return &TableRequestBuilder2{
		intCore.NewRequestBuilder2(client, tableURLTemplate, pathParameters),
	}, nil
}

// ByID creates a TableItemRequestBuilder for a specific record in the table identified by sysID.
func (rB *TableRequestBuilder2) ByID(sysID string) (*TableItemRequestBuilder2, error) {
	pathParameters := rB.GetPathParameters()
	client := rB.GetClient()
	pathParameters["sysId"] = sysID
	return NewTableItemRequestBuilder2(client, pathParameters)
}

// Get retrieves a collection of table items based on the provided query parameters.
func (rB *TableRequestBuilder2) Get(ctx context.Context, opts ...intCore.RequestConfigurationOption) (*TableCollectionResponse, error) {
	opts = append(opts, intCore.WithResponse(&TableCollectionResponse2[TableEntry]{}))

	resp, err := rB.Send(ctx, intCore.GET, opts...)
	if err != nil {
		return nil, err
	}

	return resp.(*TableCollectionResponse), nil
}

// Post creates a new table item with the provided data and query parameters.
func (rB *TableRequestBuilder2) Post(ctx context.Context, opts ...intCore.RequestConfigurationOption) (*TableItemResponse2[TableEntry], error) {
	opts = append(opts, intCore.WithResponse(&TableItemResponse2[TableEntry]{}))

	resp, err := rB.Send(ctx, intCore.POST, opts...)
	if err != nil {
		return nil, err
	}

	return resp.(*TableItemResponse2[TableEntry]), nil
}

// Head creates a new table request with the provided query parameters and returns the header values
func (rB *TableRequestBuilder2) Head(ctx context.Context, opts ...intCore.RequestConfigurationOption) (*TableCollectionResponse, error) {
	opts = append(opts, intCore.WithResponse(&TableCollectionResponse2[TableEntry]{}))

	resp, err := rB.Send(ctx, intCore.HEAD, opts...)
	if err != nil {
		return nil, err
	}

	return resp.(*TableCollectionResponse), nil
}
