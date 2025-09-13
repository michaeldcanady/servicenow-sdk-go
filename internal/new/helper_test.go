package internal

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNil(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "untyped nil",
			test: func(t *testing.T) {
				assert.True(t, IsNil(nil))
			},
		},
		{
			name: "nil string pointer",
			test: func(t *testing.T) {
				assert.True(t, IsNil((*string)(nil)))
			},
		},
		{
			name: "nil interface",
			test: func(t *testing.T) {
				assert.True(t, IsNil((interface{})(nil)))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestThrowErrors(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestToPointer(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "String",
			test: func(t *testing.T) {
				val := "string"
				assert.Equal(t, &val, ToPointer("string"))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestDereference(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "String pointer",
			test: func(t *testing.T) {
				input := ToPointer("testing")
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, interface{}("testing"), output.Interface())
			},
		},
		{
			name: "String pointer pointer",
			test: func(t *testing.T) {
				input := ToPointer(ToPointer("testing"))
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, interface{}("testing"), output.Interface())
			},
		},
		{
			name: "String",
			test: func(t *testing.T) {
				input := "testing"
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, interface{}("testing"), output.Interface())
			},
		},
		{
			name: "nil",
			test: func(t *testing.T) {
				input := (*string)(nil)
				output := Dereference(reflect.ValueOf(input))

				assert.Equal(t, "", output.Interface())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: tests for all types
func TestConvert(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "String pointer",
			test: func(t *testing.T) {
				var input interface{} = ToPointer("test")
				var output *string

				err := Convert(input, &output)

				assert.Nil(t, err)
				assert.Equal(t, ToPointer("test"), output)
			},
		},
		{
			name: "int8",
			test: func(t *testing.T) {
				var input interface{} = int8(8)
				var output int8

				err := Convert(input, &output)

				assert.Nil(t, err)
				assert.Equal(t, int8(8), output)
			},
		},
		{
			name: "String pointer pointer",
			test: func(t *testing.T) {
				var input interface{} = ToPointer(ToPointer("test"))
				var output **string

				err := Convert(input, &output)

				assert.Nil(t, err)
				assert.Equal(t, ToPointer(ToPointer("test")), output)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
