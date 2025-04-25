package batchapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewRestRequest(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				parsable := NewRestRequest()

				assert.NotNil(t, parsable)
				assert.IsType(t, &RestRequestModel{}, parsable)

				assert.NotNil(t, parsable.Model)
				assert.IsType(t, &internal.BaseModel{}, parsable.Model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCreateRestRequestFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				parseNode := mocking.NewMockParseNode()

				parsable, err := CreateRestRequestFromDiscriminatorValue(parseNode)

				assert.Nil(t, err)
				assert.NotNil(t, parsable)
				assert.IsType(t, &RestRequestModel{}, parsable)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestCreateRestRequestFromRequestInformation(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestRestRequest_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequest_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				model := &RestRequestModel{}

				deserializers := model.GetFieldDeserializers()

				assert.Nil(t, deserializers)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*RestRequestModel)(nil)

				deserializers := model.GetFieldDeserializers()

				assert.Nil(t, deserializers)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequest_GetBody(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := []byte{}
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", bodyKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetBody()
				assert.Nil(t, err)
				assert.Equal(t, ret, id)
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

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetBody()
				assert.Equal(t, errors.New("body is not []byte"), err)
				assert.Nil(t, id)
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

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetBody()
				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetBody()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*RestRequestModel)(nil)

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

// TODO: add tests
func TestRestRequest_SetBodyFromParsable(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequest_SetBody(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := []byte{}
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", bodyKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetBody(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := []byte{}
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", bodyKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetBody(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := []byte{}

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetBody(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := []byte{}

				resp := (*RestRequestModel)(nil)

				err := resp.SetBody(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequest_GetExcludeResponseHeaders(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := internal.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", excludeResponseHeadersKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetExcludeResponseHeaders()
				assert.Nil(t, err)
				assert.Equal(t, ret, id)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := internal.ToPointer("test")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", excludeResponseHeadersKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetExcludeResponseHeaders()
				assert.Equal(t, errors.New("excludeResponseHeaders is not *bool"), err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", excludeResponseHeadersKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetExcludeResponseHeaders()
				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetExcludeResponseHeaders()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*RestRequestModel)(nil)

				id, err := resp.GetExcludeResponseHeaders()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequest_SetExcludeResponseHeaders(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := internal.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", excludeResponseHeadersKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetExcludeResponseHeaders(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := internal.ToPointer(true)
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", excludeResponseHeadersKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetExcludeResponseHeaders(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := internal.ToPointer(true)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetExcludeResponseHeaders(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := internal.ToPointer(true)

				resp := (*RestRequestModel)(nil)

				err := resp.SetExcludeResponseHeaders(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequest_GetHeaders(t *testing.T) {
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

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetHeaders()
				assert.Nil(t, err)
				assert.Equal(t, ret, id)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := internal.ToPointer("test")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", headersKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetHeaders()
				assert.Equal(t, errors.New("headers is not []RestRequestHeader"), err)
				assert.Nil(t, id)
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

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetHeaders()
				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetHeaders()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*RestRequestModel)(nil)

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

func TestRestRequest_SetHeaders(t *testing.T) {
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

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetHeaders(input)
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

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetHeaders(input)
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

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetHeaders(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := make([]RestRequestHeader, 0)

				resp := (*RestRequestModel)(nil)

				err := resp.SetHeaders(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequest_GetID(t *testing.T) {
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

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetID()
				assert.Nil(t, err)
				assert.Equal(t, ret, id)
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

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetID()
				assert.Equal(t, errors.New("id is not *string"), err)
				assert.Nil(t, id)
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

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetID()
				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetID()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*RestRequestModel)(nil)

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

func TestRestRequest_SetID(t *testing.T) {
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

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetID(input)
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

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetID(input)
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

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetID(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := internal.ToPointer("id")

				resp := (*RestRequestModel)(nil)

				err := resp.SetID(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequest_GetMethod(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := internal.ToPointer(abstractions.GET)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", methodKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetMethod()
				assert.Nil(t, err)
				assert.Equal(t, ret, id)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := internal.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", methodKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetMethod()
				assert.Equal(t, errors.New("method is not *abstractions.HttpMethod"), err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", methodKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetMethod()
				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetMethod()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*RestRequestModel)(nil)

				id, err := resp.GetMethod()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequest_SetMethod(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := internal.ToPointer(abstractions.GET)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", methodKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetMethod(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := internal.ToPointer(abstractions.GET)
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", methodKey, input).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetMethod(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := internal.ToPointer(abstractions.GET)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetMethod(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := internal.ToPointer(abstractions.GET)

				resp := (*RestRequestModel)(nil)

				err := resp.SetMethod(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequest_GetURL(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				ret := internal.ToPointer("url")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", urlKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetURL()
				assert.Nil(t, err)
				assert.Equal(t, ret, id)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				ret := internal.ToPointer(true)
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", urlKey).Return(ret, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetURL()
				assert.Equal(t, errors.New("url is not *string"), err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Error getting value",
			test: func(t *testing.T) {
				retErr := errors.New("failed to retrieve value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", urlKey).Return(nil, retErr)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetURL()
				assert.Equal(t, retErr, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &RestRequestModel{
					intModel,
				}

				id, err := resp.GetURL()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				resp := (*RestRequestModel)(nil)

				id, err := resp.GetURL()
				assert.Nil(t, err)
				assert.Nil(t, id)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequest_SetURL(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				input := internal.ToPointer("https://service-now.com/api/now/table")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", urlKey, mock.AnythingOfType("*string")).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetURL(input)
				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Path doesn't start with /api",
			test: func(t *testing.T) {
				input := internal.ToPointer("https://service-now.com/now/table")

				intModel := mocking.NewMockModel()

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetURL(input)
				assert.Equal(t, errors.New("invalid URL: path doesn't begin with \"/api\""), err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Error setting object",
			test: func(t *testing.T) {
				input := internal.ToPointer("https://service-now.com/api/now/table")
				ret := errors.New("failed to set value")
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", urlKey, mock.AnythingOfType("*string")).Return(ret)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetURL(input)
				assert.Equal(t, ret, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				input := internal.ToPointer("https://service-now.com/api/now/table")

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				resp := &RestRequestModel{
					intModel,
				}

				err := resp.SetURL(input)
				assert.Nil(t, err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				input := internal.ToPointer("https://service-now.com/api/now/table")

				resp := (*RestRequestModel)(nil)

				err := resp.SetURL(input)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
