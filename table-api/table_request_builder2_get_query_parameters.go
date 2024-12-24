package tableapi

// TableRequestBuilder2GetQueryParameters represents table get request query parameters
type TableRequestBuilder2GetQueryParameters struct {
	// SysparmDisplayValue Determines the type of data returned, either the actual values from the database or the display values of the fields.
	// Display values are manipulated based on the actual value in the database and user or system settings and preferences.
	SysparmDisplayValue *DisplayValue2 `url:"sysparm_display_value,omitempty"`
	// SysparmExcludeReferenceLink Flag that indicates whether to exclude Table API links for reference fields.
	SysparmExcludeReferenceLink *bool `url:"sysparm_exclude_reference_link,omitempty"`
	// SysparmFields Comma-separated list of fields to return in the response.
	SysparmFields []string `url:"sysparm_fields,omitempty"`
	// SysparmQueryNoDomain Flag that indicates whether to restrict the record search to only the domains for which the logged in user is configured.
	SysparmQueryNoDomain *bool `url:"sysparm_query_no_domain,omitempty"`
	// SysparmView UI view for which to render the data. Determines the fields returned in the response.
	SysparmView *View2 `url:"sysparm_view,omitempty"`
	// SysparmLimit Maximum number of records to return. For requests that exceed this number of records, use the sysparm_offset parameter to paginate record retrieval.
	SysparmLimit *int64 `url:"sysparm_limit,omitempty"`
	// SysparmNoCount Flag that indicates whether to execute a select count(*) query on the table to return the number of rows in the associated table.
	SysparmNoCount *bool `url:"sysparm_no_count,omitempty"`
	// SysparmOffset Starting record index for which to begin retrieving records. Use this value to paginate record retrieval.
	// This functionality enables the retrieval of all records, regardless of the number of records, in small manageable chunks.
	SysparmOffset *int64 `url:"sysparm_offset,omitempty"`
	// SysparmQuery Encoded query used to filter the result set. You can use a UI filter to obtain a properly encoded query.
	SysparmQuery *string `url:"sysparm_query,omitempty"`
	// SysparmQueryCategory Name of the category to use for queries.
	SysparmQueryCategory *string `url:"sysparm_query_category,omitempty"`
	// SysparmSuppressPaginationHeader Flag that indicates whether to remove the Link header from the response.
	// The Link header provides various URLs to relative pages in the record set which you can use to paginate the returned record set.
	SysparmSuppressPaginationHeader *bool `url:"sysparm_suppress_pagination_header,omitempty"`
}
