package batchapi

import (
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	internal "github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRestRequestHeader(t *testing.T) {
	header := NewRestRequestHeader()
	assert.NotNil(t, header)
}

func TestCreateRestRequestHeaderFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateRestRequestHeaderFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestRestRequestHeader_Serialize(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successfully",
			test: func(t *testing.T) {
				expName := internal.ToPointer("name")
				expValue := internal.ToPointer("value")

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteStringValue", nameKey, expName).Return(nil)
				writer.On("WriteStringValue", valueKey, expValue).Return(nil)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(expName, nil)
				backingStore.On("Get", valueKey).Return(expValue, nil)

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				err := header.Serialize(writer)

				require.NoError(t, err)
				writer.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Name retrieval error",
			test: func(t *testing.T) {
				writer := mocking.NewMockSerializationWriter()

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(nil, errors.New("retrieval error"))

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				err := header.Serialize(writer)

				assert.Equal(t, errors.New("retrieval error"), err)
				writer.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Value retrieval error",
			test: func(t *testing.T) {
				expName := internal.ToPointer("name")

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteStringValue", nameKey, expName).Return(nil)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(expName, nil)
				backingStore.On("Get", valueKey).Return(nil, errors.New("retrieval error"))

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				err := header.Serialize(writer)

				assert.Equal(t, errors.New("retrieval error"), err)
				writer.AssertExpectations(t)
				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "nil writer",
			test: func(t *testing.T) {
				bR := NewRestRequestHeader()
				err := bR.Serialize(nil)
				assert.ErrorIs(t, err, snerrors.ErrNilWriter)
			},
		},
		{
			name: "nil_model",
			test: func(t *testing.T) {
				var header *RestRequestHeaderModel

				err := header.Serialize(mocking.NewMockSerializationWriter())

				assert.NoError(t, err)
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
				expName := internal.ToPointer("name")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(expName, nil)

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				name, err := header.GetName()

				require.NoError(t, err)
				assert.Equal(t, expName, name)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				expName := internal.ToPointer(true)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(expName, nil)

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				name, err := header.GetName()

				assert.Equal(t, errors.New("cannot convert 'true' to type *string"), err)
				assert.Nil(t, name)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", nameKey).Return(nil, errors.New("error retrieving"))

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				name, err := header.GetName()

				assert.Equal(t, errors.New("error retrieving"), err)
				assert.Nil(t, name)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return nil })

				name, err := header.GetName()

				require.ErrorIs(t, err, snerrors.ErrNilStore)
				assert.Nil(t, name)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				header := (*RestRequestHeaderModel)(nil)

				name, err := header.GetName()

				require.ErrorIs(t, err, snerrors.ErrNilModel)
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
				input := internal.ToPointer("name")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", nameKey, input).Return(nil)

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				err := header.SetName(input)

				require.NoError(t, err)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Store error",
			test: func(t *testing.T) {
				input := internal.ToPointer("name")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", nameKey, input).Return(errors.New("store error"))

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				err := header.SetName(input)

				assert.Equal(t, errors.New("store error"), err)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				input := internal.ToPointer("name")

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return nil })

				err := header.SetName(input)

				require.ErrorIs(t, err, snerrors.ErrNilStore)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				header := (*RestRequestHeaderModel)(nil)

				err := header.SetName(internal.ToPointer("name"))

				require.ErrorIs(t, err, snerrors.ErrNilModel)
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
				expValue := internal.ToPointer("value")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", valueKey).Return(expValue, nil)

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				value, err := header.GetValue()

				require.NoError(t, err)
				assert.Equal(t, expValue, value)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				expValue := internal.ToPointer(true)

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", valueKey).Return(expValue, nil)

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				value, err := header.GetValue()

				assert.Equal(t, errors.New("cannot convert 'true' to type *string"), err)
				assert.Nil(t, value)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Retrieval error",
			test: func(t *testing.T) {
				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Get", valueKey).Return(nil, errors.New("error retrieving"))

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				value, err := header.GetValue()

				assert.Equal(t, errors.New("error retrieving"), err)
				assert.Nil(t, value)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return nil })

				value, err := header.GetValue()

				require.ErrorIs(t, err, snerrors.ErrNilStore)
				assert.Nil(t, value)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				header := (*RestRequestHeaderModel)(nil)

				value, err := header.GetValue()

				require.ErrorIs(t, err, snerrors.ErrNilModel)
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
				input := internal.ToPointer("value")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", valueKey, input).Return(nil)

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				err := header.SetValue(input)

				require.NoError(t, err)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Store error",
			test: func(t *testing.T) {
				input := internal.ToPointer("value")

				backingStore := mocking.NewMockBackingStore()
				backingStore.On("Set", valueKey, input).Return(errors.New("store error"))

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return backingStore })

				err := header.SetValue(input)

				assert.Equal(t, errors.New("store error"), err)

				backingStore.AssertExpectations(t)
			},
		},
		{
			name: "Nil backing store",
			test: func(t *testing.T) {
				input := internal.ToPointer("value")

				header := NewRestRequestHeader()
				header.SetBackingStoreFactory(func() store.BackingStore { return nil })

				err := header.SetValue(input)

				require.ErrorIs(t, err, snerrors.ErrNilStore)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				header := (*RestRequestHeaderModel)(nil)

				err := header.SetValue(internal.ToPointer("value"))

				require.ErrorIs(t, err, snerrors.ErrNilModel)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
