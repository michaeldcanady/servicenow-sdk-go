package tableapi

import (
	"testing"
)

func TestNewTableEntry(t *testing.T) {
	entry := NewTableEntry()
	if entry == nil {
		t.Error("NewTableEntry returned nil")
	}
}

func TestTableEntry_Set(t *testing.T) {
	entry := NewTableEntry()
	entry.Set("k", "v")
	if entry["k"] != "v" {
		t.Errorf("expected v, got %v", entry["k"])
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
				if res != nil {
					t.Error("expected nil")
				}
			} else {
				if res.value != tt.expected {
					t.Errorf("got %v, expected %v", res.value, tt.expected)
				}
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
				if res != nil {
					t.Error("expected nil")
				}
			} else {
				if res.value != tt.expected {
					t.Errorf("got %v, expected %v", res.value, tt.expected)
				}
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
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if tt.expected == nil {
				if res != nil {
					t.Error("expected nil res")
				}
			} else if *res != *tt.expected {
				t.Errorf("got %s, expected %s", *res, *tt.expected)
			}
		})
	}
}

func TestTableEntry_Keys(t *testing.T) {
	entry := TableEntry{"a": 1, "b": 2}
	keys := entry.Keys()
	if len(keys) != 2 {
		t.Error("wrong number of keys")
	}
}

func TestTableEntry_Len(t *testing.T) {
	entry := TableEntry{"a": 1}
	if entry.Len() != 1 {
		t.Error("wrong len")
	}
}

func strPtr(s string) *string { return &s }
