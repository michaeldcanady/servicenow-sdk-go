package cdmchangesetapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// ChangesetsRequestBuilderGetQueryParameters represents query parameters for GET /changesets.
type ChangesetsRequestBuilderGetQueryParameters struct {
	AppName *string `url:"appName,omitempty"`
	Number  *string `url:"number,omitempty"`
	State   *string `url:"state,omitempty"`
}

// ChangesetsRequestBuilderDeleteQueryParameters represents query parameters for DELETE /changesets.
type ChangesetsRequestBuilderDeleteQueryParameters struct {
	ChangesetNumber *string `url:"changesetNumber,omitempty"`
}

// ChangesetActivityRequestBuilderGetQueryParameters represents query parameters for GET /changesets/activity.
type ChangesetActivityRequestBuilderGetQueryParameters struct {
	ChangesetNumber *string  `url:"changesetNumber,omitempty"`
	ReturnFields    []string `url:"returnFields,omitempty"`
}

// ImpactedSharedComponentsRequestBuilderGetQueryParameters represents query parameters for GET /changesets/impacted-shared-components.
type ImpactedSharedComponentsRequestBuilderGetQueryParameters struct {
	ChangesetNumber *string  `url:"changesetNumber,omitempty"`
	ReturnFields    []string `url:"returnFields,omitempty"`
}

// ImpactedDeployablesRequestBuilderGetQueryParameters represents query parameters for GET /changesets/impacted-deployables.
type ImpactedDeployablesRequestBuilderGetQueryParameters struct {
	ChangesetNumber *string  `url:"changesetNumber,omitempty"`
	ReturnFields    []string `url:"returnFields,omitempty"`
}

// Request Configurations
type ChangesetsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ChangesetsRequestBuilderGetQueryParameters]
type ChangesetsRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[ChangesetsRequestBuilderDeleteQueryParameters]
type ChangesetActivityRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ChangesetActivityRequestBuilderGetQueryParameters]
type CommitStatusRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type ImpactedSharedComponentsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ImpactedSharedComponentsRequestBuilderGetQueryParameters]
type ImpactedDeployablesRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ImpactedDeployablesRequestBuilderGetQueryParameters]
type ImpactedDeployablesBySysIdRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
