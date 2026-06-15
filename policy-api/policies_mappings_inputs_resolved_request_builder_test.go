package policyapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewPoliciesMappingsInputsResolvedRequestBuilderInternal(t *testing.T) {
	pathParameters := map[string]string{"baseurl": "https://example.com"}
	requestAdapter := mocking.NewMockRequestAdapter()
	builder := NewPoliciesMappingsInputsResolvedRequestBuilderRequestBuilderInternal(pathParameters, requestAdapter)
	assert.NotNil(t, builder)
}

func NewBasePoliciesMappingsInputCollectionResponse() *internal.BaseServiceNowCollectionResponse[*PoliciesMapping] {
	return internal.NewBaseServiceNowCollectionResponse[*PoliciesMapping](CreatePoliciesMappingsInputFromDiscriminatorValue)
}
