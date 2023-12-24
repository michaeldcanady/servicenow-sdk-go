package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

type TableItemRequestBuilder struct {
	core.RequestBuilder
}

// NewTableItemRequestBuilder creates a new instance of the TableItemRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableItemRequestBuilder.
func NewTableItemRequestBuilder(client core.Client, pathParameters map[string]string) *TableItemRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		"{+baseurl}/table{/table}{/sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_input_display_value,sysparm_query_no_domain,sysparm_view,sysparm_query_no_domain}",
		pathParameters,
	)
	return &TableItemRequestBuilder{
		*requestBuilder,
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
func (rB *TableItemRequestBuilder) Get(params *TableItemRequestBuilderGetQueryParameters) (*TableItemResponse, error) {
	config := &TableItemGetRequestConfiguration{
		Header:          nil,
		QueryParameters: params,
		Data:            nil,
		response:        &TableItemResponse{},
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
func (rB *TableItemRequestBuilder) Delete(params *TableItemRequestBuilderDeleteQueryParameters) error {
	config := &TableItemDeleteRequestConfiguration{
		Header:          nil,
		QueryParameters: params,
		Data:            nil,
		response:        nil,
	}

	return rB.SendDelete2(config.toConfiguration())
}

// Deprecated: deprecated since v{version} please use Put2 instead.
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

	var tableEntry2 = TableEntry{}

	for key, value := range tableEntry {
		tableEntry2.Set(key, value)
	}

	return rB.Put2(tableEntry2, params)
}

// Put2 updates a table item using an HTTP PUT request.
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
func (rB *TableItemRequestBuilder) Put2(tableEntry TableEntry, params *TableItemRequestBuilderPutQueryParameters) (*TableItemResponse, error) {

	if tableEntry.Len() == 0 {
		return nil, ErrEmptyTableEntry
	}

	config := &TableItemPutRequestConfiguration{
		Header:          nil,
		QueryParameters: params,
		Data:            tableEntry,
		response:        &TableItemResponse{},
	}

	err := rB.SendPut2(config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}
