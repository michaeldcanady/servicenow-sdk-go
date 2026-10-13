// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmeditorapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// These tests exercise the comma-ok type-assertion guard: when the adapter
// returns a non-nil value whose concrete type doesn't match the expected
// response type, the verb method must return an error instead of panicking.

func TestNodesRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewNodesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestNodesRequestBuilder_Post_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewNodesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Post(context.Background(), NewNodeCreateRequest(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestNodeItemRequestBuilder_Put_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewNodeItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "node_sys_id": "node123"}, adapter)

	resp, err := builder.Put(context.Background(), NewNodeUpdateRequest(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestValidationRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewValidationRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}
