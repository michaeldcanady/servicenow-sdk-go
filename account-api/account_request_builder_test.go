package accountapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestAccountRequestBuilder_Builders(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewAccountRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	tests := []struct {
		name     string
		builder  any
		expected map[string]string
	}{
		{
			name:     "ByID",
			builder:  builder.ByID("test-id"),
			expected: map[string]string{"account_id": "test-id"},
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

func TestAccountItemRequestBuilder_GetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewAccountItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "account_id": "test-id"}, adapter)

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Successful Request Information Generation",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, requestInfo)
				assert.Equal(t, accountItemURLTemplate, requestInfo.UrlTemplate)
				assert.Equal(t, "test-id", requestInfo.PathParameters["account_id"])
			}
		})
	}
}
