// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi // nolint:dupl // shares field-count shape with UserTimeFormatOptionsModel by coincidence, not copy-paste; distinct API concept, not worth sacrificing named accessors for

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// RPVariable represents the RPVariable nested object.
type RPVariable interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetDisplayName() (*string, error)
	SetDisplayName(*string) error
	GetLabel() (*string, error)
	SetLabel(*string) error
	GetName() (*string, error)
	SetName(*string) error
}

// RPVariableModel represents the rp variable model.
type RPVariableModel struct {
	core.BaseModel
}

// NewRPVariable creates a new instance of RPVariableModel.
func NewRPVariable() *RPVariableModel {
	return &RPVariableModel{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateRPVariableFromDiscriminatorValue creates a new RPVariable from a ParseNode.
func CreateRPVariableFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRPVariable(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *RPVariableModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(displayNameKey, m.GetDisplayName),
		internalSerialization.SerializeStringFunc(labelKey, m.GetLabel),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *RPVariableModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		displayNameKey: internalSerialization.DeserializeStringFunc(m.SetDisplayName),
		labelKey:       internalSerialization.DeserializeStringFunc(m.SetLabel),
		nameKey:        internalSerialization.DeserializeStringFunc(m.SetName),
	}
}

// GetDisplayName returns the display name value.
func (m *RPVariableModel) GetDisplayName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RPVariableModel, *string](m, displayNameKey)
}

// SetDisplayName sets the display name value.
func (m *RPVariableModel) SetDisplayName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, displayNameKey, val)
}

// GetLabel returns the label value.
func (m *RPVariableModel) GetLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RPVariableModel, *string](m, labelKey)
}

// SetLabel sets the label value.
func (m *RPVariableModel) SetLabel(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, labelKey, val)
}

// GetName returns the name value.
func (m *RPVariableModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RPVariableModel, *string](m, nameKey)
}

// SetName sets the name value.
func (m *RPVariableModel) SetName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}
