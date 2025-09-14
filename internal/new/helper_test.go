package internal

import (
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
