package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// UploadStatusOutput represents the output metadata from an upload.
type UploadStatusOutput struct {
	core.BaseModel
}

// NewUploadStatusOutput instantiates a new UploadStatusOutput.
func NewUploadStatusOutput() *UploadStatusOutput {
	return &UploadStatusOutput{BaseModel: *core.NewBaseModel()}
}

// Serialize serializes information the current object.
func (m *UploadStatusOutput) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(numberKey, m.GetNumber),
	)
}

// GetFieldDeserializers the deserialization information for the current model.
func (m *UploadStatusOutput) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:  internalSerialization.DeserializeStringFunc(m.setSysID),
		numberKey: internalSerialization.DeserializeStringFunc(m.setNumber),
	}
}

// GetSysID gets the sys_id property value.
func (m *UploadStatusOutput) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UploadStatusOutput, *string](m, sysIDKey)
}

// setSysID sets the sys_id property value.
func (m *UploadStatusOutput) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetNumber gets the number property value.
func (m *UploadStatusOutput) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*UploadStatusOutput, *string](m, numberKey)
}

// setNumber sets the number property value.
func (m *UploadStatusOutput) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}

// CreateUploadStatusOutputFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value.
func CreateUploadStatusOutputFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewUploadStatusOutput(), nil
}
