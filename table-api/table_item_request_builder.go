package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/abstraction"

type TableItemRequestBuilder struct {
	abstraction.RequestBuilder
}

type TableItemRequestBuilderDeleteQueryParameters struct {
	//Flag that indicates whether to restrict the record search to only the domains for which the logged in user is configured.
	//
	//Valid values:
	//
	//- false: Exclude the record if it is in a domain that the currently logged in user is not configured to access.
	//
	//- true: Include the record even if it is in a domain that the currently logged in user is not configured to access.
	QueryNoDomain bool `uriparametername:"sysparm_query_no_domain"`
}

type TableItemRequestBuilderPutQueryParameters struct {
	DisplayValue         DisplayValue `uriparametername:"sysparm_display_value"`
	ExcludeReferenceLink bool         `uriparametername:"sysparm_exclude_reference_link"`
	Fields               []string     `uriparametername:"sysparm_fields"`
	InputDisplayValue    bool         `uriparametername:"sysparm_input_display_value"`
	QueryNoDomain        bool         `uriparametername:"sysparm_query_no_domain"`
	View                 View         `uriparametername:"sysparm_view"`
}

// NewTableItemRequestBuilder creates a new instance of the TableItemRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableItemRequestBuilder.
func NewTableItemRequestBuilder(client abstraction.Client, pathParameters map[string]string) *TableItemRequestBuilder {
	requestBuilder := abstraction.NewRequestBuilder(client, "{+baseurl}/table{/table}{/sysId}", pathParameters)
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

	_, err = T.RequestBuilder.Client.Send(requestInfo, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
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
