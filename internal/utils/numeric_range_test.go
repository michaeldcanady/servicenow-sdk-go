package utils

import (
	"testing"
)

func TestNumericRange2_Compatible(t *testing.T) {
	tests := []struct {
		name     string
		rng      *numericRange2
		val      float64
		expected bool
	}{
		{"Int8Ok", int8Range, 10, true},
		{"Int8Overflow", int8Range, 200, false},
		{"Int8Decimal", int8Range, 10.5, false},
		{"FloatOk", float64Range, 10.5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.rng.Compatible(tt.val) != tt.expected {
				t.Errorf("got %v, expected %v", tt.rng.Compatible(tt.val), tt.expected)
			}
		})
	}
}

func TestNewNumericRange2(t *testing.T) {
	rng := newNumericRange2(100, 0, false)
	if rng.max != 100 || rng.min != 0 || rng.allowDecimal != false {
		t.Error("newNumericRange2 failed")
	}
}

func TestNumericRange2_Within(t *testing.T) {
	rng := newNumericRange2(10, 0, false)
	if !rng.Within(5) || rng.Within(15) || rng.Within(-5) {
		t.Error("Within failed")
	}
}
