package tableapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
)

func TestServiceNowResponse_Serialization(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Successful",
			test: func(t *testing.T) {
				collectionResponse := &ServiceNowResponseImpl[any]{}
				serializationWriter := mocking.NewMockSerializationWriter()

				err := collectionResponse.Serialize(serializationWriter)
				assert.Equal(t, errors.New("doesn't support serialization"), err)
			},
		},
		{
			"Nil collection",
			func(t *testing.T) {
				collection := (*ServiceNowResponseImpl[any])(nil)
				serializationWriter := mocking.NewMockSerializationWriter()
				assert.Equal(t, nil, collection.Serialize(serializationWriter))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

func TestServiceNowResponse_GetBackingStore(t *testing.T) {
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
				collectionResponse := &ServiceNowResponseImpl[any]{
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
				collectionResponse := &ServiceNowResponseImpl[any]{
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
				collectionResponse := &ServiceNowResponseImpl[any]{
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
				collectionResponse := (*ServiceNowResponseImpl[any])(nil)

				store := collectionResponse.GetBackingStore()
				assert.Nil(t, store)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

func TestServiceNowResponse_GetResult(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Nil Collection",
			test: func(t *testing.T) {
				collection := (*ServiceNowResponseImpl[any])(nil)
				res, err := collection.GetResult()
				assert.Nil(t, res)
				assert.Nil(t, err)
			},
		},
		{
			title: "Successful",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				expResult := []serialization.Parsable{}
				store.On("Get", resultKey).Return(expResult, nil)

				collection := &ServiceNowResponseImpl[any]{
					backingStore: store,
				}

				res, err := collection.GetResult()
				assert.Equal(t, expResult, res)
				assert.Nil(t, err)
				store.AssertExpectations(t)
			},
		},
		{
			title: "missing result",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Get", resultKey).Return(nil, nil)

				collection := &ServiceNowResponseImpl[any]{
					backingStore: store,
				}

				res, err := collection.GetResult()
				assert.Equal(t, nil, res)
				assert.Nil(t, err)
				store.AssertExpectations(t)
			},
		},
		{
			title: "get store error",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				expErr := errors.New("empty store")
				store.On("Get", resultKey).Return(nil, expErr)

				collection := &ServiceNowResponseImpl[any]{
					backingStore: store,
				}

				res, err := collection.GetResult()
				assert.Equal(t, nil, res)
				assert.Equal(t, err, expErr)
				store.AssertExpectations(t)
			},
		},
		{
			title: "wrong type",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				expErr := errors.New("val is not *string")
				store.On("Get", resultKey).Return(true, nil)

				collection := &ServiceNowResponseImpl[*string]{
					backingStore: store,
				}

				res, err := collection.GetResult()
				assert.Equal(t, (*string)(nil), res)
				assert.Equal(t, err, expErr)
				store.AssertExpectations(t)
			},
		},
		// TODO: needs 'not expected type'
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

func TestServiceNowResponse_SetResult(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Nil Collection",
			test: func(t *testing.T) {
				collection := (*ServiceNowResponseImpl[any])(nil)
				err := collection.SetResult(nil)
				assert.Nil(t, err)
			},
		},
		{
			title: "Successful",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Set", resultKey, nil).Return(nil)

				collection := &ServiceNowResponseImpl[any]{
					backingStore: store,
				}

				err := collection.SetResult(nil)
				assert.Nil(t, err)
				store.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}
