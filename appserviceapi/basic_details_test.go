// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// errWrite is a stand-in for an error a serialization.SerializationWriter can
// return from a Write* call.
var errWrite = errors.New("write error")

func TestBasicDetailsModel_GettersSetters(t *testing.T) {
	model := NewBasicDetails()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"Environment", func(v any) error { return model.SetEnvironment(v.(*string)) }, func() (any, error) { return model.GetEnvironment() }, internal.ToPointer("Production")},
		{"Name", func(v any) error { return model.SetName(v.(*string)) }, func() (any, error) { return model.GetName() }, internal.ToPointer("Service Name")},
		{"Version", func(v any) error { return model.SetVersion(v.(*string)) }, func() (any, error) { return model.GetVersion() }, internal.ToPointer("1.0")},
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

func TestBasicDetailsModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *BasicDetails
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewBasicDetails(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *BasicDetails {
				m := NewBasicDetails()
				_ = m.SetEnvironment(internal.ToPointer("Production"))
				_ = m.SetName(internal.ToPointer("Service Name"))
				_ = m.SetVersion(internal.ToPointer("1.0"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *BasicDetails {
				m := NewBasicDetails()
				_ = m.SetEnvironment(internal.ToPointer("Production"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", environmentKey, mock.Anything).Return(errWrite)
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

func TestBasicDetailsModel_GetFieldDeserializers(t *testing.T) {
	model := NewBasicDetails()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{environmentKey, nameKey, versionKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 3)
}

func TestCreateBasicDetailsFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateBasicDetailsFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}
