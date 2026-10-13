// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package documentsapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSyncDownRequestBuilder_Post(t *testing.T) {
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
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("sync down failed"))
			},
			expectedErr: errors.New("sync down failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewSyncDownRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "documentSysId": "doc-id"}, adapter)
			resp, err := builder.Post(context.Background(), nil)

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

func TestSyncDownRequestBuilder_Post_NilBuilder(t *testing.T) {
	var builder *SyncDownRequestBuilder

	resp, err := builder.Post(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestSyncDownRequestBuilder_Post_NilRequestAdapter(t *testing.T) {
	builder := NewSyncDownRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "documentSysId": "doc-id"}, nil)

	resp, err := builder.Post(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestSyncDownRequestBuilder_ToPostRequestInformation(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		adapter := &mocking.MockRequestAdapter{}
		builder := NewSyncDownRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "documentSysId": "doc-id"}, adapter)

		reqInfo, err := builder.ToPostRequestInformation(context.Background(), nil)

		require.NoError(t, err)
		assert.NotNil(t, reqInfo)
	})

	t.Run("nil receiver returns ErrNilRequestBuilder", func(t *testing.T) {
		var builder *SyncDownRequestBuilder

		reqInfo, err := builder.ToPostRequestInformation(context.Background(), nil)

		require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, reqInfo)
	})
}
