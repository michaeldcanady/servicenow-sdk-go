package actsubapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestActivitiesRequestBuilder_Get(t *testing.T) {
	type testCase struct {
		name      string
		config    *ActivitiesRequestBuilderGetRequestConfiguration
		mockRes   interface{}
		mockErr   error
		expectErr bool
	}

	tests := []testCase{
		{
			name:      "Success",
			config:    nil,
			mockRes:   core.NewBaseServiceNowCollectionResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue),
			mockErr:   nil,
			expectErr: false,
		},
		{
			name:      "Error",
			config:    nil,
			mockRes:   nil,
			mockErr:   assert.AnError,
			expectErr: true,
		},
		{
			name:      "NilResponse",
			config:    nil,
			mockRes:   nil,
			mockErr:   nil,
			expectErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			builder := NewActivitiesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tc.mockRes, tc.mockErr)

			resp, err := builder.Get(context.Background(), tc.config)

			if tc.expectErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				if tc.mockRes != nil {
					assert.Equal(t, tc.mockRes, resp)
				} else {
					assert.Nil(t, resp)
				}
			}
		})
	}
}

func TestActivitiesRequestBuilder_ToGetRequestInformation_Extra(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewActivitiesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "{+baseurl}/api/now/v1/actsub/activities", requestInfo.UrlTemplate)
	assert.Equal(t, "https://example.com", requestInfo.PathParameters["baseurl"])
}
