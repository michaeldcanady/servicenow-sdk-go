package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExportResult represents an export result.
type ExportResult struct {
	core.BaseModel
}

func NewExportResult() *ExportResult {
	return &ExportResult{BaseModel: *core.NewBaseModel()}
}

func (m *ExportResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey, m.GetSysId),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(stateKey, m.GetState),
		internalSerialization.SerializeStringFunc(statusKey, m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey, m.GetMessage),
	)
}

func (m *ExportResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:   internalSerialization.DeserializeStringFunc(m.setSysId),
		nameKey:    internalSerialization.DeserializeStringFunc(m.setName),
		stateKey:   internalSerialization.DeserializeStringFunc(m.setState),
		statusKey:  internalSerialization.DeserializeStringFunc(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc(m.setMessage),
	}
}

func (m *ExportResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportResult, *string](m, sysIdKey)
}
func (m *ExportResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIdKey, val)
}
func (m *ExportResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportResult, *string](m, nameKey)
}
func (m *ExportResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}
func (m *ExportResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportResult, *string](m, stateKey)
}
func (m *ExportResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}
func (m *ExportResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportResult, *string](m, statusKey)
}
func (m *ExportResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}
func (m *ExportResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportResult, *string](m, messageKey)
}
func (m *ExportResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, val)
}

func CreateExportResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewExportResult(), nil
}
