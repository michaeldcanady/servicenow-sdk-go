package tableapi

// TableItemRequestBuilder2PutQueryParameters represents table item put request query parameters
type TableItemRequestBuilder2PutQueryParameters struct {
	// SysparmDisplayValue Determines the type of data returned, either the actual values from the database or the display values of the fields.
	// Display values are manipulated based on the actual value in the database and user or system settings and preferences.
	SysparmDisplayValue *DisplayValue2 `url:"sysparm_display_value,omitempty"`
	// SysparmExcludeReferenceLink Flag that indicates whether to exclude Table API links for reference fields.
	SysparmExcludeReferenceLink *bool `url:"sysparm_exclude_reference_link,omitempty"`
	// SysparmFields Comma-separated list of fields to return in the response.
	SysparmFields []string `url:"sysparm_fields,omitempty"`
	// SysparmInputDisplayValue Flag that indicates whether to set field values using the display value or the actual value.
	SysparmInputDisplayValue *bool `url:"sysparm_input_display_value,omitempty"`
	// SysparmQueryNoDomain Flag that indicates whether to restrict the record search to only the domains for which the logged in user is configured.
	SysparmQueryNoDomain *bool `url:"sysparm_query_no_domain,omitempty"`
	// SysparmView UI view for which to render the data. Determines the fields returned in the response.
	SysparmView *View2 `url:"sysparm_view,omitempty"`
}
