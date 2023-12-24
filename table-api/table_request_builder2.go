package tableapi

type TableRequestBuilder2[T TableEntry2] struct {
  core.RequestBuilder
}

// NewTableRequestBuilder creates a new instance of the TableRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableRequestBuilder.
func NewTableRequestBuilder[T TableEntry2](client core.Client, pathParameters map[string]string) *TableRequestBuilder[T] {
	requestBuilder := core.NewRequestBuilder(
		client,
		"{+baseurl}/table{/table}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_query_no_domain,sysparm_view,sysparm_limit,sysparm_no_count,sysparm_offset,sysparm_query,sysparm_query_category,sysparm_suppress_pagination_header}",
		pathParameters,
	)
	return &TableRequestBuilder[T]{
		*requestBuilder,
	}
}

// Get sends an HTTP GET request using the specified query parameters and returns a TableCollectionResponse.
//
// Parameters:
//   - params: An instance of TableRequestBuilderGetQueryParameters to include in the GET request.
//
// Returns:
//   - *TableCollectionResponse: The response data as a TableCollectionResponse.
//   - error: An error if there was an issue with the request or response.
func (rB *TableRequestBuilder2[T]) Get(params *TableRequestBuilderGetQueryParameters) (*TableCollectionResponse2[T], error) {
	config := &TableGetRequestConfiguration2{
		Header:          nil,
		QueryParameters: params,
		Data:            nil,
		ErrorMapping:    nil,
		response:        &TableCollectionResponse[T]{},
	}

	err := rB.SendGet2(config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}