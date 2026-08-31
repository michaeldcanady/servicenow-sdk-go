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

func TestDiscoveryPopulationMethodAttributeModel_GettersSetters(t *testing.T) {
	model := NewDiscoveryPopulationMethodAttributeModel()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Name", func(v any) error { return model.SetName(v.(*string)) }, func() (any, error) { return model.GetName() }, internal.ToPointer("attr-name")},
		{"Value", func(v any) error { return model.SetValue(v.(*string)) }, func() (any, error) { return model.GetValue() }, internal.ToPointer("attr-value")},
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

func TestDiscoveryPopulationMethodAttributeModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *DiscoveryPopulationMethodAttributeModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewDiscoveryPopulationMethodAttributeModel(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *DiscoveryPopulationMethodAttributeModel {
				m := NewDiscoveryPopulationMethodAttributeModel()
				_ = m.SetName(internal.ToPointer("attr-name"))
				_ = m.SetValue(internal.ToPointer("attr-value"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *DiscoveryPopulationMethodAttributeModel {
				m := NewDiscoveryPopulationMethodAttributeModel()
				_ = m.SetName(internal.ToPointer("attr-name"))
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

func TestDiscoveryPopulationMethodAttributeModel_GetFieldDeserializers(t *testing.T) {
	model := NewDiscoveryPopulationMethodAttributeModel()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{nameKey, valueKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}

func TestCreateDiscoveryPopulationMethodAttributeModelFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateDiscoveryPopulationMethodAttributeModelFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
