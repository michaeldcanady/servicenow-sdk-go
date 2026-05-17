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

	t.Run("Explore", func(t *testing.T) {
		exploreBuilder := builder.Explore()
		assert.NotNil(t, exploreBuilder)
	})

	t.Run("Create", func(t *testing.T) {
		createBuilder := builder.Create()
		assert.NotNil(t, createBuilder)
	})

	t.Run("CreateDocument", func(t *testing.T) {
		createDocumentBuilder := builder.CreateDocument()
		assert.NotNil(t, createDocumentBuilder)
	})

	t.Run("VersionState", func(t *testing.T) {
		versionStateBuilder := builder.VersionState("version_sys_id")
		assert.NotNil(t, versionStateBuilder)
		assert.Equal(t, "version_sys_id", versionStateBuilder.GetPathParameters()["version_sys_id"])
	})

	t.Run("Attach", func(t *testing.T) {
		attachBuilder := builder.Attach("provider_id")
		assert.NotNil(t, attachBuilder)
		assert.Equal(t, "provider_id", attachBuilder.GetPathParameters()["provider_id"])
	})

	t.Run("Delete", func(t *testing.T) {
		deleteBuilder := builder.Delete()
		assert.NotNil(t, deleteBuilder)
	})

	t.Run("Versions", func(t *testing.T) {
		versionsBuilder := builder.Versions("doc_sys_id")
		assert.NotNil(t, versionsBuilder)
		assert.Equal(t, "doc_sys_id", versionsBuilder.GetPathParameters()["document_sys_id"])
	})

	t.Run("Content", func(t *testing.T) {
		contentBuilder := builder.Content("doc_sys_id")
		assert.NotNil(t, contentBuilder)
		assert.Equal(t, "doc_sys_id", contentBuilder.GetPathParameters()["document_sys_id"])
	})

	t.Run("SyncDown", func(t *testing.T) {
		syncDownBuilder := builder.SyncDown("doc_sys_id")
		assert.NotNil(t, syncDownBuilder)
		assert.Equal(t, "doc_sys_id", syncDownBuilder.GetPathParameters()["documentSysId"])
	})

	t.Run("Action", func(t *testing.T) {
		actionBuilder := builder.Action("move")
		assert.NotNil(t, actionBuilder)
		assert.Equal(t, "move", actionBuilder.GetPathParameters()["action"])
	})
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
