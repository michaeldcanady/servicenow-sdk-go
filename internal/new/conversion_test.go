package internal

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
