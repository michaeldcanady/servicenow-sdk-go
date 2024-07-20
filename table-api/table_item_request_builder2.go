package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intTable "github.com/michaeldcanady/servicenow-sdk-go/table-api/internal"
)

const (
	tableItemURLTemplate = "{+baseurl}/table{/table}{/sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_input_display_value,sysparm_query_no_domain,sysparm_view,sysparm_query_no_domain}"
)

type TableItemRequestBuilder2 struct {
	intTable.RequestBuilder
}

// NewTableItemRequestBuilder2 creates a new instance of TableItemRequestBuilder2.
func NewTableItemRequestBuilder2(client core.Client, pathParameters map[string]string) (*TableItemRequestBuilder2, error) {
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

	return &TableItemRequestBuilder2{
		RequestBuilder: core.NewRequestBuilder(client, tableItemURLTemplate, pathParameters), //nolint:staticcheck
	}, nil
}

// TableItemService provides an interface for the service methods, adhering to the Interface Segregation Principle.
type TableItemService interface {
	Get(params *TableItemRequestBuilderGetQueryParameters) (*TableItemResponse, error)
	Delete(params *TableItemRequestBuilderDeleteQueryParameters) error
	Put(tableEntry interface{}, params *TableItemRequestBuilderPutQueryParameters) (*TableItemResponse, error)
}

// Ensure that TableItemRequestBuilder implements TableItemService.
var _ TableItemService = (*TableItemRequestBuilder2)(nil)

// Get sends an HTTP GET request using the specified query parameters and returns a TableItemResponse.
//
// Parameters:
//   - params: An instance of TableItemRequestBuilderGetQueryParameters to include in the GET request.
//
// Returns:
//   - *TableItemResponse: The response data as a TableItemResponse.
//   - error: An error if there was an issue with the request or response.
func (rB *TableItemRequestBuilder2) Get(params *TableItemRequestBuilderGetQueryParameters) (*TableItemResponse, error) {
	config := &tableItemGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		response: &TableItemResponse2[TableEntry]{},
	}

	err := rB.SendGet2(config.toConfiguration()) //nolint:staticcheck
	if err != nil {
		return nil, err
	}

	return config.response, nil
}

// Delete sends an HTTP DELETE request using the specified query parameters and returns an error if the request or response encounters any issues.
//
// Parameters:
//   - params: An instance of TableItemRequestBuilderDeleteQueryParameters to include in the DELETE request.
//
// Returns:
//   - error: An error if there was an issue with the request or response, or nil if the request was successful.
func (rB *TableItemRequestBuilder2) Delete(params *TableItemRequestBuilderDeleteQueryParameters) error {
	config := &tableItemDeleteRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		response: nil,
	}

	return rB.SendDelete2(config.toConfiguration())
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
func (rB *TableItemRequestBuilder2) Put(tableEntry interface{}, params *TableItemRequestBuilderPutQueryParameters) (*TableItemResponse, error) {
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

	err = rB.SendPut2(config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}
