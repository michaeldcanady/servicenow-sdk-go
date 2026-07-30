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

func TestVersionStateRequestBuilder_Get(t *testing.T) {
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
		},
		{
			name: "Error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("version state failed"))
			},
			expectedErr: errors.New("version state failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewVersionStateRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "version_sys_id": "version-id"}, adapter)
			resp, err := builder.Get(context.Background(), nil)

			if tt.expectedErr != nil {
				require.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestVersionStateRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *VersionStateRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestVersionStateRequestBuilder_Get_NilRequestAdapter(t *testing.T) {
	builder := NewVersionStateRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "version_sys_id": "version-id"}, nil)

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestVersionStateRequestBuilder_ToGetRequestInformation(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		adapter := &mocking.MockRequestAdapter{}
		builder := NewVersionStateRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "version_sys_id": "version-id"}, adapter)

		reqInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

		require.NoError(t, err)
		assert.NotNil(t, reqInfo)
	})

	t.Run("nil receiver returns ErrNilRequestBuilder", func(t *testing.T) {
		var builder *VersionStateRequestBuilder

		reqInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

		require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, reqInfo)
	})
}
