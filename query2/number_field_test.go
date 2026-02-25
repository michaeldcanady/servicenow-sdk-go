//go:build preview.query

package query2

import (
	"testing"
)

func TestNumberField_Is(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f=1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).Is(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).Is(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_IsNot(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f!=1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).IsNot(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).IsNot(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_LessThan(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f<1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).LessThan(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).LessThan(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_GreaterThan(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f>1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).GreaterThan(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).GreaterThan(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_LessThanOrIs(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f<=1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).LessThanOrIs(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).LessThanOrIs(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_GreaterThanOrIs(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "f>=1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).GreaterThanOrIs(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).GreaterThanOrIs(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_Between(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		l, u     float64
		expected string
		isErr    bool
	}{
		{"Valid", "f", 1, 2, "fBETWEEN1@2", false},
		{"Invalid", "f", 2, 1, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Number(tt.field).Between(tt.l, tt.u)
			if tt.isErr {
				if res.Error() == nil {
					t.Error("Expected error")
				}
			} else {
				if res.String() != tt.expected {
					t.Errorf("got %s, expected %s", res.String(), tt.expected)
				}
			}
		})
	}
}

func TestNumberField_IsOneOf(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		vals     []float64
		expected string
	}{
		{"Multiple", "f", []float64{1, 2}, "fIN1,2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).IsOneOf(tt.vals...).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).IsOneOf(tt.vals...).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_IsNotOneOf(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		vals     []float64
		expected string
	}{
		{"Multiple", "f", []float64{1, 2}, "fNOT IN1,2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).IsNotOneOf(tt.vals...).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).IsNotOneOf(tt.vals...).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_GreaterThanField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fGT_FIELDo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).GreaterThanField(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).GreaterThanField(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_LessThanField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fLT_FIELDo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).LessThanField(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).LessThanField(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_GreaterThanOrIsField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fGT_OR_EQUALS_FIELDo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).GreaterThanOrIsField(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).GreaterThanOrIsField(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_LessThanOrIsField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		other    string
		expected string
	}{
		{"Standard", "f", "o", "fLT_OR_EQUALS_FIELDo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).LessThanOrIsField(tt.other).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).LessThanOrIsField(tt.other).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_IsMoreThan(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "fMORETHAN1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).IsMoreThan(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).IsMoreThan(tt.val).String(), tt.expected)
			}
		})
	}
}

func TestNumberField_IsLessThan(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		val      float64
		expected string
	}{
		{"Standard", "f", 1, "fLESSTHAN1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Number(tt.field).IsLessThan(tt.val).String() != tt.expected {
				t.Errorf("got %s, expected %s", Number(tt.field).IsLessThan(tt.val).String(), tt.expected)
			}
		})
	}
}
