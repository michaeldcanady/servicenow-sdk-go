package tableapi

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestElementVis_VisitSlice(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.VisitSlice(reflect.ValueOf([]string{"test", "testing"}))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: []*ElementValue{{val: "test"}, {val: "testing"}}}, elem)
			},
		},
		{
			name: "Nil",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.VisitSlice(reflect.ValueOf(([]string)(nil)))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: []*ElementValue{}}, elem)
			},
		},
		{
			name: "Invalid type",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.VisitSlice(reflect.ValueOf([]func(){func() {}, func() {}}))

				assert.Equal(t, errors.New("unsupported kind func"), err)
				assert.Nil(t, elem)
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
				visitor := &ElementVisitor{}

				elem, err := visitor.VisitMap(reflect.ValueOf(map[string]string{"test": "testing"}))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: map[string]*ElementValue{"test": {val: "testing"}}}, elem)
			},
		},
		{
			name: "Nil",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.VisitMap(reflect.ValueOf((map[string]string)(nil)))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: map[string]*ElementValue{}}, elem)
			},
		},
		{
			name: "Invalid type",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.VisitMap(reflect.ValueOf(map[string]func(){"test": func() {}, "test1": func() {}}))

				assert.Equal(t, errors.New("unsupported kind func"), err)
				assert.Nil(t, elem)
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
				visitor := &ElementVisitor{}

				elem, err := visitor.VisitPointer(reflect.ValueOf(internal.ToPointer("test")))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: "test"}, elem)
			},
		},
		{
			name: "Nil",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.VisitPointer(reflect.ValueOf((*string)(nil)))

				require.NoError(t, err)
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
				visitor := &ElementVisitor{}

				elem, err := visitor.VisitPrimitive(reflect.ValueOf("test"))

				require.NoError(t, err)
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
				visitor := &ElementVisitor{}

				elem, err := visitor.Visit("test")

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: "test"}, elem)
			},
		},
		{
			name: "Int64",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.Visit(int64(1))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: int64(1)}, elem)
			},
		},
		{
			name: "Int32",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.Visit(int32(1))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: int32(1)}, elem)
			},
		},
		{
			name: "Float64",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.Visit(float64(1.00))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: float64(1.00)}, elem)
			},
		},
		{
			name: "Float32",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.Visit(float32(1.00))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: float32(1.00)}, elem)
			},
		},
		{
			name: "Byte",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.Visit(byte(1))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: byte(1)}, elem)
			},
		},
		{
			name: "Int8",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.Visit(int8(1))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: int8(1)}, elem)
			},
		},
		{
			name: "Bool",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.Visit(true)

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: true}, elem)
			},
		},
		{
			name: "reflect.Value",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.Visit(reflect.ValueOf(true))

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: true}, elem)
			},
		},
		{
			name: "Nil",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				elem, err := visitor.Visit(nil)

				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: nil}, elem)
			},
		},
		{
			name: "Map",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				in := map[string]string{"test": "test"}

				elem, err := visitor.Visit(in)
				require.NoError(t, err)
				assert.Equal(t, &ElementValue{val: map[string]*ElementValue{"test": {val: "test"}}}, elem)
			},
		},
		{
			name: "Unsupported kind",
			test: func(t *testing.T) {
				visitor := &ElementVisitor{}

				in := make(chan string)

				elem, err := visitor.Visit(in)
				assert.Equal(t, fmt.Errorf("unsupported kind %s", reflect.Chan), err)
				assert.Equal(t, (*ElementValue)(nil), elem)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
