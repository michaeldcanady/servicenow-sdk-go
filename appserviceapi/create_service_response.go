package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreateServiceResponse represents the response containing the created application service details.
type CreateServiceResponse interface {
	core.ServiceNowItemResponse[*CreateServiceResult]
}

func CreateCreateServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*CreateServiceResult](CreateCreateServiceResultFromDiscriminatorValue), nil
}

// FindServiceResult represents the result details retrieved from the find_service endpoint.
type FindServiceResult struct {
	core.BaseModel
}

func NewFindServiceResult() *FindServiceResult {
	return &FindServiceResult{BaseModel: *core.NewBaseModel()}
}

func (m *FindServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(numberKey)(m.GetNumber),
		internalSerialization.SerializeStringFunc(environmentKey)(m.GetEnvironment),
		internalSerialization.SerializeStringFunc(versionKey)(m.GetVersion),
	)
}

func (m *FindServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:       internalSerialization.DeserializeStringFunc()(m.setSysId),
		nameKey:        internalSerialization.DeserializeStringFunc()(m.setName),
		numberKey:      internalSerialization.DeserializeStringFunc()(m.setNumber),
		environmentKey: internalSerialization.DeserializeStringFunc()(m.setEnvironment),
		versionKey:     internalSerialization.DeserializeStringFunc()(m.setVersion),
	}
}

func (m *FindServiceResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FindServiceResult, *string](m, sysIdKey)
}

func (m *FindServiceResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIdKey, val)
}

func (m *FindServiceResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FindServiceResult, *string](m, nameKey)
}

func (m *FindServiceResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

func (m *FindServiceResult) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FindServiceResult, *string](m, numberKey)
}

func (m *FindServiceResult) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}

func (m *FindServiceResult) GetEnvironment() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FindServiceResult, *string](m, environmentKey)
}

func (m *FindServiceResult) setEnvironment(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, environmentKey, val)
}

func (m *FindServiceResult) GetVersion() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FindServiceResult, *string](m, versionKey)
}

func (m *FindServiceResult) setVersion(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, versionKey, val)
}

func CreateFindServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFindServiceResult(), nil
}
