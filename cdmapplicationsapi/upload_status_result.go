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

var _ UploadStatusResult = (*UploadStatusResultModel)(nil)

// UploadStatusResult is the status payload returned for CDM upload operations.
type UploadStatusResult interface {
	serialization.Parsable

	GetType() (*string, error)
	GetState() (*string, error)
	GetOutput() (*UploadStatusOutput, error)
}

// UploadStatusResultModel represents the status response of an upload.
type UploadStatusResultModel struct {
	core.BaseModel
}

// NewUploadStatusResult instantiates a new UploadStatusResult.
func NewUploadStatusResult() *UploadStatusResultModel {
	return &UploadStatusResultModel{BaseModel: *core.NewBaseModel()}
}

// Serialize serializes information the current object.
func (m *UploadStatusResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
		internalSerialization.SerializeStringFunc(stateKey, m.GetState),
		internalSerialization.SerializeObjectValueFunc[*UploadStatusOutput](outputKey, m.GetOutput),
	)
}

// GetFieldDeserializers the deserialization information for the current model.
func (m *UploadStatusResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		typeKey:   internalSerialization.DeserializeStringFunc(m.setType),
		stateKey:  internalSerialization.DeserializeStringFunc(m.setState),
		outputKey: internalSerialization.DeserializeObjectValueFunc[*UploadStatusOutput](CreateUploadStatusOutputFromDiscriminatorValue, m.setOutput),
	}
}

// GetType gets the type property value.
func (m *UploadStatusResultModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UploadStatusResultModel, *string](m, typeKey)
}

// setType sets the type property value.
func (m *UploadStatusResultModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}

// GetState gets the state property value.
func (m *UploadStatusResultModel) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UploadStatusResultModel, *string](m, stateKey)
}

// setState sets the state property value.
func (m *UploadStatusResultModel) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}

// GetOutput gets the output property value.
func (m *UploadStatusResultModel) GetOutput() (*UploadStatusOutput, error) {
	return store.DefaultBackedModelAccessorFunc[*UploadStatusResultModel, *UploadStatusOutput](m, outputKey)
}

// setOutput sets the output property value.
func (m *UploadStatusResultModel) setOutput(val *UploadStatusOutput) error {
	return store.DefaultBackedModelMutatorFunc(m, outputKey, val)
}

// CreateUploadStatusResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value.
func CreateUploadStatusResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUploadStatusResult(), nil
}
