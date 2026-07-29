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
			name:  "empty model writes only the any-typed field unconditionally",
			model: NewAvailabilityResult(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteAnyValue", nextAvailableSlotKey, mock.Anything).Return(nil)
			},
		},
		{
			name: "happy path - writes all fields",
			model: func() *AvailabilityResultModel {
				m := NewAvailabilityResult()
				_ = m.SetAvailability([]AvailabilitySlot{NewAvailabilitySlot()})
				_ = m.SetHasMore(internal.ToPointer(true))
				_ = m.SetNextAvailableSlot("next")
				_ = m.SetNoApptAvailable(internal.ToPointer(false))
				_ = m.SetSuccess(internal.ToPointer(true))
				_ = m.SetTimeZone(internal.ToPointer("UTC"))
				_ = m.SetTimeZoneDisplayValue(internal.ToPointer("Coordinated Universal Time"))
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteCollectionOfObjectValues", availabilityKey, mock.Anything).Return(nil)
				w.On("WriteBoolValue", mock.Anything, mock.Anything).Return(nil)
				w.On("WriteAnyValue", nextAvailableSlotKey, mock.Anything).Return(nil)
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
			name:  "happy path - writes additional data",
			model: NewAvailabilitySlot(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteAdditionalData", mock.Anything).Return(nil)
			},
		},
		{
			name: "write error propagates",
			model: func() *AvailabilitySlotModel {
				m := NewAvailabilitySlot()
				m.SetAdditionalData(map[string]interface{}{"start": "09:00"})
				return m
			}(),
			setupMock: func(w *mocking.MockSerializationWriter) {
				w.On("WriteAdditionalData", mock.Anything).Return(errWrite)
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
	assert.Empty(t, deserializers)
}
