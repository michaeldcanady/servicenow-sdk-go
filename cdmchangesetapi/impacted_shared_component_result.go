// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmchangesetapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImpactedSharedComponentResult represents an impacted shared component.
type ImpactedSharedComponentResult struct {
	core.BaseModel
}

// NewImpactedSharedComponentResult instantiates a new ImpactedSharedComponentResult.
func NewImpactedSharedComponentResult() *ImpactedSharedComponentResult {
	return &ImpactedSharedComponentResult{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *ImpactedSharedComponentResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(cdmSharedLibraryKey, m.GetCdmSharedLibrary),
		internalSerialization.SerializeStringFunc(descriptionKey, m.GetDescription),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(nodeKey, m.GetNode),
		internalSerialization.SerializeStringFunc(nodeMainKey, m.GetNodeMain),
		internalSerialization.SerializeStringFunc(stateKey, m.GetState),
		internalSerialization.SerializeStringFunc(sysCreatedByKey, m.GetSysCreatedBy),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey, m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(sysUpdatedByKey, m.GetSysUpdatedBy),
		internalSerialization.SerializeStringFunc(sysUpdatedOnKey, m.GetSysUpdatedOn),
		internalSerialization.SerializeInt32Func(versionCounterKey, m.GetVersionCounter),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *ImpactedSharedComponentResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		cdmSharedLibraryKey: internalSerialization.DeserializeStringFunc(m.setCdmSharedLibrary),
		descriptionKey:      internalSerialization.DeserializeStringFunc(m.setDescription),
		nameKey:             internalSerialization.DeserializeStringFunc(m.setName),
		nodeKey:             internalSerialization.DeserializeStringFunc(m.setNode),
		nodeMainKey:         internalSerialization.DeserializeStringFunc(m.setNodeMain),
		stateKey:            internalSerialization.DeserializeStringFunc(m.setState),
		sysCreatedByKey:     internalSerialization.DeserializeStringFunc(m.setSysCreatedBy),
		sysCreatedOnKey:     internalSerialization.DeserializeStringFunc(m.setSysCreatedOn),
		sysIDKey:            internalSerialization.DeserializeStringFunc(m.setSysID),
		sysUpdatedByKey:     internalSerialization.DeserializeStringFunc(m.setSysUpdatedBy),
		sysUpdatedOnKey:     internalSerialization.DeserializeStringFunc(m.setSysUpdatedOn),
		versionCounterKey:   internalSerialization.DeserializeInt32Func(m.setVersionCounter),
	}
}

// GetCdmSharedLibrary returns the cdm shared library.
func (m *ImpactedSharedComponentResult) GetCdmSharedLibrary() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, cdmSharedLibraryKey)
}
func (m *ImpactedSharedComponentResult) setCdmSharedLibrary(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, cdmSharedLibraryKey, val)
}

// GetDescription returns the description.
func (m *ImpactedSharedComponentResult) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, descriptionKey)
}
func (m *ImpactedSharedComponentResult) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, descriptionKey, val)
}

// GetName returns the name.
func (m *ImpactedSharedComponentResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, nameKey)
}
func (m *ImpactedSharedComponentResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetNode returns the node.
func (m *ImpactedSharedComponentResult) GetNode() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, nodeKey)
}
func (m *ImpactedSharedComponentResult) setNode(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nodeKey, val)
}

// GetNodeMain returns the node main.
func (m *ImpactedSharedComponentResult) GetNodeMain() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, nodeMainKey)
}
func (m *ImpactedSharedComponentResult) setNodeMain(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nodeMainKey, val)
}

// GetState returns the state.
func (m *ImpactedSharedComponentResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, stateKey)
}
func (m *ImpactedSharedComponentResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}

// GetSysCreatedBy returns the sys created by.
func (m *ImpactedSharedComponentResult) GetSysCreatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, sysCreatedByKey)
}
func (m *ImpactedSharedComponentResult) setSysCreatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedByKey, val)
}

// GetSysCreatedOn returns the sys created on.
func (m *ImpactedSharedComponentResult) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, sysCreatedOnKey)
}
func (m *ImpactedSharedComponentResult) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnKey, val)
}

// GetSysID returns the sys id.
func (m *ImpactedSharedComponentResult) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, sysIDKey)
}
func (m *ImpactedSharedComponentResult) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetSysUpdatedBy returns the sys updated by.
func (m *ImpactedSharedComponentResult) GetSysUpdatedBy() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, sysUpdatedByKey)
}
func (m *ImpactedSharedComponentResult) setSysUpdatedBy(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysUpdatedByKey, val)
}

// GetSysUpdatedOn returns the sys updated on.
func (m *ImpactedSharedComponentResult) GetSysUpdatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *string](m, sysUpdatedOnKey)
}
func (m *ImpactedSharedComponentResult) setSysUpdatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysUpdatedOnKey, val)
}

// GetVersionCounter returns the version counter.
func (m *ImpactedSharedComponentResult) GetVersionCounter() (*int32, error) {
	return store.DefaultBackedModelAccessorFunc[*ImpactedSharedComponentResult, *int32](m, versionCounterKey)
}
func (m *ImpactedSharedComponentResult) setVersionCounter(val *int32) error {
	return store.DefaultBackedModelMutatorFunc(m, versionCounterKey, val)
}

// CreateImpactedSharedComponentResultFromDiscriminatorValue creates a new ImpactedSharedComponentResult from a ParseNode.
func CreateImpactedSharedComponentResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewImpactedSharedComponentResult(), nil
}
