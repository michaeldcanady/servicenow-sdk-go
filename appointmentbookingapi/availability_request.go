package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// AvailabilityRequest represents the availability request.
type AvailabilityRequest interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetCatalogID() (*string, error)
	SetCatalogID(*string) error
	GetEndDate() (*string, error)
	SetEndDate(*string) error
	GetFetchDaysSlot() (*bool, error)
	SetFetchDaysSlot(*bool) error
	GetFullDay() (*bool, error)
	SetFullDay(*bool) error
	GetGetNextAvailableSlot() (*bool, error)
	SetGetNextAvailableSlot(*bool) error
	GetLimit() (*int32, error)
	SetLimit(*int32) error
	GetLocation() (*string, error)
	SetLocation(*string) error
	GetOpenedFor() (*string, error)
	SetOpenedFor(*string) error
	GetOtherInputs() (any, error)
	SetOtherInputs(any) error
	GetServiceConfigRule() (*string, error)
	SetServiceConfigRule(*string) error
	GetStartDate() (*string, error)
	SetStartDate(*string) error
	GetTaskID() (*string, error)
	SetTaskID(*string) error
	GetTaskTable() (*string, error)
	SetTaskTable(*string) error
	GetUseReadReplica() (*bool, error)
	SetUseReadReplica(*bool) error
	GetView() (*string, error)
	SetView(*string) error
}

// AvailabilityRequestModel represents the availability request model.
type AvailabilityRequestModel struct {
	core.BaseModel
}

// NewAvailabilityRequest creates a new instance of AvailabilityRequestModel.
func NewAvailabilityRequest() *AvailabilityRequestModel {
	return &AvailabilityRequestModel{BaseModel: *core.NewBaseModel()}
}

// CreateAvailabilityRequestFromDiscriminatorValue creates a new AvailabilityRequest from a ParseNode.
func CreateAvailabilityRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAvailabilityRequest(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *AvailabilityRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(catalogIDKey, m.GetCatalogID),
		internalSerialization.SerializeStringFunc(endDateKey, m.GetEndDate),
		internalSerialization.SerializeBoolFunc(fetchDaysSlotKey, m.GetFetchDaysSlot),
		internalSerialization.SerializeBoolFunc(fullDayKey, m.GetFullDay),
		internalSerialization.SerializeBoolFunc(getNextAvailableSlotKey, m.GetGetNextAvailableSlot),
		internalSerialization.SerializeInt32Func(limitKey, m.GetLimit),
		internalSerialization.SerializeStringFunc(locationKey, m.GetLocation),
		internalSerialization.SerializeStringFunc(openedForKey, m.GetOpenedFor),
		internalSerialization.SerializeAnyFunc(otherInputsKey, m.GetOtherInputs),
		internalSerialization.SerializeStringFunc(serviceConfigRuleKey, m.GetServiceConfigRule),
		internalSerialization.SerializeStringFunc(startDateKey, m.GetStartDate),
		internalSerialization.SerializeStringFunc(taskIDKey, m.GetTaskID),
		internalSerialization.SerializeStringFunc(taskTableKey, m.GetTaskTable),
		internalSerialization.SerializeBoolFunc(useReadReplicaKey, m.GetUseReadReplica),
		internalSerialization.SerializeStringFunc(viewKey, m.GetView),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *AvailabilityRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		catalogIDKey:            internalSerialization.DeserializeStringFunc(m.SetCatalogID),
		endDateKey:              internalSerialization.DeserializeStringFunc(m.SetEndDate),
		fetchDaysSlotKey:        internalSerialization.DeserializeBoolFunc(m.SetFetchDaysSlot),
		fullDayKey:              internalSerialization.DeserializeBoolFunc(m.SetFullDay),
		getNextAvailableSlotKey: internalSerialization.DeserializeBoolFunc(m.SetGetNextAvailableSlot),
		limitKey:                internalSerialization.DeserializeInt32Func(m.SetLimit),
		locationKey:             internalSerialization.DeserializeStringFunc(m.SetLocation),
		openedForKey:            internalSerialization.DeserializeStringFunc(m.SetOpenedFor),
		otherInputsKey:          internalSerialization.DeserializeAnyFunc(m.SetOtherInputs),
		serviceConfigRuleKey:    internalSerialization.DeserializeStringFunc(m.SetServiceConfigRule),
		startDateKey:            internalSerialization.DeserializeStringFunc(m.SetStartDate),
		taskIDKey:               internalSerialization.DeserializeStringFunc(m.SetTaskID),
		taskTableKey:            internalSerialization.DeserializeStringFunc(m.SetTaskTable),
		useReadReplicaKey:       internalSerialization.DeserializeBoolFunc(m.SetUseReadReplica),
		viewKey:                 internalSerialization.DeserializeStringFunc(m.SetView),
	}
}

// GetCatalogID returns the catalog id value.
func (m *AvailabilityRequestModel) GetCatalogID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, catalogIDKey)
}

// SetCatalogID sets the catalog id value.
func (m *AvailabilityRequestModel) SetCatalogID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, catalogIDKey, val)
}

// GetEndDate returns the end date value.
func (m *AvailabilityRequestModel) GetEndDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, endDateKey)
}

// SetEndDate sets the end date value.
func (m *AvailabilityRequestModel) SetEndDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, endDateKey, val)
}

// GetFetchDaysSlot returns the fetch days slot value.
func (m *AvailabilityRequestModel) GetFetchDaysSlot() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *bool](m, fetchDaysSlotKey)
}

// SetFetchDaysSlot sets the fetch days slot value.
func (m *AvailabilityRequestModel) SetFetchDaysSlot(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, fetchDaysSlotKey, val)
}

// GetFullDay returns the full day value.
func (m *AvailabilityRequestModel) GetFullDay() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *bool](m, fullDayKey)
}

// SetFullDay sets the full day value.
func (m *AvailabilityRequestModel) SetFullDay(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, fullDayKey, val)
}

// GetGetNextAvailableSlot returns the get next available slot value.
func (m *AvailabilityRequestModel) GetGetNextAvailableSlot() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *bool](m, getNextAvailableSlotKey)
}

// SetGetNextAvailableSlot sets the get next available slot value.
func (m *AvailabilityRequestModel) SetGetNextAvailableSlot(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, getNextAvailableSlotKey, val)
}

// GetLimit returns the limit value.
func (m *AvailabilityRequestModel) GetLimit() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *int32](m, limitKey)
}

// SetLimit sets the limit value.
func (m *AvailabilityRequestModel) SetLimit(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m, limitKey, val)
}

// GetLocation returns the location value.
func (m *AvailabilityRequestModel) GetLocation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, locationKey)
}

// SetLocation sets the location value.
func (m *AvailabilityRequestModel) SetLocation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, locationKey, val)
}

// GetOpenedFor returns the opened for value.
func (m *AvailabilityRequestModel) GetOpenedFor() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, openedForKey)
}

// SetOpenedFor sets the opened for value.
func (m *AvailabilityRequestModel) SetOpenedFor(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, openedForKey, val)
}

// GetOtherInputs returns the other inputs value.
func (m *AvailabilityRequestModel) GetOtherInputs() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, any](m, otherInputsKey)
}

// SetOtherInputs sets the other inputs value.
func (m *AvailabilityRequestModel) SetOtherInputs(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, otherInputsKey, val)
}

// GetServiceConfigRule returns the service config rule value.
func (m *AvailabilityRequestModel) GetServiceConfigRule() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, serviceConfigRuleKey)
}

// SetServiceConfigRule sets the service config rule value.
func (m *AvailabilityRequestModel) SetServiceConfigRule(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, serviceConfigRuleKey, val)
}

// GetStartDate returns the start date value.
func (m *AvailabilityRequestModel) GetStartDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, startDateKey)
}

// SetStartDate sets the start date value.
func (m *AvailabilityRequestModel) SetStartDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, startDateKey, val)
}

// GetTaskID returns the task id value.
func (m *AvailabilityRequestModel) GetTaskID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, taskIDKey)
}

// SetTaskID sets the task id value.
func (m *AvailabilityRequestModel) SetTaskID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskIDKey, val)
}

// GetTaskTable returns the task table value.
func (m *AvailabilityRequestModel) GetTaskTable() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, taskTableKey)
}

// SetTaskTable sets the task table value.
func (m *AvailabilityRequestModel) SetTaskTable(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskTableKey, val)
}

// GetUseReadReplica returns the use read replica value.
func (m *AvailabilityRequestModel) GetUseReadReplica() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *bool](m, useReadReplicaKey)
}

// SetUseReadReplica sets the use read replica value.
func (m *AvailabilityRequestModel) SetUseReadReplica(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, useReadReplicaKey, val)
}

// GetView returns the view value.
func (m *AvailabilityRequestModel) GetView() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, viewKey)
}

// SetView sets the view value.
func (m *AvailabilityRequestModel) SetView(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, viewKey, val)
}
