package caseapi

// CaseFieldValuesRequestBuilderGetQueryParameters represents query parameters for field_values.
type CaseFieldValuesRequestBuilderGetQueryParameters struct {
	// DependentValue Value to select in the Case `[sn_customerservice_case]` table choice field that the requested field is dependent on. Use only when requesting a choice field that is dependent on another field.
	DependentValue *string `uriparametername:"sysparm_dependent_value"`
	// Limit Maximum number of records to return.
	Limit *int64 `uriparametername:"sysparm_limit"`
	// Offest Starting record index for which to begin retrieving records.
	Offset *int64 `uriparametername:"sysparm_offset"`
	// TODO: Comma-separated
	// ReferenceFieldColumns list of column names, from the table of the specified reference field, to return in the response.
	ReferenceFieldColumns []string `uriparametername:"sysparm_reference_field_columns"`
	// Query Encoded query used to filter the result set.
	Query *string `uriparametername:"sysparm_query"`
	// RefQualInput Encoded set of field values representing a current object to pass to reference qualifiers that use JavaScript functions.
	RefQualInput *string `uriparametername:"sysparm_ref_qual_input"`
}
