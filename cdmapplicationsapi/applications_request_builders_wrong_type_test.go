// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

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

func TestDeployablesRequestBuilder_Put_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewDeployablesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Put(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestSharedComponentsRequestBuilder_Put_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewSharedComponentsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Put(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUploadStatusItemRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewUploadStatusItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestExportsRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewExportsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestExportItemStatusRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewExportItemStatusRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestSharedLibrariesComponentsApplicationsRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewSharedLibrariesComponentsApplicationsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUploadsComponentsRequestBuilder_Post_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewUploadsComponentsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Post(context.Background(), NewComponentUploadRequest(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUploadsComponentsVarsRequestBuilder_Post_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewUploadsComponentsVarsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Post(context.Background(), NewComponentVarsUploadRequest(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUploadsCollectionsRequestBuilder_Post_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewUploadsCollectionsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Post(context.Background(), NewCollectionUploadRequest(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUploadsCollectionsFileRequestBuilder_Post_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewUploadsCollectionsFileRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Post(context.Background(), NewMedia("application/octet-stream", []byte("data")), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUploadsDeployablesFileRequestBuilder_Post_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(mocking.NewMockParsable(), nil)
	builder := NewUploadsDeployablesFileRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Post(context.Background(), NewMedia("application/octet-stream", []byte("data")), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestExportItemContentRequestBuilder_Get_WrongTypeResponse(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("SendPrimitive", mock.Anything, mock.Anything, "[]byte", mock.Anything).
		Return("not-bytes", nil)
	builder := NewExportItemContentRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	resp, err := builder.Get(context.Background(), nil)

	require.Error(t, err)
	assert.Nil(t, resp)
}
