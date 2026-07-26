package accountapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAccountRequestBuilder_Get_New(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", context.Background(), mock.Anything, mock.Anything, mock.Anything).Return(&AccountCollectionResponseMock{}, nil)

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
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, res)
			}

			// Verify request info was captured correctly
			adapter.AssertCalled(t, "Send", context.Background(), mock.Anything, mock.Anything, mock.Anything)
		})
	}
}
