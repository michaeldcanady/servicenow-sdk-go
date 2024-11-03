package tableapi

const (
	DisplayValue2True DisplayValue2 = iota
	DisplayValue2False
	DisplayValue2All
)

type DisplayValue2 int

func (dV DisplayValue2) String() string {
	return map[DisplayValue2]string{
		DisplayValue2True:  "true",
		DisplayValue2False: "false",
		DisplayValue2All:   "all",
	}[dV]
}

type TableItemRequestBuilder2GetQueryParameters struct {
	SysparmDisplayValue             *DisplayValue2 `url:"sysparm_display_value,omitempty"`
	SysparmExcludeReferenceLink     *bool          `url:"sysparm_exclude_reference_link,omitempty"`
	SysparmFields                   []string       `url:"sysparm_fields,omitempty"`
	SysparmQueryNoDomain            *bool          `url:"sysparm_query_no_domain,omitempty"`
	SysparmView                     *View2         `url:"sysparm_view,omitempty"`
	SysparmLimit                    *int           `url:"sysparm_limit,omitempty"`
	SysparmNoCount                  *bool          `url:"sysparm_no_count,omitempty"`
	SysparmOffset                   *int           `url:"sysparm_offset,omitempty"`
	SysparmQuery                    *string        `url:"sysparm_query,omitempty"`
	SysparmQueryCategory            *string        `url:"sysparm_query_category,omitempty"`
	SysparmSuppressPaginationHeader *bool          `url:"sysparm_suppress_pagination_header,omitempty"`
}
