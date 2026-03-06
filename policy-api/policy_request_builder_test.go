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

func TestPolicyRequestBuilder_Definitions(t *testing.T) {
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
			definitionsRB := rB.Definitions()

			assert.NotNil(t, definitionsRB)
			assert.Equal(t, definitionsURLTemplate, definitionsRB.GetURLTemplate())
			assert.Equal(t, tt.pathParameters, definitionsRB.GetPathParameters())
			assert.Equal(t, requestAdapter, definitionsRB.GetRequestAdapter())
		})
	}
}
