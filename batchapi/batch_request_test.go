package batchapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewBatchRequestModel(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				parsable := NewBatchRequestModel()

				assert.NotNil(t, parsable)
				assert.IsType(t, &BatchRequestModel{}, parsable)

				assert.NotNil(t, parsable.BaseModel)
				assert.IsType(t, &core.BaseModel{}, parsable.BaseModel)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCreateBatchRequestFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				parseNode := mocking.NewMockParseNode()

				parsable, err := CreateBatchRequestFromDiscriminatorValue(parseNode)

				require.NoError(t, err)
				assert.NotNil(t, parsable)
				assert.IsType(t, &BatchRequestModel{}, parsable)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchRequest_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				m := NewBatchRequestModel()
				req := newMockRestRequest()
				_ = m.AddRequest(req)

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				writer.On("WriteCollectionOfObjectValues", mock.Anything, mock.Anything).Return(nil)

				err := m.Serialize(writer)
				assert.NoError(t, err)
			},
		},
		{
			name: "Nil Model",
			test: func(t *testing.T) {
				var m *BatchRequestModel
				err := m.Serialize(nil)
				assert.NoError(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchRequest_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				m := NewBatchRequestModel()
				deser := m.GetFieldDeserializers()
				assert.NotNil(t, deser)
				assert.NotNil(t, deser[batchRequestIDKey])
				assert.NotNil(t, deser[restRequestsKey])

				// Test batchRequestIDKey
				nodeID := mocking.NewMockParseNode()
				s := "id"
				nodeID.On("GetStringValue").Return(&s, nil)
				err := deser[batchRequestIDKey](nodeID)
				require.NoError(t, err)

				// Test restRequestsKey
				nodeReqs := mocking.NewMockParseNode()
				nodeReqs.On("GetCollectionOfObjectValues", mock.Anything).Return([]serialization.Parsable{newMockRestRequest()}, nil)
				err = deser[restRequestsKey](nodeReqs)
				assert.NoError(t, err)
			},
		},
		{
			name: "Nil Model",
			test: func(t *testing.T) {
				var m *BatchRequestModel
				assert.Nil(t, m.GetFieldDeserializers())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchRequest_AddRequest(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()

				bs := core.NewBaseModel()
				bs.SetBackingStoreFactory(func() store.BackingStore { return backingStore })
				parsable := &BatchRequestModel{BaseModel: bs}
				req := newMockRestRequest()

				// Expectation that backingStore.Set is called
				backingStore.On("Set", restRequestsKey, []RestRequest{req}).Return(nil)
				backingStore.On("Get", restRequestsKey).Return([]RestRequest{}, nil)

				err := parsable.AddRequest(req)

				require.NoError(t, err)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				request := newMockRestRequest()

				parsable := (*BatchRequestModel)(nil)

				err := parsable.AddRequest(request)

				assert.NoError(t, err)
			},
		},
		{
			name: "Nil request",
			test: func(t *testing.T) {
				parsable := NewBatchRequestModel()

				err := parsable.AddRequest(nil)

				assert.NoError(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchRequest_GetBatchRequestID(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				m := NewBatchRequestModel()
				s := "id"
				_ = m.SetBatchRequestID(&s)
				res, err := m.GetBatchRequestID()
				require.NoError(t, err)
				assert.Equal(t, s, *res)
			},
		},
		{
			name: "Nil Model",
			test: func(t *testing.T) {
				var m *BatchRequestModel
				res, err := m.GetBatchRequestID()
				require.ErrorIs(t, err, snerrors.ErrNilModel)
				assert.Nil(t, res)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchRequest_SetBatchRequestID(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				m := NewBatchRequestModel()
				s := "id"
				err := m.SetBatchRequestID(&s)
				assert.NoError(t, err)
			},
		},
		{
			name: "Nil Model",
			test: func(t *testing.T) {
				var m *BatchRequestModel
				err := m.SetBatchRequestID(nil)
				require.ErrorIs(t, err, snerrors.ErrNilModel)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchRequest_GetRestRequests(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				m := NewBatchRequestModel()
				reqs := []RestRequest{newMockRestRequest()}
				_ = m.SetRestRequests(reqs)
				res, err := m.GetRestRequests()
				require.NoError(t, err)
				assert.Equal(t, reqs, res)
			},
		},
		{
			name: "Nil Model",
			test: func(t *testing.T) {
				var m *BatchRequestModel
				res, err := m.GetRestRequests()
				require.ErrorIs(t, err, snerrors.ErrNilModel)
				assert.Nil(t, res)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchRequest_SetRestRequests(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				m := NewBatchRequestModel()
				reqs := []RestRequest{newMockRestRequest()}
				err := m.SetRestRequests(reqs)
				assert.NoError(t, err)
			},
		},
		{
			name: "Nil Model",
			test: func(t *testing.T) {
				var m *BatchRequestModel
				err := m.SetRestRequests(nil)
				require.ErrorIs(t, err, snerrors.ErrNilModel)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
