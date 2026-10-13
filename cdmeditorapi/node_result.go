// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// NodeResult represents a node in the configuration tree.
type NodeResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetSysID() (*string, error)
	setSysID(*string) error
	GetName() (*string, error)
	setName(*string) error
	GetType() (*string, error)
	setType(*string) error
	GetValue() (*string, error)
	setValue(*string) error
	GetParent() (*string, error)
	setParent(*string) error
	GetCdmID() (*string, error)
	setCdmID(*string) error
}

// NodeResultModel represents a node result model.
type NodeResultModel struct {
	core.BaseModel
}

// NewNodeResult instantiates a new NodeResult.
func NewNodeResult() *NodeResultModel {
	return &NodeResultModel{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *NodeResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
		internalSerialization.SerializeStringFunc(parentKey, m.GetParent),
		internalSerialization.SerializeStringFunc(cdmIDKey, m.GetCdmID),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *NodeResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:  internalSerialization.DeserializeStringFunc(m.setSysID),
		nameKey:   internalSerialization.DeserializeStringFunc(m.setName),
		typeKey:   internalSerialization.DeserializeStringFunc(m.setType),
		valueKey:  internalSerialization.DeserializeStringFunc(m.setValue),
		parentKey: internalSerialization.DeserializeStringFunc(m.setParent),
		cdmIDKey:  internalSerialization.DeserializeStringFunc(m.setCdmID),
	}
}

// GetSysID returns the sys id.
func (m *NodeResultModel) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, sysIDKey)
}
func (m *NodeResultModel) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetName returns the name.
func (m *NodeResultModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, nameKey)
}
func (m *NodeResultModel) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetType returns the type.
func (m *NodeResultModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, typeKey)
}
func (m *NodeResultModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}

// GetValue returns the value.
func (m *NodeResultModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, valueKey)
}
func (m *NodeResultModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}

// GetParent returns the parent.
func (m *NodeResultModel) GetParent() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, parentKey)
}
func (m *NodeResultModel) setParent(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, parentKey, val)
}

// GetCdmID returns the cdm id.
func (m *NodeResultModel) GetCdmID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, cdmIDKey)
}
func (m *NodeResultModel) setCdmID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, cdmIDKey, val)
}

// CreateNodeResultFromDiscriminatorValue creates a new NodeResult from a ParseNode.
func CreateNodeResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewNodeResult(), nil
}
