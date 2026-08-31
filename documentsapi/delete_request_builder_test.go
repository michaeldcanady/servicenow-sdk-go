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

func TestDeleteRequestBuilder_Delete_NilBuilder(t *testing.T) {
	var builder *DeleteRequestBuilder

	err := builder.Delete(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
}

func TestDeleteRequestBuilder_Delete_NilRequestAdapter(t *testing.T) {
	builder := NewDeleteRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, nil)

	err := builder.Delete(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
}

func TestDeleteRequestBuilder_Delete(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("delete failed"))
			},
			expectedErr: errors.New("delete failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewDeleteRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)
			err := builder.Delete(context.Background(), nil)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
