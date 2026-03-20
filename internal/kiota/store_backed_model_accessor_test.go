package kiota

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestDefaultBackedModelAccessorFunc(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "nil model",
			test: func(t *testing.T) {
				model := (*mocking.MockBackingStore)(nil)

				value, err := DefaultBackedModelAccessorFunc[*mocking.MockBackingStore, string](model, "test")

				assert.Empty(t, value)
				assert.Equal(t, errors.New("backingStore is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
