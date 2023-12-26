package tableapi

// TableItemRequestBuilderDeleteQueryParameters represents the Query Parameters for a DELETE Table Item request.
type TableItemRequestBuilderDeleteQueryParameters struct {
	//Flag that indicates whether to restrict the record search to only the domains for which the logged in user is configured.
	//
	//Valid values:
	//
	//- false: Exclude the record if it is in a domain that the currently logged in user is not configured to access.
	//
	//- true: Include the record even if it is in a domain that the currently logged in user is not configured to access.
	QueryNoDomain bool `query:"sysparm_query_no_domain"`
}
