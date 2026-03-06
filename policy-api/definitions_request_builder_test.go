package policyapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestDefinitionsRequestBuilder(t *testing.T) {
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
			rB := NewDefinitionsRequestBuilderInternal(tt.pathParameters, requestAdapter)

			assert.NotNil(t, rB)
			assert.Equal(t, definitionsURLTemplate, rB.GetURLTemplate())
			assert.Equal(t, tt.pathParameters, rB.GetPathParameters())
			assert.Equal(t, requestAdapter, rB.GetRequestAdapter())
		})
	}
}
