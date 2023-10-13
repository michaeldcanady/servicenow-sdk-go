package tableapi

type TableItemRequestBuilderPutQueryParameters struct {
	DisplayValue         DisplayValue `uriparametername:"sysparm_display_value"`
	ExcludeReferenceLink bool         `uriparametername:"sysparm_exclude_reference_link"`
	Fields               []string     `uriparametername:"sysparm_fields"`
	InputDisplayValue    bool         `uriparametername:"sysparm_input_display_value"`
	QueryNoDomain        bool         `uriparametername:"sysparm_query_no_domain"`
	View                 View         `uriparametername:"sysparm_view"`
}