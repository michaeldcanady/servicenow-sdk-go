package tableapi

// TableItemRequestBuilderGetQueryParameters represents GET query parameters for a Table Item Request.
type TableItemRequestBuilderGetQueryParameters struct {
	//DisplayValue determines the type of data returned, either the actual values from the database or the display values of the fields.
	//Display values are manipulated based on the actual value in the database and user or system settings and preferences.
	//If returning display values, the value that is returned is dependent on the field type.
	//- Choice fields: The database value may be a number, but the display value will be more descriptive.
	//
	//- Date fields: The database value is in UTC format, while the display value is based on the user's time zone.
	//
	//- Encrypted text: The database value is encrypted, while the displayed value is unencrypted based on the user's encryption context.
	//
	//- Reference fields: The database value is sys_id, but the display value is a display field of the referenced record.
	DisplayValue DisplayValue `url:"sysparm_display_value"`
	//ExcludeReferenceLink flag that indicates whether to exclude Table API links for reference fields.
	//
	//Valid values:
	//
	//- true: Exclude Table API links for reference fields.
	//
	//- false: Include Table API links for reference fields.
	ExcludeReferenceLink bool `url:"sysparm_exclude_reference_link"`
	// Fields list of fields to return in the response.
	Fields []string `url:"sysparm_fields"`
	//Flag that indicates whether to restrict the record search to only the domains for which the logged in user is configured.
	//
	//Valid values:
	//
	//- false: Exclude the record if it is in a domain that the currently logged in user is not configured to access.
	//
	//- true: Include the record even if it is in a domain that the currently logged in user is not configured to access.
	QueryNoDomain bool `url:"sysparm_query_no_domain"`
	// View	UI view for which to render the data. Determines the fields returned in the response.
	//
	//Valid values:
	//
	//- desktop
	//- mobile
	//- both
	//If you also specify the sysparm_fields parameter, it takes precedent.
	View View `url:"sysparm_view"`
}
