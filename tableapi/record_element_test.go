package tableapi

import (
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func newRecordElementWithMockStore(bs store.BackingStore) *RecordElement {
	model := NewRecordElement()
	model.SetBackingStoreFactory(func() store.BackingStore { return bs })
	return model
}

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

// TODO: improve test table design
func TestRecordElementModel_GetDisplayValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				value := &ElementValue{val: internal.ToPointer("")}

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", "display_value").Return(value, nil)

				model := newRecordElementWithMockStore(backingStore)

				elementValue, err := model.GetDisplayValue()

				require.NoError(t, err)
				assert.NotNil(t, elementValue)
				assert.IsType(t, ElementValue{}, elementValue)
				assert.Equal(t, interface{}(internal.ToPointer("")), elementValue.val)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", "display_value").Return(nil, errors.New("retrieval error"))

				model := newRecordElementWithMockStore(backingStore)

				elementValue, err := model.GetDisplayValue()

				assert.Equal(t, errors.New("retrieval error"), err)
				assert.Equal(t, ElementValue{}, elementValue)
			},
		},
		{
			name: "Nil store",
			test: func(t *testing.T) {
				model := NewRecordElement()
				model.SetBackingStoreFactory(func() store.BackingStore { return nil })

				elementValue, err := model.GetDisplayValue()

				require.ErrorIs(t, err, snerrors.ErrNilStore)
				assert.Equal(t, ElementValue{}, elementValue)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*RecordElement)(nil)

				elementValue, err := model.GetDisplayValue()

				require.ErrorIs(t, err, snerrors.ErrNilModel)
				assert.Equal(t, ElementValue{}, elementValue)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRecordElementModel_SetDisplayValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", "display_value", mock.AnythingOfType("tableapi.ElementValue")).Return(nil)

				record := newRecordElementWithMockStore(backingStore)

				err := record.SetDisplayValue("displayValue")

				require.NoError(t, err)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Store error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", "display_value", mock.AnythingOfType("tableapi.ElementValue")).Return(errors.New("store error"))

				record := newRecordElementWithMockStore(backingStore)

				err := record.SetDisplayValue("displayValue")

				assert.Equal(t, errors.New("store error"), err)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Invalid value",
			test: func(t *testing.T) {
				record := &RecordElement{}

				err := record.SetDisplayValue(make(chan int))

				assert.Equal(t, errors.New("unsupported kind chan"), err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: improve test table design
func TestRecordElementModel_GetValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				value := &ElementValue{val: internal.ToPointer("")}

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", "value").Return(value, nil)

				model := newRecordElementWithMockStore(backingStore)

				elementValue, err := model.GetValue()

				require.NoError(t, err)
				assert.NotNil(t, elementValue)
				assert.IsType(t, ElementValue{}, elementValue)
				assert.Equal(t, interface{}(internal.ToPointer("")), elementValue.val)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", "value").Return(nil, errors.New("retrieval error"))

				model := newRecordElementWithMockStore(backingStore)

				elementValue, err := model.GetValue()

				assert.Equal(t, errors.New("retrieval error"), err)
				assert.Equal(t, ElementValue{}, elementValue)
			},
		},
		{
			name: "Nil store",
			test: func(t *testing.T) {
				model := NewRecordElement()
				model.SetBackingStoreFactory(func() store.BackingStore { return nil })

				elementValue, err := model.GetValue()

				require.ErrorIs(t, err, snerrors.ErrNilStore)
				assert.Equal(t, ElementValue{}, elementValue)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*RecordElement)(nil)

				elementValue, err := model.GetValue()

				require.ErrorIs(t, err, snerrors.ErrNilModel)
				assert.Equal(t, ElementValue{}, elementValue)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRecordElementModel_SetValue(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", "value", mock.AnythingOfType("tableapi.ElementValue")).Return(nil)

				record := newRecordElementWithMockStore(backingStore)

				err := record.SetValue("value")

				require.NoError(t, err)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Store error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", "value", mock.AnythingOfType("tableapi.ElementValue")).Return(errors.New("store error"))

				record := newRecordElementWithMockStore(backingStore)

				err := record.SetValue("value")

				assert.Equal(t, errors.New("store error"), err)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Invalid value",
			test: func(t *testing.T) {
				record := &RecordElement{}

				err := record.SetValue(make(chan int))

				assert.Equal(t, errors.New("unsupported kind chan"), err)
			},
		},
	}

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
				backingStore.On("Get", "link").Return(value, nil)

				model := newRecordElementWithMockStore(backingStore)

				elementValue, err := model.GetLink()

				require.NoError(t, err)
				assert.NotNil(t, elementValue)
				assert.Equal(t, value, elementValue)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", "link").Return(nil, errors.New("retrieval error"))

				model := newRecordElementWithMockStore(backingStore)

				elementValue, err := model.GetLink()

				assert.Equal(t, errors.New("retrieval error"), err)
				assert.Nil(t, elementValue)
			},
		},
		{
			name: "Nil store",
			test: func(t *testing.T) {
				model := NewRecordElement()
				model.SetBackingStoreFactory(func() store.BackingStore { return nil })

				elementValue, err := model.GetLink()

				require.ErrorIs(t, err, snerrors.ErrNilStore)
				assert.Nil(t, elementValue)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				model := (*RecordElement)(nil)

				elementValue, err := model.GetLink()

				require.ErrorIs(t, err, snerrors.ErrNilModel)
				assert.Nil(t, elementValue)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestRecordElementModel_SetLink(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				link := internal.ToPointer("value")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", "link", link).Return(nil)

				record := newRecordElementWithMockStore(backingStore)

				err := record.SetLink(link)

				require.NoError(t, err)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Store error",
			test: func(t *testing.T) {
				link := internal.ToPointer("value")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", "link", link).Return(errors.New("store error"))

				record := newRecordElementWithMockStore(backingStore)

				err := record.SetLink(link)

				assert.Equal(t, errors.New("store error"), err)
				backingStore.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
