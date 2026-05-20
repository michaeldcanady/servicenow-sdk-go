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
