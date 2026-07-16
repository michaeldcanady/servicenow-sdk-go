package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangesetActivityResult represents a changeset activity.
type ChangesetActivityResult struct {
	core.BaseModel
}

func NewChangesetActivityResult() *ChangesetActivityResult {
	return &ChangesetActivityResult{BaseModel: *core.NewBaseModel()}
}

func (m *ChangesetActivityResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[*Reference](changesetIdKey)(m.GetChangesetId),
		internalSerialization.SerializeBoolFunc(conflictKey)(m.GetConflict),
		internalSerialization.SerializeStringFunc(namePathKey)(m.GetNamePath),
		internalSerialization.SerializeStringFunc(newNameKey)(m.GetNewName),
		internalSerialization.SerializeStringFunc(oldNameKey)(m.GetOldName),
		internalSerialization.SerializeStringFunc(newValueKey)(m.GetNewValue),
		internalSerialization.SerializeStringFunc(oldValueKey)(m.GetOldValue),
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
		internalSerialization.SerializeBoolFunc(secureKey)(m.GetSecure),
	)
}

func (m *ChangesetActivityResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		changesetIdKey: internalSerialization.DeserializeObjectValueFunc[*Reference](CreateReferenceFromDiscriminatorValue)(m.setChangesetId),
		conflictKey:    internalSerialization.DeserializeBoolFunc()(m.setConflict),
		namePathKey:    internalSerialization.DeserializeStringFunc()(m.setNamePath),
		newNameKey:     internalSerialization.DeserializeStringFunc()(m.setNewName),
		oldNameKey:     internalSerialization.DeserializeStringFunc()(m.setOldName),
		newValueKey:    internalSerialization.DeserializeStringFunc()(m.setNewValue),
		oldValueKey:    internalSerialization.DeserializeStringFunc()(m.setOldValue),
		typeKey:        internalSerialization.DeserializeStringFunc()(m.setType),
		secureKey:      internalSerialization.DeserializeBoolFunc()(m.setSecure),
	}
}

func (m *ChangesetActivityResult) GetChangesetId() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *Reference](m, changesetIdKey)
}
func (m *ChangesetActivityResult) setChangesetId(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, changesetIdKey, val)
}
func (m *ChangesetActivityResult) GetConflict() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *bool](m, conflictKey)
}
func (m *ChangesetActivityResult) setConflict(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, conflictKey, val)
}
func (m *ChangesetActivityResult) GetNamePath() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, namePathKey)
}
func (m *ChangesetActivityResult) setNamePath(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, namePathKey, val)
}
func (m *ChangesetActivityResult) GetNewName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, newNameKey)
}
func (m *ChangesetActivityResult) setNewName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, newNameKey, val)
}
func (m *ChangesetActivityResult) GetOldName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, oldNameKey)
}
func (m *ChangesetActivityResult) setOldName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, oldNameKey, val)
}
func (m *ChangesetActivityResult) GetNewValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, newValueKey)
}
func (m *ChangesetActivityResult) setNewValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, newValueKey, val)
}
func (m *ChangesetActivityResult) GetOldValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, oldValueKey)
}
func (m *ChangesetActivityResult) setOldValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, oldValueKey, val)
}
func (m *ChangesetActivityResult) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, typeKey)
}
func (m *ChangesetActivityResult) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}
func (m *ChangesetActivityResult) GetSecure() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *bool](m, secureKey)
}
func (m *ChangesetActivityResult) setSecure(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, secureKey, val)
}

func CreateChangesetActivityResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewChangesetActivityResult(), nil
}
