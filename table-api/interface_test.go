package tableapi

import (
	"math"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInterface_UnmarshalJSON(t *testing.T) {
	tests := []test[*DataValue]{
		{
			title:     "Int",
			value:     []byte(`"42"`),
			expected:  &DataValue{value: int64(42)},
			expectErr: false,
		},
		{
			title:     "String",
			value:     []byte(`"hello"`),
			expected:  &DataValue{value: "hello"},
			expectErr: false,
		},
		{
			title:     "Bool",
			value:     []byte(`"true"`),
			expected:  &DataValue{value: true},
			expectErr: false,
		},
		{
			title:     "Invalid",
			value:     []byte(`\"true"`),
			expected:  nil,
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {

			temp := &DataValue{}

			err := temp.UnmarshalJSON(tt.value.([]byte))
			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, temp)
		})
	}
}

func TestInterface_Int(t *testing.T) {
	tests := []test[int64]{
		{"Int64", &DataValue{value: int64(42)}, 42, false, nil},
		{"Int16", &DataValue{value: int16(12)}, 12, false, nil},
		{"Int8", &DataValue{value: int8(1)}, 1, false, nil},
		{"Int", &DataValue{value: int(1)}, 1, false, nil},
		{"Int32", &DataValue{value: int32(123)}, 123, false, nil},
		{"Float32", &DataValue{value: float32(3.14)}, 0, true, nil},
		{"String", &DataValue{value: "not an integer"}, 0, true, nil},
		{"Nil", &DataValue{value: nil}, 0, true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			got, err := tt.value.(*DataValue).Int()
			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInterface_Float(t *testing.T) {
	tests := []test[float64]{
		{"Float32", &DataValue{value: float32(3.14)}, 3.1400000, false, nil},
		{"Float64", &DataValue{value: float64(2.71828)}, 2.71828, false, nil},
		{"Int", &DataValue{value: int(42)}, 0, true, nil},
		{"String", &DataValue{value: "not a float"}, 0, true, nil},
		{"Nil", &DataValue{value: nil}, 0, true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			got, err := tt.value.(*DataValue).Float()
			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
			if math.Abs(got-tt.expected) > tolerance {
				assert.Equal(t, tt.expected, got)
			}
		})
	}
}

func TestInterface_String(t *testing.T) {
	tests := []test[string]{
		{"String", &DataValue{value: "hello"}, "hello", false, nil},
		{"String", &DataValue{value: "world"}, "world", false, nil},
		{"Int", &DataValue{value: 42}, "", true, nil},
		{"Nil", &DataValue{value: nil}, "", true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			got, err := tt.value.(*DataValue).String()
			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInterface_Bool(t *testing.T) {
	tests := []test[bool]{
		{"Bool", &DataValue{value: true}, true, false, nil},
		{"Int", &DataValue{value: 0}, false, true, nil},
		{"String", &DataValue{value: "true"}, false, true, nil},
		{"Nil", &DataValue{value: nil}, false, true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			got, err := tt.value.(*DataValue).Bool()
			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInterface_Type(t *testing.T) {
	tests := []test[reflect.Type]{
		{"Int", &DataValue{value: int(42)}, reflect.TypeOf(int(0)), false, nil},
		{"Float", &DataValue{value: float64(3.14)}, reflect.TypeOf(float64(0)), false, nil},
		{"String", &DataValue{value: "hello"}, reflect.TypeOf(""), false, nil},
		{"Bool", &DataValue{value: true}, reflect.TypeOf(true), false, nil},
		{"Nil", &DataValue{value: nil}, nil, false, nil},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			got := tt.value.(*DataValue).Type()
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestInterface_Time(t *testing.T) {

	now := time.Now()

	tests := []test[time.Time]{
		{"DateTime", &DataValue{value: now}, now, false, nil},
		{"Date", &DataValue{value: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())}, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()), false, nil},
		{"Time", &DataValue{value: time.Date(0, 0, 0, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())}, time.Date(0, 0, 0, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location()), false, nil},
		{"nil", &DataValue{value: time.Time{}}, time.Time{}, false, nil},
		{"string", &DataValue{value: "TIme"}, time.Time{}, true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			got, err := tt.value.(*DataValue).Time()
			if tt.expectErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}
