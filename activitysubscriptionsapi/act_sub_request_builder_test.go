// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package activitysubscriptionsapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestActSubRequestBuilder_Hierarchy(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewActSubRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	assert.NotNil(t, builder.Activities())
	assert.NotNil(t, builder.Facets())
}

func TestFacetsRequestBuilder_Hierarchy(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewActSubRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	facets := builder.Facets().ByContext("ctx1").ByInstance("inst1")
	assert.Equal(t, "ctx1", facets.GetPathParameters()["activity_context"])
	assert.Equal(t, "inst1", facets.GetPathParameters()["context_instance"])
}
