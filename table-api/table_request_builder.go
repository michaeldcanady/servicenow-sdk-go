package tableapi

import (
	"strconv"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

type TableRequestBuilder struct {
	core.RequestBuilder
}

// NewTableRequestBuilder creates a new instance of the TableRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableRequestBuilder.
func NewTableRequestBuilder(client core.Client, pathParameters map[string]string) *TableRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		"{+baseurl}/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}",
		pathParameters,
	)
	return &TableRequestBuilder{
		*requestBuilder,
	}
}

// ById returns a TableItemRequestBuilder for a specific record in the table.
// It accepts the sysId of the record as a parameter and constructs the URL for the record.
// The returned TableItemRequestBuilder can be used to build and execute requests for the specific record.
func (T *TableRequestBuilder) ById(sysId string) *TableItemRequestBuilder {
	pathParameters := T.RequestBuilder.PathParameters
	pathParameters["sysId"] = sysId
	return NewTableItemRequestBuilder(T.RequestBuilder.Client, pathParameters)
}

// Get sends an HTTP GET request using the specified query parameters and returns a TableCollectionResponse.
//
// Parameters:
//   - params: An instance of TableRequestBuilderGetQueryParameters to include in the GET request.
//
// Returns:
//   - *TableCollectionResponse: The response data as a TableCollectionResponse.
//   - error: An error if there was an issue with the request or response.
func (rB *TableRequestBuilder) Get(params *TableRequestBuilderGetQueryParameters) (*TableCollectionResponse, error) {

	config := &TableGetRequestConfiguration{
		Header:          nil,
		QueryParameters: params,
		Data:            nil,
		ErrorMapping:    nil,
		response:        &TableCollectionResponse{},
	}

	err := rB.SendGet2(config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}

// Deprecated: deprecated since v{version}. Use `Post2` instead.
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

// Post2 sends an HTTP Post request with the provided data and query paramters and returns an `TableItemResponse`.
//
// Parameters:
//   - data: A map[string]string representing data to be included in the request body.
//   - params: An instance of `*TableRequestBuilderPostQueryParameters` for query parameters
func (rB *TableRequestBuilder) Post2(data map[string]string, params *TableRequestBuilderPostQueryParameters) (*TableItemResponse, error) {

	config := &TablePostRequestConfiguration{
		Header:          nil,
		QueryParameters: params,
		Data:            data,
		ErrorMapping:    nil,
		response:        &TableItemResponse{},
	}

	err := rB.SendPost3(config.toConfiguration())
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
func (T *TableRequestBuilder) Count() (int, error) {
	requestInfo, err := T.RequestBuilder.ToHeadRequestInformation()
	if err != nil {
		return -1, err
	}

	errorMapping := core.ErrorMapping{"4XX": "hi"}

	response, err := T.RequestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return -1, err
	}

	count, err := strconv.Atoi(response.Header.Get("X-Total-Count"))
	if err != nil {
		count = 0
	}

	return count, nil
}
