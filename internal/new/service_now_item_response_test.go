package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/assert"
)

func TestNewBaseServiceNowItemResponse(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				strct := mocking.NewMockParsableFactory()
				parsableFactory := strct.Factory

				parsable := NewBaseServiceNowItemResponse[*mocking.MockParsable](parsableFactory)
				assert.IsType(t, &BaseServiceNowItemResponse[*mocking.MockParsable]{}, parsable)
				assert.IsType(t, &store.InMemoryBackingStore{}, parsable.backingStore)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBaseServiceNowItemResponse_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Nil writer",
			test: func(t *testing.T) {
				parsable := &BaseServiceNowItemResponse[*mocking.MockParsable]{}
				err := parsable.Serialize(nil)

				assert.Equal(t, errors.New("Serialization is not supported"), err)
			},
		},
		{
			name: "Writer",
			test: func(t *testing.T) {
				writer := mocking.NewMockSerializationWriter()

				parsable := &BaseServiceNowItemResponse[*mocking.MockParsable]{}
				err := parsable.Serialize(writer)

				assert.Equal(t, errors.New("Serialization is not supported"), err)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				parsable := (*BaseServiceNowItemResponse[*mocking.MockParsable])(nil)
				err := parsable.Serialize(nil)

				assert.Equal(t, errors.New("Serialization is not supported"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: implement tests
func TestBaseServiceNowItemResponse_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBaseServiceNowItemResponse_GetBackingStore(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Existing store",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()

				parsable := &BaseServiceNowItemResponse[*mocking.MockParsable]{
					backingStore: backingStore,
				}

				store, err := parsable.GetBackingStore()

				assert.Nil(t, err)
				assert.Equal(t, backingStore, store)
			},
		},
		{
			name: "Existing factory, Nil store",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()

				strct := mocking.NewMockBackingStoreFactory()
				strct.On("MockBackingStoreFactory").Return(backingStore)
				factory := strct.MockBackingStoreFactory

				parsable := &BaseServiceNowItemResponse[*mocking.MockParsable]{
					backingStoreFactory: factory,
					backingStore:        nil,
				}

				store, err := parsable.GetBackingStore()

				assert.Nil(t, err)
				assert.Equal(t, backingStore, store)
				assert.Equal(t, backingStore, parsable.backingStore)
				strct.AssertExpectations(t)
			},
		},
		{
			name: "Nil factory, Nil store",
			test: func(t *testing.T) {
				parsable := &BaseServiceNowItemResponse[*mocking.MockParsable]{
					backingStoreFactory: nil,
					backingStore:        nil,
				}

				store, err := parsable.GetBackingStore()

				assert.Equal(t, errors.New("store is nil"), err)
				assert.Nil(t, store)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				parsable := (*BaseServiceNowItemResponse[*mocking.MockParsable])(nil)

				store, err := parsable.GetBackingStore()

				assert.Nil(t, err)
				assert.Nil(t, store)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBaseServiceNowItemResponse_GetResult(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Expected type",
			test: func(t *testing.T) {
				mockResult := mocking.NewMockParsable()

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", resultKey).Return(mockResult, nil)

				parsable := &BaseServiceNowItemResponse[*mocking.MockParsable]{
					backingStore: backingStore,
				}

				result, err := parsable.GetResult()

				assert.Nil(t, err)
				assert.Equal(t, mockResult, result)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				mockResult := struct{}{}

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", resultKey).Return(mockResult, nil)

				parsable := &BaseServiceNowItemResponse[*mocking.MockParsable]{
					backingStore: backingStore,
				}

				result, err := parsable.GetResult()

				assert.Equal(t, errors.New("value is not *mocking.MockParsable"), err)
				assert.Nil(t, result)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", resultKey).Return(nil, errors.New("retrieval error"))

				parsable := &BaseServiceNowItemResponse[*mocking.MockParsable]{
					backingStore: backingStore,
				}

				result, err := parsable.GetResult()

				assert.Equal(t, errors.New("retrieval error"), err)
				assert.Nil(t, result)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "BackingStore error",
			test: func(t *testing.T) {
				parsable := &BaseServiceNowItemResponse[*mocking.MockParsable]{}

				result, err := parsable.GetResult()

				assert.Equal(t, errors.New("store is nil"), err)
				assert.Nil(t, result)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				parsable := (*BaseServiceNowItemResponse[*mocking.MockParsable])(nil)

				result, err := parsable.GetResult()

				assert.Nil(t, err)
				assert.Nil(t, result)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: implement tests
func TestBaseServiceNowItemResponse_setResult(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
