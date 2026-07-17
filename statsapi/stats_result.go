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
	statsResultStatsKey         = "stats"
	statsResultGroupByFieldsKey = "groupby_fields"
)

// StatsResult represents a single element of a Stats API ("/api/now/stats/{table}") response.
//
// GroupbyFields is only populated when the request set sysparm_group_by; this SDK models the
// ungrouped shape (a single StatsResult under "result") — grouped responses return an array of
// these under "result" instead, which is not yet supported by StatsRequestBuilder.Get.
type StatsResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetStats() (Stats, error)
	setStats(Stats) error
	GetGroupbyFields() ([]GroupByField, error)
	setGroupbyFields([]GroupByField) error
}

// StatsResultModel is the default implementation of StatsResult.
type StatsResultModel struct {
	core.BackedModel
}

// NewStatsResult creates a new instance of StatsResultModel.
func NewStatsResult() *StatsResultModel {
	return &StatsResultModel{
		BackedModel: core.NewBaseModel(),
	}
}

// CreateStatsResultFromDiscriminatorValue creates a new StatsResult from a ParseNode.
func CreateStatsResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewStatsResult(), nil
}

// GetFieldDeserializers implements serialization.Parsable.
func (m *StatsResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		statsResultStatsKey: internalSerialization.DeserializeObjectValueFunc[*StatsModel](CreateStatsFromDiscriminatorValue)(
			func(val *StatsModel) error { return m.setStats(val) }),
		statsResultGroupByFieldsKey: internalSerialization.DeserializeCollectionOfObjectValuesFunc[*GroupByFieldModel](CreateGroupByFieldFromDiscriminatorValue)(
			func(val []*GroupByFieldModel) error {
				fields := make([]GroupByField, len(val))
				for i, v := range val {
					fields[i] = v
				}
				return m.setGroupbyFields(fields)
			}),
	}
}

// Serialize implements serialization.Parsable.
func (m *StatsResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[Stats](statsResultStatsKey)(m.GetStats),
		internalSerialization.SerializeCollectionOfObjectValuesFunc[GroupByField](statsResultGroupByFieldsKey)(m.GetGroupbyFields),
	)
}

// GetStats returns the aggregate values of this result.
func (m *StatsResultModel) GetStats() (Stats, error) {
	return store.DefaultBackedModelAccessorFunc[*StatsResultModel, Stats](m, statsResultStatsKey)
}

func (m *StatsResultModel) setStats(val Stats) error {
	return store.DefaultBackedModelMutatorFunc(m, statsResultStatsKey, val)
}

// GetGroupbyFields returns the grouping dimensions of this result, when sysparm_group_by was requested.
func (m *StatsResultModel) GetGroupbyFields() ([]GroupByField, error) {
	return store.DefaultBackedModelAccessorFunc[*StatsResultModel, []GroupByField](m, statsResultGroupByFieldsKey)
}

func (m *StatsResultModel) setGroupbyFields(val []GroupByField) error {
	return store.DefaultBackedModelMutatorFunc(m, statsResultGroupByFieldsKey, val)
}
