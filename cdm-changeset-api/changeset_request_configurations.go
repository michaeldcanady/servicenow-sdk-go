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

// Request Configurations
type ChangesetsRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ChangesetsRequestBuilderGetQueryParameters]
type ChangesetsRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[ChangesetsRequestBuilderDeleteQueryParameters]
type ChangesetActivityRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ChangesetActivityRequestBuilderGetQueryParameters]
type CommitStatusRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
