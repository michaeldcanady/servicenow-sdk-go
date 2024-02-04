package tableapi

import (
	"context"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

type TableItemRequestBuilder struct {
	context.Context
	core.RequestBuilder
}

// NewTableItemRequestBuilder creates a new instance of the TableItemRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableItemRequestBuilder.
func NewTableItemRequestBuilder(ctx context.Context, client core.Client,
	pathParameters map[string]string) *TableItemRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		"{+baseurl}/table{/table}{/sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_input_display_value,sysparm_query_no_domain,sysparm_view,sysparm_query_no_domain}",
		pathParameters,
	)
	return &TableItemRequestBuilder{
		Context:        ctx,
		RequestBuilder: *requestBuilder,
	}
}

// Get sends an HTTP GET request using the specified query parameters and returns a TableItemResponse.
//
// Parameters:
//   - params: An instance of TableItemRequestBuilderGetQueryParameters to include in the GET request.
//
// Returns:
//   - *TableItemResponse: The response data as a TableItemResponse.
//   - error: An error if there was an issue with the request or response.
func (rB *TableItemRequestBuilder) Get(ctx context.Context, params *TableItemRequestBuilderGetQueryParameters) (
	*TableItemResponse, error) {
	config := &tableItemGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		response: &TableItemResponse2[TableEntry]{},
	}

	err := rB.SendGet2(ctx, config.toConfiguration()) //nolint:staticcheck
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
func (rB *TableItemRequestBuilder) Delete(ctx context.Context,
	params *TableItemRequestBuilderDeleteQueryParameters) error {
	config := &tableItemDeleteRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		response: nil,
	}

	return rB.SendDelete2(ctx, config.toConfiguration())
}

// Deprecated: deprecated since v1.4.0. Use `Put2` instead.
//
// Put updates a table item using an HTTP PUT request.
// It takes a map of table entry data and optional query parameters to send in the request.
// The method returns a TableItemResponse representing the updated item or an error if the request fails.
//
// Parameters:
//   - tableEntry: A map containing the data to update the table item.
//   - params: An optional pointer to TableItemRequestBuilderPutQueryParameters, which can be used to specify query parameters for the request.
//
// Returns:
//   - *TableItemResponse: A TableItemResponse containing the updated item data.
//   - error: An error, if the request fails at any point, such as request information creation or JSON deserialization.
func (rB *TableItemRequestBuilder) Put(tableEntry map[string]string, params *TableItemRequestBuilderPutQueryParameters) (*TableItemResponse, error) {
	return rB.Put2(tableEntry, params)
}

// Put2 updates a table item using an HTTP PUT request.
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
func (rB *TableItemRequestBuilder) Put2(tableEntry interface{}, params *TableItemRequestBuilderPutQueryParameters) (*TableItemResponse, error) {
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

	err = rB.SendPut2(rB.Context, config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}
