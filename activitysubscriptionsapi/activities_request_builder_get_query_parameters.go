package actsubapi

type ActivitiesRequestBuilderGetQueryParameters struct {
	// TODO: is a date YYYY-MM-DD hh:mm:ss or YYYY-MM-DD
	Before *string `uriparametername:"before"`

	// Context Sys_id of an activity context.
	Context *string `uriparametername:"context"`

	// ContextInstance Sys_id of an instance of the specified activity context, representing the initiator of the activities you want to retrieve.
	ContextInstance *string `uriparametername:"context_instance"`

	// TODO: is a date YYYY-MM-DD hh:mm:ss or YYYY-MM-DD
	EndDate *string `uriparametername:"end_date"`

	// TODO: comma separated
	// Facets list of sys_ids of activity facet types to retrieve for the specified activity context.
	Facets []string `uriparametername:"facets"`

	// Last Index value of the first result row omitted from the response body.
	Last *int64 `uriparametername:"last"`

	// RecordID Sys_id of a record to use when rendering dynamic facets.
	RecordID *string `uriparametername:"record_id"`

	// TODO: is a date YYYY-MM-DD hh:mm:ss or YYYY-MM-DD
	// StartDate The request returns only records created during the time period defined by this parameter and end_date. Must be set along with end_date.
	StartDate *string `uriparametername:"start_date"`

	// From Index value of the first result row to include in the response body.
	From *int64 `uriparametername:"stFrom"`
}
