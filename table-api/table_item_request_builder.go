package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/abstraction"

type TableItemRequestBuilder struct {
	abstraction.RequestBuilder
}

type TableItemResponse struct {
	Result *TableEntry
}

type TableItemRequestBuilderGetQueryParameters struct {
	//Determines the type of data returned, either the actual values from the database or the display values of the fields.
	//Display values are manipulated based on the actual value in the database and user or system settings and preferences.
	//If returning display values, the value that is returned is dependent on the field type.
	//- Choice fields: The database value may be a number, but the display value will be more descriptive.
	//
	//- Date fields: The database value is in UTC format, while the display value is based on the user's time zone.
	//
	//- Encrypted text: The database value is encrypted, while the displayed value is unencrypted based on the user's encryption context.
	//
	//- Reference fields: The database value is sys_id, but the display value is a display field of the referenced record.
	DisplayValue DisplayValue `uriparametername:"sysparm_display_value"`
	//Flag that indicates whether to exclude Table API links for reference fields.
	//
	//Valid values:
	//
	//- true: Exclude Table API links for reference fields.
	//
	//- false: Include Table API links for reference fields.
	ExcludeReferenceLink bool `uriparametername:"sysparm_exclude_reference_link"`
	//list of fields to return in the response.
	Fields []string `uriparametername:"sysparm_fields"`
	//Flag that indicates whether to restrict the record search to only the domains for which the logged in user is configured.
	//
	//Valid values:
	//
	//- false: Exclude the record if it is in a domain that the currently logged in user is not configured to access.
	//
	//- true: Include the record even if it is in a domain that the currently logged in user is not configured to access.
	QueryNoDomain bool `uriparametername:"sysparm_query_no_domain"`
	//	UI view for which to render the data. Determines the fields returned in the response.
	//
	//Valid values:
	//
	//- desktop
	//- mobile
	//- both
	//If you also specify the sysparm_fields parameter, it takes precedent.
	View View `uriparametername:"sysparm_view"`
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
	requestInfo, err := T.ToGetRequestInformation(params)
	if err != nil {
		return nil, err
	}

	_, err = T.Client.Send(requestInfo, nil)
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
	requestInfo, err := T.ToDeleteRequestInformation(params)
	if err != nil {
		return err
	}

	_, err = T.Client.Send(requestInfo, nil)
	if err != nil {
		return err
	}

	return nil
}
