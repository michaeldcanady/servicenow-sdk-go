package cmdbinstanceapi

import (
	"context"
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCmdbRelationItemRequestBuilder_Delete(t *testing.T) {
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

			builder := NewCmdbRelationItemRequestBuilderInternal(map[string]string{
				"baseurl":    "https://example.com",
				"className":  "test",
				"sys_id":     "123",
				"rel_sys_id": "456",
			}, adapter)
			err := builder.Delete(context.Background(), nil)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCmdbRelationItemRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*CmdbRelationItemRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			err := builder.Delete(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		})
	}
}

func TestCmdbRelationItemRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewCmdbRelationItemRequestBuilderInternal(map[string]string{}, nil)

	err := builder.Delete(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
}

func TestCmdbRelationItemRequestBuilder_ToDeleteRequestInformation(t *testing.T) {
	builder := NewCmdbRelationItemRequestBuilderInternal(map[string]string{
		"baseurl":    "https://example.com",
		"className":  "test",
		"sys_id":     "123",
		"rel_sys_id": "456",
	}, &mocking.MockRequestAdapter{})

	t.Run("nil config", func(t *testing.T) {
		requestInfo, err := builder.ToDeleteRequestInformation(context.Background(), nil)
		require.NoError(t, err)
		assert.Equal(t, abstractions.DELETE, requestInfo.Method)
	})

	t.Run("config with headers and options", func(t *testing.T) {
		headers := abstractions.NewRequestHeaders()
		headers.Add("X-Test", "value")

		option := mocking.NewMockRequestOption()
		option.On("GetKey").Return(abstractions.RequestOptionKey{Key: "testOption"})

		config := &CmdbRelationItemRequestBuilderDeleteRequestConfiguration{
			Headers: headers,
			Options: []abstractions.RequestOption{option},
		}

		requestInfo, err := builder.ToDeleteRequestInformation(context.Background(), config)
		require.NoError(t, err)
		assert.Equal(t, []string{"value"}, requestInfo.Headers.Get("X-Test"))
		assert.Len(t, requestInfo.GetRequestOptions(), 1)
	})
}
