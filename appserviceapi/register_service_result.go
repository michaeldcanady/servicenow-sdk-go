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

// RegisterServiceResult represents the result details of a registered CSDM service.
type RegisterServiceResult struct {
	core.BaseModel
}

// NewRegisterServiceResult creates a new instance of [RegisterServiceResult].
func NewRegisterServiceResult() *RegisterServiceResult {
	return &RegisterServiceResult{BaseModel: *core.NewBaseModel()}
}

// CreateRegisterServiceResultFromDiscriminatorValue creates a new [RegisterServiceResult] from a [serialization.ParseNode].
func CreateRegisterServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRegisterServiceResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *RegisterServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(numberKey, m.GetNumber),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *RegisterServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:  internalSerialization.DeserializeStringFunc(m.SetSysID),
		numberKey: internalSerialization.DeserializeStringFunc(m.SetNumber),
	}
}

// GetSysID returns the sys id value.
func (m *RegisterServiceResult) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RegisterServiceResult, *string](m, sysIDKey)
}

func (m *RegisterServiceResult) SetSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetNumber returns the number value.
func (m *RegisterServiceResult) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RegisterServiceResult, *string](m, numberKey)
}

func (m *RegisterServiceResult) SetNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}
