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

// PopulateServiceRequest represents the request body for populating a CSDM service.
type PopulateServiceRequest struct {
	core.BaseModel
}

// NewPopulateServiceRequest creates a new instance of [PopulateServiceRequest].
func NewPopulateServiceRequest() *PopulateServiceRequest {
	return &PopulateServiceRequest{BaseModel: *core.NewBaseModel()}
}

// CreatePopulateServiceRequestFromDiscriminatorValue creates a new [PopulateServiceRequest] from a [serialization.ParseNode].
func CreatePopulateServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPopulateServiceRequest(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *PopulateServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[PopulationMethod](populationMethodKey, m.GetPopulationMethod),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *PopulateServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		populationMethodKey: internalSerialization.DeserializeObjectValueFunc[PopulationMethod](CreatePopulationMethodFromDiscriminatorValue, m.SetPopulationMethod),
	}
}

// GetServiceRelations returns the service relations value.
func (m *PopulateServiceRequest) GetPopulationMethod() (PopulationMethod, error) {
	return store.DefaultBackedModelAccessorFunc[*PopulateServiceRequest, PopulationMethod](m, populationMethodKey)
}

// SetPopulationMethod
func (m *PopulateServiceRequest) SetPopulationMethod(val PopulationMethod) error {
	return store.DefaultBackedModelMutatorFunc(m, populationMethodKey, val)
}
