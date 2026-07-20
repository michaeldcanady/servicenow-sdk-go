package aggregationapi

// StatsRequestBuilderGetQueryParameters represents the query parameters for a Stats API GET request.
type StatsRequestBuilderGetQueryParameters struct {
	// Count indicates whether to return the record count.
	Count *bool `uriparametername:"sysparm_count"`
	// SumFields is the list of fields to sum.
	SumFields []string `uriparametername:"sysparm_sum_fields"`
	// AvgFields is the list of fields to average.
	AvgFields []string `uriparametername:"sysparm_avg_fields"`
	// MinFields is the list of fields to find the minimum value of.
	MinFields []string `uriparametername:"sysparm_min_fields"`
	// MaxFields is the list of fields to find the maximum value of.
	MaxFields []string `uriparametername:"sysparm_max_fields"`
	// Query is an encoded query string used to filter the records the stats are computed over.
	Query *string `uriparametername:"sysparm_query"`
	// DisplayValue determines whether to return display values, actual values, or both.
	DisplayValue *DisplayValue `uriparametername:"sysparm_display_value"`
}
