package tableapi

type TableRequestBuilder2[T TableEntry2] struct {
  core.RequestBuilder
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