package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// Reference represents a reference field object.
type Reference interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetLink() (*string, error)
	setLink(*string) error
	GetValue() (*string, error)
	setValue(*string) error
}

type ReferenceModel struct {
	internal.BaseModel
}

func NewReference() *ReferenceModel {
	return &ReferenceModel{BaseModel: *internal.NewBaseModel()}
}

func (m *ReferenceModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(linkKey)(m.GetLink),
		internalSerialization.SerializeStringFunc(valueKey)(m.GetValue),
	)
}

func (m *ReferenceModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		linkKey:  internalSerialization.DeserializeStringFunc()(m.setLink),
		valueKey: internalSerialization.DeserializeStringFunc()(m.setValue),
	}
}

func (m *ReferenceModel) GetLink() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), linkKey)
}
func (m *ReferenceModel) setLink(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), linkKey, val)
}
func (m *ReferenceModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), valueKey)
}
func (m *ReferenceModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), valueKey, val)
}

func CreateReferenceFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewReference(), nil
}

// CaseResult represents a single case object.
type CaseResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetSysId() (*string, error)
	setSysId(*string) error
	GetNumber() (*string, error)
	setNumber(*string) error
	GetShortDescription() (*string, error)
	setShortDescription(*string) error
	GetDescription() (*string, error)
	setDescription(*string) error
	GetState() (*string, error)
	setState(*string) error
	GetPriority() (*string, error)
	setPriority(*string) error
	GetCategory() (*string, error)
	setCategory(*string) error
	GetAssignmentGroup() (Reference, error)
	setAssignmentGroup(Reference) error
	GetAssignedTo() (Reference, error)
	setAssignedTo(Reference) error
	GetContact() (Reference, error)
	setContact(Reference) error
	GetAccount() (Reference, error)
	setAccount(Reference) error
	GetSysCreatedOn() (*string, error)
	setSysCreatedOn(*string) error
	GetSysUpdatedOn() (*string, error)
	setSysUpdatedOn(*string) error
}

type CaseResultModel struct {
	internal.BaseModel
}

func NewCaseResult() *CaseResultModel {
	return &CaseResultModel{BaseModel: *internal.NewBaseModel()}
}

func (m *CaseResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(numberKey)(m.GetNumber),
		internalSerialization.SerializeStringFunc(shortDescriptionKey)(m.GetShortDescription),
		internalSerialization.SerializeStringFunc(descriptionKey)(m.GetDescription),
		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),
		internalSerialization.SerializeStringFunc(priorityKey)(m.GetPriority),
		internalSerialization.SerializeStringFunc(categoryKey)(m.GetCategory),
		internalSerialization.SerializeObjectValueFunc[Reference](assignmentGroupKey)(m.GetAssignmentGroup),
		internalSerialization.SerializeObjectValueFunc[Reference](assignedToKey)(m.GetAssignedTo),
		internalSerialization.SerializeObjectValueFunc[Reference](contactKey)(m.GetContact),
		internalSerialization.SerializeObjectValueFunc[Reference](accountKey)(m.GetAccount),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey)(m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(sysUpdatedOnKey)(m.GetSysUpdatedOn),
	)
}

func (m *CaseResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:            internalSerialization.DeserializeStringFunc()(m.setSysId),
		numberKey:           internalSerialization.DeserializeStringFunc()(m.setNumber),
		shortDescriptionKey: internalSerialization.DeserializeStringFunc()(m.setShortDescription),
		descriptionKey:      internalSerialization.DeserializeStringFunc()(m.setDescription),
		stateKey:            internalSerialization.DeserializeStringFunc()(m.setState),
		priorityKey:         internalSerialization.DeserializeStringFunc()(m.setPriority),
		categoryKey:         internalSerialization.DeserializeStringFunc()(m.setCategory),
		assignmentGroupKey:  internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue)(m.setAssignmentGroup),
		assignedToKey:       internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue)(m.setAssignedTo),
		contactKey:          internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue)(m.setContact),
		accountKey:          internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue)(m.setAccount),
		sysCreatedOnKey:     internalSerialization.DeserializeStringFunc()(m.setSysCreatedOn),
		sysUpdatedOnKey:     internalSerialization.DeserializeStringFunc()(m.setSysUpdatedOn),
	}
}

func (m *CaseResultModel) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}
func (m *CaseResultModel) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}
func (m *CaseResultModel) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), numberKey)
}
func (m *CaseResultModel) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), numberKey, val)
}
func (m *CaseResultModel) GetShortDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), shortDescriptionKey)
}
func (m *CaseResultModel) setShortDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), shortDescriptionKey, val)
}
func (m *CaseResultModel) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), descriptionKey)
}
func (m *CaseResultModel) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), descriptionKey, val)
}
func (m *CaseResultModel) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stateKey)
}
func (m *CaseResultModel) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stateKey, val)
}
func (m *CaseResultModel) GetPriority() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), priorityKey)
}
func (m *CaseResultModel) setPriority(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), priorityKey, val)
}
func (m *CaseResultModel) GetCategory() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), categoryKey)
}
func (m *CaseResultModel) setCategory(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), categoryKey, val)
}
func (m *CaseResultModel) GetAssignmentGroup() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, Reference](m.GetBackingStore(), assignmentGroupKey)
}
func (m *CaseResultModel) setAssignmentGroup(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), assignmentGroupKey, val)
}
func (m *CaseResultModel) GetAssignedTo() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, Reference](m.GetBackingStore(), assignedToKey)
}
func (m *CaseResultModel) setAssignedTo(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), assignedToKey, val)
}
func (m *CaseResultModel) GetContact() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, Reference](m.GetBackingStore(), contactKey)
}
func (m *CaseResultModel) setContact(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), contactKey, val)
}
func (m *CaseResultModel) GetAccount() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, Reference](m.GetBackingStore(), accountKey)
}
func (m *CaseResultModel) setAccount(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), accountKey, val)
}
func (m *CaseResultModel) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysCreatedOnKey)
}
func (m *CaseResultModel) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysCreatedOnKey, val)
}
func (m *CaseResultModel) GetSysUpdatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysUpdatedOnKey)
}
func (m *CaseResultModel) setSysUpdatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysUpdatedOnKey, val)
}

func CreateCaseResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCaseResult(), nil
}

// ActivitiesResult represents case activities.
type ActivitiesResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetSysId() (*string, error)
	setSysId(*string) error
	GetType() (*string, error)
	setType(*string) error
	GetValue() (*string, error)
	setValue(*string) error
	GetUser() (*string, error)
	setUser(*string) error
	GetSysCreatedOn() (*string, error)
	setSysCreatedOn(*string) error
	GetFieldName() (*string, error)
	setFieldName(*string) error
}

type ActivitiesResultModel struct {
	internal.BaseModel
}

func NewActivitiesResult() *ActivitiesResultModel {
	return &ActivitiesResultModel{BaseModel: *internal.NewBaseModel()}
}

func (m *ActivitiesResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
		internalSerialization.SerializeStringFunc(valueKey)(m.GetValue),
		internalSerialization.SerializeStringFunc(userKey)(m.GetUser),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey)(m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(fieldNameKey)(m.GetFieldName),
	)
}

func (m *ActivitiesResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:        internalSerialization.DeserializeStringFunc()(m.setSysId),
		typeKey:         internalSerialization.DeserializeStringFunc()(m.setType),
		valueKey:        internalSerialization.DeserializeStringFunc()(m.setValue),
		userKey:         internalSerialization.DeserializeStringFunc()(m.setUser),
		sysCreatedOnKey: internalSerialization.DeserializeStringFunc()(m.setSysCreatedOn),
		fieldNameKey:    internalSerialization.DeserializeStringFunc()(m.setFieldName),
	}
}

func (m *ActivitiesResultModel) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}
func (m *ActivitiesResultModel) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}
func (m *ActivitiesResultModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
}
func (m *ActivitiesResultModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, val)
}
func (m *ActivitiesResultModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), valueKey)
}
func (m *ActivitiesResultModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), valueKey, val)
}
func (m *ActivitiesResultModel) GetUser() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), userKey)
}
func (m *ActivitiesResultModel) setUser(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), userKey, val)
}
func (m *ActivitiesResultModel) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysCreatedOnKey)
}
func (m *ActivitiesResultModel) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysCreatedOnKey, val)
}
func (m *ActivitiesResultModel) GetFieldName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), fieldNameKey)
}
func (m *ActivitiesResultModel) setFieldName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), fieldNameKey, val)
}

func CreateActivitiesResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitiesResult(), nil
}

// FieldValuesResult represents field values.
type FieldValuesResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetLabel() (*string, error)
	setLabel(*string) error
	GetValue() (*string, error)
	setValue(*string) error
	GetSequence() (*int32, error)
	setSequence(*int32) error
	GetDependentValue() (*string, error)
	setDependentValue(*string) error
}

type FieldValuesResultModel struct {
	internal.BaseModel
}

func NewFieldValuesResult() *FieldValuesResultModel {
	return &FieldValuesResultModel{BaseModel: *internal.NewBaseModel()}
}

func (m *FieldValuesResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(labelKey)(m.GetLabel),
		internalSerialization.SerializeStringFunc(valueKey)(m.GetValue),
		internalSerialization.SerializeInt32Func(sequenceKey)(m.GetSequence),
		internalSerialization.SerializeStringFunc(dependentValueKey)(m.GetDependentValue),
	)
}

func (m *FieldValuesResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		labelKey:          internalSerialization.DeserializeStringFunc()(m.setLabel),
		valueKey:          internalSerialization.DeserializeStringFunc()(m.setValue),
		sequenceKey:       internalSerialization.DeserializeInt32Func()(m.setSequence),
		dependentValueKey: internalSerialization.DeserializeStringFunc()(m.setDependentValue),
	}
}

func (m *FieldValuesResultModel) GetLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), labelKey)
}
func (m *FieldValuesResultModel) setLabel(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), labelKey, val)
}
func (m *FieldValuesResultModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), valueKey)
}
func (m *FieldValuesResultModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), valueKey, val)
}
func (m *FieldValuesResultModel) GetSequence() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *int32](m.GetBackingStore(), sequenceKey)
}
func (m *FieldValuesResultModel) setSequence(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sequenceKey, val)
}
func (m *FieldValuesResultModel) GetDependentValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), dependentValueKey)
}
func (m *FieldValuesResultModel) setDependentValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), dependentValueKey, val)
}

func CreateFieldValuesResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFieldValuesResult(), nil
}
