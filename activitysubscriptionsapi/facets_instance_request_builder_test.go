// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package activitysubscriptionsapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFacetsInstanceRequestBuilder_ToGetRequestInformation_IncludesQueryParameters(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewFacetsInstanceRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "{+baseurl}/api/now/v1/actsub/facets/{activity_context}/{context_instance}{?end_date,facets,get_activity_count,lazy_load,start_date}", requestInfo.UrlTemplate)
	assert.Equal(t, "https://example.com", requestInfo.PathParameters["baseurl"])
}
