// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAppServiceItemRequestBuilderInternal(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewAppServiceItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com", "sys_id": "service123"}, adapter)

	require.NotNil(t, builder)
	assert.Equal(t, appServiceItemURLTemplate, builder.GetURLTemplate())
}

func TestAppServiceItemRequestBuilder_GetContent(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewAppServiceItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com", "sys_id": "service123"}, adapter)

	getContentBuilder := builder.GetContent()
	assert.NotNil(t, getContentBuilder)
	assert.Equal(t, getContentURLTemplate, getContentBuilder.GetURLTemplate())
}
