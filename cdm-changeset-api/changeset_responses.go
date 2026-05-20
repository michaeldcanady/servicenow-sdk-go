package cdmchangesetapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangesetsResponse represents a response for changesets.
type ChangesetsResponse interface {
	newInternal.ServiceNowCollectionResponse[*ChangesetResult]
}

func CreateChangesetsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowCollectionResponse[*ChangesetResult](CreateChangesetResultFromDiscriminatorValue), nil
}

// ChangesetActivityResponse represents a response for changeset activity.
type ChangesetActivityResponse interface {
	newInternal.ServiceNowCollectionResponse[*ChangesetActivityResult]
}

func CreateChangesetActivityResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowCollectionResponse[*ChangesetActivityResult](CreateChangesetActivityResultFromDiscriminatorValue), nil
}

// CommitStatusResponse represents a response for commit status.
type CommitStatusResponse interface {
	newInternal.ServiceNowItemResponse[*CommitStatusResult]
}

func CreateCommitStatusResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*CommitStatusResult](CreateCommitStatusResultFromDiscriminatorValue), nil
}

// ImpactedSharedComponentsResponse represents a response for impacted shared components.
type ImpactedSharedComponentsResponse interface {
	newInternal.ServiceNowCollectionResponse[*ImpactedSharedComponentResult]
}

func CreateImpactedSharedComponentsResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowCollectionResponse[*ImpactedSharedComponentResult](CreateImpactedSharedComponentResultFromDiscriminatorValue), nil
}

// ImpactedDeployablesResponse represents a response for impacted deployables.
type ImpactedDeployablesResponse interface {
	newInternal.ServiceNowCollectionResponse[*ImpactedDeployableResult]
}

func CreateImpactedDeployablesResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowCollectionResponse[*ImpactedDeployableResult](CreateImpactedDeployableResultFromDiscriminatorValue), nil
}

// ImpactedDeployablesBySysIdResponse represents a response for impacted deployables (path-based).
type ImpactedDeployablesBySysIdResponse interface {
	newInternal.ServiceNowCollectionResponse[*ImpactedDeployableBySysIdResult]
}

func CreateImpactedDeployablesBySysIdResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowCollectionResponse[*ImpactedDeployableBySysIdResult](CreateImpactedDeployableBySysIdResultFromDiscriminatorValue), nil
}
