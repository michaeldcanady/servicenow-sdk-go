package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// UploadStatusOutput represents the output metadata from an upload.
type UploadStatusOutput struct {
	newInternal.BaseModel
}

// NewUploadStatusOutput instantiates a new UploadStatusOutput.
func NewUploadStatusOutput() *UploadStatusOutput {
	return &UploadStatusOutput{BaseModel: *newInternal.NewBaseModel()}
}

// Serialize serializes information the current object.
func (m *UploadStatusOutput) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(numberKey)(m.GetNumber),
	)
}

// GetFieldDeserializers the deserialization information for the current model.
func (m *UploadStatusOutput) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:  internalSerialization.DeserializeStringFunc()(m.setSysId),
		numberKey: internalSerialization.DeserializeStringFunc()(m.setNumber),
	}
}

// GetSysId gets the sys_id property value.
func (m *UploadStatusOutput) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}

// setSysId sets the sys_id property value.
func (m *UploadStatusOutput) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}

// GetNumber gets the number property value.
func (m *UploadStatusOutput) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), numberKey)
}

// setNumber sets the number property value.
func (m *UploadStatusOutput) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), numberKey, val)
}

// CreateUploadStatusOutputFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value.
func CreateUploadStatusOutputFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUploadStatusOutput(), nil
}

// UploadStatusResult represents the status response of an upload.
type UploadStatusResult struct {
	newInternal.BaseModel
}

// NewUploadStatusResult instantiates a new UploadStatusResult.
func NewUploadStatusResult() *UploadStatusResult {
	return &UploadStatusResult{BaseModel: *newInternal.NewBaseModel()}
}

// Serialize serializes information the current object.
func (m *UploadStatusResult) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
		internalSerialization.SerializeStringFunc(stateKey)(m.GetState),
		internalSerialization.SerializeObjectValueFunc[*UploadStatusOutput](outputKey)(m.GetOutput),
	)
}

// GetFieldDeserializers the deserialization information for the current model.
func (m *UploadStatusResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		typeKey:   internalSerialization.DeserializeStringFunc()(m.setType),
		stateKey:  internalSerialization.DeserializeStringFunc()(m.setState),
		outputKey: internalSerialization.DeserializeObjectValueFunc[*UploadStatusOutput](CreateUploadStatusOutputFromDiscriminatorValue)(m.setOutput),
	}
}

// GetType gets the type property value.
func (m *UploadStatusResult) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
}

// setType sets the type property value.
func (m *UploadStatusResult) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, val)
}

// GetState gets the state property value.
func (m *UploadStatusResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), stateKey)
}

// setState sets the state property value.
func (m *UploadStatusResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), stateKey, val)
}

// GetOutput gets the output property value.
func (m *UploadStatusResult) GetOutput() (*UploadStatusOutput, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *UploadStatusOutput](m.GetBackingStore(), outputKey)
}

// setOutput sets the output property value.
func (m *UploadStatusResult) setOutput(val *UploadStatusOutput) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), outputKey, val)
}

// CreateUploadStatusResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value.
func CreateUploadStatusResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUploadStatusResult(), nil
}
