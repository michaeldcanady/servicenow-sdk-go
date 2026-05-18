package cmdbinstanceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestCmdbInstanceRequestBuilder_Builders(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCmdbInstanceRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	tests := []struct {
		name     string
		builder  any
		expected map[string]string
	}{
		{
			name:     "ByClass",
			builder:  builder.ByClass("cmdb_ci_linux_server"),
			expected: map[string]string{"className": "cmdb_ci_linux_server"},
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
