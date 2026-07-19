package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

type AvailabilityRequest interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetCatalogId() (*string, error)
	SetCatalogId(*string) error
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
	GetTaskId() (*string, error)
	SetTaskId(*string) error
	GetTaskTable() (*string, error)
	SetTaskTable(*string) error
	GetUseReadReplica() (*bool, error)
	SetUseReadReplica(*bool) error
	GetView() (*string, error)
	SetView(*string) error
}

type AvailabilityRequestModel struct {
	core.BaseModel
}

func NewAvailabilityRequest() *AvailabilityRequestModel {
	return &AvailabilityRequestModel{BaseModel: *core.NewBaseModel()}
}

func CreateAvailabilityRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAvailabilityRequest(), nil
}

func (m *AvailabilityRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(catalogIDKey, m.GetCatalogId),
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
		internalSerialization.SerializeStringFunc(taskIDKey, m.GetTaskId),
		internalSerialization.SerializeStringFunc(taskTableKey, m.GetTaskTable),
		internalSerialization.SerializeBoolFunc(useReadReplicaKey, m.GetUseReadReplica),
		internalSerialization.SerializeStringFunc(viewKey, m.GetView),
	)
}

func (m *AvailabilityRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		catalogIDKey:            internalSerialization.DeserializeStringFunc(m.SetCatalogId),
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
		taskIDKey:               internalSerialization.DeserializeStringFunc(m.SetTaskId),
		taskTableKey:            internalSerialization.DeserializeStringFunc(m.SetTaskTable),
		useReadReplicaKey:       internalSerialization.DeserializeBoolFunc(m.SetUseReadReplica),
		viewKey:                 internalSerialization.DeserializeStringFunc(m.SetView),
	}
}

func (m *AvailabilityRequestModel) GetCatalogId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, catalogIDKey)
}
func (m *AvailabilityRequestModel) SetCatalogId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, catalogIDKey, val)
}
func (m *AvailabilityRequestModel) GetEndDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, endDateKey)
}
func (m *AvailabilityRequestModel) SetEndDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, endDateKey, val)
}
func (m *AvailabilityRequestModel) GetFetchDaysSlot() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *bool](m, fetchDaysSlotKey)
}
func (m *AvailabilityRequestModel) SetFetchDaysSlot(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, fetchDaysSlotKey, val)
}
func (m *AvailabilityRequestModel) GetFullDay() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *bool](m, fullDayKey)
}
func (m *AvailabilityRequestModel) SetFullDay(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, fullDayKey, val)
}
func (m *AvailabilityRequestModel) GetGetNextAvailableSlot() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *bool](m, getNextAvailableSlotKey)
}
func (m *AvailabilityRequestModel) SetGetNextAvailableSlot(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, getNextAvailableSlotKey, val)
}
func (m *AvailabilityRequestModel) GetLimit() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *int32](m, limitKey)
}
func (m *AvailabilityRequestModel) SetLimit(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m, limitKey, val)
}
func (m *AvailabilityRequestModel) GetLocation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, locationKey)
}
func (m *AvailabilityRequestModel) SetLocation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, locationKey, val)
}
func (m *AvailabilityRequestModel) GetOpenedFor() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, openedForKey)
}
func (m *AvailabilityRequestModel) SetOpenedFor(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, openedForKey, val)
}
func (m *AvailabilityRequestModel) GetOtherInputs() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, any](m, otherInputsKey)
}
func (m *AvailabilityRequestModel) SetOtherInputs(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, otherInputsKey, val)
}
func (m *AvailabilityRequestModel) GetServiceConfigRule() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, serviceConfigRuleKey)
}
func (m *AvailabilityRequestModel) SetServiceConfigRule(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, serviceConfigRuleKey, val)
}
func (m *AvailabilityRequestModel) GetStartDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, startDateKey)
}
func (m *AvailabilityRequestModel) SetStartDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, startDateKey, val)
}
func (m *AvailabilityRequestModel) GetTaskId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, taskIDKey)
}
func (m *AvailabilityRequestModel) SetTaskId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskIDKey, val)
}
func (m *AvailabilityRequestModel) GetTaskTable() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, taskTableKey)
}
func (m *AvailabilityRequestModel) SetTaskTable(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskTableKey, val)
}
func (m *AvailabilityRequestModel) GetUseReadReplica() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *bool](m, useReadReplicaKey)
}
func (m *AvailabilityRequestModel) SetUseReadReplica(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, useReadReplicaKey, val)
}
func (m *AvailabilityRequestModel) GetView() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AvailabilityRequestModel, *string](m, viewKey)
}
func (m *AvailabilityRequestModel) SetView(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, viewKey, val)
}
