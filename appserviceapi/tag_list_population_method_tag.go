package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	tagKey = "tag"
)

var _ TagListPopulationMethodTag = (*TagListPopulationMethodTagModel)(nil)

// TagListPopulationMethodTag represents a tag key-value pair used by a tag-list population method to discover service relationships.
type TagListPopulationMethodTag interface {
	core.Model
	GetTag() (*string, error)
	SetTag(*string) error
	GetValue() (*string, error)
	SetValue(*string) error
}

// TagListPopulationMethodTagModel is the backing-store-backed implementation of [TagListPopulationMethodTag].
type TagListPopulationMethodTagModel struct {
	*core.BaseModel
}

// NewTagListPopulationMethodTag
func NewTagListPopulationMethodTag() *TagListPopulationMethodTagModel {
	return &TagListPopulationMethodTagModel{
		BaseModel: core.NewBaseModel(),
	}
}

// CreateTagListPopulationMethodTagFromDiscriminatorValue
func CreateTagListPopulationMethodTagFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewTagListPopulationMethodTag(), nil
}

// GetFieldDeserializers
func (t *TagListPopulationMethodTagModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		tagKey:   internalSerialization.DeserializeStringFunc(t.SetTag),
		valueKey: internalSerialization.DeserializeStringFunc(t.SetValue),
	}
}

// Serialize
func (t *TagListPopulationMethodTagModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(t) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(tagKey, t.GetTag),
		internalSerialization.SerializeStringFunc(valueKey, t.GetValue),
	)
}

// GetTag
func (t *TagListPopulationMethodTagModel) GetTag() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*TagListPopulationMethodTagModel, *string](t, tagKey)
}

// SetTag
func (t *TagListPopulationMethodTagModel) SetTag(val *string) error {
	return store.DefaultBackedModelMutatorFunc(t, tagKey, val)
}

// GetValue
func (t *TagListPopulationMethodTagModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*TagListPopulationMethodTagModel, *string](t, valueKey)
}

// SetValue
func (t *TagListPopulationMethodTagModel) SetValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(t, valueKey, val)
}
