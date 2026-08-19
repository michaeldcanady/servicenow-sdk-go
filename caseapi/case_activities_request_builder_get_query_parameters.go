package caseapi

// CaseActivitiesRequestBuilderGetQueryParameters defines query parameters for retrieving case activities.
type CaseActivitiesRequestBuilderGetQueryParameters struct {
	// ActivityType
	ActivityType []*string `uriparameters:"sysparm_activity_type"`
	// Limit
	Limit *int64 `uriparameters:"sysparm_limit"`
	// Offset
	Offset *int64 `uriparameters:"sysparm_offset"`
}
