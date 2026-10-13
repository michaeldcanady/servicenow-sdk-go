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

func TestVersionsRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(*mocking.MockRequestAdapter)
		wantErr   error
	}{
		{
			name: "happy path - returns collection response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				resp := core.NewBaseServiceNowCollectionResponse[Document](CreateDocumentFromDiscriminatorValue)
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
			},
		},
		{
			name: "adapter error propagates",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("versions failed"))
			},
			wantErr: errors.New("versions failed"),
		},
		{
			name: "nil response returns snerrors.ErrNilResponse",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
			},
			wantErr: snerrors.ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)
			builder := NewVersionsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "document_sys_id": "doc-id"}, adapter)

			resp, err := builder.Get(context.Background(), nil)

			if tt.wantErr != nil {
				if tt.name == "nil response returns snerrors.ErrNilResponse" {
					require.ErrorIs(t, err, tt.wantErr)
				} else {
					require.EqualError(t, err, tt.wantErr.Error())
				}
				assert.Nil(t, resp)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestVersionsRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *VersionsRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestVersionsRequestBuilder_Get_NilRequestAdapter(t *testing.T) {
	builder := NewVersionsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "document_sys_id": "doc-id"}, nil)

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestVersionsRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewVersionsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "document_sys_id": "doc-id"}, adapter)

	reqInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.NoError(t, err)
	assert.NotNil(t, reqInfo)
}
