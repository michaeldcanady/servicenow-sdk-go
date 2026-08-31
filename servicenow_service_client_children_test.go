// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package servicenowsdkgo

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestServiceClient builds a client directly off a base request builder, bypassing the
// option plumbing so these tests exercise only the fluent chaining.
func newTestServiceClient(pathParameters map[string]string) (*ServiceNowServiceClient, *mocking.MockRequestAdapter) {
	requestAdapter := mocking.NewMockRequestAdapter()

	return &ServiceNowServiceClient{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, baseURLVariable, pathParameters),
	}, requestAdapter
}

// TestServiceNowServiceClient_Children checks that each top-level namespace accessor forwards
// the request adapter and hands the child a copy of the client's path parameters.
func TestServiceNowServiceClient_Children(t *testing.T) {
	tests := []struct {
		name  string
		build func(client *ServiceNowServiceClient) core.RequestBuilder
	}{
		{
			name:  "Now",
			build: func(client *ServiceNowServiceClient) core.RequestBuilder { return client.Now() },
		},
		{
			name:  "Cdm",
			build: func(client *ServiceNowServiceClient) core.RequestBuilder { return client.Cdm() },
		},
		{
			name:  "AppointmentBooking",
			build: func(client *ServiceNowServiceClient) core.RequestBuilder { return client.AppointmentBooking() },
		},
		{
			name:  "CustomerService",
			build: func(client *ServiceNowServiceClient) core.RequestBuilder { return client.CustomerService() },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			url := "https://example.service-now.com"
			client, requestAdapter := newTestServiceClient(map[string]string{"baseurl": url})

			child := test.build(client)

			require.NotNil(t, child)
			assert.Equal(t, map[string]string{"baseurl": url}, child.GetPathParameters())
			assert.Equal(t, requestAdapter, child.GetRequestAdapter())
		})
	}
}

// TestServiceNowServiceClient_ChildrenDoNotShareState guards the clone in each accessor:
// writing to one namespace's path parameters must not be visible to the client or a sibling.
func TestServiceNowServiceClient_ChildrenDoNotShareState(t *testing.T) {
	url := "https://example.service-now.com"
	client, _ := newTestServiceClient(map[string]string{"baseurl": url})

	now := client.Now()
	cdm := client.Cdm()

	now.GetPathParameters()["scratch"] = "value"

	assert.NotContains(t, client.GetPathParameters(), "scratch")
	assert.NotContains(t, cdm.GetPathParameters(), "scratch")
}

// TestServiceNowServiceClient_ChainDepth walks a full chain to confirm path parameters
// accumulate down the chain rather than being replaced at each hop.
func TestServiceNowServiceClient_ChainDepth(t *testing.T) {
	url := "https://example.service-now.com"
	client, _ := newTestServiceClient(map[string]string{"baseurl": url})

	table := client.Now().Table("incident")

	assert.Equal(t, map[string]string{"baseurl": url, "table": "incident"}, table.GetPathParameters())
}
