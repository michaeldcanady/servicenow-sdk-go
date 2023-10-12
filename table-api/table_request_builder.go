package tableapi

import (
	"strconv"

	"github.com/michaeldcanady/servicenow-sdk-go/abstraction"
)

type TableRequestBuilder struct {
	abstraction.RequestBuilder
}

type TableCollectionResponse struct {
	Result []*TableEntry
}

type TableResponse struct {
	Result TableEntry
}

type TableEntry map[string]interface{}

type DisplayValue string
type View string

const (
	TRUE  DisplayValue = "true"
	FALSE DisplayValue = "false"
	ALL   DisplayValue = "all"

	DESKTOP View = "desktop"
	MOBILE  View = "mobile"
	BOTH    View = "both"
)

type TableRequestBuilderGetQueryParameters struct {
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
	View                     View   `uriparametername:"sysparm_view"`
	Limit                    int32  `uriparametername:"sysparm_limit"`
	NoCount                  bool   `uriparametername:"sysparm_no_count"`
	Offset                   int    `uriparametername:"sysparm_offset"`
	Query                    string `uriparametername:"sysparm_query"`
	QueryCategory            string `uriparametername:"sysparm_query_category"`
	SuppressPaginationHeader bool   `uriparameter:"sysparm_suppress_pagination_header"`
}

type TableRequestBuilderPostQueryParamters struct {
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
	Fields            []string `uriparametername:"sysparm_fields"`
	InputDisplayValue bool     `uriparametername:"sysparm_input_display_value"`
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

// NewTableRequestBuilder creates a new instance of the TableRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableRequestBuilder.
func NewTableRequestBuilder(client abstraction.Client, pathParameters map[string]string) *TableRequestBuilder {
	requestBuilder := abstraction.NewRequestBuilder(client, "{+baseurl}/table{/table}{?sysparm_limit}", pathParameters)
	return &TableRequestBuilder{
		*requestBuilder,
	}
}

// ById returns a TableItemRequestBuilder for a specific record in the table.
// It accepts the sysId of the record as a parameter and constructs the URL for the record.
// The returned TableItemRequestBuilder can be used to build and execute requests for the specific record.
func (T *TableRequestBuilder) ById(sysId string) *TableItemRequestBuilder {
	pathParameters := T.PathParameters
	pathParameters["sysId"] = sysId
	return NewTableItemRequestBuilder(T.Client, pathParameters)
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
	requestInfo, err := T.ToGetRequestInformation(params)
	if err != nil {
		return nil, err
	}

	errorMapping := abstraction.ErrorMapping{"4XX": "hi"}

	response, err := T.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return nil, err
	}

	value, err := abstraction.FromJson[TableCollectionResponse](response)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// POST sends an HTTP POST request with the provided data and query parameters and returns a TableResponse.
//
// Parameters:
//   - data: A map[string]interface{} representing data to be included in the request body.
//   - params: An instance of TableRequestBuilderPostQueryParamters for query parameters.
//
// Returns:
//   - *TableResponse: The response data as a TableResponse.
//   - error: An error if there was an issue with the request or response.
func (T *TableRequestBuilder) POST(data map[string]interface{}, params *TableRequestBuilderPostQueryParamters) (*TableResponse, error) {
	requestInfo, err := T.ToPostRequestInformation(data, params)
	if err != nil {
		return nil, err
	}

	errorMapping := abstraction.ErrorMapping{"4XX": "hi"}

	response, err := T.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return nil, err
	}

	value, err := abstraction.FromJson[TableResponse](response)
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
	requestInfo, err := T.ToHeadRequestInformation()
	if err != nil {
		return -1, err
	}

	errorMapping := abstraction.ErrorMapping{"4XX": "hi"}

	response, err := T.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return -1, err
	}

	count, err := strconv.Atoi(response.Header.Get("X-Total-Count"))
	if err != nil {
		count = 0
	}

	return count, nil
}
