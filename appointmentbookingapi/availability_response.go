package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// AvailabilityResponse represents the availability response.
type AvailabilityResponse = core.ServiceNowItemResponse[*AvailabilityResultModel]

// CreateAvailabilityResponseFromDiscriminatorValue is a factory for creating an AvailabilityResponse.
func CreateAvailabilityResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*AvailabilityResultModel](CreateAvailabilityResultFromDiscriminatorValue), nil
}

// AvailabilityResult represents the result object in availability response.
type AvailabilityResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetAvailability() ([]AvailabilitySlot, error)
	setAvailability([]AvailabilitySlot) error
	GetHasMore() (*bool, error)
	setHasMore(*bool) error
	GetNextAvailableSlot() (any, error)
	setNextAvailableSlot(any) error
	GetNoApptAvailable() (*bool, error)
	setNoApptAvailable(*bool) error
	GetSuccess() (*bool, error)
	setSuccess(*bool) error
	GetTimeZone() (*string, error)
	setTimeZone(*string) error
	GetTimeZoneDisplayValue() (*string, error)
	setTimeZoneDisplayValue(*string) error
}

type AvailabilityResultModel struct {
	core.BaseModel
}

func NewAvailabilityResult() *AvailabilityResultModel {
	return &AvailabilityResultModel{BaseModel: *core.NewBaseModel()}
}

func CreateAvailabilityResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAvailabilityResult(), nil
}

func (m *AvailabilityResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeCollectionOfObjectValuesFunc[AvailabilitySlot](availabilityKey)(m.GetAvailability),
		internalSerialization.SerializeBoolFunc(hasMoreKey)(m.GetHasMore),
		internalSerialization.SerializeAnyFunc(nextAvailableSlotKey)(m.GetNextAvailableSlot),
		internalSerialization.SerializeBoolFunc(noApptAvailableKey)(m.GetNoApptAvailable),
		internalSerialization.SerializeBoolFunc(successKey)(m.GetSuccess),
		internalSerialization.SerializeStringFunc("time_zone")(m.GetTimeZone),
		internalSerialization.SerializeStringFunc(timeZoneDisplayValueKey)(m.GetTimeZoneDisplayValue),
	)
}

func (m *AvailabilityResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		availabilityKey:         internalSerialization.DeserializeCollectionOfObjectValuesFunc[AvailabilitySlot](CreateAvailabilitySlotFromDiscriminatorValue)(m.setAvailability),
		hasMoreKey:              internalSerialization.DeserializeBoolFunc()(m.setHasMore),
		nextAvailableSlotKey:    internalSerialization.DeserializeAnyFunc()(m.setNextAvailableSlot),
		noApptAvailableKey:      internalSerialization.DeserializeBoolFunc()(m.setNoApptAvailable),
		successKey:              internalSerialization.DeserializeBoolFunc()(m.setSuccess),
		"time_zone":             internalSerialization.DeserializeStringFunc()(m.setTimeZone),
		timeZoneDisplayValueKey: internalSerialization.DeserializeStringFunc()(m.setTimeZoneDisplayValue),
	}
}

func (m *AvailabilityResultModel) GetAvailability() ([]AvailabilitySlot, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, []AvailabilitySlot](m.GetBackingStore(), availabilityKey)
}
func (m *AvailabilityResultModel) setAvailability(val []AvailabilitySlot) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), availabilityKey, val)
}
func (m *AvailabilityResultModel) GetHasMore() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), hasMoreKey)
}
func (m *AvailabilityResultModel) setHasMore(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), hasMoreKey, val)
}
func (m *AvailabilityResultModel) GetNextAvailableSlot() (any, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), nextAvailableSlotKey)
}
func (m *AvailabilityResultModel) setNextAvailableSlot(val any) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nextAvailableSlotKey, val)
}
func (m *AvailabilityResultModel) GetNoApptAvailable() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), noApptAvailableKey)
}
func (m *AvailabilityResultModel) setNoApptAvailable(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), noApptAvailableKey, val)
}
func (m *AvailabilityResultModel) GetSuccess() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), successKey)
}
func (m *AvailabilityResultModel) setSuccess(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), successKey, val)
}
func (m *AvailabilityResultModel) GetTimeZone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), "time_zone")
}
func (m *AvailabilityResultModel) setTimeZone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), "time_zone", val)
}
func (m *AvailabilityResultModel) GetTimeZoneDisplayValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), timeZoneDisplayValueKey)
}
func (m *AvailabilityResultModel) setTimeZoneDisplayValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), timeZoneDisplayValueKey, val)
}

// AvailabilitySlot represents an available slot.
type AvailabilitySlot interface {
	serialization.Parsable
	kiotaStore.BackedModel
}

type AvailabilitySlotModel struct {
	core.BaseModel
}

func NewAvailabilitySlot() *AvailabilitySlotModel {
	return &AvailabilitySlotModel{BaseModel: *core.NewBaseModel()}
}

func CreateAvailabilitySlotFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAvailabilitySlot(), nil
}

func (m *AvailabilitySlotModel) Serialize(writer serialization.SerializationWriter) error {
	val, _ := m.GetBackingStore().Get("additionalData")
	if val != nil {
		return writer.WriteAdditionalData(val.(map[string]interface{}))
	}
	return nil
}

func (m *AvailabilitySlotModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		"*": func(n serialization.ParseNode) error {
			val, _ := n.GetRawValue()
			return m.GetBackingStore().Set("additionalData", val)
		},
	}
}
