package attachmentapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestAttachmentRequestBuilderGetQueryParameters(t *testing.T) {
	tests := []struct {
		name     string
		params   AttachmentRequestBuilderGetQueryParameters
		expected map[string]string
	}{
		{
			name: "Standard parameters",
			params: AttachmentRequestBuilderGetQueryParameters{
				Limit:  1000,
				Offset: 500,
				Query:  "field1=value1",
			},
			expected: map[string]string{
				"sysparm_limit":  "1000",
				"sysparm_offset": "500",
				"sysparm_query":  "field1=value1",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := core.ToQueryMap(test.params)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
