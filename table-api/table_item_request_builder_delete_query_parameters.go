package tableapi

// Deprecated: deprecated since v1.9.0. Please use [TableItemRequestBuilder2DeleteQueryParameters]
type TableItemRequestBuilderDeleteQueryParameters struct {
	// QueryNoDomain flag that indicates whether to restrict the record search to only the domains for which the logged in user is configured.
	//
	//Valid values:
	//
	//- false: Exclude the record if it is in a domain that the currently logged in user is not configured to access.
	//
	//- true: Include the record even if it is in a domain that the currently logged in user is not configured to access.
	QueryNoDomain bool `url:"sysparm_query_no_domain"`
}
