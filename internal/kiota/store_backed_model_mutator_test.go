package kiota

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestDefaultBackedModelMutatorFunc(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				key := "key"
				value := "value"

				store := mocking.NewMockBackingStore()
				store.On("Set", key, value).Return(nil)

				err := DefaultBackedModelMutatorFunc(store, key, value)
				assert.Nil(t, err)
				store.AssertExpectations(t)
			},
		},
		{
			name: "Nil store",
			test: func(t *testing.T) {
				key := "key"
				value := "value"

				model := (*mocking.MockBackingStore)(nil)

				err := DefaultBackedModelMutatorFunc(model, key, value)
				assert.Equal(t, errors.New("backingStore is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
