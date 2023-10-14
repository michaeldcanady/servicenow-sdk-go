package tableapi

import (
	"strconv"

	"github.com/michaeldcanady/servicenow-sdk-go/abstraction"
)

type TableRequestBuilder struct {
	abstraction.RequestBuilder
}

// NewTableRequestBuilder creates a new instance of the TableRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableRequestBuilder.
func NewTableRequestBuilder(client abstraction.Client, pathParameters map[string]string) *TableRequestBuilder {
	requestBuilder := abstraction.NewRequestBuilder(client, "{+baseurl}/table{/table}{?sysparm_limit,sysparm_no_count}", pathParameters)
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
func (T *TableRequestBuilder) Get(params *TableRequestBuilderGetQueryParameters) (*TableCollectionResponse, error) {

	requestInfo, err := T.RequestBuilder.ToGetRequestInformation(params)
	if err != nil {
		return nil, err
	}

	errorMapping := abstraction.ErrorMapping{"4XX": "hi"}

	response, err := T.RequestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return nil, err
	}

	value, err := abstraction.FromJson[TableCollectionResponse](response)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Post sends an HTTP Post request with the provided data and query parameters and returns a TableResponse.
//
// Parameters:
//   - data: A map[string]interface{} representing data to be included in the request body.
//   - params: An instance of TableRequestBuilderPostQueryParamters for query parameters.
//
// Returns:
//   - *TableResponse: The response data as a TableResponse.
//   - error: An error if there was an issue with the request or response.
func (T *TableRequestBuilder) Post(data map[string]interface{}, params *TableRequestBuilderPostQueryParamters) (*TableItemResponse, error) {
	requestInfo, err := T.RequestBuilder.ToPostRequestInformation(data, params)
	if err != nil {
		return nil, err
	}

	errorMapping := abstraction.ErrorMapping{"4XX": "hi"}

	response, err := T.RequestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return nil, err
	}

	value, err := abstraction.FromJson[TableItemResponse](response)
	if err != nil {
		return nil, err
	}

	return value, nil
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

	errorMapping := abstraction.ErrorMapping{"4XX": "hi"}

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
