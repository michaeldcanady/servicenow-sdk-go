package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// UserTimeFormat represents userTimeFormat nested object.
type UserTimeFormat interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetType() (*string, error)
	setType(*string) error
	GetValue() (*string, error)
	setValue(*string) error
}

type UserTimeFormatModel struct {
	core.BaseModel
}

func NewUserTimeFormat() *UserTimeFormatModel {
	return &UserTimeFormatModel{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateUserTimeFormatFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUserTimeFormat(), nil
}

func (m *UserTimeFormatModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
	)
}

func (m *UserTimeFormatModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		typeKey:  internalSerialization.DeserializeStringFunc(m.setType),
		valueKey: internalSerialization.DeserializeStringFunc(m.setValue),
	}
}

func (m *UserTimeFormatModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserTimeFormatModel, *string](m, typeKey)
}
func (m *UserTimeFormatModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}
func (m *UserTimeFormatModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UserTimeFormatModel, *string](m, valueKey)
}
func (m *UserTimeFormatModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}
