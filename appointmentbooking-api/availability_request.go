package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

type AvailabilityRequest interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetCatalogId() (*string, error)
	setCatalogId(*string) error
	GetEndDate() (*string, error)
	setEndDate(*string) error
	GetFetchDaysSlot() (*bool, error)
	setFetchDaysSlot(*bool) error
	GetFullDay() (*bool, error)
	setFullDay(*bool) error
	GetGetNextAvailableSlot() (*bool, error)
	setGetNextAvailableSlot(*bool) error
	GetLimit() (*int32, error)
	setLimit(*int32) error
	GetLocation() (*string, error)
	setLocation(*string) error
	GetOpenedFor() (*string, error)
	setOpenedFor(*string) error
	GetOtherInputs() (any, error)
	setOtherInputs(any) error
	GetServiceConfigRule() (*string, error)
	setServiceConfigRule(*string) error
	GetStartDate() (*string, error)
	setStartDate(*string) error
	GetTaskId() (*string, error)
	setTaskId(*string) error
	GetTaskTable() (*string, error)
	setTaskTable(*string) error
	GetUseReadReplica() (*bool, error)
	setUseReadReplica(*bool) error
	GetView() (*string, error)
	setView(*string) error
}

type AvailabilityRequestModel struct {
	internal.BaseModel
}

func NewAvailabilityRequest() *AvailabilityRequestModel {
	return &AvailabilityRequestModel{BaseModel: *internal.NewBaseModel()}
}

func (m *AvailabilityRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc("catalog_id")(m.GetCatalogId),
		internalSerialization.SerializeStringFunc(endDateKey)(m.GetEndDate),
		internalSerialization.SerializeBoolFunc(fetchDaysSlotKey)(m.GetFetchDaysSlot),
		internalSerialization.SerializeBoolFunc(fullDayKey)(m.GetFullDay),
		internalSerialization.SerializeBoolFunc(getNextAvailableSlotKey)(m.GetGetNextAvailableSlot),
		internalSerialization.SerializeInt32Func(limitKey)(m.GetLimit),
		internalSerialization.SerializeStringFunc("location")(m.GetLocation),
		internalSerialization.SerializeStringFunc(openedForKey)(m.GetOpenedFor),
		internalSerialization.SerializeAnyFunc(otherInputsKey)(m.GetOtherInputs),
		internalSerialization.SerializeStringFunc(serviceConfigRuleKey)(m.GetServiceConfigRule),
		internalSerialization.SerializeStringFunc(startDateKey)(m.GetStartDate),
		internalSerialization.SerializeStringFunc("taskId")(m.GetTaskId),
		internalSerialization.SerializeStringFunc("task_table")(m.GetTaskTable),
		internalSerialization.SerializeBoolFunc(useReadReplicaKey)(m.GetUseReadReplica),
		internalSerialization.SerializeStringFunc(viewKey)(m.GetView),
	)
}

func (m *AvailabilityRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		"catalog_id":            internalSerialization.DeserializeStringFunc()(m.setCatalogId),
		endDateKey:              internalSerialization.DeserializeStringFunc()(m.setEndDate),
		fetchDaysSlotKey:        internalSerialization.DeserializeBoolFunc()(m.setFetchDaysSlot),
		fullDayKey:              internalSerialization.DeserializeBoolFunc()(m.setFullDay),
		getNextAvailableSlotKey: internalSerialization.DeserializeBoolFunc()(m.setGetNextAvailableSlot),
		limitKey:                internalSerialization.DeserializeInt32Func()(m.setLimit),
		"location":              internalSerialization.DeserializeStringFunc()(m.setLocation),
		openedForKey:            internalSerialization.DeserializeStringFunc()(m.setOpenedFor),
		otherInputsKey:          internalSerialization.DeserializeAnyFunc()(m.setOtherInputs),
		serviceConfigRuleKey:    internalSerialization.DeserializeStringFunc()(m.setServiceConfigRule),
		startDateKey:            internalSerialization.DeserializeStringFunc()(m.setStartDate),
		"taskId":                internalSerialization.DeserializeStringFunc()(m.setTaskId),
		"task_table":            internalSerialization.DeserializeStringFunc()(m.setTaskTable),
		useReadReplicaKey:       internalSerialization.DeserializeBoolFunc()(m.setUseReadReplica),
		viewKey:                 internalSerialization.DeserializeStringFunc()(m.setView),
	}
}

func (m *AvailabilityRequestModel) GetCatalogId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), "catalog_id")
}
func (m *AvailabilityRequestModel) setCatalogId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), "catalog_id", val)
}
func (m *AvailabilityRequestModel) GetEndDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), endDateKey)
}
func (m *AvailabilityRequestModel) setEndDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), endDateKey, val)
}
func (m *AvailabilityRequestModel) GetFetchDaysSlot() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), fetchDaysSlotKey)
}
func (m *AvailabilityRequestModel) setFetchDaysSlot(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), fetchDaysSlotKey, val)
}
func (m *AvailabilityRequestModel) GetFullDay() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), fullDayKey)
}
func (m *AvailabilityRequestModel) setFullDay(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), fullDayKey, val)
}
func (m *AvailabilityRequestModel) GetGetNextAvailableSlot() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), getNextAvailableSlotKey)
}
func (m *AvailabilityRequestModel) setGetNextAvailableSlot(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), getNextAvailableSlotKey, val)
}
func (m *AvailabilityRequestModel) GetLimit() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int32](m.GetBackingStore(), limitKey)
}
func (m *AvailabilityRequestModel) setLimit(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), limitKey, val)
}
func (m *AvailabilityRequestModel) GetLocation() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), "location")
}
func (m *AvailabilityRequestModel) setLocation(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), "location", val)
}
func (m *AvailabilityRequestModel) GetOpenedFor() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), openedForKey)
}
func (m *AvailabilityRequestModel) setOpenedFor(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), openedForKey, val)
}
func (m *AvailabilityRequestModel) GetOtherInputs() (any, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), otherInputsKey)
}
func (m *AvailabilityRequestModel) setOtherInputs(val any) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), otherInputsKey, val)
}
func (m *AvailabilityRequestModel) GetServiceConfigRule() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), serviceConfigRuleKey)
}
func (m *AvailabilityRequestModel) setServiceConfigRule(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), serviceConfigRuleKey, val)
}
func (m *AvailabilityRequestModel) GetStartDate() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), startDateKey)
}
func (m *AvailabilityRequestModel) setStartDate(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), startDateKey, val)
}
func (m *AvailabilityRequestModel) GetTaskId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), "taskId")
}
func (m *AvailabilityRequestModel) setTaskId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), "taskId", val)
}
func (m *AvailabilityRequestModel) GetTaskTable() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), "task_table")
}
func (m *AvailabilityRequestModel) setTaskTable(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), "task_table", val)
}
func (m *AvailabilityRequestModel) GetUseReadReplica() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), useReadReplicaKey)
}
func (m *AvailabilityRequestModel) setUseReadReplica(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), useReadReplicaKey, val)
}
func (m *AvailabilityRequestModel) GetView() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), viewKey)
}
func (m *AvailabilityRequestModel) setView(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), viewKey, val)
}

func CreateAvailabilityRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAvailabilityRequest(), nil
}
