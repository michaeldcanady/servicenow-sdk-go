package tableapi

import (
	"context"
	"strconv"

	"github.com/RecoLabs/servicenow-sdk-go/core"
)

// Deprecated: deprecated since v{unreleased}. Use `TableRequestBuilder2` instead.
type TableRequestBuilder struct {
	context.Context
	core.RequestBuilder
}

// Deprecated: deprecated since v{unreleased}. Use `NewTableRequestBuilder2` instead.
// NewTableRequestBuilder creates a new instance of the TableRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableRequestBuilder.
func NewTableRequestBuilder(ctx context.Context, client core.Client,
	pathParameters map[string]string) *TableRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		"{+baseurl}/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}",
		pathParameters,
	)
	return &TableRequestBuilder{
		Context:        ctx,
		RequestBuilder: *requestBuilder,
	}
}

// ById returns a TableItemRequestBuilder for a specific record in the table.
// It accepts the sysId of the record as a parameter and constructs the URL for the record.
// The returned TableItemRequestBuilder can be used to build and execute requests for the specific record.
func (rB *TableRequestBuilder) ById(ctx context.Context, sysId string) *TableItemRequestBuilder { //nolint:stylecheck
	pathParameters := rB.RequestBuilder.PathParameters
	pathParameters["sysId"] = sysId
	return NewTableItemRequestBuilder(ctx, rB.RequestBuilder.Client, pathParameters)
}

// Get sends an HTTP GET request using the specified query parameters and returns a TableCollectionResponse.
//
// Parameters:
//   - params: An instance of TableRequestBuilderGetQueryParameters to include in the GET request.
//
// Returns:
//   - *TableCollectionResponse: The response data as a TableCollectionResponse.
//   - error: An error if there was an issue with the request or response.
func (rB *TableRequestBuilder) Get(ctx context.Context, params *TableRequestBuilderGetQueryParameters) (
	*TableCollectionResponse, error) {
	config := &tableGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     nil,
		mapping:  nil,
		response: &TableCollectionResponse2[TableEntry]{},
	}

	err := rB.SendGet2(rB.Context, config.toConfiguration())
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
func (rB *TableRequestBuilder) Post2(ctx context.Context, data map[string]string,
	params *TableRequestBuilderPostQueryParameters) (*TableItemResponse, error) {
	config := &tablePostRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    params,
		data:     data,
		mapping:  nil,
		response: &TableItemResponse2[TableEntry]{},
	}

	err := rB.SendPost3(rB.Context, config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}

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

	err = rB.SendPost3(rB.Context, config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}

// Count sends an HTTP HEAD request and retrieves the value of "X-Total-Count" from the response header, which represents the count of items.
//
// Returns:
//   - int: The count of items.
//   - error: An error if there was an issue with the request or response.
func (rB *TableRequestBuilder) Count() (int, error) {
	requestInfo, err := rB.RequestBuilder.ToHeadRequestInformation()
	if err != nil {
		return -1, err
	}

	errorMapping := core.ErrorMapping{"4XX": "hi"}

	response, err := rB.RequestBuilder.Client.Send(rB.Context, requestInfo, errorMapping)
	if err != nil {
		return -1, err
	}

	count, err := strconv.Atoi(response.Header.Get("X-Total-Count"))
	if err != nil {
		count = 0
	}

	return count, nil
}
