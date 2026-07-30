package servicenowsdkgo

import (
	"maps"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const nowTestBaseURL = "https://example.service-now.com/api"

func TestNewServiceNowRequestBuilder(t *testing.T) {
	requestAdapter := mocking.NewMockRequestAdapter()

	builder := NewServiceNowRequestBuilder(nowTestBaseURL, requestAdapter)

	require.NotNil(t, builder)
	assert.Equal(t, map[string]string{"request-raw-url": nowTestBaseURL}, builder.GetPathParameters())
	assert.Equal(t, nowURLTemplate, builder.GetURLTemplate())
	assert.Equal(t, requestAdapter, builder.GetRequestAdapter())
}

func TestNewServiceNowRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name           string
		pathParameters map[string]string
		expected       map[string]string
	}{
		{
			name:           "with base url path parameter",
			pathParameters: map[string]string{"baseurl": nowTestBaseURL},
			expected:       map[string]string{"baseurl": nowTestBaseURL},
		},
		{
			name:           "with empty path parameters",
			pathParameters: map[string]string{},
			expected:       map[string]string{},
		},
		{
			// The base builder normalizes a nil map to an empty one, so callers can
			// always write to the returned map without a nil check.
			name:           "nil path parameters are normalized to an empty map",
			pathParameters: nil,
			expected:       map[string]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requestAdapter := mocking.NewMockRequestAdapter()

			builder := NewServiceNowRequestBuilderInternal(test.pathParameters, requestAdapter)

			require.NotNil(t, builder)
			assert.Equal(t, test.expected, builder.GetPathParameters())
			assert.Equal(t, nowURLTemplate, builder.GetURLTemplate())
			assert.Equal(t, requestAdapter, builder.GetRequestAdapter())
		})
	}
}

// TestNowRequestBuilder_Children checks that every child accessor forwards the request
// adapter and hands the child a copy of the parent's path parameters. The copy matters:
// each accessor clones so that adding a path parameter deeper in the chain cannot leak
// back into the parent builder.
func TestNowRequestBuilder_Children(t *testing.T) {
	tests := []struct {
		name  string
		build func(rB *NowRequestBuilder) core.RequestBuilder
		// extra holds path parameters the accessor is expected to add on top of the parent's.
		extra map[string]string
	}{
		{
			name:  "Table",
			build: func(rB *NowRequestBuilder) core.RequestBuilder { return rB.Table("incident") },
			extra: map[string]string{"table": "incident"},
		},
		{
			name:  "Stats",
			build: func(rB *NowRequestBuilder) core.RequestBuilder { return rB.Stats("incident") },
			extra: map[string]string{"table": "incident"},
		},
		{
			name:  "Attachment",
			build: func(rB *NowRequestBuilder) core.RequestBuilder { return rB.Attachment() },
		},
		{
			name:  "Batch",
			build: func(rB *NowRequestBuilder) core.RequestBuilder { return rB.Batch() },
		},
		{
			name:  "Documents",
			build: func(rB *NowRequestBuilder) core.RequestBuilder { return rB.Documents() },
		},
		{
			name:  "Cmdb",
			build: func(rB *NowRequestBuilder) core.RequestBuilder { return rB.Cmdb() },
		},
		{
			name:  "Account",
			build: func(rB *NowRequestBuilder) core.RequestBuilder { return rB.Account() },
		},
		{
			name:  "ActSub",
			build: func(rB *NowRequestBuilder) core.RequestBuilder { return rB.ActSub() },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requestAdapter := mocking.NewMockRequestAdapter()
			parent := NewServiceNowRequestBuilderInternal(map[string]string{"baseurl": nowTestBaseURL}, requestAdapter)

			child := test.build(parent)

			require.NotNil(t, child)
			assert.Equal(t, requestAdapter, child.GetRequestAdapter())

			expected := map[string]string{"baseurl": nowTestBaseURL}
			maps.Copy(expected, test.extra)
			assert.Equal(t, expected, child.GetPathParameters())

			// The parent must be untouched by whatever the child added.
			assert.Equal(t, map[string]string{"baseurl": nowTestBaseURL}, parent.GetPathParameters())
		})
	}
}

// TestNowRequestBuilder_ChildrenDoNotShareState guards the clone in each accessor: mutating
// a child's path parameters must not be visible to the parent or to a sibling.
func TestNowRequestBuilder_ChildrenDoNotShareState(t *testing.T) {
	parent := NewServiceNowRequestBuilderInternal(
		map[string]string{"baseurl": nowTestBaseURL},
		mocking.NewMockRequestAdapter(),
	)

	first := parent.Table("incident")
	second := parent.Table("problem")

	assert.Equal(t, "incident", first.GetPathParameters()["table"])
	assert.Equal(t, "problem", second.GetPathParameters()["table"])
	assert.NotContains(t, parent.GetPathParameters(), "table")
}
