package caseapi

type CaseItemRequestBuilderGetQueryParameters struct {
	// DisplayValue Determines the type of data returned, either the actual values from the database or the display values of the fields.
	DisplayValue *DisplayValue `uriparametername:"sysparm_display_value"`
}
