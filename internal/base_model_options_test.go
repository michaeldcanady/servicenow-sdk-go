package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWithBackingStoreFactory(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				factoryStrct := mocking.NewMockBackingStoreFactory()
				factory := factoryStrct.MockBackingStoreFactory

				setter := mocking.NewMockBackingStoreFactorySetter()
				setter.On("SetBackingStoreFactory", mock.AnythingOfType("store.BackingStoreFactory")).Return(nil)

				option := WithBackingStoreFactory[*mocking.MockBackingStoreFactorySetter](factory)

				err := option(setter)

				assert.Nil(t, err)
			},
		},
		{
			name: "Nil config",
			test: func(t *testing.T) {
				factoryStrct := mocking.NewMockBackingStoreFactory()
				factory := factoryStrct.MockBackingStoreFactory

				option := WithBackingStoreFactory[*mocking.MockBackingStoreFactorySetter](factory)

				err := option(nil)

				assert.Equal(t, errors.New("config is nil"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
