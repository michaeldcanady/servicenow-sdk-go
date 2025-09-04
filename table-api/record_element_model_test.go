//go:build preview.tableApiV2

package tableapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	internal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestNewRecordElement(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				model := NewRecordElement()

				assert.NotNil(t, model)
				assert.IsType(t, &RecordElement{}, model)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRecordElementModel_GetDisplayValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				value := internal.ToPointer("")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", displayValueKey).Return(value, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				model := &RecordElement{intModel}

				elementValue, err := model.GetDisplayValue()

				assert.Nil(t, err)
				assert.NotNil(t, elementValue)
				assert.IsType(t, &ElementValue{}, elementValue)
				assert.Equal(t, interface{}(value), elementValue.val)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", displayValueKey).Return(nil, errors.New("retrieval error"))

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				model := &RecordElement{intModel}

				elementValue, err := model.GetDisplayValue()

				assert.Equal(t, errors.New("retrieval error"), err)
				assert.Nil(t, elementValue)
			},
		},
		{
			name: "Nil store",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				model := &RecordElement{intModel}

				elementValue, err := model.GetDisplayValue()

				assert.Nil(t, err)
				assert.Nil(t, elementValue)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*RecordElement)(nil)

				elementValue, err := model.GetDisplayValue()

				assert.Nil(t, err)
				assert.Nil(t, elementValue)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestRecordElementModel_SetDisplayValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRecordElementModel_GetValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				value := internal.ToPointer("")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", valueKey).Return(value, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				model := &RecordElement{intModel}

				elementValue, err := model.GetValue()

				assert.Nil(t, err)
				assert.NotNil(t, elementValue)
				assert.IsType(t, &ElementValue{}, elementValue)
				assert.Equal(t, interface{}(value), elementValue.val)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", valueKey).Return(nil, errors.New("retrieval error"))

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				model := &RecordElement{intModel}

				elementValue, err := model.GetValue()

				assert.Equal(t, errors.New("retrieval error"), err)
				assert.Nil(t, elementValue)
			},
		},
		{
			name: "Nil store",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				model := &RecordElement{intModel}

				elementValue, err := model.GetValue()

				assert.Nil(t, err)
				assert.Nil(t, elementValue)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*RecordElement)(nil)

				elementValue, err := model.GetValue()

				assert.Nil(t, err)
				assert.Nil(t, elementValue)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestRecordElementModel_SetValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRecordElementModel_GetLink(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				value := internal.ToPointer("")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", linkKey).Return(value, nil)

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				model := &RecordElement{intModel}

				elementValue, err := model.GetLink()

				assert.Nil(t, err)
				assert.NotNil(t, elementValue)
				assert.Equal(t, value, elementValue)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", linkKey).Return(nil, errors.New("retrieval error"))

				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return(backingStore)

				model := &RecordElement{intModel}

				elementValue, err := model.GetLink()

				assert.Equal(t, errors.New("retrieval error"), err)
				assert.Nil(t, elementValue)
			},
		},
		{
			name: "Nil store",
			test: func(t *testing.T) {
				intModel := mocking.NewMockModel()
				intModel.On("GetBackingStore").Return((*mocking.MockBackingStore)(nil))

				model := &RecordElement{intModel}

				elementValue, err := model.GetLink()

				assert.Nil(t, err)
				assert.Nil(t, elementValue)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*RecordElement)(nil)

				elementValue, err := model.GetLink()

				assert.Nil(t, err)
				assert.Nil(t, elementValue)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
func TestRecordElementModel_SetLink(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
