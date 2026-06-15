package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangesetsResponse represents a response for changesets.
type ChangesetsResponse interface {
	internal.ServiceNowCollectionResponse[*ChangesetResult]
}

func CreateChangesetsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowCollectionResponse[*ChangesetResult](CreateChangesetResultFromDiscriminatorValue), nil
}

// ChangesetActivityResponse represents a response for changeset activity.
type ChangesetActivityResponse interface {
	internal.ServiceNowCollectionResponse[*ChangesetActivityResult]
}

func CreateChangesetActivityResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowCollectionResponse[*ChangesetActivityResult](CreateChangesetActivityResultFromDiscriminatorValue), nil
}

// CommitStatusResponse represents a response for commit status.
type CommitStatusResponse interface {
	internal.ServiceNowItemResponse[*CommitStatusResult]
}

func CreateCommitStatusResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowItemResponse[*CommitStatusResult](CreateCommitStatusResultFromDiscriminatorValue), nil
}

// ImpactedSharedComponentsResponse represents a response for impacted shared components.
type ImpactedSharedComponentsResponse interface {
	internal.ServiceNowCollectionResponse[*ImpactedSharedComponentResult]
}

func CreateImpactedSharedComponentsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowCollectionResponse[*ImpactedSharedComponentResult](CreateImpactedSharedComponentResultFromDiscriminatorValue), nil
}

// ImpactedDeployablesResponse represents a response for impacted deployables.
type ImpactedDeployablesResponse interface {
	internal.ServiceNowCollectionResponse[*ImpactedDeployableResult]
}

func CreateImpactedDeployablesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowCollectionResponse[*ImpactedDeployableResult](CreateImpactedDeployableResultFromDiscriminatorValue), nil
}

// ImpactedDeployablesBySysIdResponse represents a response for impacted deployables (path-based).
type ImpactedDeployablesBySysIdResponse interface {
	internal.ServiceNowCollectionResponse[*ImpactedDeployableBySysIdResult]
}

func CreateImpactedDeployablesBySysIdResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return internal.NewBaseServiceNowCollectionResponse[*ImpactedDeployableBySysIdResult](CreateImpactedDeployableBySysIdResultFromDiscriminatorValue), nil
}
