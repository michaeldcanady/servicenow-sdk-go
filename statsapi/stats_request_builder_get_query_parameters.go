package statsapi

// StatsRequestBuilderGetQueryParameters represents the query parameters for a Stats API GET request.
type StatsRequestBuilderGetQueryParameters struct {
	// Count indicates whether to return the record count.
	Count bool `url:"sysparm_count,omitempty"`
	// SumFields is the list of fields to sum.
	SumFields []string `url:"sysparm_sum_fields,comma,omitempty"`
	// AvgFields is the list of fields to average.
	AvgFields []string `url:"sysparm_avg_fields,comma,omitempty"`
	// MinFields is the list of fields to find the minimum value of.
	MinFields []string `url:"sysparm_min_fields,comma,omitempty"`
	// MaxFields is the list of fields to find the maximum value of.
	MaxFields []string `url:"sysparm_max_fields,comma,omitempty"`
	// Query is an encoded query string used to filter the records the stats are computed over.
	Query string `url:"sysparm_query,omitempty"`
	// DisplayValue determines whether to return display values, actual values, or both.
	DisplayValue DisplayValue `url:"sysparm_display_value,omitempty"`
}
