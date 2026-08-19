package appserviceapi

import (
	"context"
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func newTestServicePostRequestBuilder(adapter *mocking.MockRequestAdapter) *servicePostRequestBuilder[*CreateServiceRequest, CreateServiceResponse] {
	return newServicePostRequestBuilder[*CreateServiceRequest, CreateServiceResponse](
		adapter,
		createURLTemplate,
		map[string]string{"baseurl": "https://example.com"},
		CreateCreateServiceResponseFromDiscriminatorValue,
	)
}

func TestServicePostRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name        string
		nilBuilder  bool
		nilInner    bool
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(core.NewBaseServiceNowItemResponse[*CreateServiceResult](CreateCreateServiceResultFromDiscriminatorValue), nil)
			},
		},
		{
			name: "Error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("post failed"))
			},
			expectedErr: errors.New("post failed"),
		},
		{
			name: "Nil response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
			},
		},
		{
			name:       "Nil builder",
			nilBuilder: true,
		},
		{
			name:     "Nil inner request builder",
			nilInner: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *servicePostRequestBuilder[*CreateServiceRequest, CreateServiceResponse]

			switch {
			case tt.nilBuilder:
				builder = nil
			case tt.nilInner:
				builder = &servicePostRequestBuilder[*CreateServiceRequest, CreateServiceResponse]{}
			default:
				adapter := mocking.NewMockRequestAdapter()
				adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
				tt.setupMock(adapter)
				builder = newTestServicePostRequestBuilder(adapter)
			}

			resp, err := builder.post(context.Background(), NewCreateServiceRequest(), nil)

			switch {
			case tt.nilBuilder, tt.nilInner:
				require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
				assert.Nil(t, resp)
			case tt.expectedErr != nil:
				require.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
			default:
				assert.NoError(t, err)
			}
		})
	}
}

func TestServicePostRequestBuilder_ToPostRequestInformation(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())

	builder := newTestServicePostRequestBuilder(adapter)

	reqInfo, err := builder.toPostRequestInformation(context.Background(), NewCreateServiceRequest(), nil)

	require.NoError(t, err)
	assert.NotNil(t, reqInfo)
	assert.Equal(t, abstractions.POST, reqInfo.Method)
	assert.Equal(t, createURLTemplate, reqInfo.UrlTemplate)

	t.Run("Nil builder", func(t *testing.T) {
		var builder *servicePostRequestBuilder[*CreateServiceRequest, CreateServiceResponse]
		reqInfo, err := builder.toPostRequestInformation(context.Background(), NewCreateServiceRequest(), nil)
		require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, reqInfo)
	})

	t.Run("Nil inner request builder", func(t *testing.T) {
		builder := &servicePostRequestBuilder[*CreateServiceRequest, CreateServiceResponse]{}
		reqInfo, err := builder.toPostRequestInformation(context.Background(), NewCreateServiceRequest(), nil)
		require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, reqInfo)
	})

	t.Run("Serialization failure propagates", func(t *testing.T) {
		builder := newTestServicePostRequestBuilder(newFailingSerializationAdapter())

		reqInfo, err := builder.toPostRequestInformation(context.Background(), NewCreateServiceRequest(), nil)

		require.ErrorIs(t, err, errSerializationFailed)
		assert.Nil(t, reqInfo)
	})
}

// errSerializationFailed stands in for an error surfaced while writing the request body.
var errSerializationFailed = errors.New("serialization failed")

// newFailingSerializationAdapter returns an adapter whose serialization writer fails on
// write, which is the only way toPostRequestInformation can fail past its nil guard.
func newFailingSerializationAdapter() *mocking.MockRequestAdapter {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteObjectValue", "", mock.Anything, mock.Anything).Return(errSerializationFailed)
	writer.On("Close").Return(nil)

	factory := mocking.NewMockSerializationWriterFactory()
	factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(factory, nil)

	return adapter
}

// TestServicePostRequestBuilder_PostSerializationFailure covers post's propagation of a
// toPostRequestInformation error, which only a body-write failure can trigger.
func TestServicePostRequestBuilder_PostSerializationFailure(t *testing.T) {
	builder := newTestServicePostRequestBuilder(newFailingSerializationAdapter())

	resp, err := builder.post(context.Background(), NewCreateServiceRequest(), nil)

	require.ErrorIs(t, err, errSerializationFailed)
	assert.Nil(t, resp)
}
