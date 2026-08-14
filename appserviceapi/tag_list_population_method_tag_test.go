package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTagListPopulationMethodTagModel_GettersSetters(t *testing.T) {
	model := NewTagListPopulationMethodTag()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Tag", func(v any) error { return model.SetTag(v.(*string)) }, func() (any, error) { return model.GetTag() }, internal.ToPointer("env")},
		{"Value", func(v any) error { return model.SetValue(v.(*string)) }, func() (any, error) { return model.GetValue() }, internal.ToPointer("production")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.setter(tt.value)
			require.NoError(t, err)
			got, err := tt.getter()
			require.NoError(t, err)
			assert.Equal(t, tt.value, got)
		})
	}
}

func TestTagListPopulationMethodTagModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *TagListPopulationMethodTagModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewTagListPopulationMethodTag(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *TagListPopulationMethodTagModel {
				m := NewTagListPopulationMethodTag()
				_ = m.SetTag(internal.ToPointer("env"))
				_ = m.SetValue(internal.ToPointer("production"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *TagListPopulationMethodTagModel {
				m := NewTagListPopulationMethodTag()
				_ = m.SetTag(internal.ToPointer("env"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", tagKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if tt.setupMock != nil {
				tt.setupMock(writer)
			}

			err := tt.model.Serialize(writer)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestTagListPopulationMethodTagModel_GetFieldDeserializers(t *testing.T) {
	model := NewTagListPopulationMethodTag()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{tagKey, valueKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}

func TestCreateTagListPopulationMethodTagFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateTagListPopulationMethodTagFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
