package aggregationapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMapFromRaw(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected map[string]string
		wantErr  bool
	}{
		{
			name:     "nil",
			input:    nil,
			expected: nil,
		},
		{
			name:     "string values",
			input:    map[string]any{"reassignment_count": "56"},
			expected: map[string]string{"reassignment_count": "56"},
		},
		{
			name:     "non-string values are stringified",
			input:    map[string]any{"reassignment_count": 56},
			expected: map[string]string{"reassignment_count": "56"},
		},
		{
			name:     "string pointer values are dereferenced",
			input:    map[string]any{"reassignment_count": internal.ToPointer("56")},
			expected: map[string]string{"reassignment_count": "56"},
		},
		{
			name:     "null values are omitted, not stringified to \"<nil>\"",
			input:    map[string]any{"reassignment_count": nil, "escalation": "1"},
			expected: map[string]string{"escalation": "1"},
		},
		{
			name:    "unsupported type",
			input:   "not-a-map",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := mapFromRaw(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewStats(t *testing.T) {
	stats := NewStats()
	assert.NotNil(t, stats)
}

func TestCreateStatsFromDiscriminatorValue(t *testing.T) {
	stats, err := CreateStatsFromDiscriminatorValue(nil)
	assert.NoError(t, err)
	assert.NotNil(t, stats)
}

func TestStats_CountGetterAndSetter(t *testing.T) {
	m := NewStats()
	val := internal.ToPointer("67")

	err := m.setCount(val)
	assert.NoError(t, err)

	res, err := m.GetCount()
	assert.NoError(t, err)
	assert.Equal(t, val, res)
}

func TestStats_MapGettersAndSetters(t *testing.T) {
	tests := []struct {
		name   string
		setter func(*StatsModel, map[string]string) error
		getter func(*StatsModel) (map[string]string, error)
	}{
		{
			name:   "Sum",
			setter: func(m *StatsModel, v map[string]string) error { return m.setSum(v) },
			getter: func(m *StatsModel) (map[string]string, error) { return m.GetSum() },
		},
		{
			name:   "Avg",
			setter: func(m *StatsModel, v map[string]string) error { return m.setAvg(v) },
			getter: func(m *StatsModel) (map[string]string, error) { return m.GetAvg() },
		},
		{
			name:   "Min",
			setter: func(m *StatsModel, v map[string]string) error { return m.setMin(v) },
			getter: func(m *StatsModel) (map[string]string, error) { return m.GetMin() },
		},
		{
			name:   "Max",
			setter: func(m *StatsModel, v map[string]string) error { return m.setMax(v) },
			getter: func(m *StatsModel) (map[string]string, error) { return m.GetMax() },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewStats()
			val := map[string]string{"reassignment_count": "2"}

			err := tt.setter(m, val)
			assert.NoError(t, err)

			res, err := tt.getter(m)
			assert.NoError(t, err)
			assert.Equal(t, val, res)
		})
	}
}

func TestStats_Serialize(t *testing.T) {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
	writer.On("WriteAnyValue", mock.Anything, mock.Anything).Return(nil)

	m := NewStats()
	_ = m.setCount(internal.ToPointer("67"))
	_ = m.setSum(map[string]string{"reassignment_count": "56"})

	err := m.Serialize(writer)
	assert.NoError(t, err)

	var nilStats *StatsModel
	err = nilStats.Serialize(writer)
	assert.NoError(t, err)
}

func TestStats_Serialize_SkipsNilMaps(t *testing.T) {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)

	m := NewStats()
	_ = m.setCount(internal.ToPointer("67"))

	err := m.Serialize(writer)
	assert.NoError(t, err)

	writer.AssertNotCalled(t, "WriteAnyValue", mock.Anything, mock.Anything)
}

func TestStats_GetFieldDeserializers(t *testing.T) {
	m := NewStats()
	deser := m.GetFieldDeserializers()
	assert.NotNil(t, deser[statsCountKey])
	assert.NotNil(t, deser[statsSumKey])
	assert.NotNil(t, deser[statsAvgKey])
	assert.NotNil(t, deser[statsMinKey])
	assert.NotNil(t, deser[statsMaxKey])
}
