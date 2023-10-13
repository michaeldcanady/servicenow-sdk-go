package tableapi

type TableItemRequestBuilderPutQueryParameters struct {
	DisplayValue         DisplayValue `query:"sysparm_display_value"`
	ExcludeReferenceLink bool         `query:"sysparm_exclude_reference_link"`
	Fields               []string     `query:"sysparm_fields"`
	InputDisplayValue    bool         `query:"sysparm_input_display_value"`
	QueryNoDomain        bool         `query:"sysparm_query_no_domain"`
	View                 View         `query:"sysparm_view"`
}
