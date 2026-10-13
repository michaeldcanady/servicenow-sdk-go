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

func TestCalendarResponse_GettersSetters(t *testing.T) {
	model := NewCalendarResponse()

	tests := []struct {
		name   string
		setter func(val any) error
		getter func() (any, error)
		value  any
	}{
		{"RangeEnd", func(v any) error { return model.SetRangeEnd(v.(*string)) }, func() (any, error) { return model.GetRangeEnd() }, internal.ToPointer("2023-01-01")},
		{"RangeStart", func(v any) error { return model.SetRangeStart(v.(*string)) }, func() (any, error) { return model.GetRangeStart() }, internal.ToPointer("2023-01-02")},
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

func TestCreateCalendarResponseFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateCalendarResponseFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestCalendarResponse_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *CalendarResponse
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewCalendarResponse(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *CalendarResponse {
				m := NewCalendarResponse()
				_ = m.SetRangeEnd(internal.ToPointer("2023-01-01"))
				_ = m.SetRangeStart(internal.ToPointer("2023-01-02"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *CalendarResponse {
				m := NewCalendarResponse()
				_ = m.SetRangeEnd(internal.ToPointer("2023-01-01"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteStringValue", rangeEndKey, mock.Anything).Return(errWrite)
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

func TestCalendarResponse_GetFieldDeserializers(t *testing.T) {
	model := NewCalendarResponse()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{rangeEndKey, rangeStartKey} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 2)
}
