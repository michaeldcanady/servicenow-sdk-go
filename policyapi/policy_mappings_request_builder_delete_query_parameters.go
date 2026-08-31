// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package policyapi

// PolicyMappingsRequestBuilderDeleteQueryParameters are the query parameters for the Delete method.
type PolicyMappingsRequestBuilderDeleteQueryParameters struct {
	// AppName Name of the CDM deployable for which to map the policy.
	AppName *string `uriparametername:"appName"`
	// DeployableName Name of the CDM deployable for which to map the policy.
	DeployableName *string `uriparametername:"deployableName"`
	// PolicyName Name of the associated policy.
	PolicyName *string `uriparametername:"policyName"`
}
