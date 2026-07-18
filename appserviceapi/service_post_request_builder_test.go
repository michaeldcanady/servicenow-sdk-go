package appserviceapi

import (
	"context"
	"errors"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
				assert.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
				assert.Nil(t, resp)
			case tt.expectedErr != nil:
				assert.EqualError(t, err, tt.expectedErr.Error())
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

	assert.NoError(t, err)
	assert.NotNil(t, reqInfo)
	assert.Equal(t, abstractions.POST, reqInfo.Method)
	assert.Equal(t, createURLTemplate, reqInfo.UrlTemplate)

	t.Run("Nil builder", func(t *testing.T) {
		var builder *servicePostRequestBuilder[*CreateServiceRequest, CreateServiceResponse]
		reqInfo, err := builder.toPostRequestInformation(context.Background(), NewCreateServiceRequest(), nil)
		assert.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, reqInfo)
	})

	t.Run("Nil inner request builder", func(t *testing.T) {
		builder := &servicePostRequestBuilder[*CreateServiceRequest, CreateServiceResponse]{}
		reqInfo, err := builder.toPostRequestInformation(context.Background(), NewCreateServiceRequest(), nil)
		assert.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, reqInfo)
	})
}
