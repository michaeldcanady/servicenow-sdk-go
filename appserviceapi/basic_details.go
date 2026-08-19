package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// BasicDetails represents the basic details schema inside CSDM requests.
type BasicDetails struct {
	core.BaseModel
}

// NewBasicDetails creates a new instance of BasicDetails.
func NewBasicDetails() *BasicDetails {
	return &BasicDetails{BaseModel: *core.NewBaseModel()}
}

// CreateBasicDetailsFromDiscriminatorValue creates a new BasicDetails from a ParseNode.
func CreateBasicDetailsFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewBasicDetails(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *BasicDetails) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(environmentKey, m.GetEnvironment),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(versionKey, m.GetVersion),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *BasicDetails) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		environmentKey: internalSerialization.DeserializeStringFunc(m.SetEnvironment),
		nameKey:        internalSerialization.DeserializeStringFunc(m.SetName),
		versionKey:     internalSerialization.DeserializeStringFunc(m.SetVersion),
	}
}

// GetEnvironment returns the environment value.
func (m *BasicDetails) GetEnvironment() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*BasicDetails, *string](m, environmentKey)
}

// SetEnvironment
func (m *BasicDetails) SetEnvironment(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, environmentKey, val)
}

// TODO: required
// GetName returns the name value.
func (m *BasicDetails) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*BasicDetails, *string](m, nameKey)
}

// SetName
func (m *BasicDetails) SetName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetVersion returns the version value.
func (m *BasicDetails) GetVersion() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*BasicDetails, *string](m, versionKey)
}

// SetVersion
func (m *BasicDetails) SetVersion(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, versionKey, val)
}
