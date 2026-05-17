package documentsapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDocumentsRequestBuilder2_Builders(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewDocumentsRequestBuilder2Internal(map[string]string{"baseurl": "https://example.com"}, adapter)

	tests := []struct {
		name     string
		builder  any
		expected map[string]string
	}{
		{
			name:    "Explore",
			builder: builder.Explore(),
		},
		{
			name:    "Create",
			builder: builder.Create(),
		},
		{
			name:    "CreateDocument",
			builder: builder.CreateDocument(),
		},
		{
			name:    "VersionState",
			builder: builder.VersionState("version_sys_id"),
			expected: map[string]string{"version_sys_id": "version_sys_id"},
		},
		{
			name:    "Attach",
			builder: builder.Attach("provider_id"),
			expected: map[string]string{"provider_id": "provider_id"},
		},
		{
			name:    "Delete",
			builder: builder.Delete(),
		},
		{
			name:    "Versions",
			builder: builder.Versions("doc_sys_id"),
			expected: map[string]string{"document_sys_id": "doc_sys_id"},
		},
		{
			name:    "Content",
			builder: builder.Content("doc_sys_id"),
			expected: map[string]string{"document_sys_id": "doc_sys_id"},
		},
		{
			name:    "SyncDown",
			builder: builder.SyncDown("doc_sys_id"),
			expected: map[string]string{"documentSysId": "doc_sys_id"},
		},
		{
			name:    "Action",
			builder: builder.Action("move"),
			expected: map[string]string{"action": "move"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.builder)
			if tt.expected != nil {
				rb := tt.builder.(newInternal.RequestBuilder)
				for k, v := range tt.expected {
					assert.Equal(t, v, rb.GetPathParameters()[k])
				}
			}
		})
	}
}

func TestExploreRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				resp := newInternal.NewBaseServiceNowCollectionResponse[Document](CreateDocumentFromDiscriminatorValue)
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewExploreRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)
			resp, err := builder.Get(context.Background(), nil)

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

func TestCreateRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				resp := newInternal.NewBaseServiceNowItemResponse[Document](CreateDocumentFromDiscriminatorValue)
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
