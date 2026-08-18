package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetContentResult represents the result details retrieved from the getContent endpoint.
// The spec defines no schema for this response (it is a CI/relationship content graph).
type GetContentResult struct {
	core.BaseModel
}

// NewGetContentResult creates a new instance of GetContentResult.
func NewGetContentResult() *GetContentResult {
	return &GetContentResult{BaseModel: *core.NewBaseModel()}
}

// CreateGetContentResultFromDiscriminatorValue creates a new GetContentResult from a ParseNode.
func CreateGetContentResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewGetContentResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *GetContentResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *GetContentResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey: internalSerialization.DeserializeStringFunc(m.SetSysID),
	}
}

// GetSysID returns the sys id value.
func (m *GetContentResult) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*GetContentResult, *string](m, sysIDKey)
}

// SetSysID
func (m *GetContentResult) SetSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetResult returns the result itself.
func (m *GetContentResult) GetResult() (*GetContentResult, error) {
	return m, nil
}
