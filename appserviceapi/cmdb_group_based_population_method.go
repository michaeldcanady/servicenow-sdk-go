package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	groupIDKey = "group_id"
)

var _ PopulationMethod = (*CmdbGroupBasedPopulationMethodModel)(nil)
var _ CmdbGroupBasedPopulationMethod = (*CmdbGroupBasedPopulationMethodModel)(nil)

type CmdbGroupBasedPopulationMethod interface {
	PopulationMethod
	GetGroupID() (*string, error)
	SetGroupID(*string) error
}

type CmdbGroupBasedPopulationMethodModel struct {
	PopulationMethod
}

func NewCmdbGroupBasedPopulationMethod() *CmdbGroupBasedPopulationMethodModel {
	return &CmdbGroupBasedPopulationMethodModel{
		PopulationMethod: NewPopulationMethod(),
	}
}

func CreateCmdbGroupBasedPopulationMethodFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCmdbGroupBasedPopulationMethod(), nil
}

func extendmap[K comparable, V any](m1 map[K]V, ms ...map[K]V) map[K]V {
	for _, m := range ms {
		for key, value := range m {
			m1[key] = value
		}
	}

	return m1
}

// GetFieldDeserializers
func (m *CmdbGroupBasedPopulationMethodModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return extendmap(map[string]func(serialization.ParseNode) error{
		groupIDKey: internalSerialization.DeserializeStringFunc(m.SetGroupID),
	}, m.PopulationMethod.GetFieldDeserializers())
}

// Serialize writes the objects properties to the current writer.
func (m *CmdbGroupBasedPopulationMethodModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}

	if err := m.PopulationMethod.Serialize(writer); err != nil {
		return err
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(groupIDKey, m.GetGroupID),
	)
}

// GetGroupID
func (m *CmdbGroupBasedPopulationMethodModel) GetGroupID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CmdbGroupBasedPopulationMethodModel, *string](m, groupIDKey)
}

// SetGroupID
func (m *CmdbGroupBasedPopulationMethodModel) SetGroupID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, groupIDKey, val)
}
