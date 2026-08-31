// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestServiceDetailsResultModel_GettersSetters(t *testing.T) {
	model := NewServiceDetailsResult()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Status", func(v any) error { return model.setStatus(v.(*string)) }, func() (any, error) { return model.GetStatus() }, internal.ToPointer("success")},
		{"Message", func(v any) error { return model.setMessage(v.(*string)) }, func() (any, error) { return model.GetMessage() }, internal.ToPointer("updated")},
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

func TestServiceDetailsResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *ServiceDetailsResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewServiceDetailsResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *ServiceDetailsResult {
				m := NewServiceDetailsResult()
				_ = m.setStatus(internal.ToPointer("success"))
				_ = m.setMessage(internal.ToPointer("updated"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *ServiceDetailsResult {
				m := NewServiceDetailsResult()
				_ = m.setStatus(internal.ToPointer("success"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", statusKey, mock.Anything).Return(errWrite)
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

func TestServiceDetailsResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewServiceDetailsResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{statusKey, messageKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}

func TestCreateServiceDetailsResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateServiceDetailsResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
