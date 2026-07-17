package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedComponentUpdateRequest represents the body for updating shared components.
type SharedComponentUpdateRequest struct {
	core.BaseModel
}

func NewSharedComponentUpdateRequest() *SharedComponentUpdateRequest {
	return &SharedComponentUpdateRequest{BaseModel: *core.NewBaseModel()}
}

func (m *SharedComponentUpdateRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(appNameKey, m.GetAppName),
		internalSerialization.SerializeStringFunc(sharedComponentNameKey, m.GetSharedComponentName),
		internalSerialization.SerializeStringFunc(dataKey, m.GetData),
	)
}

func (m *SharedComponentUpdateRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:             internalSerialization.DeserializeStringFunc(m.setAppName),
		sharedComponentNameKey: internalSerialization.DeserializeStringFunc(m.setSharedComponentName),
		dataKey:                internalSerialization.DeserializeStringFunc(m.setData),
	}
}

func (m *SharedComponentUpdateRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*SharedComponentUpdateRequest, *string](m, appNameKey)
}
func (m *SharedComponentUpdateRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appNameKey, val)
}
func (m *SharedComponentUpdateRequest) GetSharedComponentName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*SharedComponentUpdateRequest, *string](m, sharedComponentNameKey)
}
func (m *SharedComponentUpdateRequest) setSharedComponentName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sharedComponentNameKey, val)
}
func (m *SharedComponentUpdateRequest) GetData() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*SharedComponentUpdateRequest, *string](m, dataKey)
}
func (m *SharedComponentUpdateRequest) setData(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dataKey, val)
}

func CreateSharedComponentUpdateRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewSharedComponentUpdateRequest(), nil
}
