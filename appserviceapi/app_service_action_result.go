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

// AppServiceActionRequest is the request body shared by the app_service action endpoints
// (ConvertToDynamicService, ConvertToManualService, CreateDynamicService,
// UpdateDynamicNumberOfLevels) - the spec defines no schema for these bodies.
type AppServiceActionRequest struct {
	core.BaseModel
}

// NewAppServiceActionRequest creates a new instance of AppServiceActionRequest.
func NewAppServiceActionRequest() *AppServiceActionRequest {
	return &AppServiceActionRequest{BaseModel: *core.NewBaseModel()}
}

// CreateAppServiceActionRequestFromDiscriminatorValue creates a new AppServiceActionRequest from a ParseNode.
func CreateAppServiceActionRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAppServiceActionRequest(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *AppServiceActionRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(commentsKey, m.GetComments),
		internalSerialization.SerializeStringFunc(numberOfLevelsKey, m.GetNumberOfLevels),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *AppServiceActionRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:          internalSerialization.DeserializeStringFunc(m.setSysID),
		nameKey:           internalSerialization.DeserializeStringFunc(m.setName),
		commentsKey:       internalSerialization.DeserializeStringFunc(m.setComments),
		numberOfLevelsKey: internalSerialization.DeserializeStringFunc(m.setNumberOfLevels),
	}
}

// GetSysID returns the sys id value.
func (m *AppServiceActionRequest) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppServiceActionRequest, *string](m, sysIDKey)
}

func (m *AppServiceActionRequest) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// SetSysID sets the sys id value.
func (m *AppServiceActionRequest) SetSysID(val *string) error {
	return m.setSysID(val)
}

// GetName returns the name value.
func (m *AppServiceActionRequest) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppServiceActionRequest, *string](m, nameKey)
}

func (m *AppServiceActionRequest) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// SetName sets the name value.
func (m *AppServiceActionRequest) SetName(val *string) error {
	return m.setName(val)
}

// GetComments returns the comments value.
func (m *AppServiceActionRequest) GetComments() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppServiceActionRequest, *string](m, commentsKey)
}

func (m *AppServiceActionRequest) setComments(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, commentsKey, val)
}

// SetComments sets the comments value.
func (m *AppServiceActionRequest) SetComments(val *string) error {
	return m.setComments(val)
}

// GetNumberOfLevels returns the number of levels value.
func (m *AppServiceActionRequest) GetNumberOfLevels() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppServiceActionRequest, *string](m, numberOfLevelsKey)
}

func (m *AppServiceActionRequest) setNumberOfLevels(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberOfLevelsKey, val)
}

// SetNumberOfLevels sets the number of levels value.
func (m *AppServiceActionRequest) SetNumberOfLevels(val *string) error {
	return m.setNumberOfLevels(val)
}

// AppServiceActionResult represents the result details returned by an app_service action endpoint.
type AppServiceActionResult struct {
	core.BaseModel
}

// NewAppServiceActionResult creates a new instance of AppServiceActionResult.
func NewAppServiceActionResult() *AppServiceActionResult {
	return &AppServiceActionResult{BaseModel: *core.NewBaseModel()}
}

// CreateAppServiceActionResultFromDiscriminatorValue creates a new AppServiceActionResult from a ParseNode.
func CreateAppServiceActionResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAppServiceActionResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *AppServiceActionResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(commentsKey, m.GetComments),
		internalSerialization.SerializeStringFunc(numberOfLevelsKey, m.GetNumberOfLevels),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *AppServiceActionResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:          internalSerialization.DeserializeStringFunc(m.setSysID),
		nameKey:           internalSerialization.DeserializeStringFunc(m.setName),
		commentsKey:       internalSerialization.DeserializeStringFunc(m.setComments),
		numberOfLevelsKey: internalSerialization.DeserializeStringFunc(m.setNumberOfLevels),
	}
}

// GetSysID returns the sys id value.
func (m *AppServiceActionResult) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppServiceActionResult, *string](m, sysIDKey)
}

func (m *AppServiceActionResult) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetName returns the name value.
func (m *AppServiceActionResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppServiceActionResult, *string](m, nameKey)
}

func (m *AppServiceActionResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetComments returns the comments value.
func (m *AppServiceActionResult) GetComments() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppServiceActionResult, *string](m, commentsKey)
}

func (m *AppServiceActionResult) setComments(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, commentsKey, val)
}

// GetNumberOfLevels returns the number of levels value.
func (m *AppServiceActionResult) GetNumberOfLevels() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppServiceActionResult, *string](m, numberOfLevelsKey)
}

func (m *AppServiceActionResult) setNumberOfLevels(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberOfLevelsKey, val)
}
