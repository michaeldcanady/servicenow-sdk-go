// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package documentsapi

import (
	"context"
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestContentRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *ContentRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestContentRequestBuilder_Get_NilRequestAdapter(t *testing.T) {
	builder := NewContentRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, nil)

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestContentRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendPrimitive", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]byte("data"), nil)
			},
			expectedErr: nil,
		},
		{
			name: "Error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendPrimitive", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("stream error"))
			},
			expectedErr: errors.New("stream error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewContentRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)
			resp, err := builder.Get(context.Background(), nil)

			if tt.expectedErr != nil {
				require.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				assert.Equal(t, []byte("data"), resp)
			}
		})
	}
}
