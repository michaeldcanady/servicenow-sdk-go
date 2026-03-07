package policyapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestPolicyRequestBuilder(t *testing.T) {
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
			rB := NewPolicyRequestBuilderInternal(tt.pathParameters, requestAdapter)

			assert.NotNil(t, rB)
			assert.Equal(t, policyURLTemplate, rB.GetURLTemplate())
			assert.Equal(t, tt.pathParameters, rB.GetPathParameters())
			assert.Equal(t, requestAdapter, rB.GetRequestAdapter())
		})
	}
}

func TestPolicyRequestBuilder_Mappings(t *testing.T) {
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
			var rB *PoliciesRequestBuilder
			requestAdapter := mocking.NewMockRequestAdapter()

			if !tt.nilRB {
				rB = NewPolicyRequestBuilderInternal(tt.pathParameters, requestAdapter)
			}

			mappingsRB := rB.Mappings()

			if tt.nilRB {
				assert.Nil(t, mappingsRB)
			} else {
				assert.NotNil(t, mappingsRB)
				assert.Equal(t, policiesMappingsURLTemplate, mappingsRB.GetURLTemplate())
				assert.Equal(t, tt.pathParameters, mappingsRB.GetPathParameters())
				assert.Equal(t, requestAdapter, mappingsRB.GetRequestAdapter())
			}
		})
	}
}
