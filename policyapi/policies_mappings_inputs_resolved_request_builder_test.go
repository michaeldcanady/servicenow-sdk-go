package policyapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewPoliciesMappingsInputsResolvedRequestBuilderInternal(t *testing.T) {
	pathParameters := map[string]string{"baseurl": "https://example.com"}
	requestAdapter := mocking.NewMockRequestAdapter()
	builder := NewPoliciesMappingsInputsResolvedRequestBuilderRequestBuilderInternal(pathParameters, requestAdapter)
	assert.NotNil(t, builder)
}

func NewBasePoliciesMappingsInputCollectionResponse() *core.BaseServiceNowCollectionResponse[*PoliciesMapping] {
	return core.NewBaseServiceNowCollectionResponse[*PoliciesMapping](CreatePoliciesMappingsInputFromDiscriminatorValue)
}
