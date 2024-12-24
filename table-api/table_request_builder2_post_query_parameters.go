package tableapi

// TableRequestBuilder2PostQueryParameters represents table post request query parameters
type TableRequestBuilder2PostQueryParameters struct {
	// SysparmDisplayValue Determines the type of data returned, either the actual values from the database or the display values of the fields.
	// Display values are manipulated based on the actual value in the database and user or system settings and preferences.
	SysparmDisplayValue *DisplayValue2 `url:"sysparm_display_value,omitempty"`
	// SysparmExcludeReferenceLink Flag that indicates whether to exclude Table API links for reference fields.
	SysparmExcludeReferenceLink *bool `url:"sysparm_exclude_reference_link,omitempty"`
	// SysparmFields Comma-separated list of fields to return in the response.
	SysparmFields []string `url:"sysparm_fields,omitempty"`
	// SysparmInputDisplayValue Flag that indicates whether to set field values using the display value or the actual value.
	// Depending on the different types of fields, the endpoint may manipulate the passed in display values to store the proper values in the database.
	// For example, if you send the display name for a reference field, the endpoint stores the sys_id for that value in the database.
	// For date and time fields, when this parameter is true, the date and time value is adjusted for the current user's timezone.
	// When false, the date and time value is inserted using the GMT timezone.
	SysparmInputDisplayValue *bool `uriparametername:"sysparm_input_display_value"`
	// SysparmView UI view for which to render the data. Determines the fields returned in the response.
	SysparmView *View2 `url:"sysparm_view,omitempty"`
}
