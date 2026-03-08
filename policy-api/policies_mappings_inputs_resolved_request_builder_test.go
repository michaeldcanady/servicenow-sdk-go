package policyapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestNewPoliciesMappingsInputsResolvedRequestBuilderInternal(t *testing.T) {
	pathParameters := map[string]string{"baseurl": "https://example.com"}
	requestAdapter := mocking.NewMockRequestAdapter()
	builder := NewPoliciesMappingsInputsResolvedRequestBuilderRequestBuilderInternal(pathParameters, requestAdapter)
	assert.NotNil(t, builder)
}

func NewBasePoliciesMappingsInputCollectionResponse() *newInternal.BaseServiceNowCollectionResponse[*PoliciesMapping] {
	return newInternal.NewBaseServiceNowCollectionResponse[*PoliciesMapping](CreatePoliciesMappingsInputFromDiscriminatorValue)
}
