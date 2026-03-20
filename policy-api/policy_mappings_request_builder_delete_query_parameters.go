package policyapi

type PolicyMappingsRequestBuilderDeleteQueryParameters struct {
	// AppName Name of the CDM deployable for which to map the policy.
	AppName string `url:"appName,omitempty"`
	// DeployableName Name of the CDM deployable for which to map the policy.
	DeployableName string `url:"deployableName,omitempty"`
	// PolicyName Name of the associated policy.
	PolicyName string `url:"policyName,omitempty"`
}
