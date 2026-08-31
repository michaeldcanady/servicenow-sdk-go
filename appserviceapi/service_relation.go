// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceRelation represents a relationship between components inside Populate request.
type ServiceRelation struct {
	core.BaseModel
}

// NewServiceRelation creates a new instance of ServiceRelation.
func NewServiceRelation() *ServiceRelation {
	return &ServiceRelation{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *ServiceRelation) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(parentKey, m.GetParent),
		internalSerialization.SerializeStringFunc(childKey, m.GetChild),
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ServiceRelation) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		parentKey: internalSerialization.DeserializeStringFunc(m.setParent),
		childKey:  internalSerialization.DeserializeStringFunc(m.setChild),
		typeKey:   internalSerialization.DeserializeStringFunc(m.setType),
	}
}

// GetParent returns the parent value.
func (m *ServiceRelation) GetParent() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceRelation, *string](m, parentKey)
}

func (m *ServiceRelation) setParent(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, parentKey, val)
}

// GetChild returns the child value.
func (m *ServiceRelation) GetChild() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceRelation, *string](m, childKey)
}

func (m *ServiceRelation) setChild(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, childKey, val)
}

// GetType returns the type value.
func (m *ServiceRelation) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceRelation, *string](m, typeKey)
}

func (m *ServiceRelation) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}

// CreateServiceRelationFromDiscriminatorValue creates a new ServiceRelation from a ParseNode.
func CreateServiceRelationFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceRelation(), nil
}
