package tableapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

const tableURLTemplate = "{+baseurl}/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}"

type TableRequestBuilder struct {
	core.RequestBuilder
}

// NewTableRequestBuilder creates a new instance of the TableRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableRequestBuilder.
func NewTableRequestBuilder(client core.Client, pathParameters map[string]string) *TableRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		tableURLTemplate,
		pathParameters,
	)
	return &TableRequestBuilder{
		*requestBuilder,
	}
}

// Deprecated: deprecated since v{unreleased}. Use ById2 instead.
// ById returns a TableItemRequestBuilder for a specific record in the table.
// It accepts the sysId of the record as a parameter and constructs the URL for the record.
// The returned TableItemRequestBuilder can be used to build and execute requests for the specific record.
func (rB *TableRequestBuilder) ById(sysId string) *TableItemRequestBuilder { //nolint:stylecheck
	pathParameters := rB.PathParameters
	pathParameters["sysId"] = sysId
	return NewTableItemRequestBuilder(rB.Client, pathParameters)
}

// Deprecated: deprecated since v{unreleased}. Use Get2 instead.
// Get sends an HTTP GET request using the specified query parameters and returns a TableCollectionResponse.
//
// Parameters:
//   - params: An instance of TableRequestBuilderGetQueryParameters to include in the GET request.
//
// Returns:
//   - *TableCollectionResponse: The response data as a TableCollectionResponse.
//   - error: An error if there was an issue with the request or response.
func (rB *TableRequestBuilder) Get(params *TableRequestBuilderGetQueryParameters) (*TableCollectionResponse, error) {
	config := &tableGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		mapping:  nil,
		response: &TableCollectionResponse2[TableEntry]{},
	}

	err := rB.SendGet2(config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}

// Deprecated: deprecated since v1.4.0. Use `Post2` instead.
// Post sends an HTTP Post request with the provided data and query parameters and returns an `TableItemResponse`.
//
// Parameters:
//   - data: A map[string]string representing data to be included in the request body.
//   - params: An instance of `*TableRequestBuilderPostQueryParameters` for query parameters.
//
// Returns:
//   - *TableResponse: The response data as a TableResponse.
//   - error: An error if there was an issue with the request or response.
func (rB *TableRequestBuilder) Post(data map[string]string, params *TableRequestBuilderPostQueryParamters) (*TableItemResponse, error) {
	var response TableItemResponse

	err := rB.SendPost(data, params, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Deprecated: deprecated since v1.4.0. Use `Post3` instead.
//
// Post2 sends an HTTP Post request with the provided data and query parameters and returns an `TableItemResponse`.
//
// Parameters:
//   - data: A map[string]string representing data to be included in the request body.
//   - params: An instance of `*TableRequestBuilderPostQueryParameters` for query parameters
func (rB *TableRequestBuilder) Post2(data map[string]string, params *TableRequestBuilderPostQueryParameters) (*TableItemResponse, error) {
	config := &tablePostRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     data,
		mapping:  nil,
		response: &TableItemResponse2[TableEntry]{},
	}

	err := rB.SendPost3(config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}

// Deprecated: deprecated since v{unreleased}. Use Post4 instead.
// Post3 sends an HTTP Post request with the provided data and query parameters and returns an `TableItemResponse`.
//
// Parameters:
//   - data: A map[string]string or TableEntry representing data to be included in the request body.
//   - params: An instance of `*TableRequestBuilderPostQueryParameters` for query parameters
func (rB *TableRequestBuilder) Post3(data interface{}, params *TableRequestBuilderPostQueryParameters) (*TableItemResponse, error) {
	data, err := convertFromTableEntry(data)
	if err != nil {
		return nil, err
	}

	config := &tablePostRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     data.(map[string]string),
		mapping:  nil,
		response: &TableItemResponse2[TableEntry]{},
	}

	err = rB.SendPost3(config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}
