// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangesetActivityResult represents a changeset activity.
type ChangesetActivityResult struct {
	core.BaseModel
}

// NewChangesetActivityResult instantiates a new ChangesetActivityResult.
func NewChangesetActivityResult() *ChangesetActivityResult {
	return &ChangesetActivityResult{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *ChangesetActivityResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[*Reference](changesetIDKey, m.GetChangesetID),
		internalSerialization.SerializeBoolFunc(conflictKey, m.GetConflict),
		internalSerialization.SerializeStringFunc(namePathKey, m.GetNamePath),
		internalSerialization.SerializeStringFunc(newNameKey, m.GetNewName),
		internalSerialization.SerializeStringFunc(oldNameKey, m.GetOldName),
		internalSerialization.SerializeStringFunc(newValueKey, m.GetNewValue),
		internalSerialization.SerializeStringFunc(oldValueKey, m.GetOldValue),
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
		internalSerialization.SerializeBoolFunc(secureKey, m.GetSecure),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *ChangesetActivityResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		changesetIDKey: internalSerialization.DeserializeObjectValueFunc[*Reference](CreateReferenceFromDiscriminatorValue, m.setChangesetID),
		conflictKey:    internalSerialization.DeserializeBoolFunc(m.setConflict),
		namePathKey:    internalSerialization.DeserializeStringFunc(m.setNamePath),
		newNameKey:     internalSerialization.DeserializeStringFunc(m.setNewName),
		oldNameKey:     internalSerialization.DeserializeStringFunc(m.setOldName),
		newValueKey:    internalSerialization.DeserializeStringFunc(m.setNewValue),
		oldValueKey:    internalSerialization.DeserializeStringFunc(m.setOldValue),
		typeKey:        internalSerialization.DeserializeStringFunc(m.setType),
		secureKey:      internalSerialization.DeserializeBoolFunc(m.setSecure),
	}
}

// GetChangesetID returns the changeset id.
func (m *ChangesetActivityResult) GetChangesetID() (*Reference, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *Reference](m, changesetIDKey)
}
func (m *ChangesetActivityResult) setChangesetID(val *Reference) error {
	return store.DefaultBackedModelMutatorFunc(m, changesetIDKey, val)
}

// GetConflict returns the conflict.
func (m *ChangesetActivityResult) GetConflict() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *bool](m, conflictKey)
}
func (m *ChangesetActivityResult) setConflict(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, conflictKey, val)
}

// GetNamePath returns the name path.
func (m *ChangesetActivityResult) GetNamePath() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, namePathKey)
}
func (m *ChangesetActivityResult) setNamePath(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, namePathKey, val)
}

// GetNewName returns the new name.
func (m *ChangesetActivityResult) GetNewName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, newNameKey)
}
func (m *ChangesetActivityResult) setNewName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, newNameKey, val)
}

// GetOldName returns the old name.
func (m *ChangesetActivityResult) GetOldName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, oldNameKey)
}
func (m *ChangesetActivityResult) setOldName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, oldNameKey, val)
}

// GetNewValue returns the new value.
func (m *ChangesetActivityResult) GetNewValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, newValueKey)
}
func (m *ChangesetActivityResult) setNewValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, newValueKey, val)
}

// GetOldValue returns the old value.
func (m *ChangesetActivityResult) GetOldValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, oldValueKey)
}
func (m *ChangesetActivityResult) setOldValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, oldValueKey, val)
}

// GetType returns the type.
func (m *ChangesetActivityResult) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *string](m, typeKey)
}
func (m *ChangesetActivityResult) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}

// GetSecure returns the secure.
func (m *ChangesetActivityResult) GetSecure() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ChangesetActivityResult, *bool](m, secureKey)
}
func (m *ChangesetActivityResult) setSecure(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, secureKey, val)
}

// CreateChangesetActivityResultFromDiscriminatorValue creates a new ChangesetActivityResult from a ParseNode.
func CreateChangesetActivityResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewChangesetActivityResult(), nil
}
