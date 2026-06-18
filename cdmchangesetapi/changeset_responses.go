package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangesetsResponse represents a response for changesets.
type ChangesetsResponse interface {
	core.ServiceNowCollectionResponse[*ChangesetResult]
}

func CreateChangesetsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ChangesetResult](CreateChangesetResultFromDiscriminatorValue), nil
}

// ChangesetActivityResponse represents a response for changeset activity.
type ChangesetActivityResponse interface {
	core.ServiceNowCollectionResponse[*ChangesetActivityResult]
}

func CreateChangesetActivityResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ChangesetActivityResult](CreateChangesetActivityResultFromDiscriminatorValue), nil
}

// CommitStatusResponse represents a response for commit status.
type CommitStatusResponse interface {
	core.ServiceNowItemResponse[*CommitStatusResult]
}

func CreateCommitStatusResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*CommitStatusResult](CreateCommitStatusResultFromDiscriminatorValue), nil
}

// ImpactedSharedComponentsResponse represents a response for impacted shared components.
type ImpactedSharedComponentsResponse interface {
	core.ServiceNowCollectionResponse[*ImpactedSharedComponentResult]
}

func CreateImpactedSharedComponentsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ImpactedSharedComponentResult](CreateImpactedSharedComponentResultFromDiscriminatorValue), nil
}

// ImpactedDeployablesResponse represents a response for impacted deployables.
type ImpactedDeployablesResponse interface {
	core.ServiceNowCollectionResponse[*ImpactedDeployableResult]
}

func CreateImpactedDeployablesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ImpactedDeployableResult](CreateImpactedDeployableResultFromDiscriminatorValue), nil
}

// ImpactedDeployablesBySysIdResponse represents a response for impacted deployables (path-based).
type ImpactedDeployablesBySysIdResponse interface {
	core.ServiceNowCollectionResponse[*ImpactedDeployableBySysIdResult]
}

func CreateImpactedDeployablesBySysIdResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowCollectionResponse[*ImpactedDeployableBySysIdResult](CreateImpactedDeployableBySysIdResultFromDiscriminatorValue), nil
}
