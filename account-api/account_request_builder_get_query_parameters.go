package accountapi

// AccountRequestBuilderGetQueryParameters represents the query parameters for Get
type AccountRequestBuilderGetQueryParameters struct {

	// SysparmLimit Maximum number of records to return.
	SysparmLimit *int `url:"sysparm_limit,omitempty"`

	// SysparmOffset Starting record index for which to begin retrieving records.
	SysparmOffset *int `url:"sysparm_offset,omitempty"`

	// SysparmQuery Encoded query used to filter the result set.
	SysparmQuery *string `url:"sysparm_query,omitempty"`
}
