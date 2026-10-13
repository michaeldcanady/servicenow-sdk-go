// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComponentVarsUploadRequest represents the body for uploading component variables.
type ComponentVarsUploadRequest struct {
	core.BaseModel
}

// NewComponentVarsUploadRequest instantiates a new ComponentVarsUploadRequest.
func NewComponentVarsUploadRequest() *ComponentVarsUploadRequest {
	return &ComponentVarsUploadRequest{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *ComponentVarsUploadRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(appNameKey, m.GetAppName),
		internalSerialization.SerializeStringFunc(componentNameKey, m.GetComponentName),
		internalSerialization.SerializeAnyFunc(varsKey, m.GetVars),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *ComponentVarsUploadRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:       internalSerialization.DeserializeStringFunc(m.setAppName),
		componentNameKey: internalSerialization.DeserializeStringFunc(m.setComponentName),
		varsKey:          internalSerialization.DeserializeAnyFunc(m.setVars),
	}
}

// GetAppName returns the app name.
func (m *ComponentVarsUploadRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentVarsUploadRequest, *string](m, appNameKey)
}
func (m *ComponentVarsUploadRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appNameKey, val)
}

// GetComponentName returns the component name.
func (m *ComponentVarsUploadRequest) GetComponentName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentVarsUploadRequest, *string](m, componentNameKey)
}
func (m *ComponentVarsUploadRequest) setComponentName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, componentNameKey, val)
}

// GetVars returns the vars.
func (m *ComponentVarsUploadRequest) GetVars() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentVarsUploadRequest, any](m, varsKey)
}
func (m *ComponentVarsUploadRequest) setVars(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, varsKey, val)
}

// CreateComponentVarsUploadRequestFromDiscriminatorValue creates a new ComponentVarsUploadRequest from a ParseNode.
func CreateComponentVarsUploadRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewComponentVarsUploadRequest(), nil
}
