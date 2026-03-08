package policyapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
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
