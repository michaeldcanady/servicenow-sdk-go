// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

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

func TestCreateAppointmentResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAppointmentResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestAppointmentResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *AppointmentResultModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewAppointmentResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *AppointmentResultModel {
				m := NewAppointmentResult()
				_ = m.SetData(internal.ToPointer("data"))
				_ = m.SetMessage(internal.ToPointer("message"))
				_ = m.SetReason(internal.ToPointer("reason"))
				_ = m.SetSuccess(internal.ToPointer(true))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteBoolValue", successKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *AppointmentResultModel {
				m := NewAppointmentResult()
				_ = m.SetData(internal.ToPointer("data"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", dataKey, mock.Anything).Return(errWrite)
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

func TestAppointmentResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewAppointmentResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{dataKey, messageKey, reasonKey, successKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 4)
}
