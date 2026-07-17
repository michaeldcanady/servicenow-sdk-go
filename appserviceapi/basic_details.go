package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// BasicDetails represents the basic details schema inside CSDM requests.
type BasicDetails struct {
	core.BaseModel
}

func NewBasicDetails() *BasicDetails {
	return &BasicDetails{BaseModel: *core.NewBaseModel()}
}

func (m *BasicDetails) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(environmentKey, m.GetEnvironment),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(versionKey, m.GetVersion),
		internalSerialization.SerializeStringFunc(businessAppKey, m.GetBusinessApp),
		internalSerialization.SerializeStringFunc(businessServiceOfferingKey, m.GetBusinessServiceOffering),
		internalSerialization.SerializeStringFunc(technicalServiceOfferingKey, m.GetTechnicalServiceOffering),
	)
}

func (m *BasicDetails) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		environmentKey:              internalSerialization.DeserializeStringFunc(m.setEnvironment),
		nameKey:                     internalSerialization.DeserializeStringFunc(m.setName),
		versionKey:                  internalSerialization.DeserializeStringFunc(m.setVersion),
		businessAppKey:              internalSerialization.DeserializeStringFunc(m.setBusinessApp),
		businessServiceOfferingKey:  internalSerialization.DeserializeStringFunc(m.setBusinessServiceOffering),
		technicalServiceOfferingKey: internalSerialization.DeserializeStringFunc(m.setTechnicalServiceOffering),
	}
}

func (m *BasicDetails) GetEnvironment() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*BasicDetails, *string](m, environmentKey)
}

func (m *BasicDetails) setEnvironment(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, environmentKey, val)
}

func (m *BasicDetails) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*BasicDetails, *string](m, nameKey)
}

func (m *BasicDetails) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

func (m *BasicDetails) GetVersion() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*BasicDetails, *string](m, versionKey)
}

func (m *BasicDetails) setVersion(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, versionKey, val)
}

func (m *BasicDetails) GetBusinessApp() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*BasicDetails, *string](m, businessAppKey)
}

func (m *BasicDetails) setBusinessApp(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, businessAppKey, val)
}

func (m *BasicDetails) GetBusinessServiceOffering() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*BasicDetails, *string](m, businessServiceOfferingKey)
}

func (m *BasicDetails) setBusinessServiceOffering(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, businessServiceOfferingKey, val)
}

func (m *BasicDetails) GetTechnicalServiceOffering() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*BasicDetails, *string](m, technicalServiceOfferingKey)
}

func (m *BasicDetails) setTechnicalServiceOffering(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, technicalServiceOfferingKey, val)
}

func CreateBasicDetailsFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewBasicDetails(), nil
}
