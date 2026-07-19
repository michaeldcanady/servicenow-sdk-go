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
	SetAvailability([]AvailabilitySlot) error
	GetHasMore() (*bool, error)
	SetHasMore(*bool) error
	GetNextAvailableSlot() (any, error)
	SetNextAvailableSlot(any) error
	GetNoApptAvailable() (*bool, error)
	SetNoApptAvailable(*bool) error
	GetSuccess() (*bool, error)
	SetSuccess(*bool) error
	GetTimeZone() (*string, error)
	SetTimeZone(*string) error
	GetTimeZoneDisplayValue() (*string, error)
	SetTimeZoneDisplayValue(*string) error
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
		internalSerialization.SerializeCollectionOfObjectValuesFunc[AvailabilitySlot](availabilityKey, m.GetAvailability),
		internalSerialization.SerializeBoolFunc(hasMoreKey, m.GetHasMore),
		internalSerialization.SerializeAnyFunc(nextAvailableSlotKey, m.GetNextAvailableSlot),
		internalSerialization.SerializeBoolFunc(noApptAvailableKey, m.GetNoApptAvailable),
		internalSerialization.SerializeBoolFunc(successKey, m.GetSuccess),
		internalSerialization.SerializeStringFunc("time_zone", m.GetTimeZone),
		internalSerialization.SerializeStringFunc(timeZoneDisplayValueKey, m.GetTimeZoneDisplayValue),
	)
}

func (m *AvailabilityResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		availabilityKey:         internalSerialization.DeserializeCollectionOfObjectValuesFunc[AvailabilitySlot](CreateAvailabilitySlotFromDiscriminatorValue, m.SetAvailability),
		hasMoreKey:              internalSerialization.DeserializeBoolFunc(m.SetHasMore),
		nextAvailableSlotKey:    internalSerialization.DeserializeAnyFunc(m.SetNextAvailableSlot),
		noApptAvailableKey:      internalSerialization.DeserializeBoolFunc(m.SetNoApptAvailable),
		successKey:              internalSerialization.DeserializeBoolFunc(m.SetSuccess),
		timeZoneKey:             internalSerialization.DeserializeStringFunc(m.SetTimeZone),
		timeZoneDisplayValueKey: internalSerialization.DeserializeStringFunc(m.SetTimeZoneDisplayValue),
	}
}

func (m *AvailabilityResultModel) GetAvailability() ([]AvailabilitySlot, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, []AvailabilitySlot](m, availabilityKey)
}
func (m *AvailabilityResultModel) SetAvailability(val []AvailabilitySlot) error {
	return store.DefaultBackedModelMutatorFunc(m, availabilityKey, val)
}
func (m *AvailabilityResultModel) GetHasMore() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, *bool](m, hasMoreKey)
}
func (m *AvailabilityResultModel) SetHasMore(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, hasMoreKey, val)
}
func (m *AvailabilityResultModel) GetNextAvailableSlot() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, any](m, nextAvailableSlotKey)
}
func (m *AvailabilityResultModel) SetNextAvailableSlot(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, nextAvailableSlotKey, val)
}
func (m *AvailabilityResultModel) GetNoApptAvailable() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, *bool](m, noApptAvailableKey)
}
func (m *AvailabilityResultModel) SetNoApptAvailable(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, noApptAvailableKey, val)
}
func (m *AvailabilityResultModel) GetSuccess() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, *bool](m, successKey)
}
func (m *AvailabilityResultModel) SetSuccess(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, successKey, val)
}
func (m *AvailabilityResultModel) GetTimeZone() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, *string](m, timeZoneKey)
}
func (m *AvailabilityResultModel) SetTimeZone(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, timeZoneKey, val)
}
func (m *AvailabilityResultModel) GetTimeZoneDisplayValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityResultModel, *string](m, timeZoneDisplayValueKey)
}
func (m *AvailabilityResultModel) SetTimeZoneDisplayValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, timeZoneDisplayValueKey, val)
}

// AvailabilitySlot represents an available slot.
//
// The ServiceNow schema for this object is not documented, so it is modeled as a
// pass-through additional-data holder: every field it contains is exposed only
// through GetAdditionalData/SetAdditionalData, the same mechanism Kiota generates
// for open/dynamic objects.
type AvailabilitySlot interface {
	serialization.Parsable
	serialization.AdditionalDataHolder
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
	if conversion.IsNil(m) {
		return nil
	}
	return writer.WriteAdditionalData(m.GetAdditionalData())
}

// GetFieldDeserializers returns no known fields; the underlying kiota-serialization-json-go
// parse node automatically routes any unrecognized property into AdditionalData for models
// implementing serialization.AdditionalDataHolder, so no wildcard entry is needed here.
func (m *AvailabilitySlotModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}

func (m *AvailabilitySlotModel) GetAdditionalData() map[string]interface{} {
	val, err := store.DefaultBackedModelAccessorFunc[*AvailabilitySlotModel, map[string]interface{}](m, additionalDataKey)
	if err != nil || val == nil {
		return make(map[string]interface{})
	}
	return val
}

func (m *AvailabilitySlotModel) SetAdditionalData(value map[string]interface{}) {
	_ = store.DefaultBackedModelMutatorFunc(m, additionalDataKey, value)
}
