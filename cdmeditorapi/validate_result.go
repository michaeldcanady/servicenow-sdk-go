package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// ValidationResult represents validation status.
type ValidationResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetStatus() (*string, error)
	setStatus(*string) error
	GetErrors() (any, error)
	setErrors(any) error
	GetWarnings() (any, error)
	setWarnings(any) error
}

type ValidationResultModel struct {
	core.BaseModel
}

func NewValidationResult() *ValidationResultModel {
	return &ValidationResultModel{BaseModel: *core.NewBaseModel()}
}

func (m *ValidationResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(statusKey)(m.GetStatus),
		internalSerialization.SerializeAnyFunc(errorsKey)(m.GetErrors),
		internalSerialization.SerializeAnyFunc(warningsKey)(m.GetWarnings),
	)
}

func (m *ValidationResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		statusKey:   internalSerialization.DeserializeStringFunc()(m.setStatus),
		errorsKey:   internalSerialization.DeserializeAnyFunc()(m.setErrors),
		warningsKey: internalSerialization.DeserializeAnyFunc()(m.setWarnings),
	}
}

func (m *ValidationResultModel) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ValidationResultModel, *string](m, statusKey)
}
func (m *ValidationResultModel) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}
func (m *ValidationResultModel) GetErrors() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*ValidationResultModel, any](m, errorsKey)
}
func (m *ValidationResultModel) setErrors(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, errorsKey, val)
}
func (m *ValidationResultModel) GetWarnings() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*ValidationResultModel, any](m, warningsKey)
}
func (m *ValidationResultModel) setWarnings(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, warningsKey, val)
}

func CreateValidationResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewValidationResult(), nil
}
