package appserviceapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewPopulateServiceRequestBuilder(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewPopulateServiceRequestBuilder("https://example.com/api/now/v1/cmdb/csdm/app_service/service123/populate_service", adapter)

	require.NotNil(t, builder)
	assert.Equal(t, populateServiceURLTemplate, builder.GetURLTemplate())
}

func TestPopulateServiceRequestBuilder_Put(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(m *mocking.MockRequestAdapter)
		wantErr   error
		wantNil   bool
	}{
		{
			name: "happy path - returns response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(core.NewBaseServiceNowItemResponse[*PopulateServiceResult](CreatePopulateServiceResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errNetwork)
			},
			wantErr: errNetwork,
		},
		{
			name: "nil response returns nil, nil",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)

			builder := NewPopulateServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "sys_id": "service123"}, adapter)

			resp, err := builder.Put(context.Background(), NewPopulateServiceRequest(), nil)

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
		})
	}

	t.Run("ToPutRequestInformation error propagates", func(t *testing.T) {
		builder := NewPopulateServiceRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "sys_id": "service123"}, nil)

		resp, err := builder.Put(context.Background(), NewPopulateServiceRequest(), nil)

		require.Error(t, err)
		assert.Nil(t, resp)
	})
}
