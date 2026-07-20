package accountapi

// AccountRequestBuilderGetQueryParameters represents the query parameters for a GET request.
type AccountRequestBuilderGetQueryParameters struct {
	// SysparmLimit Limit to be applied on pagination.
	SysparmLimit *int32 `uriparametername:"sysparm_limit"`
	// SysparmOffset Number of records to exclude from the query. Use this parameter to get more records than specified in sysparm_limit.
	SysparmOffset *int32 `uriparametername:"sysparm_offset"`
	// SysparmQuery Encoded query. Queries for the Account API are relative to the Accounts [sys_user] table.
	SysparmQuery *string `uriparametername:"sysparm_query"`
}
