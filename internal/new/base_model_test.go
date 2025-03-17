package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/assert"
)

func TestNewBaseModel(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				model := NewBaseModel()
				assert.IsType(t, (store.BackingStoreFactory)(nil), model.backingStoreFactory)
				assert.NotNil(t, model.backingStoreFactory)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBaseModel_SetBackingStoreFactory(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				strct := mocking.NewMockBackingStoreFactory()
				factory := strct.MockBackingStoreFactory

				model := &BaseModel{}
				err := model.SetBackingStoreFactory(factory)
				assert.Nil(t, err)
				assert.NotNil(t, model.backingStoreFactory)
			},
		},
		{
			name: "Nil input",
			test: func(t *testing.T) {
				model := &BaseModel{}
				err := model.SetBackingStoreFactory(nil)
				assert.Equal(t, errors.New("factory is nil"), err)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*BaseModel)(nil)
				err := model.SetBackingStoreFactory(nil)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBaseModel_GetBackingStore(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful no store provided",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()

				strct := mocking.NewMockBackingStoreFactory()
				strct.On("MockBackingStoreFactory").Return(backingStore)
				factory := strct.MockBackingStoreFactory

				model := &BaseModel{
					backingStoreFactory: factory,
				}

				valBackingStore := model.GetBackingStore()

				assert.Equal(t, backingStore, valBackingStore)
				strct.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*BaseModel)(nil)

				valBackingStore := model.GetBackingStore()

				assert.Nil(t, valBackingStore)
			},
		},
		{
			name: "Successful existing store",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()

				strct := mocking.NewMockBackingStoreFactory()
				factory := strct.MockBackingStoreFactory

				model := &BaseModel{
					backingStoreFactory: factory,
					backingStore:        backingStore,
				}

				valBackingStore := model.GetBackingStore()

				assert.Equal(t, backingStore, valBackingStore)
				strct.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
