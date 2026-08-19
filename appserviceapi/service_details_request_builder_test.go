package appserviceapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewServiceDetailsRequestBuilder(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewServiceDetailsRequestBuilder("https://example.com/api/now/cmdb/csdm/app_service/service123/service_details", adapter)

	require.NotNil(t, builder)
	assert.Equal(t, serviceDetailsURLTemplate, builder.GetURLTemplate())
}

func TestServiceDetailsRequestBuilder_Put(t *testing.T) {
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
					Return(core.NewBaseServiceNowItemResponse[*ServiceDetailsResult](CreateServiceDetailsResultFromDiscriminatorValue), nil)
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
			name: "nil response returns snerrors.ErrNilResponse",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil, nil)
			},
			wantErr: snerrors.ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := mocking.NewMockRequestAdapter()
			tt.setupMock(adapter)

			builder := NewServiceDetailsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "sys_id": "service123"}, adapter)

			resp, err := builder.Put(context.Background(), NewServiceDetailsRequest(), nil)

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
		builder := NewServiceDetailsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "sys_id": "service123"}, nil)

		resp, err := builder.Put(context.Background(), NewServiceDetailsRequest(), nil)

		require.Error(t, err)
		assert.Nil(t, resp)
	})
}
