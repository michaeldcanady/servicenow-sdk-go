package tableapi

// TableItemRequestBuilder2DeleteQueryParameters represents table item  delete request query parameters
type TableItemRequestBuilder2DeleteQueryParameters struct {
	// SysparmQueryNoDomain Flag that indicates whether to restrict the record search to only the domains for which the logged in user is configured.
	SysparmQueryNoDomain *bool `url:"sysparm_query_no_domain,omitempty"`
}
