package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestDefaultStoreMutatorFunc(t *testing.T) {
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

				err := DefaultStoreMutatorFunc(store, key, value)
				assert.Nil(t, err)
				store.AssertExpectations(t)
			},
		},
		{
			name: "Nil store",
			test: func(t *testing.T) {
				key := "key"
				value := "value"

				store := (*mocking.MockBackingStore)(nil)

				err := DefaultStoreMutatorFunc(store, key, value)
				assert.Equal(t, errors.New("store is nil"), err)
			},
		},
		{
			name: "Empty key",
			test: func(t *testing.T) {
				key := " "
				value := "value"

				store := mocking.NewMockBackingStore()

				err := DefaultStoreMutatorFunc(store, key, value)
				assert.Equal(t, errors.New("key is empty"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
