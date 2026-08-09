package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNilPointerError(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				err := NewNilPointerError("value is nil")

				assert.Equal(t, "value is nil", err.s)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNilPointerError_Error(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				err := NewNilPointerError("value is nil")

				assert.Equal(t, "value is nil", err.Error())
			},
		},
		{
			name: "NilReceiver",
			test: func(t *testing.T) {
				var err *NilPointerError

				assert.NotPanics(t, func() {
					assert.Equal(t, "nil pointer error", err.Error())
				})
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
