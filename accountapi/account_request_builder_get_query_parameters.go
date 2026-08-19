package accountapi

// AccountRequestBuilderGetQueryParameters represents the query parameters for a GET request.
type AccountRequestBuilderGetQueryParameters struct {
	// Limit Limit to be applied on pagination.
	Limit *int32 `uriparametername:"sysparm_limit"`
	// Offset Number of records to exclude from the query. Use this parameter to get more records than specified in sysparm_limit.
	Offset *int32 `uriparametername:"sysparm_offset"`
	// Query Encoded query. Queries for the Account API are relative to the Accounts [sys_user] table.
	Query *string `uriparametername:"sysparm_query"`
}
