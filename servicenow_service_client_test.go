package servicenowsdkgo

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewServiceNowServiceClientWithOptions(t *testing.T) {
	tests := []struct {
		name        string
		opts        []ServiceNowServiceClientOption
		expectedErr bool
	}{
		{
			name: "With URL",
			opts: []ServiceNowServiceClientOption{
				withURL("https://example.com"),
			},
			expectedErr: false,
		},
		{
			name: "With Instance",
			opts: []ServiceNowServiceClientOption{
				withInstance("test"),
			},
			expectedErr: false,
		},
		{
			name:        "Neither URL nor Instance",
			opts:        []ServiceNowServiceClientOption{},
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			authProvider := mocking.NewMockAuthenticationProvider()
			client, err := NewServiceNowServiceClientWithOptions(authProvider, test.opts...)
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

func TestNewServiceNowServiceClient(t *testing.T) {
	tests := []struct {
		name        string
		baseURL     string
		expectedErr bool
	}{
		{
			name:        "Valid URL",
			baseURL:     "https://example.com",
			expectedErr: false,
		},
		{
			name:        "Empty URL",
			baseURL:     "",
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requestAdapter := mocking.NewMockRequestAdapter()
			if !test.expectedErr {
				requestAdapter.On("EnableBackingStore", mock.Anything).Return()
				requestAdapter.On("SetBaseUrl", mock.Anything).Return()
			}

			backingStore := store.BackingStoreFactoryInstance

			client, err := NewServiceNowServiceClient(requestAdapter, backingStore, test.baseURL)
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
