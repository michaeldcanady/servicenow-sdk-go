package attachmentapi

import (
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestAttachmentFileRequestBuilderPostQueryParameters(t *testing.T) {
	tests := []struct {
		name     string
		params   *AttachmentFileRequestBuilderPostQueryParameters
		expected map[string][]string
	}{
		{
			name:     "nil parameters",
			params:   nil,
			expected: map[string][]string{},
		},
		{
			name:     "empty parameters",
			params:   &AttachmentFileRequestBuilderPostQueryParameters{},
			expected: map[string][]string{},
		},
		{
			name: "all parameters set",
			params: &AttachmentFileRequestBuilderPostQueryParameters{
				EncryptionContext: internal.ToPointer("ctx123"),
				FileName:          internal.ToPointer("test.txt"),
				TableName:         internal.ToPointer("incident"),
				TableSysID:        internal.ToPointer("sys123"),
			},
			expected: map[string][]string{
				"encryption_context": {"ctx123"},
				"file_name":          {"test.txt"},
				"table_name":         {"incident"},
				"table_sys_id":       {"sys123"},
			},
		},
		{
			name: "some parameters set",
			params: &AttachmentFileRequestBuilderPostQueryParameters{
				FileName:   internal.ToPointer("test.txt"),
				TableSysID: internal.ToPointer("sys123"),
			},
			expected: map[string][]string{
				"file_name":    {"test.txt"},
				"table_sys_id": {"sys123"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.params == nil {
				values, err := query.Values(tt.params)
				assert.NoError(t, err)
				assert.Empty(t, values)
				return
			}

			values, err := query.Values(tt.params)
			assert.NoError(t, err)

			assert.Equal(t, len(tt.expected), len(values))
			for k, expectedVal := range tt.expected {
				assert.Equal(t, expectedVal, values[k])
			}
		})
	}
}
