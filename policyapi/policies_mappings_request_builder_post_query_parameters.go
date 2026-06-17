package policyapi

type PoliciesMappingsRequestBuilderPostQueryParameters struct {
	// AppName Name of the CDM application for which to map the policy.
	AppName string `url:"appName,omitempty"`
	// DeployableName Name of the CDM deployable for which to map the policy.
	DeployableName string `url:"deployableName,omitempty"`
	// PolicyName Name of the associated policy.
	PolicyName string `url:"policyName,omitempty"`
	// ReturnFields List of fields to return as part of the response.
	ReturnFields []string `url:"returnFields,omitempty"`
}
