package batchapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestNewRestRequestHeader(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				header := NewRestRequestHeader()

				assert.NotNil(t, header)
				assert.IsType(t, &RestRequestHeaderModel{}, header)

				assert.NotNil(t, header.Model)
				assert.IsType(t, &internal.BaseModel{}, header.Model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCreateRestRequestHeaderFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "with parse node",
			test: func(t *testing.T) {
				parseNode := mocking.NewMockParseNode()

				parsable, err := CreateRestRequestHeaderFromDiscriminatorValue(parseNode)
				assert.Nil(t, err)
				assert.NotNil(t, parsable)
				assert.IsType(t, &RestRequestHeaderModel{}, parsable)
			},
		},
		{
			name: "with nil parse node",
			test: func(t *testing.T) {
				parsable, err := CreateRestRequestHeaderFromDiscriminatorValue(nil)
				assert.Nil(t, err)
				assert.NotNil(t, parsable)
				assert.IsType(t, &RestRequestHeaderModel{}, parsable)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestRestRequestHeader_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		//{
		//	name: "Successful",
		//	test: func(t *testing.T) {
		//		writer := mocking.NewMockSerializationWriter()

		//		model := mocking.NewMockModel()

		//		header := &RestRequestHeader{model}

		//		err := header.Serialize(writer)

		//		assert.Nil(t, err)
		//	},
		//},
		{
			name: "nil writer",
			test: func(t *testing.T) {
				writer := (*mocking.MockSerializationWriter)(nil)

				model := mocking.NewMockModel()

				header := &RestRequestHeaderModel{model}

				err := header.Serialize(writer)

				assert.Equal(t, errors.New("write is nil"), err)
			},
		},
		{
			name: "nil writer",
			test: func(t *testing.T) {
				writer := (*mocking.MockSerializationWriter)(nil)

				header := (*RestRequestHeaderModel)(nil)

				err := header.Serialize(writer)

				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
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
				model := &RestRequestHeaderModel{}

				deserializers := model.GetFieldDeserializers()

				assert.Nil(t, deserializers)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*RestRequestHeaderModel)(nil)

				deserializers := model.GetFieldDeserializers()

				assert.Nil(t, deserializers)
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
			name: "Successful",
			test: func(t *testing.T) {
				expName := internal.ToPointer("name")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(expName, nil)

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{model}

				name, err := header.GetName()

				assert.Nil(t, err)
				assert.Equal(t, expName, name)

				backingStore.AssertExpectations(t)
				model.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				expName := internal.ToPointer(true)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(expName, nil)

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{model}

				name, err := header.GetName()

				assert.Equal(t, errors.New("name is not *string"), err)
				assert.Nil(t, name)

				backingStore.AssertExpectations(t)
				model.AssertExpectations(t)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(nil, errors.New("error retrieving"))

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{model}

				name, err := header.GetName()

				assert.Equal(t, errors.New("error retrieving"), err)
				assert.Nil(t, name)

				backingStore.AssertExpectations(t)
				model.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				backingStore := (*mocking.MockBackingStore)(nil)

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{model}

				name, err := header.GetName()

				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, name)

				model.AssertExpectations(t)
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
			name: "Successful",
			test: func(t *testing.T) {
				expName := internal.ToPointer("name")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", nameKey, expName).Return(nil)

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{model}

				err := header.SetName(expName)

				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				model.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				expName := internal.ToPointer("name")

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				header := &RestRequestHeaderModel{model}

				err := header.SetName(expName)

				assert.Equal(t, errors.New("backingStore is nil"), err)

				model.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				expName := internal.ToPointer("name")

				header := (*RestRequestHeaderModel)(nil)

				err := header.SetName(expName)

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
			name: "Successful",
			test: func(t *testing.T) {
				expName := internal.ToPointer("value")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", valueKey).Return(expName, nil)

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{model}

				name, err := header.GetValue()

				assert.Nil(t, err)
				assert.Equal(t, expName, name)

				backingStore.AssertExpectations(t)
				model.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				expName := internal.ToPointer(true)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", valueKey).Return(expName, nil)

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{model}

				name, err := header.GetValue()

				assert.Equal(t, errors.New("value is not *string"), err)
				assert.Nil(t, name)

				backingStore.AssertExpectations(t)
				model.AssertExpectations(t)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", valueKey).Return(nil, errors.New("error retrieving"))

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{model}

				name, err := header.GetValue()

				assert.Equal(t, errors.New("error retrieving"), err)
				assert.Nil(t, name)

				backingStore.AssertExpectations(t)
				model.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				backingStore := (*mocking.MockBackingStore)(nil)

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{model}

				name, err := header.GetValue()

				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, name)

				model.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				header := (*RestRequestHeaderModel)(nil)

				name, err := header.GetValue()

				assert.Nil(t, err)
				assert.Nil(t, name)
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
			name: "Successful",
			test: func(t *testing.T) {
				expName := internal.ToPointer("value")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", valueKey, expName).Return(nil)

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return(backingStore)

				header := &RestRequestHeaderModel{model}

				err := header.SetValue(expName)

				assert.Nil(t, err)

				backingStore.AssertExpectations(t)
				model.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				expName := internal.ToPointer("value")

				model := mocking.NewMockModel()
				model.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				header := &RestRequestHeaderModel{model}

				err := header.SetValue(expName)

				assert.Equal(t, errors.New("backingStore is nil"), err)

				model.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				expName := internal.ToPointer("value")

				header := (*RestRequestHeaderModel)(nil)

				err := header.SetValue(expName)

				assert.Nil(t, err)

			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
