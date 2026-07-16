package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// UploadStatusResult represents the status response of an upload.
type UploadStatusResult struct {
	core.BaseModel
}

// NewUploadStatusResult instantiates a new UploadStatusResult.
func NewUploadStatusResult() *UploadStatusResult {
	return &UploadStatusResult{BaseModel: *core.NewBaseModel()}
}

// Serialize serializes information the current object.
func (m *UploadStatusResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
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
	return store.DefaultBackedModelAccessorFunc[*UploadStatusResult, *string](m, typeKey)
}

// setType sets the type property value.
func (m *UploadStatusResult) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}

// GetState gets the state property value.
func (m *UploadStatusResult) GetState() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UploadStatusResult, *string](m, stateKey)
}

// setState sets the state property value.
func (m *UploadStatusResult) setState(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, stateKey, val)
}

// GetOutput gets the output property value.
func (m *UploadStatusResult) GetOutput() (*UploadStatusOutput, error) {
	return store.DefaultBackedModelAccessorFunc[*UploadStatusResult, *UploadStatusOutput](m, outputKey)
}

// setOutput sets the output property value.
func (m *UploadStatusResult) setOutput(val *UploadStatusOutput) error {
	return store.DefaultBackedModelMutatorFunc(m, outputKey, val)
}

// CreateUploadStatusResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value.
func CreateUploadStatusResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUploadStatusResult(), nil
}
