package servicenowsdkgo

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCdmRequestBuilder(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				url := "https://example.service-now.com/api"
				requestAdapter := mocking.NewMockRequestAdapter()
				builder := NewCdmRequestBuilder(url, requestAdapter)

				expected := map[string]string{
					"request-raw-url": url,
				}

				assert.NotNil(t, builder)
				assert.Equal(t, expected, builder.GetPathParameters())
				assert.Equal(t, cdmURLTemplate, builder.GetURLTemplate())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TestCdmRequestBuilder_Children checks that every child accessor forwards the request
// adapter and hands the child a copy of the parent's path parameters.
func TestCdmRequestBuilder_Children(t *testing.T) {
	tests := []struct {
		name  string
		build func(rB *CdmRequestBuilder) core.RequestBuilder
	}{
		{
			name:  "Policies",
			build: func(rB *CdmRequestBuilder) core.RequestBuilder { return rB.Policies() },
		},
		{
			name:  "Editor",
			build: func(rB *CdmRequestBuilder) core.RequestBuilder { return rB.Editor() },
		},
		{
			name:  "Changesets",
			build: func(rB *CdmRequestBuilder) core.RequestBuilder { return rB.Changesets() },
		},
		{
			name:  "Applications",
			build: func(rB *CdmRequestBuilder) core.RequestBuilder { return rB.Applications() },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			url := "https://example.service-now.com/api"
			requestAdapter := mocking.NewMockRequestAdapter()
			parent := NewCdmRequestBuilderInternal(map[string]string{"baseurl": url}, requestAdapter)

			child := test.build(parent)

			require.NotNil(t, child)
			assert.Equal(t, map[string]string{"baseurl": url}, child.GetPathParameters())
			assert.Equal(t, requestAdapter, child.GetRequestAdapter())
		})
	}
}

// TestCdmRequestBuilder_ChildrenDoNotShareState guards the clone in each accessor: writing
// to one child's path parameters must not be visible to the parent or to a sibling.
func TestCdmRequestBuilder_ChildrenDoNotShareState(t *testing.T) {
	url := "https://example.service-now.com/api"
	parent := NewCdmRequestBuilderInternal(map[string]string{"baseurl": url}, mocking.NewMockRequestAdapter())

	first := parent.Editor()
	second := parent.Changesets()

	first.GetPathParameters()["scratch"] = "value"

	assert.NotContains(t, parent.GetPathParameters(), "scratch")
	assert.NotContains(t, second.GetPathParameters(), "scratch")
}
