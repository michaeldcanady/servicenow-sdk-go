package statsapi

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	statsCountKey = "count"
	statsSumKey   = "sum"
	statsAvgKey   = "avg"
	statsMinKey   = "min"
	statsMaxKey   = "max"
)

// mapFromRaw converts a raw deserialized JSON object (map[string]any) into a map[string]string,
// which is the shape the Stats API uses for its per-field aggregate maps (sum/avg/min/max).
func mapFromRaw(rawValue any) (map[string]string, error) {
	if rawValue == nil {
		return nil, nil
	}

	typedValue, ok := rawValue.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("unsupported type %T", rawValue)
	}

	result := make(map[string]string, len(typedValue))
	for key, value := range typedValue {
		switch typedElem := value.(type) {
		case string:
			result[key] = typedElem
		case *string:
			if typedElem != nil {
				result[key] = *typedElem
			}
		default:
			result[key] = fmt.Sprintf("%v", value)
		}
	}

	return result, nil
}

// Stats represents the aggregate values returned under the "stats" key of a Stats API result.
type Stats interface {
	serialization.Parsable
	kiotaStore.BackedModel

	// GetCount returns the record count, present when sysparm_count is true.
	GetCount() (*string, error)
	setCount(*string) error
	// GetSum returns the requested sysparm_sum_fields, keyed by field name.
	GetSum() (map[string]string, error)
	setSum(map[string]string) error
	// GetAvg returns the requested sysparm_avg_fields, keyed by field name.
	GetAvg() (map[string]string, error)
	setAvg(map[string]string) error
	// GetMin returns the requested sysparm_min_fields, keyed by field name.
	GetMin() (map[string]string, error)
	setMin(map[string]string) error
	// GetMax returns the requested sysparm_max_fields, keyed by field name.
	GetMax() (map[string]string, error)
	setMax(map[string]string) error
}

// StatsModel is the default implementation of Stats.
type StatsModel struct {
	core.BackedModel
}

// NewStats creates a new instance of StatsModel.
func NewStats() *StatsModel {
	return &StatsModel{
		BackedModel: core.NewBaseModel(),
	}
}

// CreateStatsFromDiscriminatorValue creates a new Stats from a ParseNode.
func CreateStatsFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewStats(), nil
}

// GetFieldDeserializers implements serialization.Parsable.
func (m *StatsModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		statsCountKey: internalSerialization.DeserializeStringFunc()(m.setCount),
		statsSumKey:   internalSerialization.DeserializeMutatedAnyFunc(mapFromRaw)(m.setSum),
		statsAvgKey:   internalSerialization.DeserializeMutatedAnyFunc(mapFromRaw)(m.setAvg),
		statsMinKey:   internalSerialization.DeserializeMutatedAnyFunc(mapFromRaw)(m.setMin),
		statsMaxKey:   internalSerialization.DeserializeMutatedAnyFunc(mapFromRaw)(m.setMax),
	}
}

// Serialize implements serialization.Parsable.
func (m *StatsModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(statsCountKey)(m.GetCount),
		internalSerialization.SerializeAnyFunc(statsSumKey)(func() (any, error) { return m.GetSum() }),
		internalSerialization.SerializeAnyFunc(statsAvgKey)(func() (any, error) { return m.GetAvg() }),
		internalSerialization.SerializeAnyFunc(statsMinKey)(func() (any, error) { return m.GetMin() }),
		internalSerialization.SerializeAnyFunc(statsMaxKey)(func() (any, error) { return m.GetMax() }),
	)
}

// GetCount returns the record count.
func (m *StatsModel) GetCount() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*StatsModel, *string](m, statsCountKey)
}

func (m *StatsModel) setCount(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statsCountKey, val)
}

// GetSum returns the sum aggregates, keyed by field name.
func (m *StatsModel) GetSum() (map[string]string, error) {
	return store.DefaultBackedModelAccessorFunc[*StatsModel, map[string]string](m, statsSumKey)
}

func (m *StatsModel) setSum(val map[string]string) error {
	return store.DefaultBackedModelMutatorFunc(m, statsSumKey, val)
}

// GetAvg returns the average aggregates, keyed by field name.
func (m *StatsModel) GetAvg() (map[string]string, error) {
	return store.DefaultBackedModelAccessorFunc[*StatsModel, map[string]string](m, statsAvgKey)
}

func (m *StatsModel) setAvg(val map[string]string) error {
	return store.DefaultBackedModelMutatorFunc(m, statsAvgKey, val)
}

// GetMin returns the minimum aggregates, keyed by field name.
func (m *StatsModel) GetMin() (map[string]string, error) {
	return store.DefaultBackedModelAccessorFunc[*StatsModel, map[string]string](m, statsMinKey)
}

func (m *StatsModel) setMin(val map[string]string) error {
	return store.DefaultBackedModelMutatorFunc(m, statsMinKey, val)
}

// GetMax returns the maximum aggregates, keyed by field name.
func (m *StatsModel) GetMax() (map[string]string, error) {
	return store.DefaultBackedModelAccessorFunc[*StatsModel, map[string]string](m, statsMaxKey)
}

func (m *StatsModel) setMax(val map[string]string) error {
	return store.DefaultBackedModelMutatorFunc(m, statsMaxKey, val)
}
