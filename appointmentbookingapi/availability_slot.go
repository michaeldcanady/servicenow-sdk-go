package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	availableKey        = "available"
	endDateDisplayKey   = "end_date_display"
	startDateDisplayKey = "start_date_display"
)

// AvailabilitySlot represents an available slot.
//
// The ServiceNow schema for this object is not documented, so it is modeled as a
// pass-through additional-data holder: every field it contains is exposed only
// through GetAdditionalData/SetAdditionalData, the same mechanism Kiota generates
// for open/dynamic objects.
type AvailabilitySlot interface {
	serialization.Parsable
	kiotaStore.BackedModel
	GetAvailable() (*bool, error)
	SetAvailable(*bool) error
	// TODO: date
	GetEndDate() (*string, error)
	SetEndDate(*string) error
	GetEndDateDisplay() (*string, error)
	SetEndDateDisplay(*string) error
	// TODO: date
	GetEndDateUTC() (*string, error)
	SetEndDateUTC(*string) error
	// TODO: date
	GetStartDate() (*string, error)
	SetStartDate(*string) error
	GetStartDateDisplay() (*string, error)
	SetStartDateDisplay(*string) error
	// TODO: date
	GetStartDateUTC() (*string, error)
	SetStartDateUTC(*string) error
}

// AvailabilitySlotModel represents the availability slot model.
type AvailabilitySlotModel struct {
	core.BaseModel
}

// NewAvailabilitySlot creates a new instance of AvailabilitySlotModel.
func NewAvailabilitySlot() *AvailabilitySlotModel {
	return &AvailabilitySlotModel{BaseModel: *core.NewBaseModel()}
}

// CreateAvailabilitySlotFromDiscriminatorValue creates a new AvailabilitySlot from a ParseNode.
func CreateAvailabilitySlotFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAvailabilitySlot(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *AvailabilitySlotModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeBoolFunc(availableKey, m.GetAvailable),
		internalSerialization.SerializeStringFunc(endDateKey, m.GetEndDate),
		internalSerialization.SerializeStringFunc(endDateDisplayKey, m.GetEndDateDisplay),
		internalSerialization.SerializeStringFunc(endDateUTCKey, m.GetEndDateUTC),
		internalSerialization.SerializeStringFunc(startDateKey, m.GetStartDate),
		internalSerialization.SerializeStringFunc(startDateDisplayKey, m.GetStartDateDisplay),
		internalSerialization.SerializeStringFunc(startDateUTCKey, m.GetStartDateUTC),
	)
}

// GetFieldDeserializers
func (m *AvailabilitySlotModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		availableKey:        internalSerialization.DeserializeBoolFunc(m.SetAvailable),
		endDateKey:          internalSerialization.DeserializeStringFunc(m.SetEndDate),
		endDateDisplayKey:   internalSerialization.DeserializeStringFunc(m.SetEndDateDisplay),
		endDateUTCKey:       internalSerialization.DeserializeStringFunc(m.SetEndDateUTC),
		startDateKey:        internalSerialization.DeserializeStringFunc(m.SetStartDate),
		startDateDisplayKey: internalSerialization.DeserializeStringFunc(m.SetStartDateDisplay),
		startDateUTCKey:     internalSerialization.DeserializeStringFunc(m.SetStartDateUTC),
	}
}

func (m *AvailabilitySlotModel) GetAvailable() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilitySlotModel, *bool](m, availableKey)
}

func (m *AvailabilitySlotModel) SetAvailable(available *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, availableKey, available)
}

func (m *AvailabilitySlotModel) GetEndDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilitySlotModel, *string](m, endDateKey)
}

func (m *AvailabilitySlotModel) SetEndDate(endDate *string) error {
	return store.DefaultBackedModelMutatorFunc(m, endDateKey, endDate)
}

func (m *AvailabilitySlotModel) GetEndDateDisplay() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilitySlotModel, *string](m, endDateDisplayKey)
}

func (m *AvailabilitySlotModel) SetEndDateDisplay(endDateDisplay *string) error {
	return store.DefaultBackedModelMutatorFunc(m, endDateDisplayKey, endDateDisplay)
}

func (m *AvailabilitySlotModel) GetEndDateUTC() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilitySlotModel, *string](m, endDateUTCKey)
}

func (m *AvailabilitySlotModel) SetEndDateUTC(endDateUTC *string) error {
	return store.DefaultBackedModelMutatorFunc(m, endDateUTCKey, endDateUTC)
}

func (m *AvailabilitySlotModel) GetStartDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilitySlotModel, *string](m, startDateKey)
}

func (m *AvailabilitySlotModel) SetStartDate(startDate *string) error {
	return store.DefaultBackedModelMutatorFunc(m, startDateKey, startDate)
}

func (m *AvailabilitySlotModel) GetStartDateDisplay() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilitySlotModel, *string](m, startDateDisplayKey)
}

func (m *AvailabilitySlotModel) SetStartDateDisplay(startDateDisplay *string) error {
	return store.DefaultBackedModelMutatorFunc(m, startDateDisplayKey, startDateDisplay)
}

func (m *AvailabilitySlotModel) GetStartDateUTC() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilitySlotModel, *string](m, startDateUTCKey)
}

func (m *AvailabilitySlotModel) SetStartDateUTC(startDateUTC *string) error {
	return store.DefaultBackedModelMutatorFunc(m, startDateUTCKey, startDateUTC)
}
