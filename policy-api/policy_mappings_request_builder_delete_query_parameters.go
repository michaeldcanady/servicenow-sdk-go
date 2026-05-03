package policyapi

// PolicyMappingsRequestBuilderDeleteQueryParameters represents the query parameters for Delete
type PolicyMappingsRequestBuilderDeleteQueryParameters struct {

	// AppName Name of the CDM application for which to remove the mapping.
	AppName *string `url:"appName,omitempty"`

	// DeployableName Name of the CDM deployable for which to remove the mapping.
	DeployableName *string `url:"deployableName,omitempty"`

	// PolicyName Name of the associated policy.
	PolicyName *string `url:"policyName,omitempty"`
}
