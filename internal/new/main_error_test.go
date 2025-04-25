package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewMainError(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				header := NewMainError()

				assert.NotNil(t, header)
				assert.IsType(t, &MainError{}, header)

				assert.NotNil(t, header.Model)
				assert.IsType(t, &BaseModel{}, header.Model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCreateMainErrorFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "with parse node",
			test: func(t *testing.T) {
				parseNode := mocking.NewMockParseNode()

				parsable, err := CreateMainErrorFromDiscriminatorValue(parseNode)
				assert.Nil(t, err)
				assert.NotNil(t, parsable)
				assert.IsType(t, &MainError{}, parsable)
			},
		},
		{
			name: "with nil parse node",
			test: func(t *testing.T) {
				parsable, err := CreateMainErrorFromDiscriminatorValue(nil)
				assert.Nil(t, err)
				assert.NotNil(t, parsable)
				assert.IsType(t, &MainError{}, parsable)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestMainError_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				model := &MainError{}

				err := model.Serialize(nil)
				assert.Error(t, errors.New("unsupported"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestMainError_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestMainError_GetDetail(t *testing.T) { // nolint: dupl
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				actualDetails := ToPointer("detailed details")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", detailKey).Return(actualDetails, nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				details, err := model.GetDetail()
				assert.Nil(t, err)
				assert.Equal(t, actualDetails, details)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				actualDetails := ToPointer(true)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", detailKey).Return(actualDetails, nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				details, err := model.GetDetail()
				assert.Equal(t, errors.New("detail is not *string"), err)
				assert.Nil(t, details)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				model := &MainError{
					innerModel,
				}

				details, err := model.GetDetail()
				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, details)
			},
		},
		{
			name: "Error BackingStore.Get",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", detailKey).Return((*string)(nil), errors.New("failed to get"))

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				details, err := model.GetDetail()
				assert.Equal(t, errors.New("failed to get"), err)
				assert.Nil(t, details)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*MainError)(nil)

				details, err := model.GetDetail()
				assert.Nil(t, err)
				assert.Nil(t, details)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestMainError_setDetail(t *testing.T) { // nolint: dupl
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				expVal := ToPointer("detailed details")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", detailKey, expVal).Return(nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				err := model.setDetail(expVal)
				assert.Nil(t, err)
				backingStore.AssertExpectations(t)
				innerModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				expVal := ToPointer("detailed details")

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				model := &MainError{
					innerModel,
				}

				err := model.setDetail(expVal)
				assert.Equal(t, errors.New("backingStore is nil"), err)
				innerModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				expVal := ToPointer("detailed details")

				model := (*MainError)(nil)

				err := model.setDetail(expVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestMainError_GetMessage(t *testing.T) { // nolint: dupl
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				actualDetails := ToPointer("detailed details")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", messageKey).Return(actualDetails, nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				details, err := model.GetMessage()
				assert.Nil(t, err)
				assert.Equal(t, actualDetails, details)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				actualDetails := ToPointer(true)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", messageKey).Return(actualDetails, nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				details, err := model.GetMessage()
				assert.Equal(t, errors.New("message is not *string"), err)
				assert.Nil(t, details)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				model := &MainError{
					innerModel,
				}

				details, err := model.GetMessage()
				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, details)
			},
		},
		{
			name: "Error BackingStore.Get",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", messageKey).Return((*string)(nil), errors.New("failed to get"))

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				details, err := model.GetMessage()
				assert.Equal(t, errors.New("failed to get"), err)
				assert.Nil(t, details)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*MainError)(nil)

				details, err := model.GetMessage()
				assert.Nil(t, err)
				assert.Nil(t, details)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestMainError_setMessage(t *testing.T) { // nolint: dupl
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				expVal := ToPointer("detailed details")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", messageKey, expVal).Return(nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				err := model.setMessage(expVal)
				assert.Nil(t, err)
				backingStore.AssertExpectations(t)
				innerModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				expVal := ToPointer("detailed details")

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				model := &MainError{
					innerModel,
				}

				err := model.setMessage(expVal)
				assert.Equal(t, errors.New("backingStore is nil"), err)
				innerModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				expVal := ToPointer("detailed details")

				model := (*MainError)(nil)

				err := model.setMessage(expVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestMainError_GetStatus(t *testing.T) { // nolint: dupl
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				actualDetails := ToPointer("detailed details")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", statusKey).Return(actualDetails, nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				details, err := model.GetStatus()
				assert.Nil(t, err)
				assert.Equal(t, actualDetails, details)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				actualDetails := ToPointer(true)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", statusKey).Return(actualDetails, nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				details, err := model.GetStatus()
				assert.Equal(t, errors.New("status is not *string"), err)
				assert.Nil(t, details)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				model := &MainError{
					innerModel,
				}

				details, err := model.GetStatus()
				assert.Equal(t, errors.New("backingStore is nil"), err)
				assert.Nil(t, details)
			},
		},
		{
			name: "Error BackingStore.Get",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", statusKey).Return((*string)(nil), errors.New("failed to get"))

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				details, err := model.GetStatus()
				assert.Equal(t, errors.New("failed to get"), err)
				assert.Nil(t, details)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*MainError)(nil)

				details, err := model.GetStatus()
				assert.Nil(t, err)
				assert.Nil(t, details)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestMainError_setStatus(t *testing.T) { // nolint: dupl
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				expVal := ToPointer("detailed details")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", statusKey, expVal).Return(nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &MainError{
					innerModel,
				}

				err := model.setStatus(expVal)
				assert.Nil(t, err)
				backingStore.AssertExpectations(t)
				innerModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backingStore",
			test: func(t *testing.T) {
				expVal := ToPointer("detailed details")

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				model := &MainError{
					innerModel,
				}

				err := model.setStatus(expVal)
				assert.Equal(t, errors.New("backingStore is nil"), err)
				innerModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				expVal := ToPointer("detailed details")

				model := (*MainError)(nil)

				err := model.setStatus(expVal)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
