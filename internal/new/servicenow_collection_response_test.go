package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
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

				_ = NewServiceNowCollectionResponse[serialization.Parsable](factory, backingStoreFactory)
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
				sNResponse := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				sNResponse.On("GetFactory").Return(factory, nil)
				collectionResponse := &ServiceNowCollectionResponse[serialization.Parsable]{
					sNResponse,
				}
				_ = mocking.NewMockParseNode()
				// TODO: how to test map of functions??
				_ = collectionResponse.GetFieldDeserializers()
			},
		},
		{
			"Is Nil",
			func(t *testing.T) {
				collection := (*ServiceNowCollectionResponse[serialization.Parsable])(nil)
				assert.Nil(t, collection.GetFieldDeserializers())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

//nolint:dupl
func TestServiceNowCollectionResponse_GetNextLink(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Nil Collection",
			test: func(t *testing.T) {
				collection := (*ServiceNowCollectionResponse[serialization.Parsable])(nil)

				link, err := collection.GetNextLink()
				assert.Nil(t, link)
				assert.Nil(t, err)
			},
		},
		{
			title: "Successful",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				expResult := "I am a link"
				store.On("Get", nextLinkHeaderKey).Return(&expResult, nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)
				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetNextLink()
				assert.Equal(t, expResult, *link)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "Nil value",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Get", nextLinkHeaderKey).Return(nil, nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetNextLink()
				assert.Equal(t, (*string)(nil), link)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "not expected type",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Get", nextLinkHeaderKey).Return([]string{}, nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetNextLink()
				assert.Nil(t, link)
				assert.Equal(t, errors.New("val is not *string"), err)
				store.AssertExpectations(t)
			},
		},
		{
			title: "retrieval failed",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				expErr := errors.New("failed to get value")
				store.On("Get", nextLinkHeaderKey).Return(nil, expErr)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetNextLink()
				assert.Nil(t, link)
				assert.Equal(t, expErr, err)
				store.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

//nolint:dupl
func TestServiceNowCollectionResponse_setNextLink(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Successful",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Set", nextLinkHeaderKey, (*string)(nil)).Return(nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				err := collection.setNextLink(nil)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "Nil Collection",
			test: func(t *testing.T) {
				collection := (*ServiceNowCollectionResponse[serialization.Parsable])(nil)

				err := collection.setNextLink(nil)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

//nolint:dupl
func TestServiceNowCollectionResponse_GetPreviousLink(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Nil Collection",
			test: func(t *testing.T) {
				collection := (*ServiceNowCollectionResponse[serialization.Parsable])(nil)

				link, err := collection.GetPreviousLink()
				assert.Nil(t, link)
				assert.Nil(t, err)
			},
		},
		{
			title: "Successful",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				expResult := "I am a link"
				store.On("Get", prevLinkHeaderKey).Return(&expResult, nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetPreviousLink()
				assert.Equal(t, expResult, *link)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "Nil value",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Get", prevLinkHeaderKey).Return(nil, nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetPreviousLink()
				assert.Equal(t, (*string)(nil), link)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "not expected type",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Get", prevLinkHeaderKey).Return([]string{}, nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetPreviousLink()
				assert.Nil(t, link)
				assert.Equal(t, errors.New("val is not *string"), err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "retrieval failed",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				expErr := errors.New("failed to get value")
				store.On("Get", prevLinkHeaderKey).Return(nil, expErr)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetPreviousLink()
				assert.Nil(t, link)
				assert.Equal(t, expErr, err)
				store.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

//nolint:dupl
func TestServiceNowCollectionResponse_setPreviousLink(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Successful",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Set", prevLinkHeaderKey, (*string)(nil)).Return(nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				err := collection.setPreviousLink(nil)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "Nil Collection",
			test: func(t *testing.T) {
				collection := (*ServiceNowCollectionResponse[serialization.Parsable])(nil)

				err := collection.setPreviousLink(nil)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

//nolint:dupl
func TestServiceNowCollectionResponse_GetFirstLink(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Nil Collection",
			test: func(t *testing.T) {
				collection := (*ServiceNowCollectionResponse[serialization.Parsable])(nil)

				link, err := collection.GetFirstLink()
				assert.Nil(t, link)
				assert.Nil(t, err)
			},
		},
		{
			title: "Successful",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				expResult := "I am a link"
				store.On("Get", firstLinkHeaderKey).Return(&expResult, nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetFirstLink()
				assert.Equal(t, expResult, *link)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "Nil value",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Get", firstLinkHeaderKey).Return(nil, nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetFirstLink()
				assert.Equal(t, (*string)(nil), link)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "not expected type",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}
				store.On("Get", firstLinkHeaderKey).Return([]string{}, nil)

				link, err := collection.GetFirstLink()
				assert.Nil(t, link)
				assert.Equal(t, errors.New("val is not *string"), err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "retrieval failed",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				expErr := errors.New("failed to get value")
				store.On("Get", firstLinkHeaderKey).Return(nil, expErr)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetFirstLink()
				assert.Nil(t, link)
				assert.Equal(t, expErr, err)
				store.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

//nolint:dupl
func TestServiceNowCollectionResponse_setFirstLink(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Successful",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Set", firstLinkHeaderKey, (*string)(nil)).Return(nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				err := collection.setFirstLink(nil)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "Nil Collection",
			test: func(t *testing.T) {
				collection := (*ServiceNowCollectionResponse[serialization.Parsable])(nil)

				err := collection.setFirstLink(nil)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

//nolint:dupl
func TestServiceNowCollectionResponse_GetLastLink(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Nil Collection",
			test: func(t *testing.T) {
				collection := (*ServiceNowCollectionResponse[serialization.Parsable])(nil)

				link, err := collection.GetLastLink()
				assert.Nil(t, link)
				assert.Nil(t, err)
			},
		},
		{
			title: "Successful",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				expResult := "I am a link"
				store.On("Get", lastLinkHeaderKey).Return(&expResult, nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetLastLink()
				assert.Equal(t, expResult, *link)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "Nil value",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Get", lastLinkHeaderKey).Return(nil, nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetLastLink()
				assert.Equal(t, (*string)(nil), link)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "not expected type",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Get", lastLinkHeaderKey).Return([]string{}, nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetLastLink()
				assert.Nil(t, link)
				assert.Equal(t, errors.New("val is not *string"), err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "retrieval failed",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				expErr := errors.New("failed to get value")
				store.On("Get", lastLinkHeaderKey).Return(nil, expErr)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				link, err := collection.GetLastLink()
				assert.Nil(t, link)
				assert.Equal(t, expErr, err)
				store.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}

//nolint:dupl
func TestServiceNowCollectionResponse_setLastLink(t *testing.T) {
	tests := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "Successful",
			test: func(t *testing.T) {
				store := mocking.NewMockBackingStore()
				store.On("Set", lastLinkHeaderKey, (*string)(nil)).Return(nil)

				snResp := mocking.NewMockServiceNowResponse[[]serialization.Parsable]()
				snResp.On("GetBackingStore").Return(store)

				collection := &ServiceNowCollectionResponse[serialization.Parsable]{
					snResp,
				}

				err := collection.setLastLink(nil)
				assert.Nil(t, err)
				store.AssertExpectations(t)
				snResp.AssertExpectations(t)
			},
		},
		{
			title: "Nil Collection",
			test: func(t *testing.T) {
				collection := (*ServiceNowCollectionResponse[serialization.Parsable])(nil)

				err := collection.setLastLink(nil)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, test.test)
	}
}
