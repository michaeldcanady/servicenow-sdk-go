package policyapi

// DefinitionsRequestBuilderGetQueryParameters represents the query parameters for a Policy Definitions collection GET request.
type DefinitionsRequestBuilderGetQueryParameters struct {
	// Limit is the maximum number of records to return.
	Limit int `url:"sysparm_limit,omitempty"`
	// Offset is the number of records to skip before starting to return the results.
	Offset int `url:"sysparm_offset,omitempty"`
	// Query is the encoded query string used to filter the results.
	Query string `url:"sysparm_query,omitempty"`
	// Fields is a list of fields to return in the response.
	Fields []string `url:"sysparm_fields,omitempty"`
}
