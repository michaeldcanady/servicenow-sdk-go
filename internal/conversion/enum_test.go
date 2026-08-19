package conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumString(t *testing.T) {
	m := map[int]string{1: "one", 2: "two"}

	tests := []struct {
		name     string
		value    int
		fallback string
		want     string
	}{
		{"Known value", 1, "unknown", "one"},
		{"Unknown value", 99, "unknown", "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EnumString(m, tt.value, tt.fallback)
			assert.Equal(t, tt.want, got)
		})
	}
}
