package policyapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	io_definitionKey  = "io_definition"
	is_unusedKey      = "is_unused"
	mapped_valueKey   = "mapped_value"
	policy_mappingKey = "policy_mapping"
	sys_created_byKey = "sys_created_by"
	sys_created_onKey = "sys_created_on"
	sys_idKey         = "sys_id"
	sys_updated_byKey = "sys_updated_by"
	sys_updated_onKey = "sys_updated_on"
)

// MappingInputVariable defines the interface for the MappingInputVariable model.
type MappingInputVariable interface {

	// GetIoDefinition gets the io_definition property value.
	GetIoDefinition() (LinkRef, error)
	// SetIoDefinition sets the io_definition property value.
	SetIoDefinition(LinkRef) error

	// GetIsUnused gets the is_unused property value.
	GetIsUnused() (*bool, error)
	// SetIsUnused sets the is_unused property value.
	SetIsUnused(*bool) error

	// GetMappedValue gets the mapped_value property value.
	GetMappedValue() (*string, error)
	// SetMappedValue sets the mapped_value property value.
	SetMappedValue(*string) error

	// GetPolicyMapping gets the policy_mapping property value.
	GetPolicyMapping() (LinkRef, error)
	// SetPolicyMapping sets the policy_mapping property value.
	SetPolicyMapping(LinkRef) error

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

// MappingInputVariableModel is the concrete implementation of the MappingInputVariable interface.
type MappingInputVariableModel struct {
	newInternal.Model
}

// NewMappingInputVariable creates a new instance of MappingInputVariableModel with a backing store.
func NewMappingInputVariable() *MappingInputVariableModel {
	return &MappingInputVariableModel{
		newInternal.NewBaseModel(),
	}
}

// CreateMappingInputVariableFromDiscriminatorValue is a factory function for creating MappingInputVariableModel instances during deserialization.
func CreateMappingInputVariableFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewMappingInputVariable(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *MappingInputVariableModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,

		internalSerialization.SerializeObjectValueFunc[LinkRef](io_definitionKey)(m.GetIoDefinition),

		internalSerialization.SerializeStringToBoolFunc(is_unusedKey)(m.GetIsUnused),

		internalSerialization.SerializeStringFunc(mapped_valueKey)(m.GetMappedValue),

		internalSerialization.SerializeObjectValueFunc[LinkRef](policy_mappingKey)(m.GetPolicyMapping),

		internalSerialization.SerializeStringFunc(sys_created_byKey)(m.GetSysCreatedBy),

		internalSerialization.SerializeStringFunc(sys_created_onKey)(m.GetSysCreatedOn),

		internalSerialization.SerializeStringFunc(sys_idKey)(m.GetSysId),

		internalSerialization.SerializeStringFunc(sys_updated_byKey)(m.GetSysUpdatedBy),

		internalSerialization.SerializeStringFunc(sys_updated_onKey)(m.GetSysUpdatedOn),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *MappingInputVariableModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{

		io_definitionKey: internalSerialization.DeserializeObjectValueFunc[LinkRef](CreateLinkRefFromDiscriminatorValue)(m.SetIoDefinition),

		is_unusedKey: internalSerialization.DeserializeMutatedStringFunc(conversion.StringPtrToBoolPtr)(m.SetIsUnused),

		mapped_valueKey: internalSerialization.DeserializeStringFunc()(m.SetMappedValue),

		policy_mappingKey: internalSerialization.DeserializeObjectValueFunc[LinkRef](CreateLinkRefFromDiscriminatorValue)(m.SetPolicyMapping),

		sys_created_byKey: internalSerialization.DeserializeStringFunc()(m.SetSysCreatedBy),

		sys_created_onKey: internalSerialization.DeserializeStringFunc()(m.SetSysCreatedOn),

		sys_idKey: internalSerialization.DeserializeStringFunc()(m.SetSysId),

		sys_updated_byKey: internalSerialization.DeserializeStringFunc()(m.SetSysUpdatedBy),

		sys_updated_onKey: internalSerialization.DeserializeStringFunc()(m.SetSysUpdatedOn),
	}
}

// GetIoDefinition returns the io_definition property value.
func (m *MappingInputVariableModel) GetIoDefinition() (LinkRef, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, LinkRef](backingStore, io_definitionKey)
}

// SetIoDefinition sets the io_definition property value.
func (m *MappingInputVariableModel) SetIoDefinition(val LinkRef) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, io_definitionKey, val)
}

// GetIsUnused returns the is_unused property value.
func (m *MappingInputVariableModel) GetIsUnused() (*bool, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](backingStore, is_unusedKey)
}

// SetIsUnused sets the is_unused property value.
func (m *MappingInputVariableModel) SetIsUnused(val *bool) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, is_unusedKey, val)
}

// GetMappedValue returns the mapped_value property value.
func (m *MappingInputVariableModel) GetMappedValue() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, mapped_valueKey)
}

// SetMappedValue sets the mapped_value property value.
func (m *MappingInputVariableModel) SetMappedValue(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, mapped_valueKey, val)
}

// GetPolicyMapping returns the policy_mapping property value.
func (m *MappingInputVariableModel) GetPolicyMapping() (LinkRef, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, LinkRef](backingStore, policy_mappingKey)
}

// SetPolicyMapping sets the policy_mapping property value.
func (m *MappingInputVariableModel) SetPolicyMapping(val LinkRef) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, policy_mappingKey, val)
}

// GetSysCreatedBy returns the sys_created_by property value.
func (m *MappingInputVariableModel) GetSysCreatedBy() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_created_byKey)
}

// SetSysCreatedBy sets the sys_created_by property value.
func (m *MappingInputVariableModel) SetSysCreatedBy(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_created_byKey, val)
}

// GetSysCreatedOn returns the sys_created_on property value.
func (m *MappingInputVariableModel) GetSysCreatedOn() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_created_onKey)
}

// SetSysCreatedOn sets the sys_created_on property value.
func (m *MappingInputVariableModel) SetSysCreatedOn(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_created_onKey, val)
}

// GetSysId returns the sys_id property value.
func (m *MappingInputVariableModel) GetSysId() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_idKey)
}

// SetSysId sets the sys_id property value.
func (m *MappingInputVariableModel) SetSysId(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_idKey, val)
}

// GetSysUpdatedBy returns the sys_updated_by property value.
func (m *MappingInputVariableModel) GetSysUpdatedBy() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_updated_byKey)
}

// SetSysUpdatedBy sets the sys_updated_by property value.
func (m *MappingInputVariableModel) SetSysUpdatedBy(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_updated_byKey, val)
}

// GetSysUpdatedOn returns the sys_updated_on property value.
func (m *MappingInputVariableModel) GetSysUpdatedOn() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, sys_updated_onKey)
}

// SetSysUpdatedOn sets the sys_updated_on property value.
func (m *MappingInputVariableModel) SetSysUpdatedOn(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, sys_updated_onKey, val)
}
