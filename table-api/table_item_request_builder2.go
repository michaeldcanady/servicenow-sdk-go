package tableapi

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	tableItemURLTemplate = "{+baseurl}/table{/table}{/sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_input_display_value,sysparm_query_no_domain,sysparm_view,sysparm_query_no_domain}"
)

// TableItemRequestBuilder2 provides an interface for the service methods, adhering to the Interface Segregation Principle.
type TableItemRequestBuilder2 interface {
	Get(context.Context, *TableItemRequestBuilderGetQueryParameters) (*TableItemResponse, error)
	Delete(context.Context, *TableItemRequestBuilderDeleteQueryParameters) error
	Put(context.Context, interface{}, *TableItemRequestBuilderPutQueryParameters) (*TableItemResponse, error)
}

type tableItemRequestBuilder2 struct {
	intCore.Sendable
}

// NewTableItemRequestBuilder2 creates a new instance of TableItemRequestBuilder2.
func NewTableItemRequestBuilder2(client intCore.SendableWithContext, pathParameters map[string]string) (TableItemRequestBuilder2, error) {
	if internal.IsNil(client) {
		return nil, ErrNilClient
	}

	_, basePathOk := pathParameters[internal.BasePathParameter]
	if !basePathOk {
		return nil, core.ErrMissingBasePathParam
	}

	_, tableOk := pathParameters["table"]
	if !tableOk {
		return nil, ErrNilParameterTable
	}

	_, sysIDOk := pathParameters["sysId"]
	if !sysIDOk {
		return nil, ErrNilParameterSysID
	}

	return &tableItemRequestBuilder2{
		intCore.NewRequestBuilder2(client, tableURLTemplate, pathParameters),
	}, nil
}

// Get sends an HTTP GET request using the specified query parameters and returns a TableItemResponse.
//
// Parameters:
//   - params: An instance of TableItemRequestBuilderGetQueryParameters to include in the GET request.
//
// Returns:
//   - *TableItemResponse: The response data as a TableItemResponse.
//   - error: An error if there was an issue with the request or response.
func (rB *tableItemRequestBuilder2) Get(ctx context.Context, params *TableItemRequestBuilderGetQueryParameters) (*TableItemResponse, error) {
	config := &tableItemGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		response: &TableItemResponse2[TableEntry]{},
	}

	resp, err := rB.Send(ctx, intCore.MethodGet, config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return resp.(*TableItemResponse2[TableEntry]), nil
}

// Delete sends an HTTP DELETE request using the specified query parameters and returns an error if the request or response encounters any issues.
//
// Parameters:
//   - params: An instance of TableItemRequestBuilderDeleteQueryParameters to include in the DELETE request.
//
// Returns:
//   - error: An error if there was an issue with the request or response, or nil if the request was successful.
func (rB *tableItemRequestBuilder2) Delete(ctx context.Context, params *TableItemRequestBuilderDeleteQueryParameters) error {
	config := &tableItemDeleteRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		response: nil,
	}

	_, err := rB.Send(ctx, intCore.MethodDelete, config.toConfiguration())

	return err
}

// Put updates a table item using an HTTP PUT request.
// It takes a map of table entry data and optional query parameters to send in the request.
// The method returns a TableItemResponse representing the updated item or an error if the request fails.
//
// Parameters:
//   - tableEntry: A map[string]string or TableEntry containing the data to update the table item.
//   - params: An optional pointer to TableItemRequestBuilderPutQueryParameters, which can be used to specify query parameters for the request.
//
// Returns:
//   - *TableItemResponse: A TableItemResponse containing the updated item data.
//   - error: An error, if the request fails at any point, such as request information creation or JSON deserialization.
func (rB *tableItemRequestBuilder2) Put(ctx context.Context, tableEntry interface{}, params *TableItemRequestBuilderPutQueryParameters) (*TableItemResponse, error) {
	tableEntry, err := convertFromTableEntry(tableEntry)
	if err != nil {
		return nil, err
	}

	config := &tableItemPutRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     tableEntry,
		response: &TableItemResponse2[TableEntry]{},
	}

	resp, err := rB.Send(ctx, intCore.MethodPut, config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return resp.(*TableItemResponse2[TableEntry]), nil
}
