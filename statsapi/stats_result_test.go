package statsapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewStatsResult(t *testing.T) {
	result := NewStatsResult()
	assert.NotNil(t, result)
}

func TestCreateStatsResultFromDiscriminatorValue(t *testing.T) {
	result, err := CreateStatsResultFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestStatsResult_StatsGetterAndSetter(t *testing.T) {
	result := NewStatsResult()
	stats := NewStats()

	err := result.setStats(stats)
	assert.NoError(t, err)

	res, err := result.GetStats()
	assert.NoError(t, err)
	assert.Equal(t, Stats(stats), res)
}

func TestStatsResult_Serialize(t *testing.T) {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	result := NewStatsResult()
	_ = result.setStats(NewStats())

	err := result.Serialize(writer)
	assert.NoError(t, err)

	var nilResult *StatsResultModel
	err = nilResult.Serialize(writer)
	assert.NoError(t, err)
}

func TestStatsResult_GetFieldDeserializers(t *testing.T) {
	result := NewStatsResult()
	deser := result.GetFieldDeserializers()
	assert.NotNil(t, deser[statsResultStatsKey])
}
