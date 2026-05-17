package documentsapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, []byte("data"), resp)
			}
		})
	}
}
