package aggregationapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	statsResultStatsKey = "stats"
)

// StatsResult represents a Stats API ("/api/now/stats/{table}") response.
//
// This models the ungrouped shape only (a single StatsResult under "result"). A request that set
// sysparm_group_by would return an array of results under "result" instead, which
// StatsRequestBuilder.Get does not support — see statsURLTemplate.
type StatsResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetStats() (Stats, error)
	setStats(Stats) error
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
		statsResultStatsKey: internalSerialization.DeserializeObjectValueFunc[*StatsModel](CreateStatsFromDiscriminatorValue,
			func(val *StatsModel) error { return m.setStats(val) }),
	}
}

// Serialize implements serialization.Parsable.
func (m *StatsResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[Stats](statsResultStatsKey, m.GetStats),
	)
}

// GetStats returns the aggregate values of this result.
func (m *StatsResultModel) GetStats() (Stats, error) {
	return store.DefaultBackedModelAccessorFunc[*StatsResultModel, Stats](m, statsResultStatsKey)
}

func (m *StatsResultModel) setStats(val Stats) error {
	return store.DefaultBackedModelMutatorFunc(m, statsResultStatsKey, val)
}
