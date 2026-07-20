package policyapi

type PoliciesMappingsRequestBuilderPostQueryParameters struct {
	// AppName Name of the CDM application for which to map the policy.
	AppName *string `uriparametername:"appName"`
	// DeployableName Name of the CDM deployable for which to map the policy.
	DeployableName *string `uriparametername:"deployableName"`
	// PolicyName Name of the associated policy.
	PolicyName *string `uriparametername:"policyName"`
	// ReturnFields List of fields to return as part of the response.
	ReturnFields []string `uriparametername:"returnFields"`
}
