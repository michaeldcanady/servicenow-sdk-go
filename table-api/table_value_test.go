package tableapi

import (
	"reflect"
	"testing"
)

func TestTableValue_Int(t *testing.T) {
	tests := []struct {
		name     string
		val      interface{}
		expected int64
		err      bool
	}{
		{"int", 1, 1, false},
		{"int8", int8(2), 2, false},
		{"int16", int16(3), 3, false},
		{"int32", int32(4), 4, false},
		{"int64", int64(5), 5, false},
		{"string", "bad", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tv := &TableValue{value: tt.val}
			res, err := tv.Int()
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
			// Test deprecated ToInt64
			res2, _ := tv.ToInt64()
			if res2 != res {
				t.Error("ToInt64 mismatch")
			}
		})
	}
}

func TestTableValue_Float(t *testing.T) {
	tests := []struct {
		name     string
		val      interface{}
		expected float64
		err      bool
	}{
		{"float32", float32(1.5), 1.5, false},
		{"float64", float64(2.5), 2.5, false},
		{"int", 1, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tv := &TableValue{value: tt.val}
			res, err := tv.Float()
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
			// Test deprecated ToFloat64
			res2, _ := tv.ToFloat64()
			if res2 != res {
				t.Error("ToFloat64 mismatch")
			}
		})
	}
}

func TestTableValue_String(t *testing.T) {
	tests := []struct {
		name     string
		val      interface{}
		expected string
		err      bool
	}{
		{"string", "v", "v", false},
		{"int", 1, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tv := &TableValue{value: tt.val}
			res, err := tv.String()
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
			// Test deprecated ToString
			res2, _ := tv.ToString()
			if res2 != res {
				t.Error("ToString mismatch")
			}
		})
	}
}

func TestTableValue_Bool(t *testing.T) {
	tests := []struct {
		name     string
		val      interface{}
		expected bool
		err      bool
	}{
		{"bool", true, true, false},
		{"string", "true", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tv := &TableValue{value: tt.val}
			res, err := tv.Bool()
			if (err != nil) != tt.err {
				t.Errorf("err: got %v, expected %v", err, tt.err)
			}
			if res != tt.expected {
				t.Errorf("got %v, expected %v", res, tt.expected)
			}
			// Test deprecated ToBool
			res2, _ := tv.ToBool()
			if res2 != res {
				t.Error("ToBool mismatch")
			}
		})
	}
}

func TestTableValue_Type(t *testing.T) {
	tv := &TableValue{value: 1}
	if tv.Type() != reflect.TypeOf(1) {
		t.Error("Type failed")
	}
	if tv.GetType() != tv.Type() {
		t.Error("GetType mismatch")
	}
}
