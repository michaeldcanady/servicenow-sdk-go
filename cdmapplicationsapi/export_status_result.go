package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExportStatusResult represents the status of an export.
type ExportStatusResult struct {
	core.BaseModel
}

func NewExportStatusResult() *ExportStatusResult {
	return &ExportStatusResult{BaseModel: *core.NewBaseModel()}
}

func (m *ExportStatusResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(stateKey, m.GetState),
		internalSerialization.SerializeStringFunc(statusKey, m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey, m.GetMessage),
		internalSerialization.SerializeStringFunc(progressKey, m.GetProgress),
	)
}

func (m *ExportStatusResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		stateKey:    internalSerialization.DeserializeStringFunc(m.setState),
		statusKey:   internalSerialization.DeserializeStringFunc(m.setStatus),
		messageKey:  internalSerialization.DeserializeStringFunc(m.setMessage),
		progressKey: internalSerialization.DeserializeStringFunc(m.setProgress),
	}
}

func (m *ExportStatusResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportStatusResult, *string](m, stateKey)
}
func (m *ExportStatusResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}
func (m *ExportStatusResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportStatusResult, *string](m, statusKey)
}
func (m *ExportStatusResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}
func (m *ExportStatusResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportStatusResult, *string](m, messageKey)
}
func (m *ExportStatusResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, val)
}
func (m *ExportStatusResult) GetProgress() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExportStatusResult, *string](m, progressKey)
}
func (m *ExportStatusResult) setProgress(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, progressKey, val)
}

func CreateExportStatusResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewExportStatusResult(), nil
}
