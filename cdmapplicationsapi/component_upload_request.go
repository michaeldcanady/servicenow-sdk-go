package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComponentUploadRequest represents the body for uploading components.
type ComponentUploadRequest struct {
	core.BaseModel
}

func NewComponentUploadRequest() *ComponentUploadRequest {
	return &ComponentUploadRequest{BaseModel: *core.NewBaseModel()}
}

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

func (m *ComponentUploadRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:       internalSerialization.DeserializeStringFunc(m.setAppName),
		componentNameKey: internalSerialization.DeserializeStringFunc(m.setComponentName),
		dataKey:          internalSerialization.DeserializeStringFunc(m.setData),
		formatKey:        internalSerialization.DeserializeStringFunc(m.setFormat),
	}
}

func (m *ComponentUploadRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentUploadRequest, *string](m, appNameKey)
}
func (m *ComponentUploadRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appNameKey, val)
}
func (m *ComponentUploadRequest) GetComponentName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentUploadRequest, *string](m, componentNameKey)
}
func (m *ComponentUploadRequest) setComponentName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, componentNameKey, val)
}
func (m *ComponentUploadRequest) GetData() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentUploadRequest, *string](m, dataKey)
}
func (m *ComponentUploadRequest) setData(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dataKey, val)
}
func (m *ComponentUploadRequest) GetFormat() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentUploadRequest, *string](m, formatKey)
}
func (m *ComponentUploadRequest) setFormat(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, formatKey, val)
}

func CreateComponentUploadRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewComponentUploadRequest(), nil
}
