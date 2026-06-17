package policyapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestPoliciesMappingsInputsRequestBuilder(t *testing.T) {
	pathParameters := map[string]string{"baseurl": "https://example.com"}
	requestAdapter := mocking.NewMockRequestAdapter()
	builder := NewPoliciesMappingsInputsRequestBuilderInternal(pathParameters, requestAdapter)
	assert.NotNil(t, builder)
}

func TestPoliciesMappingsInputsRequestBuilder_Resolved(t *testing.T) {
	tests := []struct {
		name  string
		nilRB bool
	}{
		{"Default", false},
		{"NilRB", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rB *PoliciesMappingsInputsRequestBuilder
			if !tt.nilRB {
				pathParameters := map[string]string{"baseurl": "https://example.com"}
				requestAdapter := mocking.NewMockRequestAdapter()
				rB = NewPoliciesMappingsInputsRequestBuilderInternal(pathParameters, requestAdapter)
			}

			resolved := rB.Resolved()
			if tt.nilRB {
				assert.Nil(t, resolved)
			} else {
				assert.NotNil(t, resolved)
			}
		})
	}
}
