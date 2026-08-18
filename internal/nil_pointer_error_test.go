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

// Curveball (see write-unit-tests skill, step 6): a nil *NilPointerError
// panics on Error() instead of returning a safe value, because Error()
// dereferences the receiver directly with no nil-guard. This is the same
// "typed nil" trap that conversion.IsNil exists elsewhere in this repo to
// avoid - any caller that ends up with a nil *NilPointerError stored in an
// error-typed variable (e.g. a function returning (*NilPointerError)(nil), nil
// mistakenly, or a struct field left at its zero value) will crash the first
// time something calls .Error() on it, such as fmt.Errorf("%w", err) or a
// logger. Left as a skipped, documented case rather than "fixed" here since
// silently changing production code isn't this test pass's call - reported to
// the user instead.
func TestNilPointerError_Error_NilReceiver(t *testing.T) {
	t.Skip("BUG: (*NilPointerError)(nil).Error() panics with a nil pointer dereference instead of returning a safe string")

	var err *NilPointerError

	assert.NotPanics(t, func() {
		_ = err.Error()
	})
}
