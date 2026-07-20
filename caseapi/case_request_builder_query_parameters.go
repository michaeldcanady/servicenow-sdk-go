package caseapi

// CaseRequestBuilderGetQueryParameters represents the query parameters for GET /case.
type CaseRequestBuilderGetQueryParameters struct {
	// Add sysparm_ parameters if known, otherwise generic ones.
	// Based on "Search Case by Case Attributes", it likely supports standard sysparm_query.
	SysparmQuery *string `uriparametername:"sysparm_query"`
}

// CaseFieldValuesRequestBuilderGetQueryParameters represents query parameters for field_values.
type CaseFieldValuesRequestBuilderGetQueryParameters struct {
	// Placeholder
}
