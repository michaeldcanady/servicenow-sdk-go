package batchapi

import (
	"crypto/rand"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// TODO: add tests
func TestNewServicedRequest(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestCreateServicedRequestFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequest_Serialize(t *testing.T) {
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

// TODO: add tests
func TestServicedRequest_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestServicedRequest_GetBodyAsParsable(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequest_GetBody(t *testing.T) {
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

func TestServicedRequest_setBody(t *testing.T) {
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

func TestServicedRequest_GetErrorMessage(t *testing.T) {
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

func TestServicedRequest_setErrorMessage(t *testing.T) {
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

// TODO: add tests
func TestServicedRequest_GetExecutionTime(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestServicedRequest_setExecutionTime(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicedRequest_GetHeaders(t *testing.T) {
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

func TestServicedRequest_setHeaders(t *testing.T) {
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

func TestServicedRequest_GetID(t *testing.T) {
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

func TestServicedRequest_setID(t *testing.T) {
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

func TestServicedRequest_GetRedirectURL(t *testing.T) {
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

// TODO: add tests
func TestServicedRequest_setRedirectURL(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestServicedRequest_GetStatusCode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestServicedRequest_setStatusCode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestServicedRequest_GetStatusText(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestServicedRequest_setStatusText(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
