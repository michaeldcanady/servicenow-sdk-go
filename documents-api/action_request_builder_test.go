package documentsapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVersionActionRequestBuilder_Patch(t *testing.T) {
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
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("patch failed"))
			},
			expectedErr: errors.New("patch failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewVersionActionRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)
			err := builder.Patch(context.Background(), nil)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
