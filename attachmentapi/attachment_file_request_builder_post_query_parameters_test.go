package attachmentapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestAttachmentFileRequestBuilderPostQueryParameters(t *testing.T) {
	tests := []struct {
		name     string
		params   *AttachmentFileRequestBuilderPostQueryParameters
		expected map[string]string
	}{
		{
			name:     "nil parameters",
			params:   nil,
			expected: map[string]string{},
		},
		{
			name:     "empty parameters",
			params:   &AttachmentFileRequestBuilderPostQueryParameters{},
			expected: map[string]string{},
		},
		{
			name: "all parameters set",
			params: &AttachmentFileRequestBuilderPostQueryParameters{
				EncryptionContext: internal.ToPointer("ctx123"),
				FileName:          internal.ToPointer("test.txt"),
				TableName:         internal.ToPointer("incident"),
				TableSysID:        internal.ToPointer("sys123"),
			},
			expected: map[string]string{
				"encryption_context": "ctx123",
				"file_name":          "test.txt",
				"table_name":         "incident",
				"table_sys_id":       "sys123",
			},
		},
		{
			name: "some parameters set",
			params: &AttachmentFileRequestBuilderPostQueryParameters{
				FileName:   internal.ToPointer("test.txt"),
				TableSysID: internal.ToPointer("sys123"),
			},
			expected: map[string]string{
				"file_name":    "test.txt",
				"table_sys_id": "sys123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestInfo := abstractions.NewRequestInformation()
			if tt.params == nil {
				requestInfo.AddQueryParameters(nil)
			} else {
				requestInfo.AddQueryParameters(*tt.params)
			}

			assert.Equal(t, tt.expected, requestInfo.QueryParameters)
		})
	}
}
