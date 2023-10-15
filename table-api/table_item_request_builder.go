package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

type TableItemRequestBuilder struct {
	core.RequestBuilder
}

// NewTableItemRequestBuilder creates a new instance of the TableItemRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableItemRequestBuilder.
func NewTableItemRequestBuilder(client core.Client, pathParameters map[string]string) *TableItemRequestBuilder {
	requestBuilder := core.NewRequestBuilder(client, "{+baseurl}/table{/table}{/sysId}", pathParameters)
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
func (T *TableItemRequestBuilder) Get(params *TableItemRequestBuilderGetQueryParameters) (*TableItemResponse, error) {
	requestInfo, err := T.RequestBuilder.ToGetRequestInformation(params)
	if err != nil {
		return nil, err
	}

	response, err := T.RequestBuilder.Client.Send(requestInfo, nil)
	if err != nil {
		return nil, err
	}

	value, err := core.FromJson[TableItemResponse](response)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Delete sends an HTTP DELETE request using the specified query parameters and returns an error if the request or response encounters any issues.
//
// Parameters:
//   - params: An instance of TableItemRequestBuilderDeleteQueryParameters to include in the DELETE request.
//
// Returns:
//   - error: An error if there was an issue with the request or response, or nil if the request was successful.
func (T *TableItemRequestBuilder) Delete(params *TableItemRequestBuilderDeleteQueryParameters) error {
	requestInfo, err := T.RequestBuilder.ToDeleteRequestInformation(params)
	if err != nil {
		return err
	}

	_, err = T.RequestBuilder.Client.Send(requestInfo, nil)
	if err != nil {
		return err
	}

	return nil
}
