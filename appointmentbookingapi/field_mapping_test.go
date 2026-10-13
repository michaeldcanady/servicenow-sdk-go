// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestFieldMappingModel_GettersSetters(t *testing.T) {
	model := NewFieldMapping()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Contact", func(v any) error { return model.SetContact(v.(*string)) }, func() (any, error) { return model.GetContact() }, internal.ToPointer("contact")},
		{"ContactRPVariable", func(v any) error { return model.SetContactRPVariable(v.(RPVariable)) }, func() (any, error) { return model.GetContactRPVariable() }, RPVariable(NewRPVariable())},
		{"Location", func(v any) error { return model.SetLocation(v.(*string)) }, func() (any, error) { return model.GetLocation() }, internal.ToPointer("location")},
		{"LocationRPVariable", func(v any) error { return model.SetLocationRPVariable(v.(RPVariable)) }, func() (any, error) { return model.GetLocationRPVariable() }, RPVariable(NewRPVariable())},
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

func TestCreateFieldMappingFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateFieldMappingFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestFieldMappingModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *FieldMappingModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewFieldMapping(),
		},
		{
			name: "happy path - writes all fields including nested RPVariables",
			model: func() *FieldMappingModel {
				m := NewFieldMapping()
				_ = m.SetContact(internal.ToPointer("contact"))
				_ = m.SetContactRPVariable(NewRPVariable())
				_ = m.SetLocation(internal.ToPointer("location"))
				_ = m.SetLocationRPVariable(NewRPVariable())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *FieldMappingModel {
				m := NewFieldMapping()
				_ = m.SetContact(internal.ToPointer("contact"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", contactKey, mock.Anything).Return(errWrite)
			},
			wantErr: errWrite,
		},
		{
			name: "nested object write error propagates",
			model: func() *FieldMappingModel {
				m := NewFieldMapping()
				_ = m.SetContactRPVariable(NewRPVariable())
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteObjectValue", contactRPVariableKey, mock.Anything, mock.Anything).Return(errWrite)
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

func TestFieldMappingModel_GetFieldDeserializers(t *testing.T) {
	model := NewFieldMapping()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{contactKey, contactRPVariableKey, locationKey, locationRPVariableKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 4)
}
