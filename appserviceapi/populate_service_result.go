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

// PopulateServiceResult represents the result details of populating a service.
type PopulateServiceResult struct {
	core.BaseModel
}

// NewPopulateServiceResult creates a new instance of [PopulateServiceResult].
func NewPopulateServiceResult() *PopulateServiceResult {
	return &PopulateServiceResult{BaseModel: *core.NewBaseModel()}
}

// CreatePopulateServiceResultFromDiscriminatorValue creates a new [PopulateServiceResult] from a [serialization.ParseNode].
func CreatePopulateServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPopulateServiceResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *PopulateServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(statusKey, m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey, m.GetMessage),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *PopulateServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		statusKey:  internalSerialization.DeserializeStringFunc(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc(m.setMessage),
	}
}

// GetStatus returns the status value.
func (m *PopulateServiceResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PopulateServiceResult, *string](m, statusKey)
}

func (m *PopulateServiceResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}

// GetMessage returns the message value.
func (m *PopulateServiceResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PopulateServiceResult, *string](m, messageKey)
}

func (m *PopulateServiceResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, val)
}
