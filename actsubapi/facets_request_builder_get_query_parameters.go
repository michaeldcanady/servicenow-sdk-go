package actsubapi

type FacetsRequestBuilderGetQueryParameters struct {
	// TODO: is a date YYYY-MM-DD hh:mm:ss or YYYY-MM-DD
	EndDate *string `uriparametername:"end_date"`

	// TODO: comma separated
	// Facets List of sys_ids of activity facet types to retrieve for the specified activity context.
	Facets []string `uriparametername:"facets"`

	// GetActivityCount Whether or not to include activity counts for each facet in the response body.
	GetActivityCount *bool `uriparametername:"get_activity_count"`

	// LazyLoad whether or not to improve performance by omitting facet activity data and activity counts from the response body.
	LazyLoad *bool `uriparametername:"lazy_load"`

	// TODO: is a date YYYY-MM-DD hh:mm:ss or YYYY-MM-DD
	// StartDate The request returns only records created during the time period defined by this parameter and end_date. Must be set along with end_date.
	StartDate *string `uriparametername:"start_date"`
}
