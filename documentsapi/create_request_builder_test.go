package documentsapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateRequestBuilder_Post_NilBuilder(t *testing.T) {
	var builder *CreateRequestBuilder

	resp, err := builder.Post(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestCreateRequestBuilder_ToPostRequestInformation(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		adapter := &mocking.MockRequestAdapter{}
		builder := NewCreateRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

		reqInfo, err := builder.ToPostRequestInformation(context.Background(), nil)

		require.NoError(t, err)
		assert.NotNil(t, reqInfo)
	})

	t.Run("nil receiver returns ErrNilRequestBuilder", func(t *testing.T) {
		var builder *CreateRequestBuilder

		reqInfo, err := builder.ToPostRequestInformation(context.Background(), nil)

		require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, reqInfo)
	})
}

func TestCreateRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				resp := core.NewBaseServiceNowItemResponse[Document](CreateDocumentFromDiscriminatorValue)
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
			},
			expectedErr: nil,
		},
		{
			name: "Error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("creation failed"))
			},
			expectedErr: errors.New("creation failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewCreateRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)
			resp, err := builder.Post(context.Background(), nil)

			if tt.expectedErr != nil {
				require.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}
