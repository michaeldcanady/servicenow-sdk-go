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

func TestCreateServiceRequestModel_GettersSetters(t *testing.T) {
	model := NewCreateServiceRequest()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Name", func(v any) error { return model.setName(v.(*string)) }, func() (any, error) { return model.GetName() }, internal.ToPointer("AppService-CreateTest")},
		{"Comments", func(v any) error { return model.setComments(v.(*string)) }, func() (any, error) { return model.GetComments() }, internal.ToPointer("Testing creation endpoint")},
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

func TestCreateServiceRequestModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *CreateServiceRequest
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewCreateServiceRequest(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *CreateServiceRequest {
				m := NewCreateServiceRequest()
				_ = m.setName(internal.ToPointer("AppService-CreateTest"))
				_ = m.setComments(internal.ToPointer("Testing creation endpoint"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *CreateServiceRequest {
				m := NewCreateServiceRequest()
				_ = m.setName(internal.ToPointer("AppService-CreateTest"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", nameKey, mock.Anything).Return(errWrite)
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

func TestCreateServiceRequestModel_GetFieldDeserializers(t *testing.T) {
	model := NewCreateServiceRequest()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{nameKey, commentsKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}

func TestCreateCreateServiceRequestFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateCreateServiceRequestFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCreateServiceResultModel_GettersSetters(t *testing.T) {
	model := NewCreateServiceResult()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"SysID", func(v any) error { return model.setSysID(v.(*string)) }, func() (any, error) { return model.GetSysID() }, internal.ToPointer("sys-id")},
		{"Name", func(v any) error { return model.setName(v.(*string)) }, func() (any, error) { return model.GetName() }, internal.ToPointer("AppService-CreateTest")},
		{"Comments", func(v any) error { return model.setComments(v.(*string)) }, func() (any, error) { return model.GetComments() }, internal.ToPointer("Testing creation endpoint")},
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

func TestCreateServiceResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *CreateServiceResult
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewCreateServiceResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *CreateServiceResult {
				m := NewCreateServiceResult()
				_ = m.setSysID(internal.ToPointer("sys-id"))
				_ = m.setName(internal.ToPointer("AppService-CreateTest"))
				_ = m.setComments(internal.ToPointer("Testing creation endpoint"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *CreateServiceResult {
				m := NewCreateServiceResult()
				_ = m.setSysID(internal.ToPointer("sys-id"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", sysIDKey, mock.Anything).Return(errWrite)
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

func TestCreateServiceResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewCreateServiceResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{sysIDKey, nameKey, commentsKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 6)
}

func TestCreateCreateServiceResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateCreateServiceResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
