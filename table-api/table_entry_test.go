package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTableEntry(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Create entry",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			entry := NewTableEntry()
			assert.NotNil(t, entry)
		})
	}
}

func TestTableEntry_Set(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value interface{}
	}{
		{
			name:  "Set string",
			key:   "k",
			value: "v",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			entry := NewTableEntry()
			entry.Set(test.key, test.value)
			assert.Equal(t, test.value, entry[test.key])
		})
	}
}

func TestTableEntry_Value(t *testing.T) {
	entry := TableEntry{
		"simple": "v1",
		"nested": map[string]interface{}{"value": "v2"},
	}
	tests := []struct {
		name     string
		key      string
		expected interface{}
	}{
		{"Simple", "simple", "v1"},
		{"Nested", "nested", "v2"},
		{"Missing", "bad", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := entry.Value(tt.key)
			if tt.expected == nil {
				assert.Nil(t, res)
			} else {
				assert.NotNil(t, res)
				assert.Equal(t, tt.expected, res.value)
			}
		})
	}
}

func TestTableEntry_DisplayValue(t *testing.T) {
	entry := TableEntry{
		"simple": "v1",
		"nested": map[string]interface{}{"displayValue": "v2"},
	}
	tests := []struct {
		name     string
		key      string
		expected interface{}
	}{
		{"Simple", "simple", "v1"},
		{"Nested", "nested", "v2"},
		{"Missing", "bad", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := entry.DisplayValue(tt.key)
			if tt.expected == nil {
				assert.Nil(t, res)
			} else {
				assert.NotNil(t, res)
				assert.Equal(t, tt.expected, res.value)
			}
		})
	}
}

func TestTableEntry_Link(t *testing.T) {
	entry := TableEntry{
		"nested":    map[string]interface{}{"link": "url"},
		"badNested": map[string]interface{}{"link": 123},
		"simple":    "v",
	}
	tests := []struct {
		name     string
		key      string
		expected *string
		err      bool
	}{
		{"Nested", "nested", strPtr("url"), false},
		{"BadNested", "badNested", nil, true},
		{"Simple", "simple", nil, false},
		{"Missing", "bad", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := entry.Link(tt.key)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			if tt.expected == nil {
				assert.Nil(t, res)
			} else {
				assert.NotNil(t, res)
				assert.Equal(t, *tt.expected, *res)
			}
		})
	}
}

func TestTableEntry_Keys(t *testing.T) {
	tests := []struct {
		name          string
		entry         TableEntry
		expectedCount int
	}{
		{
			name:          "Two keys",
			entry:         TableEntry{"a": 1, "b": 2},
			expectedCount: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			keys := test.entry.Keys()
			assert.Len(t, keys, test.expectedCount)
		})
	}
}

func TestTableEntry_Len(t *testing.T) {
	tests := []struct {
		name        string
		entry       TableEntry
		expectedLen int
	}{
		{
			name:        "One entry",
			entry:       TableEntry{"a": 1},
			expectedLen: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedLen, test.entry.Len())
		})
	}
}

func strPtr(s string) *string { return &s }
