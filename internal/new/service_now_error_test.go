package internal

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

// TODO: add tests
func TestNewServicenowError(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestCreateServiceNowErrorFromDiscriminatorValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				model, err := CreateServiceNowErrorFromDiscriminatorValue(nil)
				assert.Nil(t, err)
				assert.IsType(t, &ServicenowError{}, model)
				assert.NotNil(t, model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicenowError_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				model := &ServicenowError{}

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
func TestServicenowError_GetFieldDeserializers(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicenowError_GetError(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mainError := mocking.NewMockMainError()

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", errorKey).Return(mainError, nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &ServicenowError{
					innerModel,
				}

				apiErr, err := model.GetError()
				assert.Nil(t, err)
				assert.Equal(t, mainError, apiErr)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				mainError := "mocking.NewMockMainError()"

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", errorKey).Return(mainError, nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &ServicenowError{
					innerModel,
				}

				apiErr, err := model.GetError()
				assert.Equal(t, errors.New("rawMainErr is not MainErrorable"), err)
				assert.Nil(t, apiErr)
			},
		},
		{
			name: "Error backingStore.Get",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", errorKey).Return(nil, errors.New("unable to retrieve"))

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &ServicenowError{
					innerModel,
				}

				apiErr, err := model.GetError()
				assert.Equal(t, errors.New("unable to retrieve"), err)
				assert.Nil(t, apiErr)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*ServicenowError)(nil)

				apiErr, err := model.GetError()
				assert.Nil(t, err)
				assert.Nil(t, apiErr)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestServicenowError_setError(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mainError := mocking.NewMockMainError()

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", errorKey, mainError).Return(nil)

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return(backingStore)

				model := &ServicenowError{
					innerModel,
				}

				err := model.setError(mainError)
				assert.Nil(t, err)
				backingStore.AssertExpectations(t)
				innerModel.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				mainError := mocking.NewMockMainError()

				innerModel := mocking.NewMockModel()
				innerModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				model := &ServicenowError{
					innerModel,
				}

				err := model.setError(mainError)
				assert.Equal(t, errors.New("backingStore is nil"), err)
				innerModel.AssertExpectations(t)
			},
		},
		{
			name: "nil model",
			test: func(t *testing.T) {
				mainError := mocking.NewMockMainError()

				model := (*ServicenowError)(nil)

				err := model.setError(mainError)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
