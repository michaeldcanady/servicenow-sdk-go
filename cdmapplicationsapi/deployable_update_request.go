package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeployableUpdateRequest represents the body for updating deployables.
type DeployableUpdateRequest struct {
	core.BaseModel
}

func NewDeployableUpdateRequest() *DeployableUpdateRequest {
	return &DeployableUpdateRequest{BaseModel: *core.NewBaseModel()}
}

func (m *DeployableUpdateRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(appNameKey)(m.GetAppName),
		internalSerialization.SerializeStringFunc(deployableNameKey)(m.GetDeployableName),
		internalSerialization.SerializeStringFunc(dataKey)(m.GetData),
	)
}

func (m *DeployableUpdateRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		appNameKey:        internalSerialization.DeserializeStringFunc()(m.setAppName),
		deployableNameKey: internalSerialization.DeserializeStringFunc()(m.setDeployableName),
		dataKey:           internalSerialization.DeserializeStringFunc()(m.setData),
	}
}

func (m *DeployableUpdateRequest) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*DeployableUpdateRequest, *string](m, appNameKey)
}
func (m *DeployableUpdateRequest) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appNameKey, val)
}
func (m *DeployableUpdateRequest) GetDeployableName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*DeployableUpdateRequest, *string](m, deployableNameKey)
}
func (m *DeployableUpdateRequest) setDeployableName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, deployableNameKey, val)
}
func (m *DeployableUpdateRequest) GetData() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*DeployableUpdateRequest, *string](m, dataKey)
}
func (m *DeployableUpdateRequest) setData(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dataKey, val)
}

func CreateDeployableUpdateRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewDeployableUpdateRequest(), nil
}
