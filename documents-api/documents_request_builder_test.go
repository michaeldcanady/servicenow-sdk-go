package documentsapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestDocumentsRequestBuilder_Builders(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewDocumentsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

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
			name:     "VersionState",
			builder:  builder.VersionState("version_sys_id"),
			expected: map[string]string{"version_sys_id": "version_sys_id"},
		},
		{
			name:     "Attach",
			builder:  builder.Attach("provider_id"),
			expected: map[string]string{"provider_id": "provider_id"},
		},
		{
			name:    "Delete",
			builder: builder.Delete(),
		},
		{
			name:     "Versions",
			builder:  builder.Versions("doc_sys_id"),
			expected: map[string]string{"document_sys_id": "doc_sys_id"},
		},
		{
			name:     "Content",
			builder:  builder.Content("doc_sys_id"),
			expected: map[string]string{"document_sys_id": "doc_sys_id"},
		},
		{
			name:     "SyncDown",
			builder:  builder.SyncDown("doc_sys_id"),
			expected: map[string]string{"documentSysId": "doc_sys_id"},
		},
		{
			name:     "Action",
			builder:  builder.Action("move"),
			expected: map[string]string{"action": "move"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.builder)
			if tt.expected != nil {
				rb := tt.builder.(internal.RequestBuilder)
				for k, v := range tt.expected {
					assert.Equal(t, v, rb.GetPathParameters()[k])
				}
			}
		})
	}
}
