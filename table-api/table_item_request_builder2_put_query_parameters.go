package tableapi

type TableItemRequestBuilder2PutQueryParameters struct {
	SysparmDisplayValue         *DisplayValue2 `url:"sysparm_display_value,omitempty"`
	SysparmExcludeReferenceLink *bool          `url:"sysparm_exclude_reference_link,omitempty"`
	SysparmFields               []string       `url:"sysparm_fields,omitempty"`
	SysparmInputDisplayValue    interface{}    `url:"sysparm_input_display_value,omitempty"`
	SysparmQueryNoDomain        *bool          `url:"sysparm_query_no_domain,omitempty"`
	SysparmView                 *View2         `url:"sysparm_view,omitempty"`
}
