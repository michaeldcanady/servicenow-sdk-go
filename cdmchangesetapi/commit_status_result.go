package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommitStatusResult represents a commit status.
type CommitStatusResult struct {
	core.BaseModel
}

// NewCommitStatusResult instantiates a new CommitStatusResult.
func NewCommitStatusResult() *CommitStatusResult {
	return &CommitStatusResult{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *CommitStatusResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(stateKey, m.GetState),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *CommitStatusResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		stateKey: internalSerialization.DeserializeStringFunc(m.setState),
	}
}

// GetState returns the state.
func (m *CommitStatusResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CommitStatusResult, *string](m, stateKey)
}
func (m *CommitStatusResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}

// CreateCommitStatusResultFromDiscriminatorValue creates a new CommitStatusResult from a ParseNode.
func CreateCommitStatusResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCommitStatusResult(), nil
}
