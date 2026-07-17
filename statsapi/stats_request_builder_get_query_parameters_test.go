package statsapi

import (
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatsRequestBuilderGetQueryParameters_Encoding(t *testing.T) {
	tests := []struct {
		name     string
		params   StatsRequestBuilderGetQueryParameters
		expected string
	}{
		{
			name:     "zero value omits everything",
			params:   StatsRequestBuilderGetQueryParameters{},
			expected: "",
		},
		{
			name: "count only",
			params: StatsRequestBuilderGetQueryParameters{
				Count: true,
			},
			expected: "sysparm_count=true",
		},
		{
			name: "list fields are comma-joined into a single value",
			params: StatsRequestBuilderGetQueryParameters{
				SumFields: []string{"reassignment_count", "escalation"},
			},
			expected: "sysparm_sum_fields=reassignment_count%2Cescalation",
		},
		{
			name: "display value enum encodes its string form",
			params: StatsRequestBuilderGetQueryParameters{
				DisplayValue: DisplayValueAll,
			},
			expected: "sysparm_display_value=all",
		},
		{
			name:     "unset display value is omitted, not sent as \"unknown\"",
			params:   StatsRequestBuilderGetQueryParameters{DisplayValue: DisplayValueUnknown},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, err := query.Values(tt.params)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, values.Encode())
		})
	}
}
