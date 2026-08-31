// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi // nolint:dupl // shares field-count shape with CollectionUploadRequest/ExportStatusResult by coincidence, not copy-paste; distinct API concept, not worth sacrificing named accessors for

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComponentUploadRequest represents the body for uploading components.
type ComponentUploadRequest struct {
	core.BaseModel
}

// NewComponentUploadRequest instantiates a new ComponentUploadRequest.
func NewComponentUploadRequest() *ComponentUploadRequest {
	return &ComponentUploadRequest{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *ComponentUploadRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(appNameKey, m.GetAppName),
		internalSerialization.SerializeStringFunc(componentNameKey, m.GetComponentName),
		internalSerialization.SerializeStringFunc(dataKey, m.GetData),
		internalSerialization.SerializeStringFunc(formatKey, m.GetFormat),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *ComponentUploadRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:       internalSerialization.DeserializeStringFunc(m.setAppName),
		componentNameKey: internalSerialization.DeserializeStringFunc(m.setComponentName),
		dataKey:          internalSerialization.DeserializeStringFunc(m.setData),
		formatKey:        internalSerialization.DeserializeStringFunc(m.setFormat),
	}
}

// GetAppName returns the app name.
func (m *ComponentUploadRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentUploadRequest, *string](m, appNameKey)
}
func (m *ComponentUploadRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appNameKey, val)
}

// GetComponentName returns the component name.
func (m *ComponentUploadRequest) GetComponentName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentUploadRequest, *string](m, componentNameKey)
}
func (m *ComponentUploadRequest) setComponentName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, componentNameKey, val)
}

// GetData returns the data.
func (m *ComponentUploadRequest) GetData() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentUploadRequest, *string](m, dataKey)
}
func (m *ComponentUploadRequest) setData(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dataKey, val)
}

// GetFormat returns the format.
func (m *ComponentUploadRequest) GetFormat() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentUploadRequest, *string](m, formatKey)
}
func (m *ComponentUploadRequest) setFormat(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, formatKey, val)
}

// CreateComponentUploadRequestFromDiscriminatorValue creates a new ComponentUploadRequest from a ParseNode.
func CreateComponentUploadRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewComponentUploadRequest(), nil
}
