package cdmchangesetapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// ChangesetsRequestBuilderGetQueryParameters represents query parameters for GET /changesets.
type ChangesetsRequestBuilderGetQueryParameters struct {
	AppName *string `uriparametername:"appName"`
	Number  *string `uriparametername:"number"`
	State   *string `uriparametername:"state"`
}

// ChangesetsRequestBuilderDeleteQueryParameters represents query parameters for DELETE /changesets.
type ChangesetsRequestBuilderDeleteQueryParameters struct {
	ChangesetNumber *string `uriparametername:"changesetNumber"`
}

// ChangesetActivityRequestBuilderGetQueryParameters represents query parameters for GET /changesets/activity.
type ChangesetActivityRequestBuilderGetQueryParameters struct {
	ChangesetNumber *string  `uriparametername:"changesetNumber"`
	ReturnFields    []string `uriparametername:"returnFields"`
}

// ImpactedSharedComponentsRequestBuilderGetQueryParameters represents query parameters for GET /changesets/impacted-shared-components.
type ImpactedSharedComponentsRequestBuilderGetQueryParameters struct {
	ChangesetNumber *string  `uriparametername:"changesetNumber"`
	ReturnFields    []string `uriparametername:"returnFields"`
}

// ImpactedDeployablesRequestBuilderGetQueryParameters represents query parameters for GET /changesets/impacted-deployables.
type ImpactedDeployablesRequestBuilderGetQueryParameters struct {
	ChangesetNumber *string  `uriparametername:"changesetNumber"`
	ReturnFields    []string `uriparametername:"returnFields"`
}

// Request Configurations
type ChangesetsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ChangesetsRequestBuilderGetQueryParameters]
type ChangesetsRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[ChangesetsRequestBuilderDeleteQueryParameters]
type ChangesetActivityRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ChangesetActivityRequestBuilderGetQueryParameters]
type CommitStatusRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type ImpactedSharedComponentsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ImpactedSharedComponentsRequestBuilderGetQueryParameters]
type ImpactedDeployablesRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ImpactedDeployablesRequestBuilderGetQueryParameters]
type ImpactedDeployablesBySysIdRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
