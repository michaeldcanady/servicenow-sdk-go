package accountapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/new/mocking"
	"github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestAccountRequestBuilder_Get_New(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{
		Response: &AccountCollectionResponseMock{},
	}
	builder := NewAccountRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Successful Get",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := builder.Get(context.Background(), nil)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}

			// Verify request info was captured correctly
			assert.NotNil(t, adapter.LastRequest)
			assert.Equal(t, abstractions.GET, adapter.LastRequest.Method)
		})
	}
}
