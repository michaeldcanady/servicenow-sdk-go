package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangesetsResponse represents a response for changesets.
type ChangesetsResponse interface {
	core.ServiceNowCollectionResponse[*ChangesetResult]
}

// CreateChangesetsResponseFromDiscriminatorValue creates a new ChangesetsResponse from a ParseNode.
func CreateChangesetsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ChangesetResult](CreateChangesetResultFromDiscriminatorValue), nil
}

// ChangesetActivityResponse represents a response for changeset activity.
type ChangesetActivityResponse interface {
	core.ServiceNowCollectionResponse[*ChangesetActivityResult]
}

// CreateChangesetActivityResponseFromDiscriminatorValue creates a new ChangesetActivityResponse from a ParseNode.
func CreateChangesetActivityResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ChangesetActivityResult](CreateChangesetActivityResultFromDiscriminatorValue), nil
}

// CommitStatusResponse represents a response for commit status.
type CommitStatusResponse interface {
	core.ServiceNowItemResponse[*CommitStatusResult]
}

// CreateCommitStatusResponseFromDiscriminatorValue creates a new CommitStatusResponse from a ParseNode.
func CreateCommitStatusResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*CommitStatusResult](CreateCommitStatusResultFromDiscriminatorValue), nil
}

// ImpactedSharedComponentsResponse represents a response for impacted shared components.
type ImpactedSharedComponentsResponse interface {
	core.ServiceNowCollectionResponse[*ImpactedSharedComponentResult]
}

// CreateImpactedSharedComponentsResponseFromDiscriminatorValue creates a new ImpactedSharedComponentsResponse from a ParseNode.
func CreateImpactedSharedComponentsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ImpactedSharedComponentResult](CreateImpactedSharedComponentResultFromDiscriminatorValue), nil
}

// ImpactedDeployablesResponse represents a response for impacted deployables.
type ImpactedDeployablesResponse interface {
	core.ServiceNowCollectionResponse[*ImpactedDeployableResult]
}

// CreateImpactedDeployablesResponseFromDiscriminatorValue creates a new ImpactedDeployablesResponse from a ParseNode.
func CreateImpactedDeployablesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ImpactedDeployableResult](CreateImpactedDeployableResultFromDiscriminatorValue), nil
}

// ImpactedDeployablesBySysIDResponse represents a response for impacted deployables (path-based).
type ImpactedDeployablesBySysIDResponse interface {
	core.ServiceNowCollectionResponse[*ImpactedDeployableBySysIDResult]
}

// CreateImpactedDeployablesBySysIDResponseFromDiscriminatorValue creates a new ImpactedDeployablesBySysIDResponse from a ParseNode.
func CreateImpactedDeployablesBySysIDResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ImpactedDeployableBySysIDResult](CreateImpactedDeployableBySysIDResultFromDiscriminatorValue), nil
}
