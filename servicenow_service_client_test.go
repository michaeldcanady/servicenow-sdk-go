package servicenowsdkgo

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewServiceNowServiceClient(t *testing.T) {
	tests := []struct {
		name        string
		opts        func() []ServiceNowServiceClientOption
		expectedErr bool
	}{
		{
			name: "With URL and Auth",
			opts: func() []ServiceNowServiceClientOption {
				return []ServiceNowServiceClientOption{
					WithURL("https://example.com"),
					WithAuthenticationProvider(mocking.NewMockAuthenticationProvider()),
				}
			},
			expectedErr: false,
		},
		{
			name: "With Instance and Auth",
			opts: func() []ServiceNowServiceClientOption {
				return []ServiceNowServiceClientOption{
					WithInstance("test"),
					WithAuthenticationProvider(mocking.NewMockAuthenticationProvider()),
				}
			},
			expectedErr: false,
		},
		{
			name: "With RequestAdapter and URL",
			opts: func() []ServiceNowServiceClientOption {
				adapter := mocking.NewMockRequestAdapter()
				adapter.On("EnableBackingStore", mock.Anything).Return()
				adapter.On("SetBaseUrl", mock.Anything).Return()
				return []ServiceNowServiceClientOption{
					WithRequestAdapter(adapter),
					WithURL("https://example.com"),
				}
			},
			expectedErr: false,
		},
		{
			name: "Missing Auth and Adapter",
			opts: func() []ServiceNowServiceClientOption {
				return []ServiceNowServiceClientOption{
					WithURL("https://example.com"),
				}
			},
			expectedErr: true,
		},
		{
			name: "Missing URL and Instance",
			opts: func() []ServiceNowServiceClientOption {
				return []ServiceNowServiceClientOption{
					WithAuthenticationProvider(mocking.NewMockAuthenticationProvider()),
				}
			},
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := NewServiceNowServiceClient(test.opts()...)
			if test.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, client)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)
			}
		})
	}
}

func TestNewServiceNowServiceClient_LowLevel(t *testing.T) {
	// Verifying that passing a pre-configured adapter works
	requestAdapter := mocking.NewMockRequestAdapter()
	requestAdapter.On("EnableBackingStore", mock.Anything).Return()
	requestAdapter.On("SetBaseUrl", mock.Anything).Return()

	backingStore := store.BackingStoreFactoryInstance

	client, err := NewServiceNowServiceClient(
		WithRequestAdapter(requestAdapter),
		WithBackingStoreFactory(backingStore),
		WithURL("https://example.com"),
	)

	assert.NoError(t, err)
	assert.NotNil(t, client)
	requestAdapter.AssertExpectations(t)
}
