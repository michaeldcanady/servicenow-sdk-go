package policyapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	appNameKey        = "appName"
	deployableNameKey = "deployableName"
)

// AdditionalDeployable defines the interface for the AdditionalDeployable model.
type AdditionalDeployable interface {

	// GetAppName gets the appName property value.
	GetAppName() (*string, error)
	// SetAppName sets the appName property value.
	SetAppName(*string) error

	// GetDeployableName gets the deployableName property value.
	GetDeployableName() (*string, error)
	// SetDeployableName sets the deployableName property value.
	SetDeployableName(*string) error

	serialization.Parsable
	kiotaStore.BackedModel
}

// AdditionalDeployableModel is the concrete implementation of the AdditionalDeployable interface.
type AdditionalDeployableModel struct {
	newInternal.Model
}

// NewAdditionalDeployable creates a new instance of AdditionalDeployableModel with a backing store.
func NewAdditionalDeployable() *AdditionalDeployableModel {
	return &AdditionalDeployableModel{
		newInternal.NewBaseModel(),
	}
}

// CreateAdditionalDeployableFromDiscriminatorValue is a factory function for creating AdditionalDeployableModel instances during deserialization.
func CreateAdditionalDeployableFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAdditionalDeployable(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *AdditionalDeployableModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,

		internalSerialization.SerializeStringFunc(appNameKey)(m.GetAppName),

		internalSerialization.SerializeStringFunc(deployableNameKey)(m.GetDeployableName),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *AdditionalDeployableModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{

		appNameKey: internalSerialization.DeserializeStringFunc()(m.SetAppName),

		deployableNameKey: internalSerialization.DeserializeStringFunc()(m.SetDeployableName),
	}
}

// GetAppName returns the appName property value.
func (m *AdditionalDeployableModel) GetAppName() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, appNameKey)
}

// SetAppName sets the appName property value.
func (m *AdditionalDeployableModel) SetAppName(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, appNameKey, val)
}

// GetDeployableName returns the deployableName property value.
func (m *AdditionalDeployableModel) GetDeployableName() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, deployableNameKey)
}

// SetDeployableName sets the deployableName property value.
func (m *AdditionalDeployableModel) SetDeployableName(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, deployableNameKey, val)
}
