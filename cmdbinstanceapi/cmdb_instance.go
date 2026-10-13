// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cmdbinstanceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	sysIDKey      = "sys_id"
	nameKey       = "name"
	classNameKey  = "className"
	attributesKey = "attributes"
)

// CmdbInstance represents a ServiceNow CMDB record.
//
// Collection GET responses place fields at the top level. Item GET responses nest
// the same fields under "attributes" (alongside relation arrays). Deserializing
// "attributes" promotes sys_id/name onto the top-level accessors when those
// top-level values are unset, so GetSysID()/GetName() work for both shapes.
type CmdbInstance interface {
	GetSysID() (*string, error)
	SetSysID(val *string) error
	GetName() (*string, error)
	SetName(val *string) error
	GetClassName() (*string, error)
	SetClassName(val *string) error
	GetAttributes() (*CmdbInstanceModel, error)

	serialization.Parsable
	kiotaStore.BackedModel
}

// CmdbInstanceModel implementation of CmdbInstance
type CmdbInstanceModel struct {
	core.BaseModel
}

// NewCmdbInstance creates a new instance of CmdbInstanceModel
func NewCmdbInstance() *CmdbInstanceModel {
	return &CmdbInstanceModel{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateCmdbInstanceFromDiscriminatorValue creates a new instance of CmdbInstance.
func CreateCmdbInstanceFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCmdbInstance(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *CmdbInstanceModel) Serialize(writer serialization.SerializationWriter) error {
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(classNameKey, m.GetClassName),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *CmdbInstanceModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:      internalSerialization.DeserializeStringFunc(m.SetSysID),
		nameKey:       internalSerialization.DeserializeStringFunc(m.SetName),
		classNameKey:  internalSerialization.DeserializeStringFunc(m.SetClassName),
		attributesKey: internalSerialization.DeserializeObjectValueFunc[*CmdbInstanceModel](CreateCmdbInstanceFromDiscriminatorValue, m.setAttributes),
	}
}

// setAttributes stores the nested attributes object and promotes sys_id/name when
// the top-level counterparts were not already set (collection vs item payload).
func (m *CmdbInstanceModel) setAttributes(val *CmdbInstanceModel) error {
	if err := store.DefaultBackedModelMutatorFunc(m, attributesKey, val); err != nil {
		return err
	}
	if conversion.IsNil(val) {
		return nil
	}
	if err := promoteStringIfUnset(m, sysIDKey, val.GetSysID); err != nil {
		return err
	}
	return promoteStringIfUnset(m, nameKey, val.GetName)
}

func promoteStringIfUnset(m *CmdbInstanceModel, key string, getter func() (*string, error)) error {
	current, err := store.DefaultBackedModelAccessorFunc[*CmdbInstanceModel, *string](m, key)
	if err != nil {
		return err
	}
	if current != nil {
		return nil
	}
	val, err := getter()
	if err != nil {
		return err
	}
	if val == nil {
		return nil
	}
	return store.DefaultBackedModelMutatorFunc(m, key, val)
}

// GetSysID ...
func (m *CmdbInstanceModel) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CmdbInstanceModel, *string](m, sysIDKey)
}

// SetSysID ...
func (m *CmdbInstanceModel) SetSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetName ...
func (m *CmdbInstanceModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CmdbInstanceModel, *string](m, nameKey)
}

// SetName ...
func (m *CmdbInstanceModel) SetName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetClassName ...
func (m *CmdbInstanceModel) GetClassName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CmdbInstanceModel, *string](m, classNameKey)
}

// SetClassName ...
func (m *CmdbInstanceModel) SetClassName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, classNameKey, val)
}

// GetAttributes returns the nested attributes object from an item GET payload.
func (m *CmdbInstanceModel) GetAttributes() (*CmdbInstanceModel, error) {
	return store.DefaultBackedModelAccessorFunc[*CmdbInstanceModel, *CmdbInstanceModel](m, attributesKey)
}
