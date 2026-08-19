package appserviceapi

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// TODO: make type enum
// cmdb_group_based
// discovery
// tag_list

var _ PopulationMethod = (*PopulationMethodModel)(nil)

// PopulationMethod represents the base interface for all population methods used to discover or populate application service relationships.
type PopulationMethod interface {
	core.Model
	GetType() (*string, error)
}

// PopulationMethodModel is the backing-store-backed implementation of [PopulationMethod].
type PopulationMethodModel struct {
	core.BaseModel
}

// NewPopulationMethod creates a new instance of [PopulationMethodModel].
func NewPopulationMethod() *PopulationMethodModel {
	return &PopulationMethodModel{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreatePopulationMethodFromDiscriminatorValue creates a new [PopulationMethod] from a ParseNode, dispatching to the correct concrete type based on the type field.
func CreatePopulationMethodFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	typeNode, err := parseNode.GetChildNode(typeKey)
	if err != nil {
		return nil, err
	}

	typeValue, err := typeNode.GetStringValue()
	if err != nil {
		return nil, err
	}

	if typeValue == nil {
		return nil, errors.New("unknown type")
	}

	switch *typeValue {
	case "cmdb_group_based":
		return CreateCmdbGroupBasedPopulationMethodFromDiscriminatorValue(parseNode)
	case "discovery":
		return CreateDiscoveryPopulationMethodFromDiscriminatorValue(parseNode)
	case "tag_list":
		return CreateTagListPopulationMethodFromDiscriminatorValue(parseNode)
	default:
		return nil, errors.New("unknown type")
	}
}

// Serialize writes the objects properties to the current writer.
func (m *PopulationMethodModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
	)
}

// GetFieldDeserializers
func (m *PopulationMethodModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		typeKey: internalSerialization.DeserializeStringFunc(m.SetType),
	}
}

// GetType
func (m *PopulationMethodModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PopulationMethodModel, *string](m, typeKey)
}

// SetType
func (m *PopulationMethodModel) SetType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}
