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

// NodeUpdateRequest represents the body for updating a node.
type NodeUpdateRequest interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetName() (*string, error)
	setName(*string) error
	GetValue() (*string, error)
	setValue(*string) error
}

// NodeUpdateRequestModel represents a node update request model.
type NodeUpdateRequestModel struct {
	core.BaseModel
}

// NewNodeUpdateRequest instantiates a new NodeUpdateRequest.
func NewNodeUpdateRequest() *NodeUpdateRequestModel {
	return &NodeUpdateRequestModel{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *NodeUpdateRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *NodeUpdateRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey:  internalSerialization.DeserializeStringFunc(m.setName),
		valueKey: internalSerialization.DeserializeStringFunc(m.setValue),
	}
}

// GetName returns the name.
func (m *NodeUpdateRequestModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeUpdateRequestModel, *string](m, nameKey)
}
func (m *NodeUpdateRequestModel) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetValue returns the value.
func (m *NodeUpdateRequestModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeUpdateRequestModel, *string](m, valueKey)
}
func (m *NodeUpdateRequestModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}

// CreateNodeUpdateRequestFromDiscriminatorValue creates a new NodeUpdateRequest from a ParseNode.
func CreateNodeUpdateRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewNodeUpdateRequest(), nil
}
