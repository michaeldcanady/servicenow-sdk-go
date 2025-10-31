package tableapi

import (
	"reflect"
	"testing"

	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestElementVis_VisitSlice(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.VisitSlice(reflect.ValueOf([]string{"test", "testing"}))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: []*ElementValue{{val: "test"}, {val: "testing"}}}, elem)
			},
		},
		{
			name: "Nil",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.VisitSlice(reflect.ValueOf(([]string)(nil)))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: []*ElementValue{}}, elem)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementVis_VisitMap(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.VisitMap(reflect.ValueOf(map[string]string{"test": "testing"}))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: map[string]*ElementValue{"test": {val: "testing"}}}, elem)
			},
		},
		{
			name: "Nil",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.VisitMap(reflect.ValueOf((map[string]string)(nil)))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: map[string]*ElementValue{}}, elem)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementVis_VisitPointer(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.VisitPointer(reflect.ValueOf(internal.ToPointer("test")))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: "test"}, elem)
			},
		},
		{
			name: "Nil",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.VisitPointer(reflect.ValueOf((*string)(nil)))

				assert.Nil(t, err)
				assert.Equal(t, (*ElementValue)(nil), elem)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementVis_VisitPrimitive(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "String",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.VisitPrimitive(reflect.ValueOf("test"))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: "test"}, elem)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestElementVis_Visit(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "String",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.Visit("test")

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: "test"}, elem)
			},
		},
		{
			name: "Int64",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.Visit(int64(1))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: int64(1)}, elem)
			},
		},
		{
			name: "Int32",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.Visit(int32(1))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: int32(1)}, elem)
			},
		},
		{
			name: "Float64",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.Visit(float64(1.00))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: float64(1.00)}, elem)
			},
		},
		{
			name: "Float32",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.Visit(float32(1.00))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: float32(1.00)}, elem)
			},
		},
		{
			name: "Byte",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.Visit(byte(1))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: byte(1)}, elem)
			},
		},
		{
			name: "Int8",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.Visit(int8(1))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: int8(1)}, elem)
			},
		},
		{
			name: "Bool",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.Visit(true)

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: true}, elem)
			},
		},
		{
			name: "reflect.Value",
			test: func(t *testing.T) {
				visitor := &elementVis{}

				elem, err := visitor.Visit(reflect.ValueOf(true))

				assert.Nil(t, err)
				assert.Equal(t, &ElementValue{val: true}, elem)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
