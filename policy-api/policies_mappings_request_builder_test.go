package policyapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPoliciesMappingsRequestBuilder(t *testing.T) {
	tests := []struct {
		name           string
		pathParameters map[string]string
	}{
		{
			name:           "Default",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestAdapter := mocking.NewMockRequestAdapter()
			rB := NewPoliciesMappingsRequestBuilderInternal(tt.pathParameters, requestAdapter)

			assert.NotNil(t, rB)
			assert.Equal(t, policiesMappingsURLTemplate, rB.GetURLTemplate())
			assert.Equal(t, tt.pathParameters, rB.GetPathParameters())
			assert.Equal(t, requestAdapter, rB.GetRequestAdapter())
		})
	}
}

func TestPoliciesMappingsRequestBuilder_Inputs(t *testing.T) {
	tests := []struct {
		name           string
		pathParameters map[string]string
		nilRB          bool
	}{
		{
			name:           "Default",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
		},
		{
			name:  "Nil_RequestBuilder",
			nilRB: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rB *PoliciesMappingsRequestBuilder
			requestAdapter := mocking.NewMockRequestAdapter()

			if !tt.nilRB {
				rB = NewPoliciesMappingsRequestBuilderInternal(tt.pathParameters, requestAdapter)
			}

			inputsRB := rB.Inputs()

			if tt.nilRB {
				assert.Nil(t, inputsRB)
			} else {
				assert.NotNil(t, inputsRB)
				assert.Equal(t, policiesMappingsInputsURLTemplate, inputsRB.GetURLTemplate())
				assert.Equal(t, tt.pathParameters, inputsRB.GetPathParameters())
				assert.Equal(t, requestAdapter, inputsRB.GetRequestAdapter())
			}
		})
	}
}

func TestPoliciesMappingsRequestBuilder_Delete(t *testing.T) {
	tests := []struct {
		name           string
		pathParameters map[string]string
		nilRB          bool
		config         *PoliciesMappingsRequestBuilderDeleteRequestConfiguration
		setup          func(adapter *mocking.MockRequestAdapter)
		wantErr        bool
		errMessage     string
	}{
		{
			name:           "Successful",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config: &PoliciesMappingsRequestBuilderDeleteRequestConfiguration{
				QueryParameters: &PolicyMappingsRequestBuilderDeleteQueryParameters{
					AppName:        "app",
					DeployableName: "deployable",
					PolicyName:     "policy",
				},
			},
			setup: func(adapter *mocking.MockRequestAdapter) {
				adapter.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name:    "Nil_RequestBuilder",
			nilRB:   true,
			wantErr: false,
		},
		{
			name:           "Nil_Configuration",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config:         nil,
			wantErr:        true,
			errMessage:     "requestConfiguration is nil",
		},
		{
			name:           "Missing_AppName",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config: &PoliciesMappingsRequestBuilderDeleteRequestConfiguration{
				QueryParameters: &PolicyMappingsRequestBuilderDeleteQueryParameters{
					DeployableName: "deployable",
					PolicyName:     "policy",
				},
			},
			wantErr:    true,
			errMessage: "AppName is required",
		},
		{
			name:           "Missing_DeployableName",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config: &PoliciesMappingsRequestBuilderDeleteRequestConfiguration{
				QueryParameters: &PolicyMappingsRequestBuilderDeleteQueryParameters{
					AppName:    "app",
					PolicyName: "policy",
				},
			},
			wantErr:    true,
			errMessage: "DeployableName is required",
		},
		{
			name:           "Missing_PolicyName",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config: &PoliciesMappingsRequestBuilderDeleteRequestConfiguration{
				QueryParameters: &PolicyMappingsRequestBuilderDeleteQueryParameters{
					AppName:        "app",
					DeployableName: "deployable",
				},
			},
			wantErr:    true,
			errMessage: "PolicyName is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rB *PoliciesMappingsRequestBuilder
			requestAdapter := mocking.NewMockRequestAdapter()

			if tt.setup != nil {
				tt.setup(requestAdapter)
			}

			if !tt.nilRB {
				rB = NewPoliciesMappingsRequestBuilderInternal(tt.pathParameters, requestAdapter)
			}

			err := rB.Delete(context.Background(), tt.config)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMessage, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPoliciesMappingsRequestBuilder_ToDeleteRequestInformation(t *testing.T) {
	tests := []struct {
		name           string
		pathParameters map[string]string
		nilRB          bool
		config         *PoliciesMappingsRequestBuilderDeleteRequestConfiguration
	}{
		{
			name:           "Successful",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config: &PoliciesMappingsRequestBuilderDeleteRequestConfiguration{
				QueryParameters: &PolicyMappingsRequestBuilderDeleteQueryParameters{
					AppName:        "app",
					DeployableName: "deployable",
					PolicyName:     "policy",
				},
			},
		},
		{
			name:  "Nil_RequestBuilder",
			nilRB: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rB *PoliciesMappingsRequestBuilder
			requestAdapter := mocking.NewMockRequestAdapter()

			if !tt.nilRB {
				rB = NewPoliciesMappingsRequestBuilderInternal(tt.pathParameters, requestAdapter)
			}

			requestInfo, err := rB.ToDeleteRequestInformation(context.Background(), tt.config)

			if tt.nilRB {
				assert.Nil(t, requestInfo)
				assert.NoError(t, err)
			} else {
				assert.NotNil(t, requestInfo)
				assert.NoError(t, err)
				assert.Equal(t, tt.pathParameters, requestInfo.PathParameters)
				assert.Equal(t, policiesMappingsURLTemplate, requestInfo.UrlTemplate)
			}
		})
	}
}

func TestPoliciesMappingsRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name           string
		pathParameters map[string]string
		nilRB          bool
		config         *PoliciesMappingsRequestBuilderPostRequestConfiguration
		setup          func(adapter *mocking.MockRequestAdapter)
		wantErr        bool
		errMessage     string
	}{
		{
			name:           "Successful",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config: &PoliciesMappingsRequestBuilderPostRequestConfiguration{
				QueryParameters: &PoliciesMappingsRequestBuilderPostQueryParameters{
					AppName:        "app",
					DeployableName: "deployable",
					PolicyName:     "policy",
				},
			},
			setup: func(adapter *mocking.MockRequestAdapter) {
				adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(newInternal.NewBaseServiceNowItemResponse[*PoliciesMapping](CreatePoliciesMappingsInputFromDiscriminatorValue), nil)
			},
			wantErr: false,
		},
		{
			name:    "Nil_RequestBuilder",
			nilRB:   true,
			wantErr: false,
		},
		{
			name:           "Nil_Configuration",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config:         nil,
			wantErr:        true,
			errMessage:     "requestConfiguration is nil",
		},
		{
			name:           "Missing_AppName",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config: &PoliciesMappingsRequestBuilderPostRequestConfiguration{
				QueryParameters: &PoliciesMappingsRequestBuilderPostQueryParameters{
					DeployableName: "deployable",
					PolicyName:     "policy",
				},
			},
			wantErr:    true,
			errMessage: "AppName is required",
		},
		{
			name:           "Missing_DeployableName",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config: &PoliciesMappingsRequestBuilderPostRequestConfiguration{
				QueryParameters: &PoliciesMappingsRequestBuilderPostQueryParameters{
					AppName:    "app",
					PolicyName: "policy",
				},
			},
			wantErr:    true,
			errMessage: "DeployableName is required",
		},
		{
			name:           "Missing_PolicyName",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config: &PoliciesMappingsRequestBuilderPostRequestConfiguration{
				QueryParameters: &PoliciesMappingsRequestBuilderPostQueryParameters{
					AppName:        "app",
					DeployableName: "deployable",
				},
			},
			wantErr:    true,
			errMessage: "PolicyName is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rB *PoliciesMappingsRequestBuilder
			requestAdapter := mocking.NewMockRequestAdapter()

			if tt.setup != nil {
				tt.setup(requestAdapter)
			}

			if !tt.nilRB {
				rB = NewPoliciesMappingsRequestBuilderInternal(tt.pathParameters, requestAdapter)
			}

			resp, err := rB.Post(context.Background(), tt.config)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMessage, err.Error())
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				if !tt.nilRB {
					assert.NotNil(t, resp)
				}
			}
		})
	}
}

func TestPoliciesMappingsRequestBuilder_ToPostRequestInformation(t *testing.T) {
	tests := []struct {
		name           string
		pathParameters map[string]string
		nilRB          bool
		config         *PoliciesMappingsRequestBuilderPostRequestConfiguration
	}{
		{
			name:           "Successful",
			pathParameters: map[string]string{"baseurl": "https://instance.service-now.com"},
			config: &PoliciesMappingsRequestBuilderPostRequestConfiguration{
				QueryParameters: &PoliciesMappingsRequestBuilderPostQueryParameters{
					AppName:        "app",
					DeployableName: "deployable",
					PolicyName:     "policy",
				},
			},
		},
		{
			name:  "Nil_RequestBuilder",
			nilRB: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rB *PoliciesMappingsRequestBuilder
			requestAdapter := mocking.NewMockRequestAdapter()

			if !tt.nilRB {
				rB = NewPoliciesMappingsRequestBuilderInternal(tt.pathParameters, requestAdapter)
			}

			requestInfo, err := rB.ToPostRequestInformation(context.Background(), tt.config)

			if tt.nilRB {
				assert.Nil(t, requestInfo)
				assert.NoError(t, err)
			} else {
				assert.NotNil(t, requestInfo)
				assert.NoError(t, err)
				assert.Equal(t, tt.pathParameters, requestInfo.PathParameters)
				assert.Equal(t, policiesMappingsURLTemplate, requestInfo.UrlTemplate)
			}
		})
	}
}
