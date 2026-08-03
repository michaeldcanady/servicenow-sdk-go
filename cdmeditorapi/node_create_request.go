package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// NodeCreateRequest represents the body for creating a node.
type NodeCreateRequest interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetName() (*string, error)
	setName(*string) error
	GetType() (*string, error)
	setType(*string) error
	GetParentID() (*string, error)
	setParentID(*string) error
	GetCdmID() (*string, error)
	setCdmID(*string) error
}

// NodeCreateRequestModel represents a node create request model.
type NodeCreateRequestModel struct {
	core.BaseModel
}

// NewNodeCreateRequest instantiates a new NodeCreateRequest.
func NewNodeCreateRequest() *NodeCreateRequestModel {
	return &NodeCreateRequestModel{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *NodeCreateRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
		internalSerialization.SerializeStringFunc(parentIDKey, m.GetParentID),
		internalSerialization.SerializeStringFunc(cdmIDKey, m.GetCdmID),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *NodeCreateRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey:     internalSerialization.DeserializeStringFunc(m.setName),
		typeKey:     internalSerialization.DeserializeStringFunc(m.setType),
		parentIDKey: internalSerialization.DeserializeStringFunc(m.setParentID),
		cdmIDKey:    internalSerialization.DeserializeStringFunc(m.setCdmID),
	}
}

// GetName returns the name.
func (m *NodeCreateRequestModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeCreateRequestModel, *string](m, nameKey)
}
func (m *NodeCreateRequestModel) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetType returns the type.
func (m *NodeCreateRequestModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeCreateRequestModel, *string](m, typeKey)
}
func (m *NodeCreateRequestModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}

// GetParentID returns the parent id.
func (m *NodeCreateRequestModel) GetParentID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeCreateRequestModel, *string](m, parentIDKey)
}
func (m *NodeCreateRequestModel) setParentID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, parentIDKey, val)
}

// GetCdmID returns the cdm id.
func (m *NodeCreateRequestModel) GetCdmID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeCreateRequestModel, *string](m, cdmIDKey)
}
func (m *NodeCreateRequestModel) setCdmID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, cdmIDKey, val)
}

// CreateNodeCreateRequestFromDiscriminatorValue creates a new NodeCreateRequest from a ParseNode.
func CreateNodeCreateRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewNodeCreateRequest(), nil
}
