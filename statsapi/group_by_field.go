package statsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	groupByFieldFieldKey        = "field"
	groupByFieldValueKey        = "value"
	groupByFieldDisplayValueKey = "display_value"
)

// GroupByField represents a single grouping dimension of a grouped Stats API result,
// e.g. {"field": "priority", "value": "1"}.
type GroupByField interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetField() (*string, error)
	setField(*string) error
	GetValue() (*string, error)
	setValue(*string) error
	// GetDisplayValue is only populated when the request set sysparm_display_value to "all".
	GetDisplayValue() (*string, error)
	setDisplayValue(*string) error
}

// GroupByFieldModel is the default implementation of GroupByField.
type GroupByFieldModel struct {
	core.BackedModel
}

// NewGroupByField creates a new instance of GroupByFieldModel.
func NewGroupByField() *GroupByFieldModel {
	return &GroupByFieldModel{
		BackedModel: core.NewBaseModel(),
	}
}

// CreateGroupByFieldFromDiscriminatorValue creates a new GroupByField from a ParseNode.
func CreateGroupByFieldFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewGroupByField(), nil
}

// GetFieldDeserializers implements serialization.Parsable.
func (m *GroupByFieldModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		groupByFieldFieldKey:        internalSerialization.DeserializeStringFunc()(m.setField),
		groupByFieldValueKey:        internalSerialization.DeserializeStringFunc()(m.setValue),
		groupByFieldDisplayValueKey: internalSerialization.DeserializeStringFunc()(m.setDisplayValue),
	}
}

// Serialize implements serialization.Parsable.
func (m *GroupByFieldModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(groupByFieldFieldKey)(m.GetField),
		internalSerialization.SerializeStringFunc(groupByFieldValueKey)(m.GetValue),
		internalSerialization.SerializeStringFunc(groupByFieldDisplayValueKey)(m.GetDisplayValue),
	)
}

// GetField returns the name of the field this grouping is on.
func (m *GroupByFieldModel) GetField() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*GroupByFieldModel, *string](m, groupByFieldFieldKey)
}

func (m *GroupByFieldModel) setField(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, groupByFieldFieldKey, val)
}

// GetValue returns the database value of the group.
func (m *GroupByFieldModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*GroupByFieldModel, *string](m, groupByFieldValueKey)
}

func (m *GroupByFieldModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, groupByFieldValueKey, val)
}

// GetDisplayValue returns the display value of the group, when requested via sysparm_display_value=all.
func (m *GroupByFieldModel) GetDisplayValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*GroupByFieldModel, *string](m, groupByFieldDisplayValueKey)
}

func (m *GroupByFieldModel) setDisplayValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, groupByFieldDisplayValueKey, val)
}
