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

// ServiceRelationship represents a service relationship between business and technical service offerings.
type ServiceRelationship struct {
	core.BaseModel
}

// NewServiceRelationship creates a new instance of [ServiceRelationship].
func NewServiceRelationship() *ServiceRelationship {
	return &ServiceRelationship{BaseModel: *core.NewBaseModel()}
}

// CreateServiceFromDiscriminatorValue creates a new Service from a ParseNode.
func CreateServiceRelationshipFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceRelationship(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *ServiceRelationship) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeCollectionOfStringValuesFunc(businessAppKey, m.GetBusinessApp),
		internalSerialization.SerializeCollectionOfStringValuesFunc(businessServiceOfferingKey, m.GetBusinessServiceOffering),
		internalSerialization.SerializeCollectionOfStringValuesFunc(technicalServiceOfferingKey, m.GetTechnicalServiceOffering),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ServiceRelationship) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		businessAppKey:              internalSerialization.DeserializeGetCollectionOfPrimitiveValuesFunc("string", m.SetBusinessApp),
		businessServiceOfferingKey:  internalSerialization.DeserializeGetCollectionOfPrimitiveValuesFunc("string", m.SetBusinessServiceOffering),
		technicalServiceOfferingKey: internalSerialization.DeserializeGetCollectionOfPrimitiveValuesFunc("string", m.SetTechnicalServiceOffering),
	}
}

// GetBusinessApp
func (m *ServiceRelationship) GetBusinessApp() ([]string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceRelationship, []string](m, businessAppKey)
}

// SetBusinessApp
func (m *ServiceRelationship) SetBusinessApp(vals []string) error {
	return store.DefaultBackedModelMutatorFunc(m, businessAppKey, vals)
}

// GetBusinessServiceOffering
func (m *ServiceRelationship) GetBusinessServiceOffering() ([]string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceRelationship, []string](m, businessServiceOfferingKey)
}

// SetBusinessServiceOffering
func (m *ServiceRelationship) SetBusinessServiceOffering(vals []string) error {
	return store.DefaultBackedModelMutatorFunc(m, businessServiceOfferingKey, vals)
}

// GetTechnicalServiceOffering
func (m *ServiceRelationship) GetTechnicalServiceOffering() ([]string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceRelationship, []string](m, technicalServiceOfferingKey)
}

// SetTechnicalServiceOffering
func (m *ServiceRelationship) SetTechnicalServiceOffering(vals []string) error {
	return store.DefaultBackedModelMutatorFunc(m, technicalServiceOfferingKey, vals)
}
