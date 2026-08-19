package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExportResult represents an export result.
type ExportResult struct {
	core.BaseModel
}

// NewExportResult instantiates a new ExportResult.
func NewExportResult() *ExportResult {
	return &ExportResult{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *ExportResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(stateKey, m.GetState),
		internalSerialization.SerializeStringFunc(statusKey, m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey, m.GetMessage),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *ExportResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:   internalSerialization.DeserializeStringFunc(m.setSysID),
		nameKey:    internalSerialization.DeserializeStringFunc(m.setName),
		stateKey:   internalSerialization.DeserializeStringFunc(m.setState),
		statusKey:  internalSerialization.DeserializeStringFunc(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc(m.setMessage),
	}
}

// GetSysID returns the sys id.
func (m *ExportResult) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportResult, *string](m, sysIDKey)
}
func (m *ExportResult) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetName returns the name.
func (m *ExportResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportResult, *string](m, nameKey)
}
func (m *ExportResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetState returns the state.
func (m *ExportResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportResult, *string](m, stateKey)
}
func (m *ExportResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}

// GetStatus returns the status.
func (m *ExportResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportResult, *string](m, statusKey)
}
func (m *ExportResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}

// GetMessage returns the message.
func (m *ExportResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportResult, *string](m, messageKey)
}
func (m *ExportResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, val)
}

// CreateExportResultFromDiscriminatorValue creates a new ExportResult from a ParseNode.
func CreateExportResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewExportResult(), nil
}
