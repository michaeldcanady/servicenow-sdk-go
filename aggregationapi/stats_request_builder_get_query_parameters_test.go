package aggregationapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestStatsRequestBuilderGetQueryParameters_Encoding(t *testing.T) {
	tests := []struct {
		name     string
		params   StatsRequestBuilderGetQueryParameters
		expected map[string]any
	}{
		{
			name:     "zero value omits everything",
			params:   StatsRequestBuilderGetQueryParameters{},
			expected: map[string]any{},
		},
		{
			name: "count only",
			params: StatsRequestBuilderGetQueryParameters{
				Count: internal.ToPointer(true),
			},
			expected: map[string]any{"sysparm_count": "true"},
		},
		{
			name: "list fields are comma-joined into a single value",
			params: StatsRequestBuilderGetQueryParameters{
				SumFields: []string{"reassignment_count", "escalation"},
			},
			expected: map[string]any{"sysparm_sum_fields": "reassignment_count,escalation"},
		},
		{
			name: "display value enum encodes its string form",
			params: StatsRequestBuilderGetQueryParameters{
				DisplayValue: internal.ToPointer(DisplayValueAll),
			},
			expected: map[string]any{"sysparm_display_value": "all"},
		},
		{
			name:     "unset display value is omitted, not sent as \"unknown\"",
			params:   StatsRequestBuilderGetQueryParameters{},
			expected: map[string]any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestInfo := abstractions.NewRequestInformation()
			requestInfo.AddQueryParameters(tt.params)

			got := map[string]any{}
			for k, v := range requestInfo.QueryParametersAny {
				got[k] = v
			}
			for k, v := range requestInfo.QueryParameters {
				got[k] = v
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}
