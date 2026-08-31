// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	canReadKey  = "can_read"
	canWriteKey = "can_write"
	colorKey    = "color"
)

var _ ActivitiesResultField = (*ActivitiesResultFieldModel)(nil)

// ActivitiesResultField represents a field metadata entry in a case activity result.
type ActivitiesResultField interface {
	serialization.Parsable
	kiotaStore.BackedModel

	// GetCanRead returns the flag that indicates whether the logged-in user can read this journal field or not.
	GetCanRead() (*bool, error)
	// SetCanRead sets the flag that indicates whether the logged-in user can read this journal field or not.
	SetCanRead(*bool) error

	// GetCanWrite returns the flag that indicates whether the logged-in user can write to the journal field or not.
	GetCanWrite() (*bool, error)
	// SetCanWrite sets the flag that indicates whether the logged-in user can write to the journal field or not.
	SetCanWrite(*bool) error

	// GetColor returns the color that represents the journal field in case activity stream on ServiceNow AI Platform.
	GetColor() (*string, error)
	// SetColor sets the color that represents the journal field in case activity stream on ServiceNow AI Platform.
	SetColor(*string) error

	// GetLabel returns the display name for the journal field.
	GetLabel() (*string, error)
	// SetLabel sets the display name for the journal field.
	SetLabel(*string) error

	// GetName returns the name for the journal field.
	GetName() (*string, error)
	// SetName sets the name for the journal field.
	SetName(*string) error
}

// ActivitiesResultFieldModel is the backing-store-backed implementation of [ActivitiesResultField].
type ActivitiesResultFieldModel struct {
	core.BaseModel
}

// NewActivitiesResultField creates a new instance of [ActivitiesResultFieldModel].
func NewActivitiesResultField() *ActivitiesResultFieldModel {
	return &ActivitiesResultFieldModel{BaseModel: *core.NewBaseModel()}
}

// CreateActivitiesResultFieldFromDiscriminatorValue creates a new instance of [ActivitiesResultField].
func CreateActivitiesResultFieldFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitiesResultField(), nil
}

// Serialize implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(a) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeBoolFunc(canReadKey, a.GetCanRead),
		internalSerialization.SerializeBoolFunc(canWriteKey, a.GetCanWrite),
		internalSerialization.SerializeStringFunc(colorKey, a.GetColor),
		internalSerialization.SerializeStringFunc(labelKey, a.GetLabel),
		internalSerialization.SerializeStringFunc(nameKey, a.GetName),
	)
}

// GetFieldDeserializers implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		canReadKey:  internalSerialization.DeserializeBoolFunc(a.SetCanRead),
		canWriteKey: internalSerialization.DeserializeBoolFunc(a.SetCanWrite),
		colorKey:    internalSerialization.DeserializeStringFunc(a.SetColor),
		labelKey:    internalSerialization.DeserializeStringFunc(a.SetLabel),
		nameKey:     internalSerialization.DeserializeStringFunc(a.SetName),
	}
}

// GetCanRead implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) GetCanRead() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultFieldModel, *bool](a, canReadKey)
}

// SetCanRead implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) SetCanRead(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(a, canReadKey, val)
}

// GetCanWrite implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) GetCanWrite() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultFieldModel, *bool](a, canWriteKey)
}

// SetCanWrite implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) SetCanWrite(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(a, canWriteKey, val)
}

// GetColor implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) GetColor() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultFieldModel, *string](a, colorKey)
}

// SetColor implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) SetColor(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, colorKey, val)
}

// GetLabel implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) GetLabel() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultFieldModel, *string](a, labelKey)
}

// SetLabel implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) SetLabel(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, labelKey, val)
}

// GetName implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultFieldModel, *string](a, nameKey)
}

// SetName implements [ActivitiesResultField].
func (a *ActivitiesResultFieldModel) SetName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(a, nameKey, val)
}
