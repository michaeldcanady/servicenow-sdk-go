package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommitStatusResult represents a commit status.
type CommitStatusResult struct {
	core.BaseModel
}

func NewCommitStatusResult() *CommitStatusResult {
	return &CommitStatusResult{BaseModel: *core.NewBaseModel()}
}

func (m *CommitStatusResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),
	)
}

func (m *CommitStatusResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		stateKey: internalSerialization.DeserializeStringFunc()(m.setState),
	}
}

func (m *CommitStatusResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CommitStatusResult, *string](m, stateKey)
}
func (m *CommitStatusResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}

func CreateCommitStatusResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCommitStatusResult(), nil
}
