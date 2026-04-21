package policyapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	descriptionKey       = "description"
	documentKey          = "document"
	document_refKey      = "document_ref"
	errorKey             = "error"
	exceptionKey         = "exception"
	exception_allowedKey = "exception_allowed"
	input_statusKey      = "input_status"
	last_updated_byKey   = "last_updated_by"
	numberKey            = "number"
	policyKey            = "policy"
	reasonKey            = "reason"
	stateKey             = "state"
	sys_class_nameKey    = "sys_class_name"
)

// PolicyMapping defines the interface for the PolicyMapping model.
type PolicyMapping interface {

	// GetDescription gets the description property value.
	GetDescription() (*string, error)
	// SetDescription sets the description property value.
	SetDescription(*string) error

	// GetDocument gets the document property value.
	GetDocument() (*string, error)
	// SetDocument sets the document property value.
	SetDocument(*string) error

	// GetDocumentRef gets the document_ref property value.
	GetDocumentRef() (LinkRef, error)
	// SetDocumentRef sets the document_ref property value.
	SetDocumentRef(LinkRef) error

	// GetError gets the error property value.
	GetError() (newInternal.ServicenowError, error)
	// SetError sets the error property value.
	SetError(newInternal.ServicenowError) error

	// GetException gets the exception property value.
	GetException() (*string, error)
	// SetException sets the exception property value.
	SetException(*string) error

	// GetExceptionAllowed gets the exception_allowed property value.
	GetExceptionAllowed() (*string, error)
	// SetExceptionAllowed sets the exception_allowed property value.
	SetExceptionAllowed(*string) error

	// GetInputStatus gets the input_status property value.
	GetInputStatus() (*string, error)
	// SetInputStatus sets the input_status property value.
	SetInputStatus(*string) error

	// GetLastUpdatedBy gets the last_updated_by property value.
	GetLastUpdatedBy() (LinkRef, error)
	// SetLastUpdatedBy sets the last_updated_by property value.
	SetLastUpdatedBy(LinkRef) error

	// GetNumber gets the number property value.
	GetNumber() (*string, error)
	// SetNumber sets the number property value.
	SetNumber(*string) error

	// GetPolicy gets the policy property value.
	GetPolicy() (LinkRef, error)
	// SetPolicy sets the policy property value.
	SetPolicy(LinkRef) error

	// GetReason gets the reason property value.
	GetReason() (*string, error)
	// SetReason sets the reason property value.
	SetReason(*string) error

	// GetState gets the state property value.
	GetState() (*string, error)
	// SetState sets the state property value.
	SetState(*string) error

	// GetSysClassName gets the sys_class_name property value.
	GetSysClassName() (*string, error)
	// SetSysClassName sets the sys_class_name property value.
	SetSysClassName(*string) error

	// GetSysCreatedBy gets the sys_created_by property value.
	GetSysCreatedBy() (*string, error)
	// SetSysCreatedBy sets the sys_created_by property value.
	SetSysCreatedBy(*string) error

	// GetSysCreatedOn gets the sys_created_on property value.
	GetSysCreatedOn() (*string, error)
	// SetSysCreatedOn sets the sys_created_on property value.
	SetSysCreatedOn(*string) error

	// GetSysId gets the sys_id property value.
	GetSysId() (*string, error)
	// SetSysId sets the sys_id property value.
	SetSysId(*string) error

	// GetSysUpdatedBy gets the sys_updated_by property value.
	GetSysUpdatedBy() (*string, error)
	// SetSysUpdatedBy sets the sys_updated_by property value.
	SetSysUpdatedBy(*string) error

	// GetSysUpdatedOn gets the sys_updated_on property value.
	GetSysUpdatedOn() (*string, error)
	// SetSysUpdatedOn sets the sys_updated_on property value.
	SetSysUpdatedOn(*string) error

	serialization.Parsable
	kiotaStore.BackedModel
}

// PolicyMappingModel is the concrete implementation of the PolicyMapping interface.
type PolicyMappingModel struct {
	newInternal.Model
}

// NewPolicyMapping creates a new instance of PolicyMappingModel with a backing store.
func NewPolicyMapping() *PolicyMappingModel {
	return &PolicyMappingModel{
		newInternal.NewBaseModel(),
	}
}

// CreatePolicyMappingFromDiscriminatorValue is a factory function for creating PolicyMappingModel instances during deserialization.
func CreatePolicyMappingFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPolicyMapping(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *PolicyMappingModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,

		internalSerialization.SerializeStringFunc(descriptionKey)(m.GetDescription),

		internalSerialization.SerializeStringFunc(documentKey)(m.GetDocument),

		internalSerialization.SerializeObjectValueFunc[LinkRef](document_refKey)(m.GetDocumentRef),

		internalSerialization.SerializeObjectValueFunc[*newInternal.ServicenowError](errorKey)(m.GetError),

		internalSerialization.SerializeStringFunc(exceptionKey)(m.GetException),

		internalSerialization.SerializeStringFunc(exception_allowedKey)(m.GetExceptionAllowed),

		internalSerialization.SerializeStringFunc(input_statusKey)(m.GetInputStatus),

		internalSerialization.SerializeObjectValueFunc[LinkRef](last_updated_byKey)(m.GetLastUpdatedBy),

		internalSerialization.SerializeStringFunc(numberKey)(m.GetNumber),

		internalSerialization.SerializeObjectValueFunc[LinkRef](policyKey)(m.GetPolicy),

		internalSerialization.SerializeStringFunc(reasonKey)(m.GetReason),

		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),

		internalSerialization.SerializeStringFunc(sys_class_nameKey)(m.GetSysClassName),

		internalSerialization.SerializeStringFunc(sys_created_byKey)(m.GetSysCreatedBy),

		internalSerialization.SerializeStringFunc(sys_created_onKey)(m.GetSysCreatedOn),

		internalSerialization.SerializeStringFunc(sys_idKey)(m.GetSysId),

		internalSerialization.SerializeStringFunc(sys_updated_byKey)(m.GetSysUpdatedBy),

		internalSerialization.SerializeStringFunc(sys_updated_onKey)(m.GetSysUpdatedOn),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *PolicyMappingModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{

		descriptionKey: internalSerialization.DeserializeStringFunc()(m.SetDescription),

		documentKey: internalSerialization.DeserializeStringFunc()(m.SetDocument),

		document_refKey: internalSerialization.DeserializeObjectValueFunc[LinkRef](CreateLinkRefFromDiscriminatorValue)(m.SetDocumentRef),

		errorKey: internalSerialization.DeserializeObjectValueFunc[*newInternal.ServicenowError](newInternal.CreateServiceNowErrorFromDiscriminatorValue)(m.SetError),

		exceptionKey: internalSerialization.DeserializeStringFunc()(m.SetException),

		exception_allowedKey: internalSerialization.DeserializeStringFunc()(m.SetExceptionAllowed),

		input_statusKey: internalSerialization.DeserializeStringFunc()(m.SetInputStatus),

		last_updated_byKey: internalSerialization.DeserializeObjectValueFunc[LinkRef](CreateLinkRefFromDiscriminatorValue)(m.SetLastUpdatedBy),

		numberKey: internalSerialization.DeserializeStringFunc()(m.SetNumber),

		policyKey: internalSerialization.DeserializeObjectValueFunc[LinkRef](CreateLinkRefFromDiscriminatorValue)(m.SetPolicy),

		reasonKey: internalSerialization.DeserializeStringFunc()(m.SetReason),

		stateKey: internalSerialization.DeserializeStringFunc()(m.SetState),

		sys_class_nameKey: internalSerialization.DeserializeStringFunc()(m.SetSysClassName),

		sys_created_byKey: internalSerialization.DeserializeStringFunc()(m.SetSysCreatedBy),

		sys_created_onKey: internalSerialization.DeserializeStringFunc()(m.SetSysCreatedOn),

		sys_idKey: internalSerialization.DeserializeStringFunc()(m.SetSysId),

		sys_updated_byKey: internalSerialization.DeserializeStringFunc()(m.SetSysUpdatedBy),

		sys_updated_onKey: internalSerialization.DeserializeStringFunc()(m.SetSysUpdatedOn),
	}
}

// GetDescription returns the description property value.
func (m *PolicyMappingModel) GetDescription() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, descriptionKey)
}

// SetDescription sets the description property value.
func (m *PolicyMappingModel) SetDescription(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, descriptionKey, val)
}

// GetDocument returns the document property value.
func (m *PolicyMappingModel) GetDocument() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, documentKey)
}

// SetDocument sets the document property value.
func (m *PolicyMappingModel) SetDocument(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, documentKey, val)
}

// GetDocumentRef returns the document_ref property value.
func (m *PolicyMappingModel) GetDocumentRef() (LinkRef, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, LinkRef](backingStore, document_refKey)
}

// SetDocumentRef sets the document_ref property value.
func (m *PolicyMappingModel) SetDocumentRef(val LinkRef) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, document_refKey, val)
}

// GetError returns the error property value.
func (m *PolicyMappingModel) GetError() (*newInternal.ServicenowError, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *newInternal.ServicenowError](backingStore, errorKey)
}

// SetError sets the error property value.
func (m *PolicyMappingModel) SetError(val *newInternal.ServicenowError) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, errorKey, val)
}

// GetException returns the exception property value.
func (m *PolicyMappingModel) GetException() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, exceptionKey)
}

// SetException sets the exception property value.
func (m *PolicyMappingModel) SetException(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, exceptionKey, val)
}

// GetExceptionAllowed returns the exception_allowed property value.
func (m *PolicyMappingModel) GetExceptionAllowed() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, exception_allowedKey)
}

// SetExceptionAllowed sets the exception_allowed property value.
func (m *PolicyMappingModel) SetExceptionAllowed(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, exception_allowedKey, val)
}

// GetInputStatus returns the input_status property value.
func (m *PolicyMappingModel) GetInputStatus() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, input_statusKey)
}

// SetInputStatus sets the input_status property value.
func (m *PolicyMappingModel) SetInputStatus(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, input_statusKey, val)
}

// GetLastUpdatedBy returns the last_updated_by property value.
func (m *PolicyMappingModel) GetLastUpdatedBy() (LinkRef, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, LinkRef](backingStore, last_updated_byKey)
}

// SetLastUpdatedBy sets the last_updated_by property value.
func (m *PolicyMappingModel) SetLastUpdatedBy(val LinkRef) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, last_updated_byKey, val)
}

// GetNumber returns the number property value.
func (m *PolicyMappingModel) GetNumber() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, numberKey)
}

// SetNumber sets the number property value.
func (m *PolicyMappingModel) SetNumber(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, numberKey, val)
}

// GetPolicy returns the policy property value.
func (m *PolicyMappingModel) GetPolicy() (LinkRef, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, LinkRef](backingStore, policyKey)
}

// SetPolicy sets the policy property value.
func (m *PolicyMappingModel) SetPolicy(val LinkRef) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, policyKey, val)
}

// GetReason returns the reason property value.
func (m *PolicyMappingModel) GetReason() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, reasonKey)
}

// SetReason sets the reason property value.
func (m *PolicyMappingModel) SetReason(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, reasonKey, val)
}

// GetState returns the state property value.
func (m *PolicyMappingModel) GetState() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, stateKey)
}

// SetState sets the state property value.
func (m *PolicyMappingModel) SetState(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, stateKey, val)
}

// GetSysClassName returns the sys_class_name property value.
func (m *PolicyMappingModel) GetSysClassName() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_class_nameKey)
}

// SetSysClassName sets the sys_class_name property value.
func (m *PolicyMappingModel) SetSysClassName(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_class_nameKey, val)
}

// GetSysCreatedBy returns the sys_created_by property value.
func (m *PolicyMappingModel) GetSysCreatedBy() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_created_byKey)
}

// SetSysCreatedBy sets the sys_created_by property value.
func (m *PolicyMappingModel) SetSysCreatedBy(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_created_byKey, val)
}

// GetSysCreatedOn returns the sys_created_on property value.
func (m *PolicyMappingModel) GetSysCreatedOn() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_created_onKey)
}

// SetSysCreatedOn sets the sys_created_on property value.
func (m *PolicyMappingModel) SetSysCreatedOn(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_created_onKey, val)
}

// GetSysId returns the sys_id property value.
func (m *PolicyMappingModel) GetSysId() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_idKey)
}

// SetSysId sets the sys_id property value.
func (m *PolicyMappingModel) SetSysId(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_idKey, val)
}

// GetSysUpdatedBy returns the sys_updated_by property value.
func (m *PolicyMappingModel) GetSysUpdatedBy() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_updated_byKey)
}

// SetSysUpdatedBy sets the sys_updated_by property value.
func (m *PolicyMappingModel) SetSysUpdatedBy(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_updated_byKey, val)
}

// GetSysUpdatedOn returns the sys_updated_on property value.
func (m *PolicyMappingModel) GetSysUpdatedOn() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_updated_onKey)
}

// SetSysUpdatedOn sets the sys_updated_on property value.
func (m *PolicyMappingModel) SetSysUpdatedOn(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_updated_onKey, val)
}
