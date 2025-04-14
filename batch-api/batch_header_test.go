package batchapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestNewBatchHeader(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				header := NewBatchHeader()

				assert.NotNil(t, header)
				assert.IsType(t, &BatchHeader{}, header)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCreateBatchHeader2FromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "with parse node",
			test: func(t *testing.T) {
				parseNode := mocking.NewMockParseNode()

				parsable, err := CreateBatchHeader2FromDiscriminatorValue(parseNode)
				assert.Nil(t, err)
				assert.NotNil(t, parsable)
				assert.IsType(t, &BatchHeader{}, parsable)
			},
		},
		{
			name: "with nil parse node",
			test: func(t *testing.T) {
				parsable, err := CreateBatchHeader2FromDiscriminatorValue(nil)
				assert.Nil(t, err)
				assert.NotNil(t, parsable)
				assert.IsType(t, &BatchHeader{}, parsable)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestBatchHeader_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		//{
		//	name: "Successful",
		//	test: func(t *testing.T) {
		//		writer := mocking.NewMockSerializationWriter()

		//		model := mocking.NewMockModel()

		//		header := &BatchHeader{model}

		//		err := header.Serialize(writer)

		//		assert.Nil(t, err)
		//	},
		//},
		{
			name: "nil writer",
			test: func(t *testing.T) {
				writer := (*mocking.MockSerializationWriter)(nil)

				model := mocking.NewMockModel()

				header := &BatchHeader{model}

				err := header.Serialize(writer)

				assert.Equal(t, errors.New("write is nil"), err)
			},
		},
		{
			name: "nil writer",
			test: func(t *testing.T) {
				writer := (*mocking.MockSerializationWriter)(nil)

				header := (*BatchHeader)(nil)

				err := header.Serialize(writer)

				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestBatchHeader_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchHeader_GetName(t *testing.T) {
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

				header := &BatchHeader{model}

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

				header := &BatchHeader{model}

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

				header := &BatchHeader{model}

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

				header := &BatchHeader{model}

				name, err := header.GetName()

				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, name)

				model.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				header := (*BatchHeader)(nil)

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

func TestBatchHeader_SetName(t *testing.T) {
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

				header := &BatchHeader{model}

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

				header := &BatchHeader{model}

				err := header.SetName(expName)

				assert.Equal(t, errors.New("backingStore is nil"), err)

				model.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				expName := internal.ToPointer("name")

				header := (*BatchHeader)(nil)

				err := header.SetName(expName)

				assert.Nil(t, err)

			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchHeader_GetValue(t *testing.T) {
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

				header := &BatchHeader{model}

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

				header := &BatchHeader{model}

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

				header := &BatchHeader{model}

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

				header := &BatchHeader{model}

				name, err := header.GetValue()

				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, name)

				model.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				header := (*BatchHeader)(nil)

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

func TestBatchHeader_SetValue(t *testing.T) {
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

				header := &BatchHeader{model}

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

				header := &BatchHeader{model}

				err := header.SetValue(expName)

				assert.Equal(t, errors.New("backingStore is nil"), err)

				model.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				expName := internal.ToPointer("value")

				header := (*BatchHeader)(nil)

				err := header.SetValue(expName)

				assert.Nil(t, err)

			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
