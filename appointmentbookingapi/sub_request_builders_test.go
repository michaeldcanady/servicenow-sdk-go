package appointmentbookingapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// errNetwork is a stand-in for a transport-level error returned by the adapter.
var errNetwork = errors.New("network error")

func TestAppointmentBookingRequestBuilder_NilReceiver(t *testing.T) {
	var builder *AppointmentBookingRequestBuilder

	assert.Nil(t, builder.Appointment())
	assert.Nil(t, builder.Availability())
	assert.Nil(t, builder.Calendar())
	assert.Nil(t, builder.Configuration())
	assert.Nil(t, builder.ExecuteRuleConditions())
	assert.Nil(t, builder.UserWindow())
}

func TestCalendarRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path - returns response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*CalendarResponse](CreateCalendarResponseFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)
			builder := NewCalendarRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestConfigurationRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path - returns response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*ConfigurationResult](CreateConfigurationResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)
			builder := NewConfigurationRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, resp)
				return
			}
			assert.NotNil(t, resp)
			adapter.AssertExpectations(t)
		})
	}
}

func TestAvailabilityRequestBuilder_Post_SerializationError(t *testing.T) {
	builder := NewAvailabilityRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, nil)

	resp, err := builder.Post(context.Background(), NewAvailabilityRequest(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}
