package servicenowsdkgo

type TableRequestBuilder struct {
	RequestBuilder
}

type TableCollectionResponse struct {
	Result []*TableEntry
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
	Limit                    int    `uriparametername:"sysparm_limit"`
	NoCount                  bool   `uriparametername:"sysparm_no_count"`
	Offset                   int    `uriparametername:"sysparm_offset"`
	Query                    string `uriparametername:"sysparm_query"`
	QueryCategory            string `uriparametername:"sysparm_query_category"`
	SuppressPaginationHeader bool   `uriparameter:"sysparm_suppress_pagination_header"`
}

// NewTableRequestBuilder creates a new instance of the TableRequestBuilder associated with the given URL and Client.
// It accepts the URL and Client as parameters and returns a pointer to the created TableRequestBuilder.
func NewTableRequestBuilder(url string, client *Client) *TableRequestBuilder {
	requestBuilder := NewRequestBuilder(url, client)
	return &TableRequestBuilder{
		*requestBuilder,
	}
}

// ById returns a TableItemRequestBuilder for a specific record in the table.
// It accepts the sysId of the record as a parameter and constructs the URL for the record.
// The returned TableItemRequestBuilder can be used to build and execute requests for the specific record.
func (T *TableRequestBuilder) ById(sysId string) *TableItemRequestBuilder {
	return NewTableItemRequestBuilder(T.AppendSegment(sysId), T.Client)
}

// Get performs an HTTP GET request to the table URL using the Client's session.
// It retrieves a collection of records from the table and decodes the response into a TableCollectionResponse.
// It returns the TableCollectionResponse and any errors encountered during the request or decoding.
func (T *TableRequestBuilder) Get(params *TableRequestBuilderGetQueryParameters) (*TableCollectionResponse, error) {
	resp := &TableCollectionResponse{}
	err := T.Client.Get(T.Url, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
