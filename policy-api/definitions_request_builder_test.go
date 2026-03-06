package policyapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDefinitionsRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name      string
		nilRB     bool
		setupMock func(*mocking.MockRequestAdapter)
	}{
		{
			name: "Successful",
			setupMock: func(ra *mocking.MockRequestAdapter) {
				resp := newInternal.NewBaseServiceNowCollectionResponse[*PolicyDefinition](CreatePolicyDefinitionFromDiscriminatorValue)
				ra.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
			},
		},
		{
			name:  "Nil_RequestBuilder",
			nilRB: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rB *DefinitionsRequestBuilder
			ra := mocking.NewMockRequestAdapter()

			if !tt.nilRB {
				rB = NewDefinitionsRequestBuilderInternal(map[string]string{"baseurl": "https://instance.service-now.com"}, ra)
				if tt.setupMock != nil {
					tt.setupMock(ra)
				}
			}

			resp, err := rB.Get(context.Background(), nil)

			if tt.nilRB {
				assert.Nil(t, resp)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, resp)
				assert.Nil(t, err)
				ra.AssertExpectations(t)
			}
		})
	}
}

func TestDefinitionsRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []struct {
		name           string
		nilRB          bool
		config         *DefinitionsRequestBuilderGetRequestConfiguration
		expectedMethod abstractions.HttpMethod
	}{
		{
			name:           "Default",
			expectedMethod: abstractions.GET,
		},
		{
			name:  "Nil_RequestBuilder",
			nilRB: true,
		},
		{
			name: "With_Config",
			config: &DefinitionsRequestBuilderGetRequestConfiguration{
				QueryParameters: &DefinitionsRequestBuilderGetQueryParameters{
					Limit: 10,
				},
			},
			expectedMethod: abstractions.GET,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rB *DefinitionsRequestBuilder
			ra := mocking.NewMockRequestAdapter()

			if !tt.nilRB {
				rB = NewDefinitionsRequestBuilderInternal(map[string]string{"baseurl": "https://instance.service-now.com"}, ra)
			}

			reqInfo, err := rB.ToGetRequestInformation(context.Background(), tt.config)

			if tt.nilRB {
				assert.Nil(t, reqInfo)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, reqInfo)
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedMethod, reqInfo.Method)
				if tt.config != nil && tt.config.QueryParameters != nil {
					assert.Contains(t, reqInfo.QueryParameters, "sysparm_limit")
				}
			}
		})
	}
}
