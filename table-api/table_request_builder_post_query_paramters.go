package tableapi

// Deprecated: deprecated as of v1.4.0. use `TableRequestBuilderPostQueryParameters` instead
//
// TableRequestBuilderPostQueryParamters represents POST query parameters for a Table Item Request.
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
	DisplayValue DisplayValue `url:"sysparm_display_value"`
	//Flag that indicates whether to exclude Table API links for reference fields.
	//
	//Valid values:
	//
	//- true: Exclude Table API links for reference fields.
	//
	//- false: Include Table API links for reference fields.
	ExcludeReferenceLink bool `url:"sysparm_exclude_reference_link"`
	//list of fields to return in the response.
	Fields            []string `url:"sysparm_fields"`
	InputDisplayValue bool     `url:"sysparm_input_display_value"`
	//	UI view for which to render the data. Determines the fields returned in the response.
	//
	//Valid values:
	//
	//- desktop
	//- mobile
	//- both
	//If you also specify the sysparm_fields parameter, it takes precedent.
	View View `url:"sysparm_view"`
}
