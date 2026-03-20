package batchapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestNewRestRequestHeader(t *testing.T) {
	header := NewRestRequestHeader()
	assert.NotNil(t, header)
}

func TestRestRequestHeader_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				expName := utils.ToPointer("name")
				expValue := utils.ToPointer("value")

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteStringValue", nameKey, expName).Return(nil)
				writer.On("WriteStringValue", valueKey, expValue).Return(nil)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(expName, nil)
				backingStore.On("Get", valueKey).Return(expValue, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				err := header.Serialize(writer)

				assert.Nil(t, err)
				writer.AssertExpectations(t)
				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Name retrieval error",
			test: func(t *testing.T) {
				writer := mocking.NewMockSerializationWriter()

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(nil, errors.New("retrieval error"))

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				err := header.Serialize(writer)

				assert.Equal(t, errors.New("retrieval error"), err)
				writer.AssertExpectations(t)
				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Value retrieval error",
			test: func(t *testing.T) {
				expName := utils.ToPointer("name")

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteStringValue", nameKey, expName).Return(nil)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(expName, nil)
				backingStore.On("Get", valueKey).Return(nil, errors.New("retrieval error"))

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				err := header.Serialize(writer)

				assert.Equal(t, errors.New("retrieval error"), err)
				writer.AssertExpectations(t)
				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "nil writer",
			test: func(t *testing.T) {
				bR := NewRestRequestHeader()
				err := bR.Serialize(nil)
				assert.NoError(t, err)
			},
		},
		{
			name: "nil_model",
			test: func(t *testing.T) {
				var header *RestRequestHeaderModel

				err := header.Serialize(mocking.NewMockSerializationWriter())

				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCreateRestRequestHeaderFromHeaders(t *testing.T) {
	tests := []struct {
		name        string
		headers     *abstractions.RequestHeaders
		expectedErr bool
	}{
		{
			name: "Successful",
			headers: func() *abstractions.RequestHeaders {
				h := abstractions.NewRequestHeaders()
				h.Add("test", "value")
				return h
			}(),
			expectedErr: false,
		},
		{
			name:        "Nil Headers",
			headers:     nil,
			expectedErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := createRestRequestHeaderFromHeaders(test.headers)
			if test.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				if test.headers != nil {
					assert.NotEmpty(t, res)
				} else {
					assert.Empty(t, res)
				}
			}
		})
	}
}

func TestRestRequestHeader_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				header := NewRestRequestHeader()

				deser := header.GetFieldDeserializers()

				assert.NotNil(t, deser)
				assert.Len(t, deser, 2)
				assert.Contains(t, deser, nameKey)
				assert.Contains(t, deser, valueKey)
			},
		},
		{
			name: "Nil_model",
			test: func(t *testing.T) {
				header := (*RestRequestHeaderModel)(nil)

				deser := header.GetFieldDeserializers()

				assert.Nil(t, deser)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequestHeader_GetName(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				expName := utils.ToPointer("name")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(expName, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				name, err := header.GetName()

				assert.Nil(t, err)
				assert.Equal(t, expName, name)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				expName := utils.ToPointer(true)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(expName, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				name, err := header.GetName()

				assert.Equal(t, errors.New("cannot convert 'true' to type *string"), err)
				assert.Nil(t, name)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(nil, errors.New("error retrieving"))

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				name, err := header.GetName()

				assert.Equal(t, errors.New("error retrieving"), err)
				assert.Nil(t, name)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				backingStore := (*mocking.MockBackingStore)(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				name, err := header.GetName()

				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, name)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				header := (*RestRequestHeaderModel)(nil)

				name, err := header.GetName()

				assert.Nil(t, err)
				assert.Nil(t, name)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequestHeader_SetName(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				input := utils.ToPointer("name")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", nameKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				err := header.SetName(input)

				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Store error",
			test: func(t *testing.T) {
				input := utils.ToPointer("name")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", nameKey, input).Return(errors.New("store error"))

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				err := header.SetName(input)

				assert.Equal(t, errors.New("store error"), err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				input := utils.ToPointer("name")

				backingStore := (*mocking.MockBackingStore)(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				err := header.SetName(input)

				assert.Equal(t, errors.New("backingStore is nil"), err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				header := (*RestRequestHeaderModel)(nil)

				err := header.SetName(utils.ToPointer("name"))

				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequestHeader_GetValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				expValue := utils.ToPointer("value")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", valueKey).Return(expValue, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				value, err := header.GetValue()

				assert.Nil(t, err)
				assert.Equal(t, expValue, value)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				expValue := utils.ToPointer(true)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", valueKey).Return(expValue, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				value, err := header.GetValue()

				assert.Equal(t, errors.New("cannot convert 'true' to type *string"), err)
				assert.Nil(t, value)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", valueKey).Return(nil, errors.New("error retrieving"))

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				value, err := header.GetValue()

				assert.Equal(t, errors.New("error retrieving"), err)
				assert.Nil(t, value)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				backingStore := (*mocking.MockBackingStore)(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				value, err := header.GetValue()

				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, value)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				header := (*RestRequestHeaderModel)(nil)

				value, err := header.GetValue()

				assert.Nil(t, err)
				assert.Nil(t, value)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRestRequestHeader_SetValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				input := utils.ToPointer("value")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", valueKey, input).Return(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				err := header.SetValue(input)

				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Store error",
			test: func(t *testing.T) {
				input := utils.ToPointer("value")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", valueKey, input).Return(errors.New("store error"))

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				err := header.SetValue(input)

				assert.Equal(t, errors.New("store error"), err)

				backingStore.AssertExpectations(t)
				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				input := utils.ToPointer("value")

				backingStore := (*mocking.MockBackingStore)(nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{intModel}

				err := header.SetValue(input)

				assert.Equal(t, errors.New("backingStore is nil"), err)

				intModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				header := (*RestRequestHeaderModel)(nil)

				err := header.SetValue(utils.ToPointer("value"))

				assert.Nil(t, err)
			},
		}}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
