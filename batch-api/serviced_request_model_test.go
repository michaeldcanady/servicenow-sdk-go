package batchapi

import (
	"encoding/base64"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewServicedRequest(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				parsable := NewServicedRequest()

				assert.NotNil(t, parsable)
				assert.IsType(t, &ServicedRequestModel{}, parsable)

				assert.NotNil(t, parsable.Model)
				assert.IsType(t, &internal.BaseModel{}, parsable.Model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCreateServicedRequestFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				parseNode := mocking.NewMockParseNode()

				parsable, err := CreateServicedRequestFromDiscriminatorValue(parseNode)

				assert.Nil(t, err)
				assert.NotNil(t, parsable)
				assert.IsType(t, &ServicedRequestModel{}, parsable)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Not implemented",
			test: func(t *testing.T) {
				writer := mocking.NewMockSerializationWriter()
				intModel := mocking.NewMockModel()

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.Serialize(writer)

				assert.Equal(t, errors.New("Serialize not implemented"), err)
				intModel.AssertExpectations(t)
				writer.AssertExpectations(t)
			},
		},
		{
			name: "",
			test: func(t *testing.T) {
				writer := mocking.NewMockSerializationWriter()

				resp := (*ServicedRequestModel)(nil)

				err := resp.Serialize(writer)

				assert.Nil(t, err)
				writer.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				m := NewServicedRequest()
				deser := m.GetFieldDeserializers()
				assert.NotNil(t, deser)

				for key, fn := range deser {
					node := mocking.NewMockParseNode()
					s := "test"
					if key == bodyKey {
						s = base64.StdEncoding.EncodeToString([]byte("test"))
					}
					switch key {
					case statusCodeKey:
						i := int64(200)
						node.On("GetInt64Value").Return(&i, nil)
					case executionTimeKey:
						node.On("GetISODurationValue").Return(&serialization.ISODuration{}, nil)
					case headersKey:
						node.On("GetCollectionOfObjectValues", mock.Anything).Return([]serialization.Parsable{NewRestRequestHeader()}, nil)
					default:
						node.On("GetStringValue").Return(&s, nil)
					}
					_ = fn(node)
				}
			},
		},
		{
			name: "Nil Model",
			test: func(t *testing.T) {
				var m *ServicedRequestModel
				assert.Nil(t, m.GetFieldDeserializers())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_GetBodyAsParsable(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(m *ServicedRequestModel)
		expectedErr bool
	}{
		{
			name: "Successful",
			setup: func(m *ServicedRequestModel) {
				_ = m.setBody([]byte(`{"test":"test"}`))
				h := NewRestRequestHeader()
				n := "Content-Type"
				v := "application/json"
				_ = h.SetName(&n)
				_ = h.SetValue(&v)
				_ = m.setHeaders([]RestRequestHeader{h})
				_ = m.setStatusCode(internal.ToPointer(int64(200)))

				serialization.DefaultSerializationWriterFactoryInstance.ContentTypeAssociatedFactories["application/json"] = mocking.NewMockSerializationWriterFactory()
				serialization.DefaultParseNodeFactoryInstance.ContentTypeAssociatedFactories["application/json"] = mocking.NewMockParseNodeFactory()
			},
			expectedErr: false,
		},
		{
			name: "Error mapping",
			setup: func(m *ServicedRequestModel) {
				_ = m.setStatusCode(internal.ToPointer(int64(400)))
				_ = m.setBody([]byte(`{"error":"bad"}`))
			},
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := NewServicedRequest()
			test.setup(m)

			if test.name == "Successful" {
				factory := serialization.DefaultParseNodeFactoryInstance.ContentTypeAssociatedFactories["application/json"].(*mocking.MockParseNodeFactory)
				node := mocking.NewMockParseNode()
				node.On("GetObjectValue", mock.Anything).Return(mocking.NewMockParsable(), nil)
				factory.On("GetRootParseNode", "application/json", mock.Anything).Return(node, nil)
			}

			_, err := m.GetBodyAsParsable(func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
				return mocking.NewMockParsable(), nil
			})
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestServicedRequestModel_GetBody(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := []byte("test")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", bodyKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetBody()

				assert.Nil(t, err)
				assert.Equal(t, ret, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := internal.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", bodyKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetBody()

				assert.Equal(t, errors.New("body is not []byte"), err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", bodyKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetBody()

				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetBody()
				assert.Nil(t, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*ServicedRequestModel)(nil)

				id, err := resp.GetBody()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_setBody(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := []byte("test")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", bodyKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setBody(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := []byte("test")
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", bodyKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setBody(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := []byte("test")

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setBody(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := []byte("test")

				resp := (*ServicedRequestModel)(nil)

				err := resp.setBody(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_GetErrorMessage(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := internal.ToPointer("message")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", errorMessageKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetErrorMessage()

				assert.Nil(t, err)
				assert.Equal(t, ret, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := internal.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", errorMessageKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetErrorMessage()

				assert.Equal(t, errors.New("message is not *string"), err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", errorMessageKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetErrorMessage()

				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetErrorMessage()
				assert.Nil(t, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*ServicedRequestModel)(nil)

				id, err := resp.GetErrorMessage()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_setErrorMessage(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := internal.ToPointer("message")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", errorMessageKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setErrorMessage(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := internal.ToPointer("message")
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", errorMessageKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setErrorMessage(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := internal.ToPointer("message")

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setErrorMessage(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := internal.ToPointer("message")

				resp := (*ServicedRequestModel)(nil)

				err := resp.setErrorMessage(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_GetExecutionTime(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				m := NewServicedRequest()
				v := &serialization.ISODuration{}
				_ = m.setExecutionTime(v)
				res, err := m.GetExecutionTime()
				assert.NoError(t, err)
				assert.Equal(t, v, res)
			},
		},
		{
			name: "Nil Model",
			test: func(t *testing.T) {
				var m *ServicedRequestModel
				res, err := m.GetExecutionTime()
				assert.NoError(t, err)
				assert.Nil(t, res)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_setExecutionTime(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				m := NewServicedRequest()
				v := &serialization.ISODuration{}
				err := m.setExecutionTime(v)
				assert.NoError(t, err)
			},
		},
		{
			name: "Nil Model",
			test: func(t *testing.T) {
				var m *ServicedRequestModel
				err := m.setExecutionTime(nil)
				assert.NoError(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_GetHeaders(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := make([]RestRequestHeader, 0)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", headersKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetHeaders()

				assert.Nil(t, err)
				assert.Equal(t, ret, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := internal.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", headersKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetHeaders()

				assert.Equal(t, errors.New("headers is not []RestRequestHeader"), err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", headersKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetHeaders()

				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetHeaders()
				assert.Nil(t, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*ServicedRequestModel)(nil)

				id, err := resp.GetHeaders()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_setHeaders(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := make([]RestRequestHeader, 0)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", headersKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setHeaders(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := make([]RestRequestHeader, 0)
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", headersKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setHeaders(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := make([]RestRequestHeader, 0)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setHeaders(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := make([]RestRequestHeader, 0)

				resp := (*ServicedRequestModel)(nil)

				err := resp.setHeaders(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_GetID(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := internal.ToPointer("id")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", idKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetID()

				assert.Nil(t, err)
				assert.Equal(t, ret, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := internal.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", idKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetID()

				assert.Equal(t, errors.New("id is not *string"), err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", idKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetID()

				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetID()
				assert.Nil(t, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*ServicedRequestModel)(nil)

				id, err := resp.GetID()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_setID(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", idKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setID(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", idKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setID(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setID(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")

				resp := (*ServicedRequestModel)(nil)

				err := resp.setID(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_GetRedirectURL(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := internal.ToPointer("url")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", redirectURLKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetRedirectURL()

				assert.Nil(t, err)
				assert.Equal(t, ret, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := internal.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", redirectURLKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetRedirectURL()

				assert.Equal(t, errors.New("redirectURL is not *string"), err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", redirectURLKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetRedirectURL()

				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetRedirectURL()
				assert.Nil(t, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*ServicedRequestModel)(nil)

				id, err := resp.GetRedirectURL()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_setRedirectURL(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", redirectURLKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setRedirectURL(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", redirectURLKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setRedirectURL(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setRedirectURL(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")

				resp := (*ServicedRequestModel)(nil)

				err := resp.setRedirectURL(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_GetStatusCode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := internal.ToPointer("url")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", redirectURLKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetRedirectURL()

				assert.Nil(t, err)
				assert.Equal(t, ret, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := internal.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", statusCodeKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetStatusCode()

				assert.Equal(t, errors.New("statusCode is not *int64"), err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", statusCodeKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetStatusCode()

				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetStatusCode()
				assert.Nil(t, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*ServicedRequestModel)(nil)

				id, err := resp.GetStatusCode()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_setStatusCode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := internal.ToPointer(int64(0))
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", statusCodeKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setStatusCode(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := internal.ToPointer(int64(0))
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", statusCodeKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setStatusCode(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := internal.ToPointer(int64(0))

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setStatusCode(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := internal.ToPointer(int64(0))

				resp := (*ServicedRequestModel)(nil)

				err := resp.setStatusCode(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_GetStatusText(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := internal.ToPointer("url")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", statusTextKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetStatusText()

				assert.Nil(t, err)
				assert.Equal(t, ret, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := internal.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", statusTextKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetStatusText()

				assert.Equal(t, errors.New("statusText is not *string"), err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", statusTextKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetStatusText()

				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				id, err := resp.GetStatusText()
				assert.Nil(t, err)
				assert.Nil(t, id)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*ServicedRequestModel)(nil)

				id, err := resp.GetStatusText()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequestModel_setStatusText(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", statusTextKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setStatusText(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", statusTextKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setStatusText(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &ServicedRequestModel{
					intModel,
				}

				err := resp.setStatusText(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")

				resp := (*ServicedRequestModel)(nil)

				err := resp.setStatusText(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
