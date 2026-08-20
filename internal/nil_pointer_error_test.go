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
				text := "value is nil"

				err := NewNilPointerError(text)

				assert.NotNil(t, err)
				assert.Equal(t, text, err.s)
			},
		},
		{
			name: "Empty string",
			test: func(t *testing.T) {
				err := NewNilPointerError("")

				assert.NotNil(t, err)
				assert.Empty(t, err.s)
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
			name: "Empty string",
			test: func(t *testing.T) {
				err := NewNilPointerError("")

				assert.Empty(t, err.Error())
			},
		},
		{
			name: "Unicode text",
			test: func(t *testing.T) {
				text := "値がnilです 🚫"

				err := NewNilPointerError(text)

				assert.Equal(t, text, err.Error())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// A typed-nil *NilPointerError must return a safe string instead of panicking.
// The guard was added in #591 and regressed by the v2 rework (#600); this test
// (previously skipped as a known bug) now protects the restored guard.
func TestNilPointerError_Error_NilReceiver(t *testing.T) {
	var err *NilPointerError

	var got string
	assert.NotPanics(t, func() { got = err.Error() })
	assert.Equal(t, "nil pointer error", got)
}
