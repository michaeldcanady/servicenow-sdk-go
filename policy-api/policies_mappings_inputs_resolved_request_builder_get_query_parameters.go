package policyapi

type PoliciesMappingsInputsResolvedRequestBuilderGetQueryParameters struct {
	// DeployableName Name of the CDM deployable for which to map the policy.
	DeployableName string `url:"deployable_name,omitempty"`
	// PolicyName Name of the associated policy.
	PolicyName string `url:"policy_name,omitempty"`
	// ReturnFields List of fields to return as part of the response.
	ReturnFields []string `url:"sysparm_fields,omitempty"`
}
