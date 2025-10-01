package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestDefaultStoreAccessorFunc(t *testing.T) {
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
				store.On("Get", key).Return(value, nil)

				response, err := DefaultStoreAccessorFunc[*mocking.MockBackingStore, string](store, key)

				assert.Nil(t, err)
				assert.Equal(t, value, response)
				store.AssertExpectations(t)
			},
		},
		{
			name: "Nil store",
			test: func(t *testing.T) {
				key := "key"
				store := (*mocking.MockBackingStore)(nil)

				response, err := DefaultStoreAccessorFunc[*mocking.MockBackingStore, string](store, key)

				assert.Equal(t, errors.New("store is nil"), err)
				assert.Equal(t, "", response)
			},
		},
		{
			name: "Empty key",
			test: func(t *testing.T) {
				key := " "
				store := mocking.NewMockBackingStore()

				response, err := DefaultStoreAccessorFunc[*mocking.MockBackingStore, string](store, key)

				assert.Equal(t, errors.New("key is empty"), err)
				assert.Equal(t, "", response)
				store.AssertExpectations(t)
			},
		},
		{
			name: "Store error",
			test: func(t *testing.T) {
				key := "key"
				store := mocking.NewMockBackingStore()
				store.On("Get", key).Return("", errors.New("store error"))

				response, err := DefaultStoreAccessorFunc[*mocking.MockBackingStore, string](store, key)

				assert.Equal(t, "", response)
				assert.Equal(t, errors.New("store error"), err)
				store.AssertExpectations(t)
			},
		},
		{
			name: "Wrong response type",
			test: func(t *testing.T) {
				key := "key"
				value := true
				store := mocking.NewMockBackingStore()
				store.On("Get", key).Return(value, nil)

				response, err := DefaultStoreAccessorFunc[*mocking.MockBackingStore, string](store, key)

				assert.Equal(t, errors.New("unsupported conversion: bool to string"), err)
				assert.Equal(t, "", response)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
