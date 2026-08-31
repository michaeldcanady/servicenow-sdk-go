// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

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

// ChangesetsRequestBuilderGetRequestConfiguration represents the configuration for a Get request.
type ChangesetsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ChangesetsRequestBuilderGetQueryParameters]

// ChangesetsRequestBuilderDeleteRequestConfiguration represents the DELETE request configuration for the Changesets resource.
type ChangesetsRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[ChangesetsRequestBuilderDeleteQueryParameters]

// ChangesetActivityRequestBuilderGetRequestConfiguration represents the GET request configuration for the Changeset Activity resource.
type ChangesetActivityRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ChangesetActivityRequestBuilderGetQueryParameters]

// CommitStatusRequestBuilderGetRequestConfiguration represents the GET request configuration for the Commit Status resource.
type CommitStatusRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// ImpactedSharedComponentsRequestBuilderGetRequestConfiguration represents the GET request configuration for the Impacted Shared Components resource.
type ImpactedSharedComponentsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ImpactedSharedComponentsRequestBuilderGetQueryParameters]

// ImpactedDeployablesRequestBuilderGetRequestConfiguration represents the GET request configuration for the Impacted Deployables resource.
type ImpactedDeployablesRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ImpactedDeployablesRequestBuilderGetQueryParameters]

// ImpactedDeployablesBySysIDRequestBuilderGetRequestConfiguration represents the GET request configuration for the Impacted Deployables By Sys ID resource.
type ImpactedDeployablesBySysIDRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
