package tableapi

type requestBuilder interface {
    SendGet2(core.RequestConfigurations) error
}

type TableItemRequestBuilder2 struct {
    requestBuilder
}

func NewTableItemRequestBuilder2(client core.Client, pathParameters map[string]string) *TableItemRequestBuilder2 {
   return &TableItemRequestBuilder2{
        core.NewRequestBuilder(
		client,
		"{+baseurl}/table{/table}{/sysId}{?sysparm_display_value,sysparm_exclude_reference_link,sysparm_fields,sysparm_input_display_value,sysparm_query_no_domain,sysparm_view,sysparm_query_no_domain}",
		pathParameters,
	)
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
