package appointmentbookingapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateAvailabilityResultFromDiscriminatorValue(t *testing.T) {
	parsable, err := CreateAvailabilityResultFromDiscriminatorValue(nil)
	require.NoError(t, err)
	assert.NotNil(t, parsable)
}

func TestAvailabilityResultModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *AvailabilityResultModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewAvailabilityResult(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *AvailabilityResultModel {
				m := NewAvailabilityResult()
				_ = m.SetAvailability([]AvailabilitySlot{NewAvailabilitySlot()})
				_ = m.SetHasMore(internal.ToPointer(true))
				_ = m.SetNextAvailableSlot(NewAvailabilitySlot())
				_ = m.SetNoApptAvailable(internal.ToPointer(false))
				_ = m.SetSuccess(internal.ToPointer(true))
				_ = m.SetTimeZone(internal.ToPointer("UTC"))
				_ = m.SetTimeZoneDisplayValue(internal.ToPointer("Coordinated Universal Time"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteCollectionOfObjectValues", availabilityKey, mock.Anything).Return(nil)
				w.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteObjectValue", nextAvailableSlotKey, mock.Anything, mock.Anything).Return(nil)
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *AvailabilityResultModel {
				m := NewAvailabilityResult()
				_ = m.SetAvailability([]AvailabilitySlot{NewAvailabilitySlot()})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteCollectionOfObjectValues", availabilityKey, mock.Anything).Return(errWrite)
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

func TestAvailabilityResultModel_GetFieldDeserializers(t *testing.T) {
	model := NewAvailabilityResult()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		availabilityKey, hasMoreKey, nextAvailableSlotKey, noApptAvailableKey,
		successKey, timeZoneKey, timeZoneDisplayValueKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 7)
}

func TestAvailabilitySlotModel_Serialize(t *testing.T) {
	tests := []struct {
		name      string
		model     *AvailabilitySlotModel
		setupMock func(w *mocking.MockSerializationWriter)
		wantErr   error
	}{
		{
			name:  "nil model returns nil",
			model: nil,
		},
		{
			name:  "empty model writes nothing",
			model: NewAvailabilitySlot(),
		},
		{
			name: "happy path - writes all fields",
			model: func() *AvailabilitySlotModel {
				m := NewAvailabilitySlot()
				_ = m.SetAvailable(internal.ToPointer(true))
				_ = m.SetEndDate(internal.ToPointer("2026-07-29 10:00:00"))
				_ = m.SetEndDateDisplay(internal.ToPointer("10:00 AM"))
				_ = m.SetEndDateUTC(internal.ToPointer("2026-07-29 17:00:00"))
				_ = m.SetStartDate(internal.ToPointer("2026-07-29 09:00:00"))
				_ = m.SetStartDateDisplay(internal.ToPointer("9:00 AM"))
				_ = m.SetStartDateUTC(internal.ToPointer("2026-07-29 16:00:00"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteStringValue", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *AvailabilitySlotModel {
				m := NewAvailabilitySlot()
				_ = m.SetAvailable(internal.ToPointer(true))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteBoolValue", availableKey, mock.Anything).Return(errWrite)
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

func TestAvailabilitySlotModel_GetFieldDeserializers(t *testing.T) {
	model := NewAvailabilitySlot()
	deserializers := model.GetFieldDeserializers()
	for _, key := range []string{
		availableKey, endDateKey, endDateDisplayKey, endDateUTCKey,
		startDateKey, startDateDisplayKey, startDateUTCKey,
	} {
		assert.NotNil(t, deserializers[key], "expected deserializer for %s", key)
	}
	assert.Len(t, deserializers, 7)
}
