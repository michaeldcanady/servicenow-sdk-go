package documentsapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				resp := internal.NewBaseServiceNowItemResponse[Document](CreateDocumentFromDiscriminatorValue)
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
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}
