package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComponentVarsUploadRequest represents the body for uploading component variables.
type ComponentVarsUploadRequest struct {
	core.BaseModel
}

func NewComponentVarsUploadRequest() *ComponentVarsUploadRequest {
	return &ComponentVarsUploadRequest{BaseModel: *core.NewBaseModel()}
}

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

func (m *ComponentVarsUploadRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:       internalSerialization.DeserializeStringFunc(m.setAppName),
		componentNameKey: internalSerialization.DeserializeStringFunc(m.setComponentName),
		varsKey:          internalSerialization.DeserializeAnyFunc(m.setVars),
	}
}

func (m *ComponentVarsUploadRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentVarsUploadRequest, *string](m, appNameKey)
}
func (m *ComponentVarsUploadRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appNameKey, val)
}
func (m *ComponentVarsUploadRequest) GetComponentName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentVarsUploadRequest, *string](m, componentNameKey)
}
func (m *ComponentVarsUploadRequest) setComponentName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, componentNameKey, val)
}
func (m *ComponentVarsUploadRequest) GetVars() (any, error) {
	return store.DefaultBackedModelAccessorFunc[*ComponentVarsUploadRequest, any](m, varsKey)
}
func (m *ComponentVarsUploadRequest) setVars(val any) error {
	return store.DefaultBackedModelMutatorFunc(m, varsKey, val)
}

func CreateComponentVarsUploadRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewComponentVarsUploadRequest(), nil
}
