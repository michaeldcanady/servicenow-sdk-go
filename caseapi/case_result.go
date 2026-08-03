package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// CaseResult represents a single case object.
type CaseResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetSysID() (*string, error)
	setSysID(*string) error
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

// CaseResultModel implementation of CaseResult.
type CaseResultModel struct {
	core.BaseModel
}

// NewCaseResult creates a new instance of CaseResult.
func NewCaseResult() *CaseResultModel {
	return &CaseResultModel{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *CaseResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(numberKey, m.GetNumber),
		internalSerialization.SerializeStringFunc(shortDescriptionKey, m.GetShortDescription),
		internalSerialization.SerializeStringFunc(descriptionKey, m.GetDescription),
		internalSerialization.SerializeStringFunc(stateKey, m.GetState),
		internalSerialization.SerializeStringFunc(priorityKey, m.GetPriority),
		internalSerialization.SerializeStringFunc(categoryKey, m.GetCategory),
		internalSerialization.SerializeObjectValueFunc[Reference](assignmentGroupKey, m.GetAssignmentGroup),
		internalSerialization.SerializeObjectValueFunc[Reference](assignedToKey, m.GetAssignedTo),
		internalSerialization.SerializeObjectValueFunc[Reference](contactKey, m.GetContact),
		internalSerialization.SerializeObjectValueFunc[Reference](accountKey, m.GetAccount),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey, m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(sysUpdatedOnKey, m.GetSysUpdatedOn),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *CaseResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:            internalSerialization.DeserializeStringFunc(m.setSysID),
		numberKey:           internalSerialization.DeserializeStringFunc(m.setNumber),
		shortDescriptionKey: internalSerialization.DeserializeStringFunc(m.setShortDescription),
		descriptionKey:      internalSerialization.DeserializeStringFunc(m.setDescription),
		stateKey:            internalSerialization.DeserializeStringFunc(m.setState),
		priorityKey:         internalSerialization.DeserializeStringFunc(m.setPriority),
		categoryKey:         internalSerialization.DeserializeStringFunc(m.setCategory),
		assignmentGroupKey:  internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue, m.setAssignmentGroup),
		assignedToKey:       internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue, m.setAssignedTo),
		contactKey:          internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue, m.setContact),
		accountKey:          internalSerialization.DeserializeObjectValueFunc[Reference](CreateReferenceFromDiscriminatorValue, m.setAccount),
		sysCreatedOnKey:     internalSerialization.DeserializeStringFunc(m.setSysCreatedOn),
		sysUpdatedOnKey:     internalSerialization.DeserializeStringFunc(m.setSysUpdatedOn),
	}
}

// GetSysID ...
func (m *CaseResultModel) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, sysIDKey)
}
func (m *CaseResultModel) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetNumber ...
func (m *CaseResultModel) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, numberKey)
}
func (m *CaseResultModel) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}

// GetShortDescription ...
func (m *CaseResultModel) GetShortDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, shortDescriptionKey)
}
func (m *CaseResultModel) setShortDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, shortDescriptionKey, val)
}

// GetDescription ...
func (m *CaseResultModel) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, descriptionKey)
}
func (m *CaseResultModel) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, descriptionKey, val)
}

// GetState ...
func (m *CaseResultModel) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, stateKey)
}
func (m *CaseResultModel) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}

// GetPriority ...
func (m *CaseResultModel) GetPriority() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, priorityKey)
}
func (m *CaseResultModel) setPriority(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, priorityKey, val)
}

// GetCategory ...
func (m *CaseResultModel) GetCategory() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, categoryKey)
}
func (m *CaseResultModel) setCategory(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, categoryKey, val)
}

// GetAssignmentGroup ...
func (m *CaseResultModel) GetAssignmentGroup() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, Reference](m, assignmentGroupKey)
}
func (m *CaseResultModel) setAssignmentGroup(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, assignmentGroupKey, val)
}

// GetAssignedTo ...
func (m *CaseResultModel) GetAssignedTo() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, Reference](m, assignedToKey)
}
func (m *CaseResultModel) setAssignedTo(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, assignedToKey, val)
}

// GetContact ...
func (m *CaseResultModel) GetContact() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, Reference](m, contactKey)
}
func (m *CaseResultModel) setContact(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, contactKey, val)
}

// GetAccount ...
func (m *CaseResultModel) GetAccount() (Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, Reference](m, accountKey)
}
func (m *CaseResultModel) setAccount(val Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, accountKey, val)
}

// GetSysCreatedOn ...
func (m *CaseResultModel) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, sysCreatedOnKey)
}
func (m *CaseResultModel) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnKey, val)
}

// GetSysUpdatedOn ...
func (m *CaseResultModel) GetSysUpdatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CaseResultModel, *string](m, sysUpdatedOnKey)
}
func (m *CaseResultModel) setSysUpdatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysUpdatedOnKey, val)
}

// CreateCaseResultFromDiscriminatorValue creates a new instance of CaseResult.
func CreateCaseResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCaseResult(), nil
}
