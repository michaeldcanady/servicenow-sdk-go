// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCsdmRequestBuilderInternal(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewCsdmRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	require.NotNil(t, builder)
	assert.Equal(t, csdmAppServiceURLTemplate, builder.GetURLTemplate())
}

func TestCsdmRequestBuilder_FindService(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewCsdmRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	findBuilder := builder.FindService()
	assert.NotNil(t, findBuilder)
	assert.Equal(t, findServiceURLTemplate, findBuilder.GetURLTemplate())

	name := "Email_East"
	number := "SNSVC0001018"
	config := &FindServiceRequestConfiguration{
		QueryParameters: &FindServiceQueryParameters{
			Name:   &name,
			Number: &number,
		},
	}

	requestInfo, err := findBuilder.ToGetRequestInformation(context.Background(), config)
	require.NoError(t, err)

	uri, err := requestInfo.GetUri()
	require.NoError(t, err)
	assert.Equal(t, "https://example.service-now.com/api/now/cmdb/csdm/app_service/find_service?name=Email_East&number=SNSVC0001018", uri.String())
}

func TestCsdmRequestBuilder_RegisterService(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewCsdmRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	registerBuilder := builder.RegisterService()
	assert.NotNil(t, registerBuilder)
	assert.Equal(t, registerServiceURLTemplate, registerBuilder.GetURLTemplate())
}

func TestCsdmRequestBuilder_ByID(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	builder := NewCsdmRequestBuilderInternal(map[string]string{"baseurl": "https://example.service-now.com"}, adapter)

	itemBuilder := builder.ByID("service123")
	assert.NotNil(t, itemBuilder)
	assert.Equal(t, csdmAppServiceItemURLTemplate, itemBuilder.GetURLTemplate())
	assert.Equal(t, "service123", itemBuilder.GetPathParameters()[sysIDKey])
}
