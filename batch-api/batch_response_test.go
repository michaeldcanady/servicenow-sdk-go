package batchapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewBatchResponse(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				parsable := NewBatchResponse()

				assert.NotNil(t, parsable)
				assert.IsType(t, &BatchResponseModel{}, parsable)

				assert.NotNil(t, parsable.Model)
				assert.IsType(t, &internal.BaseModel{}, parsable.Model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCreateBatchResponseFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				parseNode := mocking.NewMockParseNode()

				parsable, err := CreateBatchResponseFromDiscriminatorValue(parseNode)

				assert.Nil(t, err)
				assert.NotNil(t, parsable)
				assert.IsType(t, &BatchResponseModel{}, parsable)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchResponse_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				writer := mocking.NewMockSerializationWriter()

				resp := NewBatchResponse()
				err := resp.Serialize(writer)
				assert.NoError(t, err)
			},
		},
		{
			name: "Null model",
			test: func(t *testing.T) {
				writer := mocking.NewMockSerializationWriter()

				resp := (*BatchResponseModel)(nil)
				err := resp.Serialize(writer)
				assert.Equal(t, nil, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestBatchResponse_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				mockParseNode := mocking.NewMockParseNode()

				resp := &BatchResponseModel{
					intModel,
				}

				deserializers := resp.GetFieldDeserializers()

				tests := []struct {
					name string
					test func(*testing.T)
				}{
					{
						name: "batchRequestIDKey",
						test: func(t *testing.T) {
							ret := utils.ToPointer("id")
							mockParseNode.On("GetStringValue").Return(ret, nil)
							backingStore.On("Set", batchRequestIDKey, ret).Return(nil)

							err := deserializers[batchRequestIDKey](mockParseNode)

							assert.Nil(t, err)
						},
					},
					{
						name: "servicedRequestsKey",
						test: func(t *testing.T) {
							ret := make([]serialization.Parsable, 0)
							mockParseNode.On("GetCollectionOfObjectValues", mock.AnythingOfType("serialization.ParsableFactory")).Return(ret, nil)
							backingStore.On("Set", servicedRequestsKey, []ServicedRequest{}).Return(nil)

							err := deserializers[servicedRequestsKey](mockParseNode)

							assert.Nil(t, err)
						},
					},
					{
						name: "unservicedRequestsKey",
						test: func(t *testing.T) {
							ret := make([]interface{}, 0)
							mockParseNode.On("GetCollectionOfPrimitiveValues", "string").Return(ret, nil)
							backingStore.On("Set", unservicedRequestsKey, []string{}).Return(nil)

							err := deserializers[unservicedRequestsKey](mockParseNode)

							assert.Nil(t, err)
						},
					},
				}

				for _, test := range tests {
					t.Run(test.name, test.test)
				}
			},
		},
		{
			name: "Nil ParseNode",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				mockParseNode := mocking.NewMockParseNode()

				resp := &BatchResponseModel{
					intModel,
				}

				deserializers := resp.GetFieldDeserializers()

				tests := []struct {
					name string
					test func(*testing.T)
				}{
					{
						name: "batchRequestIDKey",
						test: func(t *testing.T) {
							mockParseNode.On("GetStringValue").Return((*string)(nil), nil)
							backingStore.On("Set", batchRequestIDKey, (*string)(nil)).Return(nil)

							err := deserializers[batchRequestIDKey](mockParseNode)

							assert.Nil(t, err)
						},
					},
					{
						name: "servicedRequestsKey",
						test: func(t *testing.T) {
							mockParseNode.On("GetCollectionOfObjectValues", mock.AnythingOfType("serialization.ParsableFactory")).Return(([]serialization.Parsable)(nil), nil)
							backingStore.On("Set", servicedRequestsKey, []ServicedRequest{}).Return(nil)

							err := deserializers[servicedRequestsKey](mockParseNode)

							assert.Nil(t, err)
						},
					},
					{
						name: "unservicedRequestsKey",
						test: func(t *testing.T) {
							mockParseNode.On("GetCollectionOfPrimitiveValues", "string").Return(([]interface{})(nil), nil)
							backingStore.On("Set", unservicedRequestsKey, []string{}).Return(nil)

							err := deserializers[unservicedRequestsKey](mockParseNode)

							assert.Nil(t, err)
						},
					},
				}

				for _, test := range tests {
					t.Run(test.name, test.test)
				}
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*BatchResponseModel)(nil)

				deserializers := resp.GetFieldDeserializers()

				assert.Nil(t, deserializers)
			},
		},
		{
			name: "Error retrieving value",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				mockParseNode := mocking.NewMockParseNode()

				resp := &BatchResponseModel{
					intModel,
				}

				deserializers := resp.GetFieldDeserializers()

				tests := []struct {
					name string
					test func(*testing.T)
				}{
					{
						name: "batchRequestIDKey",
						test: func(t *testing.T) {
							mockParseNode.On("GetStringValue").Return((*string)(nil), errors.New("failed to get value"))

							err := deserializers[batchRequestIDKey](mockParseNode)

							assert.Equal(t, errors.New("failed to get value"), err)
						},
					},
					{
						name: "servicedRequestsKey",
						test: func(t *testing.T) {
							mockParseNode.On("GetCollectionOfObjectValues", mock.AnythingOfType("serialization.ParsableFactory")).Return(([]serialization.Parsable)(nil), errors.New("failed to get value"))

							err := deserializers[servicedRequestsKey](mockParseNode)

							assert.Equal(t, errors.New("failed to get value"), err)
						},
					},
					{
						name: "unservicedRequestsKey",
						test: func(t *testing.T) {
							mockParseNode.On("GetCollectionOfPrimitiveValues", "string").Return(([]interface{})(nil), errors.New("failed to get value"))

							err := deserializers[unservicedRequestsKey](mockParseNode)

							assert.Equal(t, errors.New("failed to get value"), err)
						},
					},
				}

				for _, test := range tests {
					t.Run(test.name, test.test)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchResponse_GetBatchRequestID(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := utils.ToPointer("id")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", batchRequestIDKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetBatchRequestID()
				assert.Nil(t, err)
				assert.Equal(t, ret, id)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := utils.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", batchRequestIDKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetBatchRequestID()
				assert.Equal(t, errors.New("cannot convert 'true' to type *string"), err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", batchRequestIDKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetBatchRequestID()
				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetBatchRequestID()
				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*BatchResponseModel)(nil)

				id, err := resp.GetBatchRequestID()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchResponse_setBatchRequestID(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := utils.ToPointer("id")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", batchRequestIDKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				err := resp.setBatchRequestID(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := utils.ToPointer("id")
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", batchRequestIDKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				err := resp.setBatchRequestID(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := utils.ToPointer("id")

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &BatchResponseModel{
					intModel,
				}

				err := resp.setBatchRequestID(input)
				assert.Equal(t, errors.New("backingStore is nil"), err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := utils.ToPointer("id")

				resp := (*BatchResponseModel)(nil)

				err := resp.setBatchRequestID(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchResponse_GetServicedRequests(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := []ServicedRequest{}
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", servicedRequestsKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetServicedRequests()
				assert.Nil(t, err)
				assert.Equal(t, ret, id)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := utils.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", servicedRequestsKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetServicedRequests()
				assert.Equal(t, errors.New("cannot convert 'true' to type []batchapi.ServicedRequest"), err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", servicedRequestsKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetServicedRequests()
				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetServicedRequests()
				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*BatchResponseModel)(nil)

				id, err := resp.GetServicedRequests()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchResponse_setServicedRequests(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := []ServicedRequest{}
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", servicedRequestsKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				err := resp.setServicedRequests(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := []ServicedRequest{}
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", servicedRequestsKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				err := resp.setServicedRequests(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := []ServicedRequest{}

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &BatchResponseModel{
					intModel,
				}

				err := resp.setServicedRequests(input)
				assert.Equal(t, errors.New("backingStore is nil"), err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := []ServicedRequest{}

				resp := (*BatchResponseModel)(nil)

				err := resp.setServicedRequests(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchResponse_GetUnservicedRequests(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := []string{}
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", unservicedRequestsKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetUnservicedRequests()
				assert.Nil(t, err)
				assert.Equal(t, ret, id)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := utils.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", unservicedRequestsKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetUnservicedRequests()
				assert.Equal(t, errors.New("cannot convert 'true' to type []string"), err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", unservicedRequestsKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetUnservicedRequests()
				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &BatchResponseModel{
					intModel,
				}

				id, err := resp.GetUnservicedRequests()
				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*BatchResponseModel)(nil)

				id, err := resp.GetUnservicedRequests()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchResponse_setUnservicedRequests(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := []string{}
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", unservicedRequestsKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				err := resp.setUnservicedRequests(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := []string{}
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", unservicedRequestsKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &BatchResponseModel{
					intModel,
				}

				err := resp.setUnservicedRequests(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := []string{}

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &BatchResponseModel{
					intModel,
				}

				err := resp.setUnservicedRequests(input)
				assert.Equal(t, errors.New("backingStore is nil"), err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := []string{}

				resp := (*BatchResponseModel)(nil)

				err := resp.setUnservicedRequests(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchResponse_GetServicedRequestByID(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		setup       func(m *BatchResponseModel)
		expectedNil bool
	}{
		{
			name: "Found",
			id:   "1",
			setup: func(m *BatchResponseModel) {
				req := NewMockServicedRequest()
				req.On("GetID").Return(internal.ToPointer("1"), nil)
				_ = m.setServicedRequests([]ServicedRequest{req})
			},
			expectedNil: false,
		},
		{
			name: "Not Found",
			id:   "2",
			setup: func(m *BatchResponseModel) {
				req := NewMockServicedRequest()
				req.On("GetID").Return(internal.ToPointer("1"), nil)
				_ = m.setServicedRequests([]ServicedRequest{req})
			},
			expectedNil: true,
		},
		{
			name: "Nil Model",
			id:   "1",
			setup: func(m *BatchResponseModel) {
				// handled by test loop
			},
			expectedNil: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var m *BatchResponseModel
			if test.name != "Nil Model" {
				m = NewBatchResponse()
				test.setup(m)
			}

			res, err := m.GetServicedRequestByID(test.id)
			assert.NoError(t, err)
			if test.expectedNil {
				assert.Nil(t, res)
			} else {
				assert.NotNil(t, res)
			}
		})
	}
}
