// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// ExecuteRuleConditionsRequest represents the execute rule conditions request.
type ExecuteRuleConditionsRequest interface {
	serialization.Parsable
	kiotaStore.BackedModel

	// TODO: required
	GetCatalogID() (*string, error)
	SetCatalogID(*string) error
	// TODO: I don't understand,  required if taskID isn't specified
	GetOtherInputs() (any, error)
	SetOtherInputs(any) error
	// TODO: required if OtherInputs isn't specified
	GetTaskID() (*string, error)
	SetTaskID(*string) error
}

// ExecuteRuleConditionsRequestModel represents the execute rule conditions request model.
type ExecuteRuleConditionsRequestModel struct {
	core.BaseModel
}

// NewExecuteRuleConditionsRequest creates a new instance of ExecuteRuleConditionsRequestModel.
func NewExecuteRuleConditionsRequest() *ExecuteRuleConditionsRequestModel {
	return &ExecuteRuleConditionsRequestModel{BaseModel: *core.NewBaseModel()}
}

// CreateExecuteRuleConditionsRequestFromDiscriminatorValue creates a new ExecuteRuleConditionsRequest from a ParseNode.
func CreateExecuteRuleConditionsRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewExecuteRuleConditionsRequest(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *ExecuteRuleConditionsRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(catalogIDKey, m.GetCatalogID),
		internalSerialization.SerializeAnyFunc(otherInputsKey, m.GetOtherInputs),
		internalSerialization.SerializeStringFunc(taskIDKey, m.GetTaskID),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ExecuteRuleConditionsRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		catalogIDKey:   internalSerialization.DeserializeStringFunc(m.SetCatalogID),
		otherInputsKey: internalSerialization.DeserializeAnyFunc(m.SetOtherInputs),
		taskIDKey:      internalSerialization.DeserializeStringFunc(m.SetTaskID),
	}
}

// GetCatalogID returns the catalog id value.
func (m *ExecuteRuleConditionsRequestModel) GetCatalogID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsRequestModel, *string](m, catalogIDKey)
}

// SetCatalogID sets the catalog id value.
func (m *ExecuteRuleConditionsRequestModel) SetCatalogID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, catalogIDKey, val)
}

// GetOtherInputs returns the other inputs value.
func (m *ExecuteRuleConditionsRequestModel) GetOtherInputs() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsRequestModel, any](m, otherInputsKey)
}

// SetOtherInputs sets the other inputs value.
func (m *ExecuteRuleConditionsRequestModel) SetOtherInputs(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, otherInputsKey, val)
}

// GetTaskID returns the task id value.
func (m *ExecuteRuleConditionsRequestModel) GetTaskID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ExecuteRuleConditionsRequestModel, *string](m, taskIDKey)
}

// SetTaskID sets the task id value.
func (m *ExecuteRuleConditionsRequestModel) SetTaskID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, taskIDKey, val)
}
