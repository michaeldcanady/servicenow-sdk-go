package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAppServiceActionRequestModel_GettersSetters(t *testing.T) {
	model := NewAppServiceActionRequest()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"SysID", func(v any) error { return model.SetSysID(v.(*string)) }, func() (any, error) { return model.GetSysID() }, internal.ToPointer("service123")},
		{"Name", func(v any) error { return model.SetName(v.(*string)) }, func() (any, error) { return model.GetName() }, internal.ToPointer("Email")},
		{"Comments", func(v any) error { return model.SetComments(v.(*string)) }, func() (any, error) { return model.GetComments() }, internal.ToPointer("converted")},
		{"NumberOfLevels", func(v any) error { return model.SetNumberOfLevels(v.(*string)) }, func() (any, error) { return model.GetNumberOfLevels() }, internal.ToPointer("2")},
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

func TestAppServiceActionRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *AppServiceActionRequest
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewAppServiceActionRequest(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *AppServiceActionRequest {
				m := NewAppServiceActionRequest()
				_ = m.SetSysID(internal.ToPointer("service123"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
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

func TestAppServiceActionRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewAppServiceActionRequest()
	deserializers := model.GetFieldDeserializers()
	assert.Len(t, deserializers, 4)
}

func TestCreateAppServiceActionRequestFromDiscriminatorValue(t *testing.T) {
	val, err := CreateAppServiceActionRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.IsType(t, &AppServiceActionRequest{}, val)
}

func TestAppServiceActionResultModel_GettersSetters(t *testing.T) {
	model := NewAppServiceActionResult()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"SysID", func(v any) error { return model.setSysID(v.(*string)) }, func() (any, error) { return model.GetSysID() }, internal.ToPointer("service123")},
		{"Name", func(v any) error { return model.setName(v.(*string)) }, func() (any, error) { return model.GetName() }, internal.ToPointer("Email")},
		{"Comments", func(v any) error { return model.setComments(v.(*string)) }, func() (any, error) { return model.GetComments() }, internal.ToPointer("converted")},
		{"NumberOfLevels", func(v any) error { return model.setNumberOfLevels(v.(*string)) }, func() (any, error) { return model.GetNumberOfLevels() }, internal.ToPointer("2")},
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

func TestAppServiceActionResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *AppServiceActionResult
		setupMock func(w *mocking.MockSerializationWriter)
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewAppServiceActionResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *AppServiceActionResult {
				m := NewAppServiceActionResult()
				_ = m.setSysID(internal.ToPointer("service123"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := mocking.NewMockSerializationWriter()
			if tt.setupMock != nil {
				tt.setupMock(writer)
			}

			err := tt.model.Serialize(writer)
			require.NoError(t, err)
		})
	}
}

func TestAppServiceActionResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewAppServiceActionResult()
	deserializers := model.GetFieldDeserializers()
	assert.Len(t, deserializers, 4)
}

func TestCreateAppServiceActionResultFromDiscriminatorValue(t *testing.T) {
	val, err := CreateAppServiceActionResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.IsType(t, &AppServiceActionResult{}, val)
}
