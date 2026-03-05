package attachmentapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAttachmentRequestBuilder2_Get(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*mocking.MockRequestAdapter)
		config      *AttachmentRequestBuilder2GetRequestConfiguration
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				resp := newInternal.NewBaseServiceNowCollectionResponse[Attachment2](CreateAttachment2FromDiscriminatorValue)
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
			},
			expectedErr: nil,
		},
		{
			name: "Error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("network error"))
			},
			expectedErr: errors.New("network error"),
		},
		{
			name: "Nil Response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
			},
			expectedErr: nil, // Current implementation returns nil, nil for nil response
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewAttachmentRequestBuilder2Internal(map[string]string{"baseurl": "https://example.com"}, adapter)
			resp, err := builder.Get(context.Background(), tt.config)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				if tt.name != "Nil Response" {
					assert.NotNil(t, resp)
				} else {
					assert.Nil(t, resp)
				}
			}
		})
	}
}

func TestAttachmentRequestBuilder2_Builders(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewAttachmentRequestBuilder2Internal(map[string]string{"baseurl": "https://example.com"}, adapter)

	t.Run("ByID", func(t *testing.T) {
		itemBuilder := builder.ByID("sys_id_123")
		assert.NotNil(t, itemBuilder)
	})

	t.Run("File", func(t *testing.T) {
		fileBuilder := builder.File()
		assert.NotNil(t, fileBuilder)
	})

	t.Run("Upload", func(t *testing.T) {
		uploadBuilder := builder.Upload()
		assert.NotNil(t, uploadBuilder)
	})
}
