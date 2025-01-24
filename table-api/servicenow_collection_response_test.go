package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewServiceNowCollectionResponse(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Successful",
			test: func(t *testing.T) {
				ParsableFactoryStruct := mocking.NewMockParsableFactory()
				factory := ParsableFactoryStruct.ParsableFactory
				backingStoreFactoryStruct := mocking.NewMockBackingStoreFactory()
				backingStore := mocking.NewMockBackingStore()
				backingStoreFactoryStruct.On("BackingStoreFactory").Return(backingStore)
				backingStoreFactory := backingStoreFactoryStruct.BackingStoreFactory

				response := NewServiceNowCollectionResponse(factory, backingStoreFactory)
				assert.Equal(t, backingStore, response.(*serviceNowCollectionResponse).backingStore)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

func TestServiceNowCollectionResponse_Serialization(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Successful",
			test: func(t *testing.T) {
				ParsableFactoryStruct := mocking.NewMockParsableFactory()
				factory := ParsableFactoryStruct.ParsableFactory
				backingStoreFactoryStruct := mocking.NewMockBackingStoreFactory()
				backingStore := mocking.NewMockBackingStore()
				backingStoreFactoryStruct.On("BackingStoreFactory").Return(backingStore)
				backingStoreFactory := backingStoreFactoryStruct.BackingStoreFactory
				collectionResponse := &serviceNowCollectionResponse{
					factory:             factory,
					backingStoreFactory: backingStoreFactory,
					backingStore:        backingStore,
				}
				serializationWriter := mocking.NewMockSerializationWriter()

				err := collectionResponse.Serialize(serializationWriter)
				assert.Equal(t, nil, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

func TestServiceNowCollectionResponse_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Successful",
			test: func(t *testing.T) {
				ParsableFactoryStruct := mocking.NewMockParsableFactory()
				factory := ParsableFactoryStruct.ParsableFactory
				backingStoreFactoryStruct := mocking.NewMockBackingStoreFactory()
				backingStore := mocking.NewMockBackingStore()
				backingStoreFactoryStruct.On("BackingStoreFactory").Return(backingStore)
				backingStoreFactory := backingStoreFactoryStruct.BackingStoreFactory
				collectionResponse := &serviceNowCollectionResponse{
					factory:             factory,
					backingStoreFactory: backingStoreFactory,
					backingStore:        backingStore,
				}
				_ = mocking.NewMockParseNode()

				_ = collectionResponse.GetFieldDeserializers()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

func TestServiceNowCollectionResponse_GetBackingStore(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Successful",
			test: func(t *testing.T) {
				ParsableFactoryStruct := mocking.NewMockParsableFactory()
				factory := ParsableFactoryStruct.ParsableFactory
				backingStoreFactoryStruct := mocking.NewMockBackingStoreFactory()
				backingStore := mocking.NewMockBackingStore()
				backingStoreFactoryStruct.On("BackingStoreFactory").Return(backingStore)
				backingStoreFactory := backingStoreFactoryStruct.BackingStoreFactory
				collectionResponse := &serviceNowCollectionResponse{
					factory:             factory,
					backingStoreFactory: backingStoreFactory,
					backingStore:        backingStore,
				}

				returnedStore := collectionResponse.GetBackingStore()
				assert.Equal(t, backingStore, returnedStore)
			},
		},
		{
			title: "No store provided",
			test: func(t *testing.T) {
				ParsableFactoryStruct := mocking.NewMockParsableFactory()
				factory := ParsableFactoryStruct.ParsableFactory
				backingStoreFactoryStruct := mocking.NewMockBackingStoreFactory()
				backingStore := mocking.NewMockBackingStore()
				backingStoreFactoryStruct.On("BackingStoreFactory").Return(backingStore)
				backingStoreFactory := backingStoreFactoryStruct.BackingStoreFactory
				collectionResponse := &serviceNowCollectionResponse{
					factory:             factory,
					backingStoreFactory: backingStoreFactory,
					backingStore:        nil,
				}

				returnedStore := collectionResponse.GetBackingStore()
				assert.NotNil(t, returnedStore)
				backingStoreFactoryStruct.AssertCalled(t, "BackingStoreFactory")
			},
		},
		{
			title: "Nil backing store factory",
			test: func(t *testing.T) {
				ParsableFactoryStruct := mocking.NewMockParsableFactory()
				factory := ParsableFactoryStruct.ParsableFactory
				collectionResponse := &serviceNowCollectionResponse{
					factory:             factory,
					backingStoreFactory: nil,
					backingStore:        nil,
				}

				store := collectionResponse.GetBackingStore()
				assert.Nil(t, store)
			},
		},
		{
			title: "Nil collection response",
			test: func(t *testing.T) {
				collectionResponse := (*serviceNowCollectionResponse)(nil)

				store := collectionResponse.GetBackingStore()
				assert.Nil(t, store)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

func TestServiceNowCollectionResponse_GetResult(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{}
}
