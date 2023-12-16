package tableapi

type TableItemRequestBuilderPutQueryParameters struct {
	DisplayValue         DisplayValue `url:"sysparm_display_value"`
	ExcludeReferenceLink bool         `url:"sysparm_exclude_reference_link"`
	Fields               []string     `url:"sysparm_fields"`
	InputDisplayValue    bool         `url:"sysparm_input_display_value"`
	QueryNoDomain        bool         `url:"sysparm_query_no_domain"`
	View                 View         `url:"sysparm_view"`
}
